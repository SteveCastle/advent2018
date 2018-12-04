package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
func splitFunction(r rune) bool {
	return r == '@' || r == ',' || r == ':' || r == 'x'
}
func parseLine(s string) claim {
	parts := strings.FieldsFunc(s, splitFunction)
	left, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	top, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	width, err := strconv.Atoi(strings.TrimSpace(parts[3]))
	height, err := strconv.Atoi(strings.TrimSpace(parts[4]))
	if err != nil {
		fmt.Println("Error parsing a claim.")
	}
	ret := claim{
		Number: strings.TrimSpace(parts[0]),
		Top:    top,
		Left:   left,
		Width:  width,
		Height: height,
	}
	return ret
}

// Given a claim, map claim to every space in grid that it occupies.
func assignClaim(c claim, layout *[1001][1001][]claim, overlaps map[string]bool) {
	for i := c.Left; i <= c.Width+c.Left-1; i++ {
		for j := c.Top; j <= c.Height+c.Top-1; j++ {
			layout[i][j] = append(layout[i][j], c)
			if len(layout[i][j]) > 1 {
				for _, key := range layout[i][j] {
					overlaps[key.Number] = true
				}
			}
		}
	}
}

// Iterate through array and count overlapping spaces.
func readDuplicates(layout *[1001][1001][]claim) int {
	duplicates := 0
	for i := 0; i <= 1000; i++ {
		for j := 0; j <= 1000; j++ {
			// fmt.Printf("Row %d and Column %d.\n", i, j)
			if len(layout[i][j]) > 1 {
				// fmt.Printf("Duplicate found on Row %d and Column %d.\n", i, j)
				// fmt.Println(layout[i][j])
				duplicates++
			}
		}
	}
	return duplicates

}

func main() {
	// Read File Into Array of Strings
	parcels, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	layout := [1001][1001][]claim{}
	overlaps := map[string]bool{}

	// Iterate through parcels and map to 2d Array.
	for _, parcel := range parcels {
		claim := parseLine(parcel)
		assignClaim(claim, &layout, overlaps)
	}

	// Print the number of lines with overlaping claims.
	fmt.Println(readDuplicates(&layout))

	// Iterate through all parcels to find one that has not been marked a duplicate ever.
	for _, parcel := range parcels {
		claim := parseLine(parcel)
		_, ok := overlaps[claim.Number]
		if !ok {
			fmt.Printf("Key not found value is: %s, \n", claim.Number)
		}
	}

}

type claim struct {
	Number string
	Top    int
	Left   int
	Width  int
	Height int
}
