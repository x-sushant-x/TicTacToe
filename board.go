package main

import (
	"fmt"
	"sort"
)

type Coordinates struct {
	X      int
	Y      int
	isUsed bool
}

type TicTacToeBoard struct {
	Grid [3][3]string

	mapping map[int]Coordinates
}

func NewTicTacToeBoard() *TicTacToeBoard {
	mapping := map[int]Coordinates{
		1: {X: 0, Y: 0, isUsed: false},
		2: {X: 0, Y: 1, isUsed: false},
		3: {X: 0, Y: 2, isUsed: false},
		4: {X: 1, Y: 0, isUsed: false},
		5: {X: 1, Y: 1, isUsed: false},
		6: {X: 1, Y: 2, isUsed: false},
		7: {X: 2, Y: 0, isUsed: false},
		8: {X: 2, Y: 1, isUsed: false},
		9: {X: 2, Y: 2, isUsed: false},
	}

	return &TicTacToeBoard{
		Grid:    [3][3]string{},
		mapping: mapping,
	}
}

func (b *TicTacToeBoard) Display() {
	fmt.Println()

	for i := range 3 {
		for j := range 3 {
			cell := b.Grid[i][j]

			if cell == "" {
				fmt.Print("   ")
			} else {
				fmt.Printf(" %s ", cell)
			}

			if j < 2 {
				fmt.Print("|")
			}
		}

		fmt.Println()

		if i < 2 {
			fmt.Println("---+---+---")
		}
	}

	fmt.Println()
}

func (b *TicTacToeBoard) Mark(place int, value string) {
	coordinates, found := b.mapping[place]

	if !found {
		fmt.Println("❌ Invalid Input")
	} else {
		coordinates.isUsed = true
		b.mapping[place] = coordinates
		b.Grid[coordinates.X][coordinates.Y] = value
	}
}

func (b *TicTacToeBoard) GetValidMoves() []int {
	validMovies := []int{}

	for k, v := range b.mapping {
		if !v.isUsed {
			validMovies = append(validMovies, k)
		}
	}

	sort.Ints(validMovies)

	return validMovies
}

type Result struct {
	Win      bool
	Draw     bool
	Winner   string
	Continue bool
}

func createResult(win bool, draw bool, winner string, Continue bool) Result {
	return Result{
		Win:      win,
		Draw:     draw,
		Winner:   winner,
		Continue: Continue,
	}
}

func (board *TicTacToeBoard) CheckWin() Result {
	var result Result

	winningLines := [][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},

		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},

		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range winningLines {
		a, b, c := line[0], line[1], line[2]

		if board.Grid[a[0]][a[1]] != "" &&
			board.Grid[a[0]][a[1]] == board.Grid[b[0]][b[1]] &&
			board.Grid[a[0]][a[1]] == board.Grid[c[0]][c[1]] {

			if board.Grid[a[0]][a[1]] == "O" {
				result = createResult(true, false, "O", false)
			} else {
				result = createResult(true, false, "X", false)
			}
		}
	}

	isDraw := true
	for i := range 3 {
		for j := range 3 {
			if board.Grid[i][j] == "" {
				isDraw = false
				break
			}
		}
	}

	if isDraw {
		return createResult(false, true, "", false)
	}

	return result
}

func (board *TicTacToeBoard) Reset() {
	for i := range board.Grid {
		for j := range board.Grid[i] {
			board.Grid[i][j] = ""
		}
	}

	for k, v := range board.mapping {
		v.isUsed = false
		board.mapping[k] = v
	}
}
