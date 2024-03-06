package holamundospoliglotas

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
)

// saludo en inglés
func EnIngles() {
	fmt.Print("Te saludo en Inglés: ")
	fmt.Println("Hello World")
}

// saludo en español
func EnEspañol() {
	fmt.Print("Te saludo en Español: ")
	fmt.Println("Hola Mundo")
}

// saludo en alemán
func EnAleman() {
	fmt.Print("Te saludo en Alemán: ")
	fmt.Println("Hallo Welt")
}

// saludo en cósigo morse
func EnMorse() {
	fmt.Print("Te saludo en Morse: ")

	h := morse.NewHacker()
	morseCode, _ := h.Encode(strings.NewReader("Hola Mundo"))
	// Morse Code is: -.-. --- ...- . .-. - / - .... .. ... / - --- / -- --- .-. ... .
	fmt.Printf("%s\n", string(morseCode))
}
