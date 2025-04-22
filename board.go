package main

import "fmt"

type TicTacToeBoard struct {
	Grid [3][3]string
}

func NewTicTacToeBoard() TicTacToeBoard {
	return TicTacToeBoard{
		Grid: [3][3]string{},
	}
}

func (b TicTacToeBoard) Display() {
	for i := range 3 {
		for j := range 3 {
			cell := b.Grid[i][j]

			if cell == "" {
				fmt.Print("   ")
			} else {
				fmt.Printf(" %s", cell)
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
}
