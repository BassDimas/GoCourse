package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

//	5. * Напишите программу, генерирующую png-файл с рисунком шахматной доски.
func main() {
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	boardsize := 800
	blockcount := 8
	blocksize := boardsize / blockcount
	fieldImg := image.NewRGBA(image.Rect(0, 0, boardsize, boardsize))
	draw.Draw(fieldImg, fieldImg.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	//rectImg := image.Rect(0, 0, 100, 100)

	for x := blocksize; x < boardsize; x = x + blocksize {

		for y := 0; y < boardsize; y = y + blocksize {

			draw.Draw(fieldImg, image.Rect(x, y, x+blocksize, y+blocksize), &image.Uniform{black}, image.ZP, draw.Src)
			y += blocksize
		}
		x += blocksize
	}
	for x := 0; x < boardsize; x = x + blocksize {

		for y := blocksize; y < boardsize; y = y + blocksize {

			draw.Draw(fieldImg, image.Rect(x, y, x+blocksize, y+blocksize), &image.Uniform{black}, image.ZP, draw.Src)
			y += blocksize
		}
		x += blocksize
	}

	file, err := os.Create("chessboard.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, fieldImg)
}
