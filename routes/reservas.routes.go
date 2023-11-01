package routes

import (
	"encoding/json"
	"net/http"

	"ms-reservas/db"
	"ms-reservas/models"

	"github.com/gorilla/mux"
)

func GetReservasHandler(w http.ResponseWriter, r *http.Request) {
	var reservas []models.Reserva
	db.DB.Find(&reservas)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&reservas)
}

func GetReservaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reserva models.Reserva
	db.DB.First(&reserva, params["id"])

	if reserva.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("reserva not found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&reserva)
}

func CreateReservaHandler(w http.ResponseWriter, r *http.Request) {
	var reservas models.Reserva
	json.NewDecoder(r.Body).Decode(&reservas)
	createdReserva := db.DB.Create(&reservas)
	err := createdReserva.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&reservas)
}

func DeleteReservaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reserva models.Reserva
	db.DB.First(&reserva, params["Id"])

	if reserva.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("reserva not found"))
		return
	}

	db.DB.Unscoped().Delete(&reserva)
	w.WriteHeader(http.StatusOK)
}
