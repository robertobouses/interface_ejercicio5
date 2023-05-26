package user

type Chaqueta struct {
	Nombre   string
	Unidades int
	Precio   float32
	Marca    string
}

func (C Chaqueta) ImporteStock() float32 {
	Stock := float32(C.Unidades) * C.Precio
	return Stock
}

func (C Chaqueta) ImporteIva(Stock float32) float32 {
	Iva := (Stock) * 0.21
	return Iva
}

func (C Chaqueta) Datos() (string, string) {
	return C.Nombre, C.Marca
}
