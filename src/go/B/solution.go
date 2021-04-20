package main

import (
	"fmt"
)

func main() {
	var n int
	for fmt.Scan(&n); n > 0; n-- {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
