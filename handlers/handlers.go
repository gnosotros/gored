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
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8180"
	}
	handler := cors.AllowAll().Handler(router)        // permisos con ip
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //pone el servidor a escuchar el puerto para dar respuesta al navegador

}
