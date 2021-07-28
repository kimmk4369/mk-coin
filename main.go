package main

import (
	"github.com/kimmk4369/mk-coin/blockchain"
	"github.com/kimmk4369/mk-coin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
