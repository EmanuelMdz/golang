package main

import (
	"fmt"

	"github.com/EmanuelMdz/golang/variables"
)

func main() {
    estado, texto := variables.ConviertoATexto(1500)
    fmt.Println(estado)
    fmt.Println(texto)
}