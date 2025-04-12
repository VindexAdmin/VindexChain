// VindexChain © 2025 Luis González – Todos los derechos reservados
// Licencia de uso restringido: Prohibido su uso por Banco Popular de Puerto Rico.
// Solo bancos y cooperativas autorizadas podrán obtener licencia previa por escrito.
// Parte del ecosistema Vindex (blockchain, wallet, DEX, explorer, DAO).

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/VindexAdmin/VindexChain/internal/blockchain"
	"github.com/VindexAdmin/VindexChain/internal/core"
)

var (
	txQueue []core.Transaction
	mutex   sync.Mutex
)

// StartServer lanza el servidor HTTP
func StartServer(bc *blockchain.Blockchain) {
	http.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(bc.GetBlocks())
	})

	http.HandleFunc("/validators", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(bc.GetValidators())
	})

	http.HandleFunc("/stake", func(w http.ResponseWriter, r *http.Request) {
		var v core.Validator
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			http.Error(w, "Validador inválido", http.StatusBadRequest)
			return
		}
		bc.AddValidator(v)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "Validador registrado exitosamente",
			"address": v.Address,
			"stake":   fmt.Sprintf("%.2f", v.Stake),
		})
	})

	http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		var tx core.Transaction
		if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
			http.Error(w, "Transacción inválida", http.StatusBadRequest)
			return
		}

		mutex.Lock()
		txQueue = append(txQueue, tx)
		mutex.Unlock()

		json.NewEncoder(w).Encode(map[string]string{
			"status": "Transacción recibida y será procesada en el próximo bloque",
		})
	})

	http.HandleFunc("/balance/", func(w http.ResponseWriter, r *http.Request) {
		address := strings.TrimPrefix(r.URL.Path, "/balance/")
		balance := 0.0
		for _, block := range bc.GetBlocks() {
			for _, tx := range block.Transactions {
				if tx.From == address {
					balance -= tx.Amount
				}
				if tx.To == address {
					balance += tx.Amount
				}
			}
		}
		json.NewEncoder(w).Encode(map[string]float64{
			"balance": balance,
		})
	})

	fmt.Println("🌐 Servidor API corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
