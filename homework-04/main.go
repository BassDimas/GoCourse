package main

import (
	"GoCourse/calculator"
	"fmt"
	"sort"
)

//	1. Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.
//	2. Создать тип, описывающий контакт в телефонной книге.
//		Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}.
//	3. Дописать функцию, которая будет выводить справку к калькулятору. Она должна вызываться
//		при вводе слова help вместо выражения.
//	4. * Написать функцию, которая будет получать позицию коня на шахматной доске, а возвращать
//		массив из точек, в которые конь сможет сделать ход.
//		a. Точку следует обозначить как структуру, содержащую x и y типа int
//		b. Получение значений и их запись в точку должны происходить только с помощью
//		отдельных методов. В них надо проводить проверку на то, что такая точка может
//		существовать на шахматной доске.

type Figure interface {
	Position() (int, int)
	AvailablePoints() []Point
}

type Point struct {
	x, y int
}

type Knight struct {
	x, y int
}

var (
	knightMoves = []Point{
		Point{
			y: 1,
			x: -2,
		},
		Point{
			y: 2,
			x: -1,
		},
		Point{
			y: 2,
			x: 1,
		},
		Point{
			y: 1,
			x: 2,
		},
		Point{
			y: -1,
			x: 2,
		},
		Point{
			y: -2,
			x: 1,
		},
		Point{
			y: -2,
			x: -1,
		},
		Point{
			y: -1,
			x: -2,
		},
	}
)

func (k *Knight) AvailablePoints() []Point {
	//panic("not implemented")
	returnResult := make([]Point, 0, 2)
	for _, v := range knightMoves {
		if canMove(k, v.x, v.y) {
			k := Point{(k.x + v.x), (k.y + v.y)}
			returnResult = append(returnResult, k)

		}
	}
	return returnResult
}

func canMove(p *Knight, whereX, whereY int) bool {
	destX := p.x + whereX
	destY := p.y + whereY
	return (destX <= 8 && destX >= 1) && (destY >= 1 && destY <= 8)
}

func (k *Knight) Position() (int, int) {
	return k.x, k.y
}

func NewKnight(x, y int) Figure {
	return &Knight{x, y}
}

//	2. Создать тип, описывающий контакт в телефонной книге.
type PhoneBook struct {
	name   string
	number string
}

//		Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}.
type ByName []PhoneBook

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}
		if input == "exit" {
			break
		}
		//	3. Дописать функцию, которая будет выводить справку к калькулятору. Она должна вызываться
		//		при вводе слова help вместо выражения.
		if input == "help" {
			help1 := "Помощь при работе с калькулятором:"
			help2 := "Введите число, далее оператор действия которое нужно произвести с числами(+,-,*,/), далее следующее число."
			help3 := "Операторы можно комбинировать и использовать скобки. Для завершения работы введите exit"
			fmt.Printf("%v \n %v \n %v \n", help1, help2, help3)
			continue
		}
		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
	//Создать псевдоним типа телефонной
	//книги (массив контактов) и реализовать для него интерфейс Sort{}.
	people := []PhoneBook{
		{"Alice", "89994569988"},
		{"Mike", "88889456543"},
		{"Bob", "84988765432"},
		{"John", "87778756712"},
	}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

	var inputx int
	var inputy int

	fmt.Print("Введите позицию коня на доске, в формате x y, где x и y число от 1 до 8: ")
	if _, err := fmt.Scanln(&inputx, &inputy); err != nil {
		fmt.Println(err)
	}
	if input == "exit" {

	}
	var p Figure = NewKnight(inputx, inputy)
	fmt.Println(p.AvailablePoints())
}

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].name < a[j].name
}
