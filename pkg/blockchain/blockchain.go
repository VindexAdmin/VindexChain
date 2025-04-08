package blockchain

import (
    "github.com/VindexAdmin/VindexChain/pkg/core"
)

// Blockchain representa toda la cadena de bloques
type Blockchain struct {
    Blocks []*core.Block
}

// NewBlockchain crea la cadena con el bloque génesis
func NewBlockchain() *Blockchain {
    genesis := core.NewBlock([]core.Transaction{}, []byte{})
    return &Blockchain{[]*core.Block{genesis}}
}

// AddBlock agrega un nuevo bloque a la cadena
func (bc *Blockchain) AddBlock(transactions []core.Transaction) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := core.NewBlock(transactions, prevBlock.Hash)
    bc.Blocks = append(bc.Blocks, newBlock)
}
