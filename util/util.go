package util

import (
	"bufio"
)

func GetLine(sc *bufio.Scanner) string {
	if !sc.Scan() {
		return ""
	}
	return sc.Text()
}
