package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	marked  [5][5]int
	numbers [5][5]int
	id      int
}

func main() {
	n, b := ParseInput("input.txt")
	score := Part1(n, b)
	fmt.Printf("score: %d\n", score)
	score2 := Part2(n, b)
	fmt.Printf("score2: %d\n", score2)
}

func Part1(numbers []int, boards []*Board) int {
	for _, n := range numbers {
		for _, b := range boards {
			b.MarkNumber(n)
			if b.HasWin() {
				return b.GetBoardScore() * n
			}
		}
	}
	return 0
}

func AlreadyWonBefore(boardId int, alreadyWonBoards []int) bool {
	for _, n := range alreadyWonBoards {
		if n == boardId {
			return true
		}
	}
	return false
}

func Part2(numbers []int, boards []*Board) int {
	var lastWinScore int
	var winningBoardIds []int
	for _, n := range numbers {
		for _, b := range boards {
			b.MarkNumber(n)
			if b.HasWin() && !AlreadyWonBefore(b.id, winningBoardIds) {
				winningBoardIds = append(winningBoardIds, b.id)
				//fmt.Printf("Board Score: %d\n", b.GetBoardScore())
				//fmt.Printf("n: %d\n", n)
				lastWinScore = b.GetBoardScore() * n
			}
		}
	}
	return lastWinScore
}

func (b *Board) MarkNumber(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] == n {
				b.marked[i][j] = 1
			}
		}
	}
}

func (b *Board) GetBoardScore() int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.marked[i][j] == 0 {
				score += b.numbers[i][j]
			}
		}
	}
	return score
}

func (b *Board) HasWin() bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.marked[i][j] == 0 {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.marked[j][i] == 0 {
				break
			}
			if j == 4 {
				return true
			}
		}
	}

	return false
}

func (b *Board) UnmarkAll() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b.marked[i][j] = 0
		}
	}
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func ParseInput(file string) (numbers []int, boards []*Board) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("There has been an error!: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	numberStr := strings.Split(scanner.Text(), ",")

	for _, val := range numberStr {
		number, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("There has been an error!: %v", err)
		}
		numbers = append(numbers, number)
	}

	bId := 1
	for scanner.Scan() {
		//5 board rows
		b := Board{}
		for i := 0; i < 5; i++ {
			scanner.Scan()
			str := scanner.Text()
			sarr := strings.Split(str, " ")
			sarr = delete_empty(sarr)
			if len(sarr) != 5 {
				log.Fatalf("board is not 5x5: %v", sarr)
			}
			for j := 0; j < 5; j++ {
				n, err := strconv.Atoi(sarr[j])
				if err != nil {
					log.Fatalf("There has been an error!: %v", err)
				}
				b.numbers[i][j] = n
			}
		}

		b.id = bId
		bId++
		boards = append(boards, &b)
	}
	return
}
