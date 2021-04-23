package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gnosotros/gored/middlew"
	"github.com/gnosotros/gored/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Menejadores seto mi puerto, el hanler
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)        // permisos con ip
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //pone el servidor a escuchar el puerto para dar respuesta al navegador

}
