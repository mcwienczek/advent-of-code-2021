package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	arr := ReadInput("input_test.txt")
	for i := 5; i < 5;i++ {
		for j := 0; j < len(arr); j++ {
			
		}
	}
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
		n, err := strconv.ParseUint(binaryStr, 2, 5))
		if err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
		arr = append(arr, n)
	}

	return arr
}
