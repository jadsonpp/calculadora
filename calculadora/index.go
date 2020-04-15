package calculadora

func soma(a, b *int) int {
	return *a + *b
}

func subt(a, b *int) int {
	return *a - *b
}

func mult(a, b *int) int {
	return (*a) * (*b)
}

func div(a, b *int) int {
	return (*a) / (*b)
}
