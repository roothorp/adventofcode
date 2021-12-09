package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	lines := readFile("input.txt")
	part1(lines)
	part2(lines)
	end := time.Now()
	fmt.Printf("Took %dns\n", end.Sub(start))
}

func readFile(path string) []string {
	// open file & scanner
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Couldn't read file: %s", path)
	}
	scanner := bufio.NewScanner(file)

	// read file into (int) slice
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1(lines []string) {
	// get the line length, which is the length of our binary
	lineLength := len(lines[0])
	result := make([]int, lineLength)

	// loop through each line; add 1 for each 1, nothing for each 0
	for _, line := range lines {
		for i, v := range line {
			result[i] += int(v - '0')
		}
	}

	// we're using math.Pow, which requires float64s
	var gamma, epsilon float64

	for i, v := range result {
		// if there were more 1s than 0s, then the sum will be greater than half the size of the input
		if v > (len(lines) / 2) {
			// ...and we can essentially flip a bit by adding a power of 2
			gamma += math.Pow(2, float64(lineLength-1-i))
		} else {
			// if it was less than half, we "flip the bit" of the other number
			epsilon += math.Pow(2, float64(lineLength-1-i))
		}
	}
	// then multiply for the result!
	fmt.Printf("Part 1 result: %d\n", int(gamma*epsilon))
}

func part2(lines []string) {
	oxy := lines
	co2 := lines

	// step through each digit
	for i := range lines[0] {
		oxy = filter(oxy, i, true)
		if len(oxy) == 1 {
			break
		}
	}

	for i := range lines[0] {
		co2 = filter(co2, i, false)
		if len(co2) == 1 {
			break
		}
	}

	// we should have 1 remaining number; convert it to an int
	o, err := strconv.ParseInt(oxy[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	c, err := strconv.ParseInt(co2[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	// multiply them for our answer
	fmt.Printf("Part 2 result: %d\n", o*c)
}

func filter(lines []string, pos int, most bool) []string {
	var zero, one []string
	// sort the lines by 0s or 1s in the current position
	for _, line := range lines {
		if []rune(line)[pos] == '1' {
			one = append(one, line)
		} else {
			zero = append(zero, line)
		}
	}
	// Decide if we want the most or least common

	if most {
		if len(zero) == len(one) || len(zero) < len(one) {
			return one
		} else {
			return zero
		}
	} else {
		if len(zero) == len(one) || len(zero) < len(one) {
			return zero
		} else {
			return one
		}
	}
}
