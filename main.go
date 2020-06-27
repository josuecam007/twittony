package main

import (
	"log"

	"github.com/josuecam007/twittony/bd"
	"github.com/josuecam007/twittony/handlers"
)

func main() {
	if bd.ChequeoConexion() == false {
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
