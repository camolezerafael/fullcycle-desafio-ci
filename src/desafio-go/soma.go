package main

import "fmt"

func Soma( a int, b int ) int{
	return a + b
}

func main() {
	resultado := Soma(5,5)

	fmt.Println(resultado)
}
