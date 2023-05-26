package user

type Sudadera struct {
	Nombre   string
	Unidades int
	Precio   float32
}

func (S Sudadera) ImporteStock() float32 {
	Stock := float32(S.Unidades) * S.Precio
	return Stock
}

func (S Sudadera) ImporteIva(Stock float32) float32 {
	Iva := (Stock) * 0.19
	return Iva
}

func (S Sudadera) Datos() (string, string) {
	return S.Nombre, ""
}
