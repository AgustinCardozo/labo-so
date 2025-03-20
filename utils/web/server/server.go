package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func InitServer(puerto int) error {
	port := ":" + strconv.Itoa(puerto)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error al escuchar en el puerto " + port)
		fmt.Println(err)
	}
	return err
}

func SendJSONResponse(writer http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, "Error al convertir datos a JSON", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
