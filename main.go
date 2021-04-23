package main

import (
	"log"

	"github.com/gnosotros/gored/bd"
	"github.com/gnosotros/gored/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexxion a BD")
		return
	}
	handlers.Manejadores()
}
