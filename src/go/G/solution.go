package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

func main() {
	var sc scanner.Scanner
	sc.Init(bufio.NewReader(os.Stdin))
	sc.Whitespace &^= 1 << '\n'
	var sum int
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		switch tok {
		case '\n':
			fmt.Println(sum)
			sum = 0
		default:
			n, _ := strconv.Atoi(sc.TokenText())
			sum += n
		}
	}
}
