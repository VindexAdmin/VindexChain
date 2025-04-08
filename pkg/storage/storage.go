// VindexChain © 2025 Luis González – Todos los derechos reservados
// Licencia de uso restringido: Prohibido su uso por Banco Popular de Puerto Rico.
// Solo bancos y cooperativas autorizadas podrán obtener licencia previa por escrito.
// Parte del ecosistema Vindex (blockchain, wallet, DEX, explorer, DAO).

package storage

import (
    "encoding/json"
    "io/ioutil"
    "os"

    "github.com/VindexAdmin/VindexChain/pkg/blockchain"
)

// Ruta del archivo donde se guarda la cadena
const BlockchainFile = "chain.json"

// Guarda la blockchain en un archivo JSON
func SaveBlockchain(bc *blockchain.Blockchain) error {
    data, err := json.MarshalIndent(bc, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(BlockchainFile, data, 0644)
}

// Carga la blockchain desde un archivo JSON (si existe)
func LoadBlockchain() (*blockchain.Blockchain, error) {
    if _, err := os.Stat(BlockchainFile); os.IsNotExist(err) {
        return nil, nil // No existe, devolver nil
    }

    data, err := ioutil.ReadFile(BlockchainFile)
    if err != nil {
        return nil, err
    }

    var bc blockchain.Blockchain
    err = json.Unmarshal(data, &bc)
    if err != nil {
        return nil, err
    }

    return &bc, nil
}
