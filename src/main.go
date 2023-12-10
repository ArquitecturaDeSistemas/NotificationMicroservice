package main

import (
	"fmt"
	"notificationmicroservice/basedatos"
)

func main() {
	fmt.Println("hola")
	basedatos.Connect()
	basedatos.IsTable()
	basedatos.Close()
}
