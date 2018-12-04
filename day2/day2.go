package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to return a slice of strings from lines in a file.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		if err != nil {
			return nil, err
		}
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

// Iterate through the characters in a string.
// Keep track of charMap state to track current reps of 2 or 3.
func countChars(s string) (bool, bool) {
	charMap := map[rune]int{}
	hasTwo := 0
	hasThree := 0
	for _, rune := range s {
		charMap[rune] = charMap[rune] + 1
		if charMap[rune] == 2 {
			hasTwo++
		} else if charMap[rune] == 3 {
			hasTwo--
			hasThree++
		} else if charMap[rune] == 4 {
			hasThree--
		}
	}
	return hasTwo > 0, hasThree > 0
}

// Count mismatches between two strings.
func countDiff(s string, s2 string) int {
	diff := 0
	for i := range s {
		if s[i] != s2[i] {
			diff++
		}
	}
	return diff
}

// Create a map of strings, where the key is the first or second half of the string.
func partition(list []string, prefix bool) map[string][]string {
	stringMap := map[string][]string{}
	for _, s := range list {
		idx := ""
		if prefix {
			idx = s[:len(s)/2]
		} else {
			idx = s[len(s)/2:]
		}
		stringMap[idx] = append(stringMap[idx], s)
	}
	return stringMap
}

func main() {
	ids, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Find count of strings with two repetitions or 3 repetitions.
	// Multiply to get answer for problem 1.
	twos := 0
	threes := 0
	for _, id := range ids {
		// Check if this id has twos or threes instances.
		hasTwo, hasThree := countChars(id)
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}
	fmt.Printf("%d times %d is: %d \n", twos, threes, twos*threes)

	// Find strings with one difference for problem 2.
	frontPartition := partition(ids, true)
	endPartition := partition(ids, false)
	results := []string{}

	// Group strings with same first half and find matches.
	for _, v := range frontPartition {
		for _, id := range v {
			for _, id2 := range v {
				if countDiff(id, id2) == 1 {
					results = append(results, id)
				}
			}
		}
	}

	// Group strings with same second half and find matches.
	for _, v := range endPartition {
		for _, id := range v {
			for _, id2 := range v {
				if countDiff(id, id2) == 1 {
					results = append(results, id)
				}
			}
		}
	}
	fmt.Println(results)
}
