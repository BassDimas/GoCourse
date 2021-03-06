package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

/*
1. Написать функцию, которая определяет, четное ли число.
2. Написать функцию, которая определяет, делится ли число без остатка на 3.
3. Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
	a. Числа Фибоначчи определяются соотношениями ​Fn=Fn-1 + Fn-2.
4. * Заполнить массив из 100 элементов различными простыми числами.
Натуральное число, которое больше единицы, называется простым, если оно делится только на себя и на единицу.
Для нахождения всех простых чисел не больше заданного числа ​n​, следуя методу Эратосфена, нужно выполнить следующие шаги:
	a. Выписать подряд все целые числа от двух до n (2, 3, 4, ..., n).
	b. Пусть переменная ​p​ изначально равна 2 — первому простому числу.
	c. Зачеркнуть в списке числа от​ 2p до​ n​, считая шагами по ​p(это будут числа, кратные​ p​: 2p, 3p, 4p, ...).
	d. Найти первое не зачеркнутое число в списке, превышающее ​p​, и присвоить значению переменной ​p​ это число.
	e. Повторять шаги ​c​ и ​d​, пока возможно.
*/

func main() {
	hwork01 := "Введите число для определения четное/нечетное "
	hwork02 := "Введите число для проверки деления без остатка на 3"
	hwork03 := "Первые 100 чисел Фибоначчи"
	var num01 int
	fmt.Println(hwork01)
	fmt.Fscan(os.Stdin, &num01)
	fmt.Println(divInt(num01))
	var num02 int
	fmt.Println(hwork02)
	fmt.Fscan(os.Stdin, &num02)
	fmt.Println(divnum(num02, 3))
	fmt.Println(hwork03)
	fibonacci(0, 1)
	natural()
}

//Функция определения четное/нечетное число
func divInt(a int) string {
	if a%2 == 0 {
		return fmt.Sprintf("Число %v четное.\n", a)
	}
	return fmt.Sprintf("Число %v нечетное.\n", a)
}

//Функция определения деления без остатка числа, на число
func divnum(a, b int) string {
	if a == 0 || b == 0 {
		return fmt.Sprintf("Ошибка числа %v, при делении на %v.\n", a, b)
	} else if a%b == 0 {
		return fmt.Sprintf("Число %v делится на %v.\n", a, b)
	}
	return fmt.Sprintf("Число %v не делится на %v.\n", a, b)
}

//Функция нахождения первых 100 чисел Фибоначчи
func fibonacci(x, y int64) {
	a := big.NewInt(x)
	b := big.NewInt(y)
	for i := 0; i < 100; i++ {
		a.Add(a, b)
		a, b = b, a
		fmt.Printf("%v\n", b)
	}
}

/*Для нахождения всех простых чисел не больше заданного числа ​n​, следуя методу Эратосфена, нужно выполнить следующие шаги:
a. Выписать подряд все целые числа от двух до n (2, 3, 4, ..., n).
b. Пусть переменная ​p​ изначально равна 2 — первому простому числу.
c. Зачеркнуть в списке числа от​ 2p до​ n​, считая шагами по ​p (это будут числа, кратные​ p​: 2p, 3p, 4p, ...).
d. Найти первое не зачеркнутое число в списке, превышающее ​p​, и присвоить значению переменной ​p​ это число.
e. Повторять шаги ​c​ и ​d​, пока возможно.
*/

//Функция создания  и наполнения массива из 100 простых чисел
func natural() {
	var arr [100]int
	for i := 0; i < len(arr); i++ {
		i1 := i + 2
		arr[i] = i1
	}
	for j := 2; j < int(math.Sqrt(100)); j++ {
		for i := 0; i < len(arr); i++ {
			if arr[i] == j*j {
				for k := i; k < len(arr); k = k + j {
					arr[k] = 0
				}
			}
			continue
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			fmt.Print(arr[i], " ")
		}
		continue
	}
}
