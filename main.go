package main

import (
	"ms-reservas/db"
	//"ms-reservas/models"

	"ms-reservas/routes"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// database connection
	db.DBConnection()
	//db.DB.Migrator().DropTable(models.Reserva{})
	//db.DB.AutoMigrate(models.Reserva{})

	r := mux.NewRouter()

	// Index route
	r.HandleFunc("/", routes.HomeHandler)

	s := r.PathPrefix("/api").Subrouter()

	// reservas routes
	s.HandleFunc("/reservas", routes.GetReservasHandler).Methods("GET")
	s.HandleFunc("/reserva/{id}", routes.GetReservaHandler).Methods("GET")
	s.HandleFunc("/reserva", routes.CreateReservaHandler).Methods("POST")
	s.HandleFunc("/reserva/{id}", routes.DeleteReservaHandler).Methods("DELETE")
	http.ListenAndServe(":4000", r)

}
