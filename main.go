package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*Estado, Texto := variables.ConviertoaTexto(1588)
	fmt.Println(Estado)
	fmt.Println(Texto) */
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
