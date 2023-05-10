package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	/*
		mensajes.Hola()
		mensajes.ImpriimirMensaje()

		cua := figuras.Cuadrado{Lado: 10}
		figuras.Medidas(&cua)
	*/

	p1 := models.Persona{}
	p1.Constructor("Reales", 32)

	fmt.Println(p1)

}
