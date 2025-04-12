package blockchain

import (
	"sync"

	"github.com/VindexAdmin/VindexChain/internal/core"
)

type Blockchain struct {
	blocks     []core.Block
	validators []core.Validator
	mu         sync.Mutex
}

// NewBlockchain crea una nueva instancia con el bloque génesis
func NewBlockchain() *Blockchain {
	genesis := core.Block{
		Index:     0,
		Timestamp: 0,
		Hash:      "genesis-hash",
	}
	return &Blockchain{
		blocks:     []core.Block{genesis},
		validators: []core.Validator{},
	}
}

// GetBlocks retorna la cadena de bloques
func (bc *Blockchain) GetBlocks() []core.Block {
	return bc.blocks
}

// GetValidators retorna los validadores registrados
func (bc *Blockchain) GetValidators() []core.Validator {
	return bc.validators
}

// AddValidator registra un nuevo validador
func (bc *Blockchain) AddValidator(v core.Validator) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.validators = append(bc.validators, v)
}

// AddBlock agrega un nuevo bloque
func (bc *Blockchain) AddBlock(txs []core.Transaction, creator string) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	prev := bc.blocks[len(bc.blocks)-1]
	newBlock := core.CreateBlock(prev, txs, creator)
	bc.blocks = append(bc.blocks, newBlock)
}

// LastBlock devuelve el último bloque
func (bc *Blockchain) LastBlock() core.Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.blocks[len(bc.blocks)-1]
}
