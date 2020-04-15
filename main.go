package main

import (
	"calculadora/controller"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	a, b := 2, 3
	fmt.Println("Usando Ponteiros")
	fmt.Println("Soma: ", controller.SomaPointer(&a, &b))
	fmt.Println("Subtração: ", controller.SubtPointer(&a, &b))
	fmt.Println("Multiplicação: ", controller.MultPointer(&a, &b))
	fmt.Println("Divisão: ", controller.DivPointer(&a, &b))
	fmt.Println("Tempo gasto: ", time.Since(start))

	fmt.Println("\nSem Ponteiros")
	start2 := time.Now()
	fmt.Println("Soma: ", controller.Soma(a, b))
	fmt.Println("Subtração: ", controller.Subt(a, b))
	fmt.Println("Multiplicação: ", controller.Mult(a, b))
	fmt.Println("Divisão: ", controller.Div(a, b))
	fmt.Println("Tempo gasto: ", time.Since(start2))
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	http.ListenAndServe(":8000", nil)
	*/
}
