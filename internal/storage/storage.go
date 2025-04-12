package storage

import (
	"encoding/json"
	"os"

	"github.com/VindexAdmin/VindexChain/internal/core"
)

// SaveValidators guarda la lista de validadores en un archivo JSON
func SaveValidators(validators []core.Validator) error {
	file, err := os.Create("validators.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(validators)
}
