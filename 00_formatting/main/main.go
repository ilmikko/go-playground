package main

import (
	"fmt"
	"github.com/tati-z/Go-Arena/00_formatting/utility"
	"time"
)

// package scope
func max(x int, y int) int {
	return map[bool]int{true: x, false:y}[x>y] // ternary operator hack but if-else=better
}
func main() {
	t := time.Now()
	num := 9924
	custom := fmt.Sprintf("Today is %v %vth", t.Month(), t.Day())
	fmt.Print(custom + "\n")
	// preced the verb with '#' for raw display : e.g %#x will add the '0x' hex prefix
	fmt.Printf("d:%[1]d b:%[1]b oc:%[1]o h:%[1]x h#:%#[1]x uniC:%[1]c uniC:%[1]U \n", num)
	fmt.Printf("%t %3.2f %[4]d %[3]s\n", (6 < 8), (10.0 / 3), "😍", 0x26C4)
	fmt.Printf("%v %[1]T", '😰')
	// parsing emoji's hex unicode values
	for i := 0x0; i < 0x10*5; i++ {
		if (i % 0x10) == 0 {
			fmt.Println()
		}
		fmt.Printf("%c ", 0x1F600+i)
	}
	for i := 2; i < 40; i++ {
		if utility.IsPrime(i) {
		fmt.Printf("%d =  %T:%[2]t\n", i, utility.IsPrime(i))
		}
	}
	max := max(2,3)
	// y := max(5,6) cannot be done as max is now an integer from here on
	fmt.Printf("max = %d\n", max)

}
