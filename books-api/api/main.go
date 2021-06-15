package main

import "fmt"

func main() {
	// fmt.Println("Run")

	// r := router.Gerar()

	// log.Fatal(http.ListenAndServe(":5000", r))

	list := []int{1, 4, 5, 7, 9}

	count := 0

	for _, n := range list[:4] {
		count += n
	}
	fmt.Println(count)
	fmt.Println(len(list))
}
