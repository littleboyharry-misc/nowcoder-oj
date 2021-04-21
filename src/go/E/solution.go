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

	readInt := func() int {
		sc.Scan()
		r, _ := strconv.Atoi(sc.Text())
		return r
	}

	for t := readInt(); t > 0; t-- {
		var sum int
		for n := readInt(); n > 0; n-- {
			sum += readInt()
		}
		fmt.Println(sum)
	}
}
