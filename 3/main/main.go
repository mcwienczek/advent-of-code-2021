package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	wordLength := 12
	arr := ReadInput("input.txt")
	result1 := Part1(wordLength, arr)

	fmt.Printf("part1: %d\n", result1)
	result2a := Part2(wordLength, arr, Higest)
	result2b := Part2(wordLength, arr, Lowest)
	fmt.Printf("part2: %d\n", result2a * result2b)
}

func Higest(oneCount int, zeroCount int) bool {
	return oneCount >= zeroCount
}

func Lowest(oneCount int, zeroCount int) bool {
	return oneCount < zeroCount
}

type CompareOperator func(int, int) bool

func Part2(wordLength int, arr []uint64, cmp CompareOperator) uint64 {
	var selectedOnes []uint64

	selectedOnes = arr
	for i := wordLength; i > 0; i-- {
		zeroCount := 0
		oneCount := 0
		var ones []uint64
		var zeroes []uint64

		var one uint64 = 1 << (i - 1)
		for j := 0; j < len(selectedOnes); j++ {
			if selectedOnes[j] & one == 0 {
				zeroes = append(zeroes, selectedOnes[j])
				zeroCount += 1
			} else {
				ones = append(ones, selectedOnes[j])
				oneCount += 1
			}
		}

		if cmp(oneCount, zeroCount) {
			selectedOnes = ones
		} else {
			selectedOnes = zeroes
		}

		if len(selectedOnes) == 1 {
			return selectedOnes[0]
		}
	}
	return selectedOnes[0]
}



func Part1(wordLength int, arr []uint64) uint64 {
	var gamma uint64 = 0

	for i := wordLength; i > 0; i-- {
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
	epsilon = epsilon << (64 - wordLength)
	epsilon = epsilon >> (64 - wordLength)
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
