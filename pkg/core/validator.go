 // VindexChain © 2025 Luis González – Todos los derechos reservados
// Licencia de uso restringido: Prohibido su uso por Banco Popular de Puerto Rico.
// Solo bancos y cooperativas autorizadas podrán obtener licencia previa por escrito.
// Parte del ecosistema Vindex (blockchain, wallet, DEX, explorer, DAO).

package core

import (
    "math/rand"
    "time"
)

// Validator representa un nodo con tokens en stake
type Validator struct {
    Address string  // Dirección del validador
    Stake   float64 // Tokens apostados (stake)
}

// SelectValidator elige un validador según su stake (probabilidad proporcional)
func SelectValidator(validators []Validator) Validator {
    rand.Seed(time.Now().UnixNano())
    totalStake := 0.0
    for _, v := range validators {
        totalStake += v.Stake
    }

    pick := rand.Float64() * totalStake
    current := 0.0

    for _, v := range validators {
        current += v.Stake
        if current >= pick {
            return v
        }
    }

    // Fallback en caso extremo
    return validators[len(validators)-1]
}
