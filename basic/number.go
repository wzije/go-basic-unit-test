package tdd

import (
	"fmt"
)

func OddOrEven(num int) string {
	value := num % 2

	if value == 1 || value == -1 {
		return fmt.Sprintf("%d is odd number", num)
	} else {
		return fmt.Sprintf("%d is even number", num)
	}
}
