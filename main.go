package main

import (
	"fmt"
	"goblockchain/block"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	wallet_M := wallet.NewWallet()
	wallet_A := wallet.NewWallet()
	wallet_B := wallet.NewWallet()

	// wallet側
	t := wallet.NewTransaction(wallet_A.PrivateKey(), wallet_A.PublicKey(),
		wallet_A.BlockchainAddress(), wallet_B.BlockchainAddress(), 1.0)
	// blockchain側
	blockchain := block.NewBlockchain(wallet_M.BlockchainAddress())
	isAdded := blockchain.AddTransaction(wallet_A.BlockchainAddress(),
		wallet_B.BlockchainAddress(), 1.0, wallet_A.PublicKey(), t.GenerateSignature())

	fmt.Println(isAdded)

	blockchain.Mining()
	blockchain.Print()
	fmt.Println("A => ", blockchain.CalculateTotalAmount(wallet_A.BlockchainAddress()))
	fmt.Println("B => ", blockchain.CalculateTotalAmount(wallet_B.BlockchainAddress()))
	fmt.Println("M => ", blockchain.CalculateTotalAmount(wallet_M.BlockchainAddress()))

}
