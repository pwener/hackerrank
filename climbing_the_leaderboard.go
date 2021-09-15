package main

import "fmt"

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

    for _, score := range player {
        copyArr := ranked
        copyArr = append(copyArr, score)
        // insert new item
        for i, value := range copyArr {
            if value < score {
                temp := copyArr[i:]
                copyArr := copyArr[:i]
                copyArr = append(copyArr, score)
                copyArr = append(copyArr, temp...)
                break
            }
        }
        playerRanks = append(playerRanks, findRank(copyArr, score))
    }

    return playerRanks
}

func findRank(ranked []int32, score int32) int32 {
    ranks := getRanks(ranked)

    rankMap := make(map[int32]int32, len(ranked) + 1)

    for i, v := range ranks {
        rankMap[ranked[i]] = v
    }

    return rankMap[score]
}

func getRanks(ranked []int32) []int32 {
    var ranks = []int32{}

    for i, v := range ranked {
        if i == 0 {
            ranks = append(ranks, 1)
            continue
        }

        index := i - 1
        lastRank := ranks[index]
        if (v == ranked[index]) {
            ranks = append(ranks, lastRank)
            continue
        }

        ranks = append(ranks, lastRank + 1)
    }
    return ranks
}

func main() {
    input := []int32{100,100,50,40,40,20,10}
    score := []int32{5, 25, 50, 120}

    fmt.Println(climbingLeaderboard(input, score))
}