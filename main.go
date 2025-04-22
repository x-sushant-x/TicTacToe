package main

import (
	"fmt"
	"os"
)

func main() {
	lobby := NewLobby()

	for {
		choice := lobby.ShowLobby()

		switch choice {
		case "1":
			gameManager := NewGameManager(lobby)
			gameManager.StartGame()
		case "2":
			fmt.Println("Showing Leaderboard")
		case "3":
			fmt.Println("👋 Exiting")
			os.Exit(0)
		default:
			fmt.Println("❌ Invalid Choice")
		}
	}

}
