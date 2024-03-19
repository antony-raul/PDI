package utils

import (
	"image"
)

type Img struct {
	image.Image
}

// NivelDeCinza Função para obter o nível de cinza de uma cor
func (i *Img) NivelDeCinza(x, y int) int {
	r, g, b, _ := i.At(x, y).RGBA()
	// Normaliza os valores de intensidade de cor para o intervalo de 0 a 255
	r >>= 8
	g >>= 8
	b >>= 8
	// Calcula a média dos valores R, G e B
	return int((r + g + b) / 3)
}

func (i *Img) Linhas() int {
	return i.Image.Bounds().Max.X
}

func (i *Img) Colunas() int {
	return i.Image.Bounds().Max.Y
}
