#SALUDOS EN GO

Este paquete proporciona una forma de slaudos simple en Go

##Instalación
Ejecuta el siguiente comando para instalar el paquete:

go get -u https://github.com/ChrisInz/greetings

Aqui un ejemplo:

package main

import (
	"fmt"
	"log"

	"github.com/ChrisInz/greeting"
)

func main() {
	log.SetPrefix("Grettings: ")
	log.SetFlags(0) //define que se quiten o agregue el timestamp a los registros del log

	names := []string{"Chris", "Yokasta", "Luz", "Yuliana", "Alejandro"}
	messages, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	message, err := greeting.Hello("Chris")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println(messages)
}