package main

import (
	"fmt"
	"sort"
)

type Leaderboard struct {
	scores  map[string]int
	history map[string][]string
}

func NewLeaderboard() *Leaderboard {
	return &Leaderboard{
		scores:  map[string]int{},
		history: map[string][]string{},
	}
}

func (l *Leaderboard) Update(username string, result string) {
	if _, exists := l.scores[username]; !exists {
		l.scores[username] = 0
	}

	if _, exists := l.history[username]; !exists {
		l.history[username] = []string{}
	}

	l.history[username] = append(l.history[username], result)

	if result == "W" {
		is5thWin := l.isFifthConsecutiveWin(username)

		if is5thWin {
			l.scores[username] += 10
			return
		}

		is3rdWin := l.isThirdConsecutiveWin(username)

		if is3rdWin {
			l.scores[username] += 5
			return
		}

		l.scores[username] += 2
	}
}

func (l *Leaderboard) IncrementScore(username string, points int) {
	l.scores[username] += points
}

func (l *Leaderboard) SaveHistory(username string, history string) {
	gameHistory := l.history[username]
	gameHistory = append(gameHistory, history)
	l.history[username] = gameHistory
}

func (l *Leaderboard) isThirdConsecutiveWin(username string) bool {
	history := l.history[username]
	if len(history) < 3 {
		return false
	}

	winStreak := 0
	for i := len(history) - 1; i >= 0 && history[i] == "W"; i-- {
		winStreak++
	}

	return winStreak >= 3 && winStreak%3 == 0
}

func (l *Leaderboard) isFifthConsecutiveWin(username string) bool {
	history := l.history[username]
	if len(history) < 5 {
		return false
	}

	winStreak := 0
	for i := len(history) - 1; i == 0; i-- {
		winStreak++
	}

	return winStreak >= 5 && winStreak%5 == 0
}

func (l *Leaderboard) PrintHighScores() {
	type pair struct {
		name  string
		score int
	}

	var pairs []pair
	for name, score := range l.scores {
		pairs = append(pairs, pair{name, score})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

	fmt.Println("+----------------+--------+")
	fmt.Println("| Player         | Score  |")
	fmt.Println("+----------------+--------+")

	for _, p := range pairs {
		fmt.Printf("| %-14s | %-6d |\n", p.name, p.score)
	}

	fmt.Println("+----------------+--------+")
}
