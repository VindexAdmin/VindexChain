package main

import (
	"fmt"
	"github.com/VindexAdmin/VindexChain/internal/api"
	"github.com/VindexAdmin/VindexChain/internal/blockchain"
	"github.com/VindexAdmin/VindexChain/internal/validator"
)

func main() {
	// Crear instancia del blockchain
	bc := blockchain.NewBlockchain()

	// Dirección del validador (puede venir de configuración en el futuro)
	validatorAddress := "lufer-validator"

	// Iniciar el validador en una goroutine
	go validator.Run(bc, validatorAddress)

	// Iniciar servidor API
	api.StartServer(bc)

	fmt.Println("✅ VindexChain iniciado con éxito.")
}
