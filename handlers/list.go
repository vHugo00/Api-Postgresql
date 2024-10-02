package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vHugo00/Api-Postgresql/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Println("Erro ao buscar todos os registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
