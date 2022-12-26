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

type handShape string

const (
	rock     handShape = "Rock"
	paper    handShape = "Paper"
	scissors handShape = "Scissors"
)

type gameResult string

const (
	win  gameResult = "Win"
	lost gameResult = "Lost"
	draw gameResult = "Draw"
)

var enemyMoves = map[rune]handShape{
	'A': rock,
	'B': paper,
	'C': scissors,
}

var playerStrategies = map[rune]gameResult{
	'X': lost,
	'Y': draw,
	'Z': win,
}

var enemyMoveToGameResult = map[handShape]map[gameResult]handShape{
	rock: {
		draw: rock,
		win:  paper,
		lost: scissors,
	},
	paper: {
		draw: paper,
		win:  scissors,
		lost: rock,
	},
	scissors: {
		draw: scissors,
		win:  rock,
		lost: paper,
	},
}

var shapePoints = map[handShape]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var gamePoints = map[gameResult]int{
	lost: 0,
	draw: 3,
	win:  6,
}

func CalculateTotalScore(in io.Reader) (int, error) {
	scanner := bufio.NewScanner(in)

	totalScore := 0
	// row := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// row++
		if line == "" {
			continue
		}

		var enemyRecord, playerRecord rune
		_, err := fmt.Sscanf(line, "%c %c", &enemyRecord, &playerRecord)
		if err != nil {
			return 0, err
		}

		enemyMove := enemyMoves[enemyRecord]
		gameResultPlayerWant := playerStrategies[playerRecord]
		playerMove := enemyMoveToGameResult[enemyMove][gameResultPlayerWant]

		pointsForRound := gamePoints[gameResultPlayerWant] + shapePoints[playerMove]
		totalScore += pointsForRound

		// debug
		// fmt.Printf("%d. %c%c %s-%s %s. Player got=%d points(game result=%d + shape=%d) for round. Total: %d\n",
		// 	row, enemyRecord, playerRecord, enemyMove, playerMove, gameResultPlayerWant,
		// 	pointsForRound, gamePoints[gameResultPlayerWant], shapePoints[playerMove],
		// 	totalScore,
		// )
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	return totalScore, nil
}
