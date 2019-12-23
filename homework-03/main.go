package main

import (
	"GoCourse/queue"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//	1. Описать несколько структур — любой легковой автомобиль и грузовик. Структуры должны
//	содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты
//	ли окна, насколько заполнен объем багажника.

type Auto struct {
	Model          string
	Year           int
	Trunk          float64
	VehicleIsStart bool
	WindowIsOpen   bool
	TrunkFilling   float64
	DriverInfo     Driver
}

type Driver struct {
	FirstName     string
	LastName      string
	DriverLicense string
}

type Message struct {
	Name string
	Body string
	Time int64
}

//	2. Инициализировать несколько экземпляров структур.

var CargoCar = Auto{
	Model:          "VOLVO",
	Year:           2018,
	Trunk:          2000,
	VehicleIsStart: true,
	WindowIsOpen:   false,
	TrunkFilling:   100,
	DriverInfo: Driver{
		FirstName:     "Иван",
		LastName:      "Иванов",
		DriverLicense: "D",
	},
}

var PassengerCar = Auto{
	Model:          "Lexus",
	Year:           2010,
	Trunk:          100,
	VehicleIsStart: false,
	WindowIsOpen:   true,
	TrunkFilling:   0,
	DriverInfo: Driver{
		FirstName:     "Петр",
		LastName:      "Петров",
		DriverLicense: "B",
	},
}

func main() {

	//	Применить к структурам различные действия. Вывести значения свойств экземпляров в консоль.

	if CargoCar.VehicleIsStart == true || PassengerCar.WindowIsOpen == true {
		fmt.Println(CargoCar.Model, "- Двигатель запущен", "Водитель - ", CargoCar.DriverInfo)
		fmt.Println(PassengerCar.Model, "- Окна открыты", "Водитель - ", PassengerCar.DriverInfo)
	}

	//	3. * Реализовать очередь. Это структура данных, работающая по принципу FIFO (First Input —
	//		first output, или «первым зашел — первым вышел»).

	queue.Push("Строка №1")
	queue.Push("Строка №2")
	queue.Push("Строка №3")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	queue.Push("Строка №4")
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	//	4. * Внести в телефонный справочник дополнительные данные. Реализовать сохранение
	//		json-файла на диске с помощью пакета ioutil при изменении данных.

	addressBook := make(map[string][]int)
	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)
	addressBookSize := len(addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}

	}

	addressBook["Alice"] = []int{89998765432}
	if len(addressBook) > addressBookSize {
		file, _ := json.MarshalIndent(addressBook, "", " ")
		err := ioutil.WriteFile("test.json", file, 0644)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
