package main

import (
	"fmt"

	"github.com/EmersonV00/godesde0/variables"
)

func main() {
	Estado, Texto := variables.ConviertoaTexto(1588)
	fmt.Println(Estado)
	fmt.Println(Texto)
}
