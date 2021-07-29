package main

import (
	"github.com/kimmk4369/mk-coin/cli"
	"github.com/kimmk4369/mk-coin/db"
)

func main() {
	defer db.Close()

	cli.Start()
}
