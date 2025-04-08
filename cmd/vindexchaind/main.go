// VindexChain © 2025 Luis González – Todos los derechos reservados
// Licencia de uso restringido: Prohibido su uso por Banco Popular de Puerto Rico.
// Solo bancos y cooperativas autorizadas podrán obtener licencia previa por escrito.
// Parte del ecosistema Vindex (blockchain, wallet, DEX, explorer, DAO).

package main

import (
	"fmt"

	"github.com/VindexAdmin/VindexChain/pkg/blockchain"
	"github.com/VindexAdmin/VindexChain/pkg/core"
)

func main() {
	bc := blockchain.NewBlockchain()

	// Agregamos algunos bloques de prueba
	tx1 := core.Transaction{From: "Lufer", To: "Esposa", Amount: 100}
	bc.AddBlock([]core.Transaction{tx1})

	tx2 := core.Transaction{From: "Esposa", To: "CriptoLibertad", Amount: 50}
	bc.AddBlock([]core.Transaction{tx2})

	for i, block := range bc.Blocks {
		fmt.Printf("\n🧱 Bloque #%d\n", i)
		fmt.Printf("⏱️ Timestamp: %d\n", block.Timestamp)
		fmt.Printf("🔗 Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("📦 Hash: %x\n", block.Hash)
		fmt.Printf("🔁 Transacciones: %+v\n", block.Transactions)
	}
}
