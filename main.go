package main

import (
	"github.com/kimmk4369/mk-coin/blockchain"
)

func main() {
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
}
