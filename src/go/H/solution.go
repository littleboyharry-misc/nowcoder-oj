package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sc.Scan()
	line := sc.Text()
	words := strings.Split(line, " ")
	sort.Strings(words)
	if len(words) > 0 {
		fmt.Print(words[0])
		for _, word := range words[1:] {
			fmt.Print(' ', word)
		}
	}
}
