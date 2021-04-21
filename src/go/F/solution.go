package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		var sum int
		for n := bytesToInt(sc.Bytes()); n > 0; n-- {
			sc.Scan()
			m := bytesToInt(sc.Bytes())
			sum += m
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
