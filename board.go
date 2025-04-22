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
