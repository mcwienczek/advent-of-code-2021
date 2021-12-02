package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput("input.txt")
	fmt.Println(Part2(input))

}

func Part1(input [][2]int) int {
	forward := 0
	depth := 0

	for i := 0; i < len(input); i++ {
		switch input[i][0] {
		case 0: // forward
			forward += input[i][1]
		case 1: // up
			depth -= input[i][1]
		case 2: // down
			depth += input[i][1]
		}

	}
	return forward * depth
}

func Part2(input [][2]int) int {
	forward := 0
	aim := 0
	depth := 0

	for i := 0; i < len(input); i++ {
		switch input[i][0] {
		case 0: // forward
			depth += input[i][1] * aim
			forward += input[i][1]
		case 1: // up
			aim -= input[i][1]
		case 2: // down
			aim += input[i][1]
		}
	}
	return forward * depth
}

func ReadInput(file string) [][2]int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("There has been an error!: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var acc [][2]int

	for scanner.Scan() {
		str := scanner.Text()
		sarr := strings.Split(str, " ")
		command := sarr[0]
		value, err := strconv.Atoi(sarr[1])
		if err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
		var v [2]int
		switch {
		case command == "forward":
			v[0] = 0
		case command == "up":
			v[0] = 1
		case command == "down":
			v[0] = 2

		}
		v[1] = value

		if err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
		acc = append(acc, v)
		if err := scanner.Err(); err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
	}
	return acc
}
