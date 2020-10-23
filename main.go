package main

import (
	"log"

	"github.com/SolBaa/TwitterLang/bd"
	"github.com/SolBaa/TwitterLang/handlers"
)

func main() {
	if bd.CheckConection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()

}
