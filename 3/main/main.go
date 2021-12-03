package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arr := ReadInput("input.txt")
	fmt.Printf("part1: %d", Solve1(arr))
}

func Part1(arr []uint64) int {
	var gamma uint64 = 0

	for i := 12; i > 0; i-- {
		zeroCount := 0
		oneCount := 0
		var one uint64 = 1 << (i - 1)
		//fmt.Printf("one: %05b\n", one)
		for j := 0; j < len(arr); j++ {
			//fmt.Printf("arr[j]: %05b\n", arr[j])
			//fmt.Printf("arr[j]&one: %05b\n", arr[j]&one)
			if arr[j]&one == 0 {
				zeroCount += 1
			} else {
				oneCount += 1
			}
		}
		if oneCount > zeroCount {
			//fmt.Printf("iteration %d, most common bit is 1\n", i)
			gamma += 1
		} else {
			//fmt.Printf("iteration %d, most common bit is 0\n", i)
		}
		gamma = gamma << 1
		//fmt.Printf("iteration %d gamma: %05b\n", i, gamma)
	}
	gamma = gamma >> 1
	var epsilon uint64 = ^gamma
	//fmt.Printf("epsilon: %b", epsilon)
	epsilon = epsilon << (64 - 12)
	epsilon = epsilon >> (64 - 12)
	//fmt.Printf("epsilon: %b\n gamma: %b\n", epsilon, gamma)
	return gamma * epsilon
}

func ReadInput(file string) []uint64 {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("There has been an error!: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var arr []uint64

	for scanner.Scan() {
		binaryStr := scanner.Text()
		n, err := strconv.ParseUint(binaryStr, 2, 12)
		if err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
		arr = append(arr, n)
	}

	return arr
}
