package main

import (
	"paquetes/figuras"
	"paquetes/mensajes"
)

func main() {
	mensajes.Hola()
	mensajes.ImpriimirMensaje()

	cua := figuras.Cuadrado{Lado: 10}
	figuras.Medidas(&cua)
}
