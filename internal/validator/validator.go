package validator

import (
	"time"

	"github.com/VindexAdmin/VindexChain/internal/blockchain"
	"github.com/VindexAdmin/VindexChain/internal/txpool"
)

// Run ejecuta el proceso del validador
func Run(bc *blockchain.Blockchain, address string) {
	for {
		txs := txpool.GetTransactions()
		if len(txs) > 0 {
			bc.AddBlock(txs, address)
			txpool.Clear()
		}
		time.Sleep(5 * time.Second)
	}
}
