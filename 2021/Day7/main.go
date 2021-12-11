package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	pos := readFile("input.txt")
	part1(pos)
	part2(pos)
	end := time.Now()
	fmt.Printf("Took %dns\n", end.Sub(start))
}

func readFile(path string) []int {
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
	temp := strings.Split(lines[0], ",")
	var positions []int
	for _, v := range temp {
		v, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, v)
	}
	return positions
}

func mean(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return int(math.Floor(float64(total) / float64(len(nums))))
}

func median(nums []int) int {
	sort.Ints(nums)

	mid := len(nums) / 2

	if len(nums)%2 != 0 {
		return nums[mid]
	}

	return (nums[mid-1] + nums[mid]) / 2
}

func part1(positions []int) {
	target := median(positions)
	fmt.Printf("Target is %d\n", target)

	fuel := 0

	for _, v := range positions {
		if v < target {
			fuel += target - v
		} else {
			fuel += v - target
		}
	}

	fmt.Printf("Part 1 result: %d\n", fuel)

}

func part2(positions []int) {
	target := mean(positions)
	fuel := 0
	fmt.Printf("Target is %d\n", target)
	for _, v := range positions {
		n := 0
		if v < target {
			n = target - v
		} else {
			n = v - target
		}
		fuel += (n * (n + 1)) / 2
	}
	fmt.Printf("Part 2 result: %d\n", fuel)
}
