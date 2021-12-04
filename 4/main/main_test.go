package main

import "testing"

func TestIsWinOnBoard1(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		numbers: [5][5]int{},
	}

	if !b.HasWin() {
		t.Fail()
	}
}

func TestIsWinOnBoard2(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		numbers: [5][5]int{},
	}

	if b.HasWin() {
		t.Fail()
	}
}

func TestIsWinOnBoard3(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0},
		},
		numbers: [5][5]int{},
	}

	res := b.HasWin()
	if res == false {
		t.Fail()
	}
}

func TestIsWinOnBoard4(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		},
		numbers: [5][5]int{},
	}

	res := b.HasWin()
	if res == true {
		t.Fail()
	}
}

func TestIsWinOnBoard5(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
		},
		numbers: [5][5]int{},
	}

	if !b.HasWin() {
		t.Fail()
	}
}

func TestIsWinOnBoard6(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 0, 0, 0, 0},
			{1, 0, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{1, 0, 0, 1, 0},
			{1, 0, 0, 0, 0},
		},
		numbers: [5][5]int{},
	}

	if !b.HasWin() {
		t.Fail()
	}
}

func TestIsWinOnBoard7(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 0, 0, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 0, 1, 1},
			{1, 0, 0, 1, 1},
			{1, 0, 0, 0, 1},
		},
		numbers: [5][5]int{},
	}

	if !b.HasWin() {
		t.Fail()
	}
}

func TestIsWinOnBoard8(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 0, 0, 0, 1},
			{0, 0, 1, 0, 1},
			{1, 0, 0, 1, 1},
			{1, 0, 0, 1, 0},
			{1, 0, 0, 0, 1},
		},
		numbers: [5][5]int{},
	}

	if b.HasWin() {
		t.Fail()
	}
}

func TestIsWinOnBoard9(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{1, 0, 0, 0, 1},
			{0, 1, 1, 0, 1},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		},
		numbers: [5][5]int{},
	}

	if b.HasWin() {
		t.Fail()
	}
}

func TestIsWinOnBoard10(t *testing.T) {
	b := Board{
		marked: [5][5]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		numbers: [5][5]int{},
	}

	if b.HasWin() {
		t.Fail()
	}
}
