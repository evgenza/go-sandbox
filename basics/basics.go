package main

import (
	"fmt"
	"time"
)

// Основная функция
func main() {
	// Объявление переменных
	var a int = 10       // Явное объявление типа
	b := 20              // Неявное определение типа
	var c, d = 30, "Go!" // Несколько переменных

	fmt.Println("Переменные:", a, b, c, d)

	// Условный оператор
	if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("b больше или равно a")
	}

	// Циклы
	for i := 0; i < 5; i++ {
		fmt.Println("Цикл:", i)
	}

	// Бесконечный цикл (выход по условию)
	x := 0
	for {
		if x >= 3 {
			break
		}
		fmt.Println("Бесконечный цикл, x =", x)
		x++
	}

	// Вызов функции
	result := add(15, 25)
	fmt.Println("Результат сложения:", result)

	// Работа со структурами
	p := Person{Name: "Alice", Age: 25}
	fmt.Println("Структура:", p.Name, p.Age)

	// Работа с массивами
	arr := [3]int{1, 2, 3}
	fmt.Println("Массив:", arr)

	// Работа со срезами
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 5)
	fmt.Println("Срез:", slice)

	// Работа с картами
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println("Карта:", m["one"])

	// Горутины и каналы
	go sayHello() // запуск горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Завершение main")
}

// Простая функция
func add(a int, b int) int {
	return a + b
}

// Работа со структурами
type Person struct {
	Name string
	Age  int
}

// Горутина
func sayHello() {
	fmt.Println("Привет из горутины!")
}
