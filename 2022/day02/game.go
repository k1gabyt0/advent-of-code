package day02

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"strings"
)

//go:embed input.txt
var input embed.FS

func SolvePuzzle() {
	f, err := input.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("--- Day 2: Rock Paper Scissors ---")
	totalScore, err := CalculateTotalScore(bufio.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The answer is: %d\n", totalScore)
}

type HandShape string

const (
	Rock     HandShape = "Rock"
	Paper    HandShape = "Paper"
	Scissors HandShape = "Scissors"
)

type GameResult string

const (
	Win  GameResult = "Win"
	Lost GameResult = "Lost"
	Draw GameResult = "Draw"
)

var playerFormToEnemyFormResults = map[HandShape]map[HandShape]GameResult{
	Rock: {
		Rock:     Draw,
		Paper:    Lost,
		Scissors: Win,
	},
	Paper: {
		Rock:     Win,
		Paper:    Draw,
		Scissors: Lost,
	},
	Scissors: {
		Rock:     Lost,
		Paper:    Win,
		Scissors: Draw,
	},
}

var shapePoints = map[HandShape]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var gamePoints = map[GameResult]int{
	Lost: 0,
	Draw: 3,
	Win:  6,
}

var dictionary = map[rune]HandShape{
	// enemy
	'A': Rock,
	'B': Paper,
	'C': Scissors,
	// player
	'Y': Paper,
	'X': Rock,
	'Z': Scissors,
}

func CalculateTotalScore(in io.Reader) (int, error) {
	scanner := bufio.NewScanner(in)

	totalScore := 0
	row := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		row++
		if line == "" {
			continue
		}

		var enemyMove, playerMove rune
		_, err := fmt.Sscanf(line, "%c %c", &enemyMove, &playerMove)
		if err != nil {
			return 0, err
		}

		enemyForm := dictionary[enemyMove]
		playerForm := dictionary[playerMove]
		gameResultForPlayer := playerFormToEnemyFormResults[playerForm][enemyForm]

		pointsForRound := gamePoints[gameResultForPlayer] + shapePoints[playerForm]
		totalScore += pointsForRound

		// debug
		// fmt.Printf("%d. %c%c %s-%s %s. Player got=%d points(game result=%d + shape=%d) for round. Total: %d\n",
		// 	row, enemyMove, playerMove, enemyForm, playerForm, gameResultForPlayer,
		// 	pointsForRound, gamePoints[gameResultForPlayer], shapePoints[playerForm],
		// 	totalScore,
		// )
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	return totalScore, nil
}
