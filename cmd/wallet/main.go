package main

import (
	"fmt"
	"log"

	"github.com/simplysabir/mpc-tss-wallet/internal/keygen"
	"github.com/simplysabir/mpc-tss-wallet/internal/transfer"
)

func main() {
    fmt.Println("MPC TSS Wallet")

    // Key Generation
    err := keygen.GenerateKeys(21, 15)
    if err != nil {
        log.Fatalf("Error generating keys: %v", err)
    }

    // Asset Transfer (example)
    err = transfer.TransferAsset("secp256k1", "0x1234...", "0x5678...", 1.0)
    if err != nil {
        log.Fatalf("Error transferring asset: %v", err)
    }
}