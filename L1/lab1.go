package main

import (
	"fmt"
	"time"
)

// task1 выводит текущее вреия и дату.
func task1() string {
	c_time := time.Now().Format("01.01.1000 10:10:10")
	fmt.Println("Задача 1")
	fmt.Printf("Текущая дата и время: %s\n", c_time)

	return c_time
}

// task2 создаёт переменные типов
// int, float64, boolean и string и
// выводит их.
func task2() {
	var v_i int = 10
	var v_f64 float64 = 15.5
	var v_b bool = true
	var v_s string = "Анамнез"

	fmt.Println("Задача 2. Объявление переменных")
	fmt.Printf(
		"Число v_i: %d" +
		"Число с плавающей запятой v_f64: %f" +
		"Булевое значение v_b: %t" +
		"Строка v_s: %s\n",
		v_i,
		v_f64,
		v_b,
		v_s)
}

// task3 делает тоже что и task2, но
// объявляет переменные кратко
func task3() {
	v_i := 10
	v_f64 := 15.5
	v_b := true
	v_s := "Синопсис"

	fmt.Println("Задача 3. Обявление перменных через :=")
	fmt.Printf(
		"Число v_i: %d\n" +
		"Число с плавающей запятой v_f64: %f\n" +
		"Булевое значение v_b: %t\n" +
		"Строка v_s: %s\n",
		v_i,
		v_f64,
		v_b,
		v_s)
}

// Задание 4 производит основные 
// арифметические операции над числами
func task4(a, b int) [5]int {
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	rem := a % b
	fmt.Println("Задача 4")
	fmt.Printf(
		"Сумма: %d\n" + 
		"Разность: %d\n" +
		"Произведение: %d\n" +
		"Деление: %d\n" +
		"Остаток от деления: %d\n",
		sum,
		sub,
		mul,
	 	div,
		rem)
	return [5]int{sum,sub,mul,div,rem}
}

func task5(a, b float64) [2]float64 {
	sum := a + b
	sub := a - b

	fmt.Println("Задача 5. Сумма и разность чисел с плавающей запятой")
	fmt.Printf(
		"Сумма %f и %f: %f\n" +
		"Разность %f и %f: %f\n",
		a, b, sum,
		a, b, sub)

	return [2]float64{sum, sub}
}

// task6 находит среднее арифметическое значение 
// 3-х чисел
func task6(a, b, c int) {
	fmt.Println("Задача 6")
	fmt.Printf("Среднее арифметическое чисел %d, %d и %d: %f\n", a, b, c, float64((a + b + c) / 3))
}

func main() {
	task1()
	task2()
	task3()
	task4(11,7)
	task5(10.11,34.55)
	task6(2,6,1)
}