package main

import (
	"math"
)

//Inicializando com letra maiúscula é PÚBLICO
// visível dentro e fora do pacote

//Iniciando com letra minúscula ou _ é PRIVADO
//visível apenas dentro do pacote

//Não existe como tornar a função privada dentro de um arquivo.
//Só é possível isolar uma função dentro do pacote.

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsavel por calcular a distância
// linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
