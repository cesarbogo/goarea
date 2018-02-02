package goarea

import (
	"math"
)

//Pi Costate de Pi
const Pi = 3.1416

//Circ aria do circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect area retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
