package handlers

import (
	"fmt"
	"github.com/AgustinCardozo/labo-so/utils/web/client"
	"io"
	"log/slog"
)

func SendMessageToMemoria(port int, ip string, message string) {
	query := fmt.Sprintf("memoria?name=%s", message)
	response, err := client.DoRequest(port, ip, "GET", query, nil)

	if err != nil {
		slog.Error(fmt.Sprintf("Ocurrió un error: %v", err))
		return
	}

	responseBody, _ := io.ReadAll(response.Body)
	fmt.Printf("Mensaje recibido de memoria: %s", string(responseBody))
}
