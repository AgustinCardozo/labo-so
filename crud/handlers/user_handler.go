package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/AgustinCardozo/labo-so/crud/models"
	"github.com/AgustinCardozo/labo-so/crud/repositories"
	"github.com/AgustinCardozo/labo-so/utils/web/server"
	"net/http"
	"strconv"
)

var userRepo repositories.IUserRepository = &repositories.UserRepository{}

func CreateHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(writer, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		var newUser models.User
		err := json.NewDecoder(request.Body).Decode(&newUser)
		if err != nil {
			http.Error(writer, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		err = userRepo.Create(newUser)
		if err != nil {
			http.Error(writer, "Error al crear el usuario", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		fmt.Fprint(writer, "Usuario creado correctamente")
	}
}

func DeleteHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := request.PathValue("id")
		id, _ := strconv.Atoi(idStr)

		err := userRepo.Delete(id)

		if err != nil {
			http.Error(writer, fmt.Sprintf("No se encontro al usuario con id %d", id), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		fmt.Fprint(writer, "Usuario borrado correctamente")
	}
}

func GetAllHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		users, err := userRepo.GetAll()

		if err != nil {
			http.Error(writer, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		server.SendJSONResponse(writer, users)
	}
}

func GetByIdHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := request.PathValue("id")
		id, _ := strconv.Atoi(idStr)

		user, err := userRepo.GetById(id)

		if err != nil && user == nil {
			http.Error(writer, fmt.Sprintf("No se encontro al usuario con id %d", id), http.StatusInternalServerError)
			return
		}

		server.SendJSONResponse(writer, user)
	}
}

func UpdateHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		idStr := request.PathValue("id")
		id, _ := strconv.Atoi(idStr)

		user, err := userRepo.GetById(id)

		if err != nil && user == nil {
			http.Error(writer, fmt.Sprintf("No se encontro al usuario con id %d", id), http.StatusInternalServerError)
			return
		}

		if request.Method != http.MethodPut {
			http.Error(writer, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		err = json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			http.Error(writer, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		err = userRepo.Update(id, *user)
		if err != nil {
			http.Error(writer, "Error al crear el usuario", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		fmt.Fprint(writer, "Usuario modificado correctamente")
	}
}
