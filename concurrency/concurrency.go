package main

import (
	"fmt"
	"sync"
	"time"
)

// Главная функция
func main() {
	// Пример параллельного выполнения
	fmt.Println("Начинаем параллельное выполнение...")

	// Запускаем горутины
	go printMessage("Привет из первой горутины!")
	go printMessage("Привет из второй горутины!")

	// Используем time.Sleep для ожидания завершения горутин
	time.Sleep(1 * time.Second)

	// Использование WaitGroup для синхронизации
	var wg sync.WaitGroup

	wg.Add(2) // Указываем количество горутин, которые должны завершиться

	go func() {
		defer wg.Done()
		printNumbers("Горутина 1")
	}()

	go func() {
		defer wg.Done()
		printNumbers("Горутина 2")
	}()

	wg.Wait() // Ждём завершения всех горутин
	fmt.Println("Все горутины завершены!")

	// Использование каналов для взаимодействия между горутинами
	messageChannel := make(chan string)

	go sendMessage("Сообщение через канал", messageChannel)

	receivedMessage := <-messageChannel // Читаем сообщение из канала
	fmt.Println("Получено из канала:", receivedMessage)

	// Использование буферизованных каналов
	bufferedChannel := make(chan int, 3)

	go sendNumbers(bufferedChannel)

	// Читаем числа из канала
	for num := range bufferedChannel {
		fmt.Println("Получено из буферизованного канала:", num)
	}
}

// Функция для вывода сообщения
func printMessage(message string) {
	fmt.Println(message)
}

// Функция для вывода чисел
func printNumbers(prefix string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(200 * time.Millisecond)
	}
}

// Функция для отправки сообщения в канал
func sendMessage(message string, ch chan string) {
	time.Sleep(500 * time.Millisecond)
	ch <- message // Отправляем сообщение в канал
}

// Функция для отправки чисел в буферизованный канал
func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Отправляем числа в канал
	}
	close(ch) // Закрываем канал после завершения
}
