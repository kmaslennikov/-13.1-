package main

import (
	"fmt"
)

func main() {
  fmt.Println("Введите значения границы диапазона:")
  var start, end int
  fmt.Scan(&start)
  fmt.Scan(&end)
  
	sum(start, end)
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
