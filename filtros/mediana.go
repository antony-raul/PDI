package filtros

import (
	"PDI/utils"
	"image"
	"image/color"
	"sort"
)

func FiltroMediana(imagemOG image.Image) (imgRGBA *image.RGBA) {
	var (
		img = utils.Img{
			Image: imagemOG,
		}
	)

	imgRGBA = image.NewRGBA(img.Bounds())

	// Aplicar o filtro de mediana
	var P = make([]int, 9)

	for i := 0; i < img.Linhas(); i++ {
		for j := 0; j < img.Colunas(); j++ {
			P[0] = img.NivelDeCinza(i, j)
			P[1] = img.NivelDeCinza(i-1, j)
			P[2] = img.NivelDeCinza(i+1, j)
			P[3] = img.NivelDeCinza(i, j+1)
			P[4] = img.NivelDeCinza(i, j-1)
			P[5] = img.NivelDeCinza(i-1, j+1)
			P[6] = img.NivelDeCinza(i-1, j-1)
			P[7] = img.NivelDeCinza(i+1, j+1)
			P[8] = img.NivelDeCinza(i+1, j-1)

			// Ordenar a matriz P e substituir o pixel atual pela mediana
			sort.Ints(P)
			mediana := P[4] // Índice 4 representa a mediana em um array de tamanho 9
			imgRGBA.Set(i, j, color.Gray{Y: uint8(mediana)})
		}

	}

	return
}
