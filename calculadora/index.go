package calculadora

func Soma(a, b *int) int {
	return *a + *b
}

func Subt(a, b *int) int {
	return *a - *b
}

func Mult(a, b *int) int {
	return (*a) * (*b)
}

func Div(a, b *int) int {
	return (*a) / (*b)
}
