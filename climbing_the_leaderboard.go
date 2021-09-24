package main

import (
	"fmt"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	playerRanks := []int32{}

	var currentRank int32 = 1
	rankIndex := 0

    lastScore := ranked[len(ranked)-1]
	lastScoreIndex := len(player) - 1

	for lastScoreIndex >= 0 {
		rankScore := ranked[rankIndex]

		if player[lastScoreIndex] >= rankScore {
            playerRanks = append(playerRanks, currentRank)
            lastScoreIndex--;
		}

		if rankIndex == len(ranked) - 1 {
			playerRanks = append(playerRanks, currentRank+1)
            break;
		}

		if rankScore != lastScore {
			currentRank++
		}

		rankIndex++

		lastScore = rankScore
	}

	return reverseInts(playerRanks)
}

func reverseInts(input []int32) []int32 {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func main() {
	input := []int32{100, 90, 90, 80, 75, 60}
	score := []int32{50, 65, 77, 90, 102}

	fmt.Println(climbingLeaderboard(input, score))
}
