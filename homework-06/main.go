package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

//	4. * Напишите функцию для вычисления корней квадратного уравнения (алгоритм можно найти в
//	Википедии ) и тесты к ней.

func main() {
	//	2. Дополните пример из раздела «Пакет img» изображением горизонтальных и вертикальных
	//	линий. Воспользуйтесь статьей « Работа с изображениями ».
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{255, 0, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			for y := 10; y < 500; y = y + 10 {
				rectImg.Set(x, y, red)
			}
		}

		for y := 0; y < 500; y++ {
			for x := 10; x < 500; x = x + 10 {
				rectImg.Set(x, y, red)
			}
		}
	}

	file, err := os.Create("static/rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":80", nil)
}

//	3. Дополните функцию hello() http-сервера так, чтобы принять и отобразить на странице один
//	GET-параметр, например name . При этом в браузере запрос будет выглядеть так:
//	http://localhost/hello&name=World . Значение этого параметра надо получить в функции и
//	отобразить при выводе html-кода.
//	В таком формате (http://localhost/hello&name=World) не работает,
//	В таком http://localhost/hello?name=World рпботает
func hello(res http.ResponseWriter, req *http.Request) {
	getparam := req.URL.Query()
	getValue := getparam.Get("name")
	res.Header().Set("Content-Type", "text/html")
	getparamvalue := ""
	if getValue == "World" {
		getparamvalue = getValue
	}
	htmlstring := "<!DOCTYPE html> <html> <head> <title>Hello World!</title> </head> <body> Hello World! <p>" + getparamvalue + "</p> <img src = " + "rectangle.png" + " /> </body> </html>"
	io.WriteString(res, htmlstring)
}
