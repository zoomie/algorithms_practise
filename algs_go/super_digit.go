package main

import (
	"fmt"
	"strconv"
)

func superDigit(digits int64) {
	stringDigits := strconv.FormatInt(digits, 10)
	fmt.Println(stringDigits)
	for pos, digit := range stringDigits {
		fmt.Println(pos, int(digit)-48)
	}
}
func main() {
	digit := 12345
	superDigit(int64(digit))
}
