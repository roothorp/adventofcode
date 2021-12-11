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
		file.Close()
		log.Fatalf("Couldn't read file: %s", path)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// read file into (int) slice
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func cycle(fish []int) []int {
	new := fish[0]
	for i := 0; i < len(fish)-1; i++ {
		fish[i] = fish[i+1]
	}
	fish[6] += new
	fish[len(fish)-1] = new

	return fish
}

func total(fish []int) int {
	count := 0
	for _, v := range fish {
		count += v
	}
	return count
}

func lanternfish(lines []string, duration int) int {
	line := strings.Split(lines[0], ",")
	fish := make([]int, 9)
	for _, v := range line {
		v, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		fish[v]++
	}

	days := 0
	for days < duration {
		fish = cycle(fish)
		days++
	}

	return total(fish)
}

func part1(lines []string) {
	fmt.Printf("Part 1 result: %d\n", lanternfish(lines, 80))
}

func part2(lines []string) {
	fmt.Printf("Part 1 result: %d\n", lanternfish(lines, 256))
}
