package main

import (
	"fmt"
	"os"
)

func main() {
	lobby := NewLobby()
	leaderboard := NewLeaderboard()

	for {
		choice := lobby.ShowLobby()

		switch choice {
		case "1":
			gameManager := NewGameManager(lobby, leaderboard)
			if gameManager.StartGame() {
				continue
			}
		case "2":
			leaderboard.PrintHighScores()
		case "3":
			fmt.Println("👋 Exiting")
			os.Exit(0)
		default:
			fmt.Println("❌ Invalid Choice")
		}
	}

}
