package day01

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
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
	maxCals, err := CountMaxCalories(bufio.NewReader(f), 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The answer is: %d\n", maxCals)
}

// CountMaxCalories returns sum of top leadersCount calories.
func CountMaxCalories(in io.Reader, leadersCount uint) (int, error) {
	if leadersCount == 0 {
		return 0, errors.New("leaders count can't be 0")
	}
	var elfCals []int

	scanner := bufio.NewScanner(in)

	curCals := 0
	for scanner.Scan() {
		calsRecord := strings.TrimSpace(scanner.Text())
		if calsRecord == "" {
			if curCals != 0 {
				elfCals = append(elfCals, curCals)
			}
			curCals = 0
			continue
		}

		calsRecordNumber, err := strconv.Atoi(calsRecord)
		if err != nil {
			return 0, err
		}

		curCals += calsRecordNumber
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}
	if curCals != 0 {
		elfCals = append(elfCals, curCals)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCals)))
	return sumLeaders(elfCals, leadersCount), nil
}

func sumLeaders(leaderboard []int, leadersCount uint) int {
	sum := 0
	handledLeadersCount := 0
	for i := 0; i < len(leaderboard); i++ {
		if handledLeadersCount == int(leadersCount) {
			break
		}

		sum += leaderboard[i]
		handledLeadersCount++
	}
	return sum
}
