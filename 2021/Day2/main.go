package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	// counters for the x and y of the "submarine"
	horizontal := 0
	depth := 0

	// split each line into [direction, distance]
	for _, v := range lines {
		line := strings.Split(v, " ")
		if len(line) < 2 {
			log.Fatalf("Couldn't split line: %s", line)
		}
		direction := line[0]
		distance, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("Couldn't case string to int: %s", line[1])
		}

		//appropriately modify counters
		switch direction {
		case "up":
			depth -= distance
		case "down":
			depth += distance
		case "forward":
			horizontal += distance
		}
	}

	fmt.Printf("Part 1 result: %d\n", horizontal*depth)
}

func part2(lines []string) {
	// counters for the x and y of the "submarine", and it's course/aim
	horizontal := 0
	depth := 0
	aim := 0

	// split each line into [direction, distance]
	for _, v := range lines {
		line := strings.Split(v, " ")
		if len(line) < 2 {
			log.Fatalf("Couldn't split line: %s", line)
		}
		direction := line[0]
		distance, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("Couldn't case string to int: %s", line[1])
		}

		//appropriately modify counters
		switch direction {
		case "up":
			aim -= distance
		case "down":
			aim += distance
		case "forward":
			horizontal += distance
			depth += aim * distance
		}
	}

	fmt.Printf("Part 2 result: %d\n", horizontal*depth)
}
