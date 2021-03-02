package app

import (
	"banking-auth/dto"
	"banking-auth/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AutHandler struct {
	service service.AuthService
}

// NotImplementedHandler
func (h AuthHandler) NotImplementedHanler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handler no implementado")
}

// Login Handlder
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Error al decodificar la peticion login | " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// pide accesso
		token, err := h.service.Login(loginRequest)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, err.Error())
		} else {
			// Success
			fmt.Fprint(w, *token)
		}
	}
}
