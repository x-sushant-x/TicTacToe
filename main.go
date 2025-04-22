package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	lobby := NewLobby()
	leaderboard := NewLeaderboard()

	for {
		choice := lobby.ShowLobby()

		switch choice {
		case "1":
			gameManager := NewGameManager(lobby, leaderboard, Multiplayer)
			if gameManager.StartGame() {
				continue
			}

		case "2":
			botGameManager := NewGameManager(lobby, leaderboard, SinglePlayer)
			if botGameManager.StartGame() {
				continue
			}
		case "3":
			leaderboard.PrintHighScores()
		case "4":
			fmt.Println("👋 Exiting")
			os.Exit(0)
		default:
			fmt.Println("❌ Invalid Choice")
		}
	}
}
