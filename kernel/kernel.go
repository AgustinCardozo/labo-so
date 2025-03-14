package main

import (
	"fmt"
	"github.com/AgustinCardozo/labo-so/kernel/globals"
	kernelHandlers "github.com/AgustinCardozo/labo-so/kernel/handlers"
	"github.com/AgustinCardozo/labo-so/utils/config"
	"github.com/AgustinCardozo/labo-so/utils/web/server"
	"github.com/AgustinCardozo/labo-so/utils/web/server/handlers"

	"net/http"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(quote.Hello())

	config.InitConfig("./configs/kernel.json", &globals.ConfigKernel)

	http.HandleFunc("/", handlers.HandshakeHandler("Memoria en funcionamiento 🚀"))

	kernelHandlers.SendMessageToMemoria(globals.ConfigKernel.Port_memory, globals.ConfigKernel.Ip_memory, "Test")

	err := server.InitServer(globals.ConfigKernel.Port)
	if err != nil {
		fmt.Errorf("error initializing server: %v", err)
		panic(err)
	}
}
