package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type GameMode int

const (
	SINGLE_PLAYER GameMode = iota
	MULTIPLAYER

	PlayerOneMove = "O"
	PlayerTwoMove = "X"
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
		err := m.takeInput()
		m.board.Display()

		if err == nil {
			if m.turn == m.playerOne {
				m.turn = m.playerTwo
			} else {
				m.turn = m.playerOne
			}
		}

		m.checkWinningCondition()
	}
}

func (m *GameManager) takeInput() error {
	fmt.Println("Current Turn: " + m.turn.Username)

	validMoves := m.board.GetValidMoves()

	fmt.Println("Valid Moves:", validMoves)
	input := ShowInputPrompt("Enter your move: ")

	i, err := strconv.Atoi(input)
	err = checkValidMove(err, validMoves, i)

	if err != nil {
		return err
	}

	if m.turn == m.playerOne {
		m.board.Mark(i, PlayerOneMove)
	} else {
		m.board.Mark(i, PlayerTwoMove)
	}

	return nil
}

func (m *GameManager) checkWinningCondition() {
	result := m.board.CheckWin()

	if result.Draw {
		fmt.Println("🤝 It's a draw!")
		os.Exit(0)
	}

	if result.Win {
		winner := result.Winner

		if winner == "O" {
			fmt.Printf("🏆 %s wins!!!\n", m.playerOne.Username)
		} else {
			fmt.Printf("🏆 %s wins!!!\n", m.playerTwo.Username)
		}

		os.Exit(0)
	}
}

func checkValidMove(err error, validMoves []int, i int) error {
	if err != nil {
		fmt.Println("❌ Invalid Input")
		return errors.New("invalid input")
	}

	valid := slices.Contains(validMoves, i)

	if !valid {
		fmt.Println("❌ Invalid Move")
		return errors.New("invalid move")
	}

	return nil
}

func (m *GameManager) switchTurn() {
	if m.turn == m.playerOne {
		m.turn = m.playerTwo
	} else {
		m.turn = m.playerOne
	}
}
