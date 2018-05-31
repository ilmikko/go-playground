package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 0 {
		number, err := strconv.Atoi(os.Args[1])
		if err == nil {
			fmt.Println(IsPrime(number))
		}
	}
}
func IsPrime(value int) bool {
	if value < 2 {
		return false
	} else {
		cache := map[int]bool{2: true, 3: true, 5: true} // use if-else instead
		if cache[value] {
			return true
		}
		for digit := range cache {
			if value%digit == 0 {
				return false
			}
		}
		for i := 7; i*i < value; i += 2 {
			if value%i == 0 {
				return false
			}
		}
	}
	return true
}
