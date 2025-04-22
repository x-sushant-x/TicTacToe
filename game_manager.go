package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"time"
)

type GameMode int

const (
	SinglePlayer GameMode = iota
	Multiplayer

	PlayerOneMove = "O"
	PlayerTwoMove = "X"
)

type GameManager struct {
	gameMode    GameMode
	playerOne   *Player
	playerTwo   *Player
	board       *TicTacToeBoard
	turn        *Player
	lobby       *Lobby
	leaderboard *Leaderboard
}

func NewGameManager(lobby *Lobby, leaderboard *Leaderboard, mode GameMode) *GameManager {
	playerOneName := ShowInputPrompt("Enter Player 1 Name: ")
	playerTwoName := "Bot"

	if mode == Multiplayer {
		playerTwoName = ShowInputPrompt("Enter Player 2 Name: ")
	}

	playerOne := Player{
		Username: playerOneName,
	}

	playerTwo := Player{
		Username: playerTwoName,
	}

	return &GameManager{
		gameMode:    mode,
		playerOne:   &playerOne,
		playerTwo:   &playerTwo,
		board:       NewTicTacToeBoard(),
		turn:        &playerOne,
		lobby:       lobby,
		leaderboard: leaderboard,
	}
}

func (m *GameManager) StartGame() bool {
	ClearTerminal()
	fmt.Printf("🚀 Match Starting: %s 🆚 %s\n", m.playerOne.Username, m.playerTwo.Username)

	m.board.Reset()
	m.board.Display()

	for {
		err := m.takeInput()
		m.board.Display()

		if err == nil {
			m.switchTurn()
		}

		isWinOrDraw := m.checkWinningCondition()

		if isWinOrDraw {
			break
		}
	}

	nextAction := ShowInputPrompt("Game Over! Press 1 to Restart Game or 2 to go to the Lobby: ")

	if nextAction == "1" {
		m.StartGame()
	} else if nextAction == "2" {
		return true
	}

	return false
}

func (m *GameManager) takeInput() error {
	fmt.Println("Current Turn: " + m.turn.Username)
	validMoves := m.board.GetValidMoves()
	fmt.Println("Valid Moves:", validMoves)

	var move int
	var err error

	if m.gameMode == SinglePlayer && m.turn == m.playerTwo {
		time.Sleep(time.Second * 1)
		move = validMoves[rand.Intn(len(validMoves))]
		fmt.Printf("🤖 Bot choose move: %d\n", move)
	} else {
		input := ShowInputPrompt("Enter your move: ")
		move, err = strconv.Atoi(input)
		err = checkValidMove(err, validMoves, move)

		if err != nil {
			return err
		}
	}

	if m.turn == m.playerOne {
		m.board.Mark(move, PlayerOneMove)
	} else {
		m.board.Mark(move, PlayerTwoMove)
	}

	return nil
}

func (m *GameManager) checkWinningCondition() bool {
	result := m.board.CheckWin()

	if result.Draw {
		fmt.Println("🤝 It's a draw!")
		return true
	}

	if result.Win {
		winner := result.Winner

		if winner == "O" {
			m.leaderboard.Update(m.playerOne.Username, "W")
			fmt.Printf("🏆 %s wins!!!\n\n", m.playerOne.Username)
		} else {
			m.leaderboard.Update(m.playerTwo.Username, "W")
			fmt.Printf("🏆 %s wins!!!\n\n", m.playerTwo.Username)
		}

		return true
	}

	return false
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
