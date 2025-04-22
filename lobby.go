package main

import (
	"fmt"
	"strings"
)

type Lobby struct{}

func NewLobby() *Lobby {
	return &Lobby{}
}

func (l *Lobby) ShowLobby() string {
	fmt.Printf("\n\n")
	fmt.Printf("🎯 Welcome To Tic Tac Toe 🎯\n")

	fmt.Println("1. Start Game")
	fmt.Println("2. Show Leaderboard")
	fmt.Println("3. Exit")

	var input string

	for {
		input = ShowInputPrompt("Choose From Options: ")

		input = strings.TrimSpace(input)

		if input != "1" && input != "2" && input != "3" {
			fmt.Printf("❌ Invalid Choice \n\n")
			continue
		}

		return input
	}
}
