package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type Block struct {
	Index        int           `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	Hash         string        `json:"hash"`
	Creator      string        `json:"creator"`
	Validator    string        `json:"validator,omitempty"`
}

type Validator struct {
	Address string  `json:"address"`
	Stake   float64 `json:"stake"`
}

// CalculateHash calcula el hash de un bloque
func (b *Block) CalculateHash() []byte {
	record := fmt.Sprintf("%d%d%s%s", b.Index, b.Timestamp, b.PrevHash, b.Creator)
	for _, tx := range b.Transactions {
		record += tx.From + tx.To + fmt.Sprintf("%f", tx.Amount)
	}
	hash := sha256.Sum256([]byte(record))
	return hash[:]
}

// CreateBlock crea un nuevo bloque
func CreateBlock(prevBlock Block, transactions []Transaction, creator string) Block {
	block := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Creator:      creator,
	}
	block.Hash = hex.EncodeToString(block.CalculateHash())
	return block
}
