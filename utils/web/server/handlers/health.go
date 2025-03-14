package handlers

import (
	"encoding/json"
	"net/http"
)

func HandshakeHandler(message string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		response, err := json.Marshal(message)
		if err != nil {
			http.Error(writer, "Error al codificar los datos como JSON", http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}
