package main

import (
	"fmt"
	"strconv"
)

type GameMode int

const (
	SINGLE_PLAYER GameMode = iota
	MULTIPLAYER
)

type GameManager struct {
	gameMode  GameMode
	playerOne *Player
	playerTwo *Player
	board     *TicTacToeBoard
	turn      *Player
}

func NewGameManager() *GameManager {
	playerOneName := ShowInputPrompt("Enter Player 1 Name: ")
	playerTwoName := ShowInputPrompt("Enter Player 2 Name: ")

	playerOne := Player{
		Username: playerOneName,
	}

	playerTwo := Player{
		Username: playerTwoName,
	}

	return &GameManager{
		gameMode:  MULTIPLAYER,
		playerOne: &playerOne,
		playerTwo: &playerTwo,
		board:     NewTicTacToeBoard(),
		turn:      &playerOne,
	}
}

func (m *GameManager) StartGame() {
	ClearTerminal()
	fmt.Printf("🚀 Match Starting: %s 🆚 %s\n", m.playerOne.Username, m.playerTwo.Username)

	m.board.Display()

	for {
		m.TakeInput()
		m.board.Display()

		if m.turn == m.playerOne {
			m.turn = m.playerTwo
		} else {
			m.turn = m.playerOne
		}
	}
}

func (m *GameManager) TakeInput() {
	fmt.Println("Current Turn: " + m.turn.Username)

	fmt.Println("Valid Moves:", m.board.GetValidMoves())
	input := ShowInputPrompt("Enter your move: ")

	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("❌ Invalid Input")
		return
	}

	if m.turn == m.playerOne {
		m.board.Mark(i, "O")
	} else {
		m.board.Mark(i, "X")
	}
}

func (m *GameManager) CheckWinningCondition() {}
