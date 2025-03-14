package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func DoRequest(port int, ip string, metodo string, query string, bodies ...[]byte) (*http.Response, error) {
	// Se declara un nuevo cliente
	cliente := &http.Client{}

	// Se declara la url a utilizar (depende de una ip y un puerto).
	url := fmt.Sprintf("http://%s:%d/%s", ip, port, query)

	body := ifBody(bodies...)

	// Se crea una request donde se "efectúa" el metodo (PUT / DELETE / GET / POST) hacia url, enviando el Body si lo hay
	req, err := http.NewRequest(metodo, url, body)

	// Error Handler de la construcción de la request
	if err != nil {
		fmt.Printf("error creando request a ip: %s puerto: %d\n", ip, port)
		return nil, err
	}

	// Se establecen los headers
	req.Header.Set("Content-Type", "application/json")

	// Se envía el request al servidor
	respuesta, err := cliente.Do(req)

	// Error handler de la request
	if err != nil {
		fmt.Printf("error enviando request a ip: %s puerto: %d\n", ip, port)
		return nil, err
	}

	// Verificar el código de estado de la respuesta del servidor a nuestra request (de no ser OK)
	if respuesta.StatusCode != http.StatusOK {
		fmt.Printf("Status Error: %d\n", respuesta.StatusCode)
		return nil, err
	}
	return respuesta, err
}

func ifBody(bodies ...[]byte) io.Reader {
	if len(bodies) == 0 {
		return nil
	}
	return bytes.NewBuffer(bodies[0])
}
