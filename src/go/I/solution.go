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
	for sc.Scan() {
		line := sc.Text()
		splits := strings.Split(line, " ")
		sort.Strings(splits)
		fmt.Print(splits[0])
		for _, str := range splits[1:] {
			fmt.Print(" ", str)
		}
		fmt.Println()
	}
}
