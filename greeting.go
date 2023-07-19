package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Создаем новый сканер для чтения ввода пользователя
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Введите ваше имя: ")
    scanner.Scan() // Читаем строку, введенную пользователем

    // Получаем имя пользователя из ввода
    userName := scanner.Text()

    // Проверяем, ввел ли пользователь имя или ничего не ввел
    if userName == "" {
        fmt.Println("Пожалуйста, введите ваше имя.")
    } else {
        fmt.Printf("Привет, %s! Хорошего дня!\n", userName)
    }
}
