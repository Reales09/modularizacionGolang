package figuras

type Circulo struct {
	Radio float64
}

func (c Circulo) area() float64 {
	return c.Radio * c.Radio * 3.14
}

func (c Circulo) perimetro() float64 {
	return c.Radio * 2 * 3.14
}
