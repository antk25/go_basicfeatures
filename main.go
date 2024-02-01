package main

import "fmt"

func main() {
	var price float32 = 275.00
	var tax float32 = 27.50
	fmt.Println(price + tax)
	// Значение 300 можно присвоить переменной типа float потому что 300 - является нетипизированной константой и преобразуется
	price = 300
	fmt.Println(price + tax)
}
