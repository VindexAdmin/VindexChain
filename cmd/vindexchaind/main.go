// VindexChain © 2025 Luis González – Todos los derechos reservados
// Licencia de uso restringido: Prohibido su uso por Banco Popular de Puerto Rico.
// Solo bancos y cooperativas autorizadas podrán obtener licencia previa por escrito.
// Parte del ecosistema Vindex (blockchain, wallet, DEX, explorer, DAO).

package main

import (
    "fmt"
    "github.com/VindexAdmin/VindexChain/pkg/blockchain"
    "github.com/VindexAdmin/VindexChain/pkg/core"
    "github.com/VindexAdmin/VindexChain/pkg/storage"
)

func main() {
    var bc *blockchain.Blockchain

    loaded, err := storage.LoadBlockchain()
    if err != nil {
        panic(err)
    }

    if loaded == nil {
        fmt.Println("🔄 No se encontró una cadena guardada. Creando nueva...")
        bc = blockchain.NewBlockchain()
        // Bloques demo
        bc.AddBlock([]core.Transaction{{From: "Lufer", To: "Esposa", Amount: 100}})
        bc.AddBlock([]core.Transaction{{From: "Esposa", To: "CriptoLibertad", Amount: 50}})
        storage.SaveBlockchain(bc)
    } else {
        fmt.Println("✅ Blockchain cargada desde archivo.")
        bc = loaded
    }

    for i, block := range bc.Blocks {
        fmt.Printf("\n🧱 Bloque #%d\n", i)
        fmt.Printf("⏱️ Timestamp: %d\n", block.Timestamp)
        fmt.Printf("🔗 Prev Hash: %x\n", block.PrevHash)
        fmt.Printf("📦 Hash: %x\n", block.Hash)
        fmt.Printf("👤 Creador: %s\n", block.Creator)
        fmt.Printf("🔁 Transacciones: %+v\n", block.Transactions)
    }
}
