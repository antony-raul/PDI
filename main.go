package main

import (
	"PDI/filtros"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	arquivo, err := os.Open("./imagens/originais/original.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	imgOG, _, err := image.Decode(arquivo)
	if err != nil {
		log.Fatal(err)
	}

	imgRGBA := filtros.FiltroMediana(imgOG)

	// Criar um novo arquivo para salvar a imagem com o filtro de mediana aplicado
	novaimg, err := os.Create("imagens/resultados/resultado_filtro_mediana.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer novaimg.Close()

	// Codificar a nova imagem em JPEG
	opt := jpeg.Options{Quality: 100}
	err = jpeg.Encode(novaimg, imgRGBA, &opt)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Filtro de mediana aplicado com sucesso!")
}
