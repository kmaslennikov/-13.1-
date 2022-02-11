package main

import (
	"fmt"
)

func main() {
	sum(2, 11)
}

func sum(start, end int) {
	var sum int

	if start > end {
		temp := end
		end = start
		start = temp
	}
	if start%2 != 0 {
		start = start + 1
	}
	for i := start; i <= end; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}
