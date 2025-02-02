package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println(Soma(10, 10))
	fmt.Println(Subtracao(10, 10))
	fmt.Println(Multiplicacao(10, 10))
	fmt.Println(Divisao(10, 10))
}

func Soma(a, b int) int {
	return a + b
}

func Subtracao(a, b int) int {
	return a - b
}

func Multiplicacao(a, b int) int {
	return a * b
}

func Divisao(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
