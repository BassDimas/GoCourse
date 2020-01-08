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
type Figure struct {
	PositionX string
	PositionY int
}

type Point struct {
	x int
	y int
}

type Move struct {
	X int
	Y int
}

var (
	knightMoves = []Move{
		Move{
			Y: 1,
			X: -2,
		},
		Move{
			Y: 2,
			X: -1,
		},
		Move{
			Y: 2,
			X: 1,
		},
		Move{
			Y: 1,
			X: 2,
		},
		Move{
			Y: -1,
			X: 2,
		},
		Move{
			Y: -2,
			X: 1,
		},
		Move{
			Y: -2,
			X: -1,
		},
		Move{
			Y: -1,
			X: -2,
		},
	}
)

//	2. Создать тип, описывающий контакт в телефонной книге.
type PhoneBook struct {
	name   string
	number string
}

func (f Figure) Method() {
	fmt.Println(f.PositionX, f.PositionY)
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

	var inputx string
	var inputy int

	fmt.Print("Введите позицию коня на доске, в формате x y, где x символ от a до h, а y число от 1 до 8: ")
	if _, err := fmt.Scanln(&inputx, &inputy); err != nil {
		fmt.Println(err)
	}
	if input == "exit" {

	}
	p := Figure{inputx, inputy}
	p.SetPoint()
	fmt.Println(findMove(p.SetPoint()))
}

func (p Figure) SetPoint() Point {
	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	a := 0
	for i, r := range arr {
		if r == p.PositionX {
			a = i + 1
		}
	}
	return Point{
		x: a,
		y: p.PositionY,
	}
}

func findMove(p Point) []Figure {
	returnResult := make([]Figure, 0, 2)
	for _, v := range knightMoves {
		if canMove(p, v.X, v.Y) {
			p := Point{(p.x + v.X), (p.y + v.Y)}
			returnResult = append(returnResult, p.GetPoint())
		}
	}
	return returnResult
}

func (p Point) GetPoint() Figure {
	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	a := ""
	for i, r := range arr {
		if i+1 == p.x {
			a = r
		}
	}
	return Figure{
		PositionX: a,
		PositionY: p.y,
	}

}

func canMove(p Point, whereX, whereY int) bool {
	destX := p.x + whereX
	destY := p.y + whereY
	return (destX <= 8 && destX >= 1) && (destY >= 1 && destY <= 8)
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
