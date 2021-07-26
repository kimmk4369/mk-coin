package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:newestHash`
	Height     int    `json:height`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	createBlock(data, b.NewestHash, b.Height)
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
