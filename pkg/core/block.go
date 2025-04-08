// VindexChain © 2025 Luis González – Todos los derechos reservados
// Licencia de uso restringido: Prohibido su uso por Banco Popular de Puerto Rico.
// Solo bancos y cooperativas autorizadas podrán obtener licencia previa por escrito.
// Parte del ecosistema Vindex (blockchain, wallet, DEX, explorer, DAO).

package core

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block representa un bloque individual en la blockchain
type Block struct {
	Timestamp    int64
	Transactions []Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int
	Creator      string // 👤 Nuevo: validador que creó el bloque
}

// NewBlock crea un nuevo bloque
func NewBlock(transactions []Transaction, prevHash []byte, creator string) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Creator:      creator,
	}
	block.Hash = block.calculateHash()
	return block
}

// calcula el hash del bloque
func (b *Block) calculateHash() []byte {
	headers := fmt.Sprintf("%x%d%d%s", b.PrevHash, b.Timestamp, b.Nonce, b.Creator)
	hash := sha256.Sum256([]byte(headers))
	return hash[:]
}
