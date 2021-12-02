package parser

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInput(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("There has been an error!: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var arr []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
		arr = append(arr, number)
		if err := scanner.Err(); err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
	}
	return arr
}