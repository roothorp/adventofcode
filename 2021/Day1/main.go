package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	part1()
	part2()
}

func readFile(path string) []int {
	// open file & scanner
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Couldn't read file: %s", path)
	}
	scanner := bufio.NewScanner(file)

	// read file into slice
	var lines []int
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf(err.Error())
		}
		lines = append(lines, line)
	}
	return lines
}

func part1() {
	// read file
	lines := readFile("input.txt")

	// count occurances of i being bigger than i-1
	count := 0
	start := time.Now()
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			count++
		}
	}
	end := time.Now()
	fmt.Printf("Part 1 time elapsed: %d\n", end.Sub(start))

	fmt.Printf("Part 1 result: %d\n", count)
}

func part2() {
	// read file
	lines := readFile("input.txt")

	//make another slice with (i + i-1 + i-2)
	start := time.Now()
	sum := make([]int, len(lines)-2)
	for i := 2; i < len(lines); i++ {
		sum[i-2] = lines[i] + lines[i-1] + lines[i-2]
	}

	// count occurances of i being bigger than i-1
	count := 0
	for i := 1; i < len(sum); i++ {
		if sum[i] > sum[i-1] {
			count++
		}
	}
	end := time.Now()
	fmt.Printf("Part 2 time elapsed: %d\n", end.Sub(start))

	fmt.Printf("Part 2 result: %d\n", count)

}
