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
	sumOfPriorities, err := ReorganizeRucksack(bufio.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The answer is: %d\n", sumOfPriorities)
}

func ReorganizeRucksack(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)

	var sum int
	for scanner.Scan() {
		line := []rune(strings.TrimSpace(scanner.Text()))
		if len(line) == 0 {
			continue
		}
		if len(line)%2 != 0 {
			return 0, errors.New("rucksack doesn't contain even quantity of items")
		}

		mid := len(line) / 2
		firstCompartment, secondCompatment := line[0:mid], line[mid:]
		commonItems := findCommonItems(firstCompartment, secondCompatment)

		sum += calcSum(commonItems)
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	return sum, nil
}

func findCommonItems[T comparable](left, right []T) []T {
	res := make(map[T]struct{})

	set := make(map[T]struct{})
	for _, lItem := range left {
		set[lItem] = struct{}{}
	}
	for _, rItem := range right {
		if _, ok := set[rItem]; ok {
			res[rItem] = struct{}{}
		}
	}
	return toSlice(res)
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
