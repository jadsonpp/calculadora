package main

import (
	"fmt"
)

func main() {
	a, b := 2, 3

	fmt.Println(calculadora.soma(&a, &b))
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	http.ListenAndServe(":8000", nil)
	*/
}
