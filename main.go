package main

import (
	"github.com/kimmk4369/mk-coin/explorer"
	"github.com/kimmk4369/mk-coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
