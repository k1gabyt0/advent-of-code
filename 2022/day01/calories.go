package day01

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"strconv"
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

	fmt.Println("--- Day 1: Calorie Counting ---")
	fmt.Printf("The answer is: %d", CountMaxCalories(bufio.NewReader(f)))
}

func CountMaxCalories(in io.Reader) int {
	scanner := bufio.NewScanner(in)

	maxCals := 0
	curCals := 0
	for scanner.Scan() {
		calsRecord := strings.TrimSpace(scanner.Text())
		if calsRecord == "" {
			if curCals > maxCals {
				maxCals = curCals

			}
			curCals = 0
			continue
		}

		curCals += must(strconv.Atoi(calsRecord))
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return maxCals
}

func must[T any](arg T, err error) T {
	if err != nil {
		panic(err)
	}
	return arg
}
