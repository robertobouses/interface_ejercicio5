package user

import "fmt"

type Pantalon struct {
	Nombre   string
	Unidades int
	Precio   float32
	Marca    string
}

func (P Pantalon) ImporteStock() int {
	Stock := P.Unidades * int(P.Precio)
	return Stock

}
func (P Pantalon) ImporteIva(float32) float32 {
	fmt.Println("Los pantalones viene de importación sin IVA")
	return 0
}

func (P Pantalon) Datos() string {
	return P.Nombre
}
