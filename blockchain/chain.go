package blockchain

import (
	"sync"

	"github.com/kimmk4369/mk-coin/db"
	"github.com/kimmk4369/mk-coin/utils"
)

type blockchain struct {
	NewestHash string `json:newestHash`
	Height     int    `json:height`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// singleton
func Blockchain() *blockchain {
	if b == nil {
		// 한번만 실행
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}