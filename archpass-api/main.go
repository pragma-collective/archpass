package main

import (
	"github.com/pragma-collective/archpass/internal/adapter"
)

func main() {
	r := adapter.Router()

	// Start the ticket minter service
	// note(jhudiel) - don't enable yet
	// need to implement a processing status
	go startTicketMinter()

	r.Logger.Fatal(r.Start(":8000"))
}
