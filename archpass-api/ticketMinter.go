package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/pragma-collective/archpass/ent"
	"github.com/pragma-collective/archpass/internal/adapter/repository"
	"github.com/pragma-collective/archpass/internal/application/contract"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

type TicketMinter struct {
	client   *ethclient.Client
	contract *contract.Event
	auth     *bind.TransactOpts
}

type MintResult struct {
	TokenID        string
	MintTxHash     string
	TransferTxHash string
}

// resetTransactOpts updates the nonce and gas settings for a new transaction
func (m *TicketMinter) resetTransactOpts(value *big.Int) error {
	nonce, err := m.client.PendingNonceAt(context.Background(), m.auth.From)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := m.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}

	m.auth.Nonce = big.NewInt(int64(nonce))
	m.auth.Value = value
	m.auth.GasPrice = gasPrice

	return nil
}

// estimateAndSetGas estimates gas for a transaction and sets it with a buffer
func (m *TicketMinter) estimateAndSetGas(msg ethereum.CallMsg) error {
	gasLimit, err := m.client.EstimateGas(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("failed to estimate gas: %w", err)
	}

	m.auth.GasLimit = gasLimit * 12 / 10 // Add 20% buffer
	return nil
}

// waitForTransaction waits for a transaction to be mined and checks its status
func (m *TicketMinter) waitForTransaction(tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), m.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed waiting for transaction to be mined: %w", err)
	}

	if receipt.Status == 0 {
		return nil, fmt.Errorf("transaction failed")
	}

	return receipt, nil
}

func NewTicketMinter(chainUrl, contractAddress, privateKey string) (*TicketMinter, error) {
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum client: %w", err)
	}

	eventContract, err := contract.NewEvent(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate contract: %w", err)
	}

	auth, err := createTransactOpts(privateKey, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction opts: %w", err)
	}

	return &TicketMinter{
		client:   client,
		contract: eventContract,
		auth:     auth,
	}, nil
}

// extractTokenID extracts the token ID from transaction receipt logs
func (m *TicketMinter) extractTokenID(receipt *types.Receipt) (*big.Int, error) {
	transferEventSignature := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	for _, log := range receipt.Logs {
		if len(log.Topics) == 4 && log.Topics[0] == transferEventSignature {
			return new(big.Int).SetBytes(log.Topics[3].Bytes()), nil
		}
	}

	return nil, fmt.Errorf("could not find tokenId in transaction logs")
}

func createTransactOpts(privateKeyHex string, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to get public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain id: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	return auth, nil
}

func (m *TicketMinter) MintTicket(order *ent.Order) (*MintResult, error) {
	ticketAddress := common.HexToAddress(order.Edges.Ticket.ContractAddress)
	recipient := common.HexToAddress(order.WalletAddress)

	// Check if ticket exists
	exists, err := m.contract.DoesTicketExist(&bind.CallOpts{}, ticketAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check ticket existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("ticket does not exist at address: %s", order.Edges.Ticket.ContractAddress)
	}

	// Prepare for minting
	if err := m.resetTransactOpts(big.NewInt(0)); err != nil {
		return nil, fmt.Errorf("failed to reset transaction options: %w", err)
	}

	// Estimate gas for minting
	parsed, err := contract.EventMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ABI: %w", err)
	}

	fmt.Println(m.auth.From)
	mintData, err := parsed.Pack("mintNFT", ticketAddress, recipient, order.Edges.Ticket.BaseTokenURI)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data: %w", err)
	}

	eventAddr := common.HexToAddress(order.Edges.Event.ContractAddress)
	if err := m.estimateAndSetGas(ethereum.CallMsg{
		From:  m.auth.From,
		To:    &eventAddr,
		Value: big.NewInt(0),
		Data:  mintData,
	}); err != nil {
		return nil, err
	}

	// Mint NFT
	mintTx, err := m.contract.MintNFT(m.auth, ticketAddress, recipient, order.Edges.Ticket.BaseTokenURI)
	if err != nil {
		return nil, fmt.Errorf("failed to mint NFT: %w", err)
	}

	// Wait for mint transaction
	mintReceipt, err := m.waitForTransaction(mintTx)
	if err != nil {
		return nil, fmt.Errorf("mint transaction: %w", err)
	}

	// Extract tokenId from logs
	tokenId, err := m.extractTokenID(mintReceipt)
	if err != nil {
		return nil, err
	}

	// Prepare for transfer
	ticketContract, err := contract.NewTicket(ticketAddress, m.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create NFT contract instance: %w", err)
	}

	// Reset transaction options for transfer
	if err := m.resetTransactOpts(big.NewInt(0)); err != nil {
		return nil, fmt.Errorf("failed to reset transaction options for transfer: %w", err)
	}

	// Estimate gas for transfer
	userAddress := common.HexToAddress(order.WalletAddress)
	ticketParsed, err := contract.TicketMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket abi: %w", err)
	}

	transferData, err := ticketParsed.Pack("transferFrom", m.auth.From, userAddress, tokenId)
	if err != nil {
		return nil, fmt.Errorf("failed to pack transfer data: %w", err)
	}

	if err := m.estimateAndSetGas(ethereum.CallMsg{
		From:  m.auth.From,
		To:    &ticketAddress,
		Value: big.NewInt(0),
		Data:  transferData,
	}); err != nil {
		return nil, err
	}

	// Transfer NFT
	transferTx, err := ticketContract.TransferFrom(m.auth, m.auth.From, userAddress, tokenId)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer NFT to user: %w", err)
	}

	fmt.Printf("Transfer transaction submitted: %s\n", transferTx.Hash().Hex())

	// Wait for transfer transaction
	_, err = m.waitForTransaction(transferTx)
	if err != nil {
		return nil, fmt.Errorf("transfer transaction: %w", err)
	}

	return &MintResult{
		TokenID:        tokenId.String(),
		MintTxHash:     mintTx.Hash().Hex(),
		TransferTxHash: transferTx.Hash().Hex(),
	}, nil
}

