package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		var sum int
		for n, _ := strconv.Atoi(sc.Text()); n > 0; n-- {
			sc.Scan()
			m, _ := strconv.Atoi(sc.Text())
			sum += m
		}
		fmt.Println(sum)
	}
}
