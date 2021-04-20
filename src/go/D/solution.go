package main

import "fmt"

func main() {
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			sum := 0
			for ; n > 0; n-- {
				var i int
				fmt.Scan(&i)
				sum += i
			}
			fmt.Println(sum)
		}
	}
}
