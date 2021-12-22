package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	//Sample: p1 = 4, p2 = 8
	//Real: p1 = 4, p2 = 1
	playerOneStart := 4
	playerTwoStart := 1

	loserScore, dieRolls := playDeterministicGame(playerOneStart, playerTwoStart)
	fmt.Printf("Losers Score: %d, Die Rolls: %d, Multiplied: %d \n", loserScore, dieRolls, loserScore*dieRolls)
	playerOneUniversesWon, playerTwoUniversesWon := startGame(playerOneStart, playerTwoStart)
	fmt.Printf("Player One Wins: %d, Player Two Wins: %d \n", playerOneUniversesWon, playerTwoUniversesWon)

	duration := time.Since(start)
	fmt.Printf("Time Since Start: %d \n", duration)
}

func playDeterministicGame(playerOneStart, playerTwoStart int) (int, int) {
	var dieRolls int
	var playerOneScore int
	var playerTwoScore int
	startRoll := -3
	gameOver := false
	for !gameOver {
		dieRolls += 3
		totalDieRoll := 3*dieRolls + startRoll
		playerOneStart += totalDieRoll
		for playerOneStart > 10 {
			playerOneStart -= 10
		}
		playerOneScore += playerOneStart
		if playerOneScore >= 1000 {
			return playerTwoScore, dieRolls
		}
		dieRolls += 3
		totalDieRoll = 3*dieRolls + startRoll
		playerTwoStart += totalDieRoll
		for playerTwoStart > 10 {
			playerTwoStart -= 10
		}
		playerTwoScore += playerTwoStart
		if playerTwoScore >= 1000 {
			return playerOneScore, dieRolls
		}
	}
	return 0, dieRolls
}

func startGame(p1Position, p2Position int) (int, int) {
	previousUniverses := make(map[Universe][]int)
	universe := Universe{P1Pos: p1Position, P2Pos: p2Position, P1Score: 0, P2Score: 0, PlayerTurn: 1}
	p1Score, p2Score, _ := playDirac(universe, previousUniverses)
	return p1Score, p2Score
}

type Universe struct {
	P1Pos      int
	P2Pos      int
	P1Score    int
	P2Score    int
	PlayerTurn int
}

func playDirac(universe Universe, previousUniverses map[Universe][]int) (int, int, map[Universe][]int) {
	if scores, ok := previousUniverses[universe]; ok {
		return scores[0], scores[1], previousUniverses
	}
	var totalP1Wins, totalP2Wins, p1Wins, p2Wins int
	totalRoll := []int{3, 4, 5, 6, 7, 8, 9}
	multiplier := []int{1, 3, 6, 7, 6, 3, 1}
	for i := 0; i < 7; i++ {
		currentUniverse := Universe{P1Pos: universe.P1Pos, P2Pos: universe.P2Pos, P1Score: universe.P1Score, P2Score: universe.P2Score, PlayerTurn: universe.PlayerTurn}
		if currentUniverse.PlayerTurn == 1 {
			if currentUniverse.P1Pos+totalRoll[i] > 10 {
				currentUniverse.P1Pos = currentUniverse.P1Pos + totalRoll[i] - 10
			} else {
				currentUniverse.P1Pos = currentUniverse.P1Pos + totalRoll[i]
			}
			if currentUniverse.P1Score+currentUniverse.P1Pos >= 21 {
				totalP1Wins += multiplier[i]
			} else {
				currentUniverse.P1Score += currentUniverse.P1Pos
				currentUniverse.PlayerTurn = 2
				p1Wins, p2Wins, previousUniverses = playDirac(currentUniverse, previousUniverses)
				totalP1Wins += p1Wins * multiplier[i]
				totalP2Wins += p2Wins * multiplier[i]
			}
		} else {
			if currentUniverse.P2Pos+totalRoll[i] > 10 {
				currentUniverse.P2Pos = currentUniverse.P2Pos + totalRoll[i] - 10
			} else {
				currentUniverse.P2Pos = currentUniverse.P2Pos + totalRoll[i]
			}
			if currentUniverse.P2Score+currentUniverse.P2Pos >= 21 {
				totalP2Wins += multiplier[i]
			} else {
				currentUniverse.P2Score += currentUniverse.P2Pos
				currentUniverse.PlayerTurn = 1
				p1Wins, p2Wins, previousUniverses = playDirac(currentUniverse, previousUniverses)
				totalP1Wins += p1Wins * multiplier[i]
				totalP2Wins += p2Wins * multiplier[i]
			}
		}
	}
	previousUniverses[universe] = []int{totalP1Wins, totalP2Wins}
	return totalP1Wins, totalP2Wins, previousUniverses
}