func startTicketMinterService(ctx context.Context) {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	chainUrl, ok := os.LookupEnv("CHAIN_URL")
	if !ok {
		log.Fatal("Cannot get url for eth-client")
	}

	walletPrivateKey, ok := os.LookupEnv("ARCHPASS_WALLET_PRIVATE_KEY")
	if !ok {
		log.Fatal("Cannot get private key for archpass wallet")
	}

	// Triggers every 5 minutes
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	processOrders(ctx, chainUrl, walletPrivateKey)

	// Then wait for ticker
	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down ticket minter service...")
			return
		case <-ticker.C:
			processOrders(ctx, chainUrl, walletPrivateKey)
		}
	}
}

func processOrders(ctx context.Context, chainUrl, walletPrivateKey string) {
	log.Println("Starting to process orders...")

	orderRepo := repository.NewOrderRepository(&ctx)
	orders, err := orderRepo.GetPaidOrders(100)
	if err != nil {
		log.Printf("Failed to load orders: %v", err)
		return
	}

	if len(orders) == 0 {
		log.Println("No new orders to process")
		return
	}

	// Group orders by event contract address
	ordersByEvent := make(map[string][]*ent.Order)
	for _, order := range orders {
		eventAddr := order.Edges.Event.ContractAddress
		ordersByEvent[eventAddr] = append(ordersByEvent[eventAddr], order)
	}

	// Process orders by event
	for eventAddr, eventOrders := range ordersByEvent {
		log.Printf("Processing %d orders for event %s", len(eventOrders), eventAddr)

		// Create one minter per event
		minter, err := NewTicketMinter(chainUrl, eventAddr, walletPrivateKey)
		if err != nil {
			log.Printf("Failed to initialize ticket minter for event %s: %v", eventAddr, err)
			continue
		}

		// Process all orders for this event
		for _, order := range eventOrders {
			log.Printf("Processing order ID: %s", order.ID)

			mintDto, err := minter.MintTicket(order)
			if err != nil {
				log.Printf("Failed to mint ticket for order %s: %v", order.ID, err)
				continue
			}

			tokenIdInt, err := strconv.Atoi(mintDto.TokenID)
			if err != nil {
				log.Printf("Failed to parse token id %s for order %s: %v", mintDto.TokenID, order.ID, err)
				continue
			}

			if err := orderRepo.UpdateOrderMint(order.ID, tokenIdInt, mintDto.MintTxHash, mintDto.TransferTxHash); err != nil {
				log.Printf("Failed to store token id %s for order %s: %v", mintDto.TokenID, order.ID, err)
				continue
			}

			log.Printf("Successfully processed order %s with token ID %s", order.ID, mintDto.TokenID)
		}
	}
}

func startTicketMinter() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start ticket minter service
	startTicketMinterService(ctx)
}
