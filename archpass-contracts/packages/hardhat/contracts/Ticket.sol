//SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract Ticket is Initializable, ERC721URIStorageUpgradeable, OwnableUpgradeable, AccessControlUpgradeable {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    uint256 private _tokenIdCounter;
    uint256 public maxSupply;
    bool public allowMint = true;

    mapping(address => uint256) public userMintCount; 

    function initialize(
        string memory name,
        string memory symbol,
        uint256 _maxSupply,
        address owner,
        address admin
    ) public initializer {
        __ERC721_init(name, symbol);
        __Ownable_init(owner);
        __AccessControl_init();

        _grantRole(DEFAULT_ADMIN_ROLE, owner);
        _grantRole(ADMIN_ROLE, admin);

        maxSupply = _maxSupply;
        allowMint = true;
        _tokenIdCounter = 0;
    }

    function mintNFT(address to, string memory tokenURI) public returns (uint256) {
       require(
            owner() == to || hasRole(ADMIN_ROLE, to),
            "UNAUTHORIZED_TO_MINT"
        );
        require(_tokenIdCounter < maxSupply, "Max supply reached");
        require(allowMint, "Minting is not allowed");

        _tokenIdCounter++;
        uint256 newItemId = _tokenIdCounter;
        
        _mint(to, newItemId);
        _setTokenURI(newItemId, tokenURI);

        userMintCount[to]++;

        return newItemId;
    }

    function setAllowMint(bool _allowMint) public onlyOwner {
        allowMint = _allowMint;
    }

    function setMaxSupply(uint256 _maxSupply) public onlyOwner {
        maxSupply = _maxSupply;
    }

    function totalSupply() public view returns (uint256) {
        return _tokenIdCounter;
    }

    // Override required function
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721URIStorageUpgradeable, AccessControlUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
