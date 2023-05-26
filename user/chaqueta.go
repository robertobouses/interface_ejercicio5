package user

type Camiseta struct {
	Nombre   string
	Unidades int
	Precio   float32
	Marca    string
}

func (C Camiseta) ImporteStock() float32 {
	Stock := float32(C.Unidades) * C.Precio
	return Stock
}

func (C Camiseta) ImporteIva(Stock float32) float32 {
	Iva := (Stock) * 0.21
	return Iva
}

func (C Camiseta) Datos() (string, string) {
	return C.Nombre, C.Color
}
