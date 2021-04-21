package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	readInt := func() int {
		sc.Scan()
		r := bytesToInt(sc.Bytes())

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

func bytesToInt(buf []byte) int {
	var n int
	for _, v := range buf {
		n += int(v - '0')
		n *= 10
	}
	return n
}
