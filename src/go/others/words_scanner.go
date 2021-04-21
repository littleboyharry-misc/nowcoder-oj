package others

import (
	"bufio"
	"fmt"
	"os"
)

//goland:noinspection GoUnusedExportedFunction
func ScanWords() {
	rd := bufio.NewScanner(os.Stdin)
	rd.Split(bufio.ScanWords)
	for rd.Scan() {
		fmt.Println(rd.Bytes())
	}
}
