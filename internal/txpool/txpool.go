package txpool

import (
	"sync"

	"github.com/VindexAdmin/VindexChain/internal/core"
)

var (
	mu         sync.Mutex
	pendingTxs []core.Transaction
)

// Add agrega una transacción a la pool
func Add(tx core.Transaction) {
	mu.Lock()
	defer mu.Unlock()
	pendingTxs = append(pendingTxs, tx)
}

// GetTransactions devuelve todas las transacciones pendientes y las limpia de la pool
func GetTransactions() []core.Transaction {
	mu.Lock()
	defer mu.Unlock()
	txs := pendingTxs
	pendingTxs = nil
	return txs
}

// Clear elimina todas las transacciones pendientes (por si deseas limpiarlas manualmente)
func Clear() {
	mu.Lock()
	defer mu.Unlock()
	pendingTxs = nil
}
