package main

import (
	"fmt"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	//f300664bacbb5512ad3fb5e29b0190f52bb6293759cdfacc5c35ddfe8481c2a6
	//aa9df154d186c65c550051017264b501a912cf97ed89a40536cb3d88d81dddcfdd58709a5366b13415a2e764485fed3ea29aeddc82e7baae4ed4ce2a5cdea963
}
