package day03

import (
	"bufio"
	"embed"
	"errors"
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

	fmt.Println("--- Day 3: Rucksack Reorganization ---")
	sumOfPriorities, err := ReorganizeRucksack(bufio.NewReader(f), 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The answer is: %d\n", sumOfPriorities)
}

func scanNonEmptyNLines(n int) bufio.SplitFunc {
	if n <= 0 {
		panic("scanning 0 or less lines makes no sense")
	}
	if n == 1 {
		return bufio.ScanLines
	}

	return func(data []byte, atEOF bool) (int, []byte, error) {
		var token []byte

		advance := 0
		for i := 0; i < n; {
			a, t, err := bufio.ScanLines(data[advance:], atEOF)
			if err != nil {
				return 0, nil, err
			}
			// if nothing to read then cancel
			if a == 0 && len(t) == 0 {
				return 0, nil, nil
			}

			advance += a
			if len(t) == 0 {
				continue
			}

			token = append(token, '\n')
			token = append(token, t...)
			i++
		}
		return advance, token, nil
	}
}

func ReorganizeRucksack(reader io.Reader, groupSize int) (int, error) {
	if groupSize <= 0 {
		return 0, errors.New("groupSize can't be <= 0")
	}
	scanner := bufio.NewScanner(reader)
	scanner.Split(scanNonEmptyNLines(groupSize))

	var sum int
	for scanner.Scan() {
		rucksackGroup := strings.TrimSpace(scanner.Text())
		if rucksackGroup == "" {
			continue
		}
		rucksacks := strings.Split(rucksackGroup, "\n")
		var rucksackItems [][]rune
		for _, r := range rucksacks {
			rucksackItems = append(rucksackItems, []rune(strings.TrimSpace(r)))
		}
		commonItems := findCommonItems(rucksackItems)

		sum += calcSum(commonItems)
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	return sum, nil
}

func findCommonItems[T comparable](targets [][]T) []T {
	if len(targets) == 0 {
		return nil
	}
	if len(targets) == 1 {
		return targets[0]
	}

	res := make(map[T]struct{})

	set := toSet(targets[0])
	for _, rItem := range targets[1] {
		if _, ok := set[rItem]; ok {
			res[rItem] = struct{}{}
		}
	}
	if len(targets) == 2 {
		return toSlice(res)
	}
	for _, target := range targets[2:] {
		set := toSet(target)
		for item := range res {
			if _, isInSet := set[item]; !isInSet {
				delete(res, item)
			}
		}
	}

	return toSlice(res)
}

func toSet[T comparable](slice []T) map[T]struct{} {
	if len(slice) == 0 {
		return nil
	}
	set := make(map[T]struct{})
	for _, item := range slice {
		set[item] = struct{}{}
	}
	return set
}

var items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var priorityMap = getPriorityMap(items)

func getPriorityMap(items string) map[rune]int {
	priorities := make(map[rune]int)
	for idx, item := range items {
		priorities[item] = idx + 1
	}
	return priorities
}

func calcSum(items []rune) int {
	var sum int
	for _, item := range items {
		sum += priorityMap[item]
	}
	return sum
}

func toSlice[K comparable, V any](m map[K]V) []K {
	if len(m) == 0 {
		return nil
	}
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}
