package main

import (
	"fmt"
	"os"
)

func main() {
	//Расчет суммы конвертации
	const current = 63.5240
	var sum float64
	var strTotal = "Сумма в рублях составит: "
	var strText = "Введите сумму: "
	fmt.Println(strText)
	fmt.Fscan(os.Stdin, &sum)
	fmt.Println(strTotal)
	fmt.Println(myCurrency(sum, current))
}

/*
	//Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу.
	var strTriang = "Введите длины катетов прямоугольного треугольника: "
	var a, b float64
	a = 10
	b = 12.5
	fmt.Println(strTriang)
	fmt.Fscan(os.Stdin, &a, &b)
	fmt.Println("Гипотенуза ", triangHypo(a, b), "Площадь ", triangSq(a, b), "Периметр ", triangPerim(a, b))

	//Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.
	const yearbank = 5
	var strBank = "Введите сумму вклада и годовой процент: "
	var sumBank, prcBank float64
	sumBank = 100
	prcBank = 8
	fmt.Println(strBank)
	fmt.Fscan(os.Stdin, &sumBank, &prcBank)
	fmt.Println(bankAmmount(sumBank, prcBank))
}
*/
func myCurrency(a, b float64) float64 {
	return a / b

}

/*
func triangHypo(a float64, b float64) float64 {
	var c float64
	a = math.Pow(a, 2)
	b = math.Pow(b, 2)
	c = math.Sqrt(a + b)
	return c
}
/*
func triangSq(a float64, b float64) float64 {
	var c float64
	c = 0.5 * a * b
	return c
}

func triangPerim(a float64, b float64) float64 {
	var c float64
	c = a + b + triangHypo(a, b)
	return c
}

func bankAmmount(a float64, b float64) float64 {
	var c, d float64
	b = b/100 + 1
	c = mult(a, b)
	d = c
	c = mult(c, b)
	d = d + c
	c = mult(c, b)
	d = d + c
	c = mult(c, b)
	d = d + c
	c = mult(c, b)
	d = d + c
	return d
}
*/
