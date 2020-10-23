package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SolBaa/TwitterLang/middlew"
	"github.com/SolBaa/TwitterLang/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto y pongo a escuchar mi servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
