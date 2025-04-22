package main

import "fmt"

type GameMode int

const (
	SINGLE_PLAYER GameMode = iota
	MULTIPLAYER
)

type GameManager struct {
	gameMode  GameMode
	playerOne Player
	playerTwo Player
	board     TicTacToeBoard
	turn      Player
}

func NewGameManager() GameManager {
	playerOneName := ShowInputPrompt("Enter Player 1 Name: ")
	playerTwoName := ShowInputPrompt("Enter Player 2 Name: ")

	playerOne := Player{
		Username: playerOneName,
	}

	playerTwo := Player{
		Username: playerTwoName,
	}

	return GameManager{
		gameMode:  MULTIPLAYER,
		playerOne: playerOne,
		playerTwo: playerTwo,
		board:     NewTicTacToeBoard(),
		turn:      playerOne,
	}
}

func (m GameManager) StartGame() {
	ClearTerminal()
	fmt.Printf("🚀 Match Starting: %s 🆚 %s\n", m.playerOne.Username, m.playerTwo.Username)

	m.board.Display()

	for {
		m.TakeInput()
	}
}

func (m GameManager) TakeInput() {
	ShowInputPrompt("Current Turn: " + m.turn.Username)
}

func (m GameManager) CheckWinningCondition() {}
