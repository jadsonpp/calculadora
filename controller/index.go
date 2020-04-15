package controller

func SomaPointer(a, b *int) int {
	return *a + *b
}

func SubtPointer(a, b *int) int {
	return *a - *b
}

func MultPointer(a, b *int) int {
	return (*a) * (*b)
}

func DivPointer(a, b *int) int {
	return (*a) / (*b)
}

func Soma(a, b int) int {
	return a + b
}

func Subt(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}
