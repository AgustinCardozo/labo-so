package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func InitServer(puerto int) error {
	// Declara el puerto
	port := ":" + strconv.Itoa(puerto)

	// Escucha y sirve con la información de config.json
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error al escuchar en el puerto " + port)
		fmt.Println(err)
	}
	return err
}
