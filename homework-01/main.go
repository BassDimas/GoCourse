package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	//Расчет суммы конвертации
	const current = 63.5240
	var sum float64
	var strTotal = "Сумма в $ составит: "
	var strText = "Введите сумму в рублях: "
	fmt.Println(strText)
	fmt.Fscan(os.Stdin, &sum)
	fmt.Println(strTotal, div(sum, current))

	//Даны катеты прямоугольного треугольника. Найти его площадь, периметр и гипотенузу.
	var strTriang = "Введите длины катетов прямоугольного треугольника: "
	var a, b float64
	fmt.Println(strTriang)
	fmt.Fscan(os.Stdin, &a, &b)
	fmt.Println("Гипотенуза ", triangHypo(a, b), "Площадь ", triangSq(a, b), "Периметр ", triangPerim(a, b))

	//Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.
	const yearbank = 5
	var strBankIn = "Введите сумму вклада и годовой процент: "
	var sumBank, prcBank float64
	//sumBank = 100
	//prcBank = 8
	fmt.Println(strBankIn)
	fmt.Fscan(os.Stdin, &sumBank, &prcBank)
	fmt.Println("Сумма вклада за", yearbank, "лет составит :", bankAmmount(sumBank, prcBank))

}

func mult(a float64, b float64) float64 {
	return a * b
}
func div(a float64, b float64) float64 {
	return a / b
}

//Функция расчета гипотенузы
func triangHypo(a float64, b float64) float64 {
	var c float64
	a = math.Pow(a, 2)
	b = math.Pow(b, 2)
	c = math.Sqrt(a + b)
	return c
}

//Функция расчета площади треугольника
func triangSq(a float64, b float64) float64 {
	var c float64
	c = 0.5 * a * b
	return c
}

//Функция расчета периметра треугольника
func triangPerim(a float64, b float64) float64 {
	var c float64
	c = a + b + triangHypo(a, b)
	return c
}

//Функция расчета суммы вклада, без использования циклов. Не знаю как можно без циклов использовать const yearbank = 5
func bankAmmount(a float64, b float64) float64 {
	var c float64
	b = b/100 + 1
	c = mult(a, b)
	a = c
	c = mult(c, b)
	a = a + c
	c = mult(c, b)
	a = a + c
	c = mult(c, b)
	a = a + c
	c = mult(c, b)
	a = a + c
	return a
}
