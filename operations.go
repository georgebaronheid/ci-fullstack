package main

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
