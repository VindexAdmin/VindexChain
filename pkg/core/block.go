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
}

// NewBlock crea un nuevo bloque
func NewBlock(transactions []Transaction, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevHash, []byte{}, 0}
	block.Hash = block.calculateHash()
	return block
}

// calcula el hash del bloque
func (b *Block) calculateHash() []byte {
	headers := fmt.Sprintf("%x%d%d", b.PrevHash, b.Timestamp, b.Nonce)
	hash := sha256.Sum256([]byte(headers))
	return hash[:]
}
