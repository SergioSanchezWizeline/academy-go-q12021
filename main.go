package main

import "fmt"

func main() {

	var yolo = "yollo"
	fmt.Println(yolo)

	switch omai := 4; {
	case omai > 3:
		fmt.Println("Omai omai")
	case omai < 3:
		fmt.Println("No Omai")
	case omai == 3:
		fmt.Println("Omai is a lie")
	}
}
