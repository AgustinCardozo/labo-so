package main

import (
	"fmt"
	"github.com/AgustinCardozo/labo-so/memoria/globals"
	handlersMemoria "github.com/AgustinCardozo/labo-so/memoria/handlers"
	"github.com/AgustinCardozo/labo-so/utils/config"
	"github.com/AgustinCardozo/labo-so/utils/list"
	"github.com/AgustinCardozo/labo-so/utils/web/server"
	"github.com/AgustinCardozo/labo-so/utils/web/server/handlers"
	"net/http"
)

const (
	ConfigPath = "./configs/memoria.json"
)

func main() {
	var numberList list.ArrayList[int]

	numberList.Add(10)
	numberList.Add(20)
	numberList.Add(30)

	fmt.Println(numberList.Size())

	config.InitConfig(ConfigPath, &globals.ConfigMemoria)
	fmt.Printf("Port: %d", globals.ConfigMemoria.Port)

	http.HandleFunc("/", handlers.HandshakeHandler("Memoria en funcionamiento 🚀"))
	http.HandleFunc("/memoria", handlersMemoria.HelloWorldHandler)

	err := server.InitServer(globals.ConfigMemoria.Port)
	if err != nil {
		fmt.Errorf("error initializing server: %v", err)
		panic(err)
	}
}
