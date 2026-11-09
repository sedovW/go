package main

import "fmt"

func main() {
	var a int
	var b int
	var result int
	fmt.Println("Введите ваше первое число: ")
	fmt.Scanln(&a)
	fmt.Println("Введите ваше второе число:")
	fmt.Scanln(&b)
	result = multiple(a, b)
	fmt.Println("Результат:", result)
}
