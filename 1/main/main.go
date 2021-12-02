package main

import (
	"1/parser"
	"fmt"
)

func main() {
	arr := parser.ReadInput("input.txt")

	count := SolvePart2(arr)

	fmt.Println(count)
}

func SolvePart1(arr []int) int {
	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i - 1] {
			count++
		}
	}
	return count
}

func SolvePart2(arr []int) int {
	if len(arr) < 4 {
		return 0
	}
	count := 0

	mprev := arr[0] + arr[1] + arr[2]
	prev := 0

	var mn int

	for i := 3; i < len(arr); i++ {
		mn = mprev - prev + arr[i]
		if mn > mprev {
			count++
		}
		mprev = mn
		prev = arr[i - 2]
	}

	return count
}

