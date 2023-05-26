package main

import (
	"fmt"

	"github.com/robertobouses/interface_ejercicio5/user"
)

type Inventariar interface {
	ImporteStock() float32
	ImporteIva(Stock float32) float32
	Datos() (string, string)
}

func ImprimirDatos(I Inventariar) {
	fmt.Println(I.Datos())

}

func main() {
	prendas := []Inventariar{
		user.Camiseta{"camiseta", 29, 1.3, "rojo"},
		user.Camiseta{"camiseta", 5, 1.6, "lila"},
		user.Camiseta{"camiseta", 29, 2.1, "verde"},
		user.Chaqueta{"chaqueta", 39, 6.8, "desconocida"},
		user.Pantalon{"pantalón", 3, 4.5, "goldjeans"},
		user.Sudadera{"sudadera", 30, 23.40},
	}

	for _, prenda := range prendas {
		stock := prenda.ImporteStock()
		fmt.Println("dato del Stock:")
		fmt.Println(stock)

		iva := prenda.ImporteIva(stock)
		fmt.Println("dato del Iva")
		fmt.Println(iva)

		fmt.Println(prenda)

		fmt.Println(prenda.Datos())
		fmt.Println("------------------------")
	}

}
