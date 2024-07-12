package main

import (
	"arch/internal/contracts/open_api"
	v1 "arch/internal/contracts/open_api/v1"
	"arch/internal/services"
	"arch/internal/stores"
	"log"
	"os"
	"os/signal"
)

// we can inject config everywhere we want
// we can inject stores and services only where we want
func main() {
	//config getting here

	// getting dbConn
	// db := sql.Open(...)
	// creating logger
	logger := log.New(log.Writer(), log.Prefix(), log.Flags())
	storeSomeObject := stores.NewSomeObjects(db)

	serviceSomeObject := services.NewSomeDomainObjects(storeSomeObject)

	contract := v1.New(logger, serviceSomeObject)

	app := open_api.New(logger, nil, contract)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	app.Run(sigChan)
}
