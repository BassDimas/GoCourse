package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//	1. Изучите статью Time in Go: A primer (ссылка — в дополнительных материалах). В письменном
//	виде кратко изложите свое мнение: что из этой статьи стоило бы добавить в методичку? Если
//	считаете, что ничего, — так и напишите, приведя аргументы.
//	2. Что бы вы изменили в программе чтения из файла? Приведите исправленный вариант,
//	обоснуйте свое решение в комментарии.

//	5. ** Напишите упрощенный аналог утилиты grep .

func main() {
	//	4. * Напишите утилиту для копирования файлов, используя пакет flag .
	src := flag.String("src", "file.txt", "copy file from")
	dst := flag.String("dst", "file_copy.txt", "copy file to")
	flag.Parse()
	bsSource, err := ioutil.ReadFile(*src)
	if err != nil {
		return
	}
	file, err := os.Create(*dst)
	if err != nil {
		return
	}
	defer file.Close()
	file.Write(bsSource)
	//	3. Самостоятельно изучите пакет encoding/csv . Напишите пример, иллюстрирующий его
	//	применение.
	writeRecords := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	filecsv, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()
	w := csv.NewWriter(filecsv)
	w.WriteAll(writeRecords)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	csvFile, err := os.Open("result.csv")
	r := csv.NewReader(bufio.NewReader(csvFile))
	readRecords, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(readRecords)

}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
