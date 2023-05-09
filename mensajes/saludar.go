package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola desde el paquete mensajes")
}

const mensaje = "Hola desde mi constante"

func ImpriimirMensaje() {
	fmt.Println(mensaje)
	funcionPrivada()
}

func funcionPrivada() {
	fmt.Println("Hola desde mi funcion privada")
}
