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

type BingoSquare struct {
	Value  int
	Marked bool
}

type BingoGrid [][]BingoSquare

func (b BingoGrid) String() string {
	s := ""
	for _, x := range b {
		s += fmt.Sprintln(x)
	}
	return s
}

func (b BingoGrid) HasWon() bool {
	for i := 0; i < len(b); i++ {
		row := true
		col := true
		for j := 0; j < len(b[i]); j++ {
			if !b[i][j].Marked {
				row = false
			}
			if !b[j][i].Marked {
				col = false
			}
		}
		if row || col {
			return true
		}
	}
	return false
}

func (b BingoGrid) Mark(target int) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j].Value == target {
				b[i][j].Marked = true
			}
		}
	}
}

func (b BingoGrid) UnmarkedTotal() int {
	value := 0
	for _, row := range b {
		for _, square := range row {
			if !square.Marked {
				value += square.Value
			}
		}
	}
	return value
}

func part1(lines []string) {
	drawn := strings.Split(lines[0], ",")
	lines = lines[2:]
	grid := [][]BingoSquare{}
	grids := []BingoGrid{grid}
	for _, v := range lines {
		if v == "" {
			grids = append(grids, grid)
			grid = [][]BingoSquare{}
			continue
		}
		line := strings.Split(v, " ")
		g := []BingoSquare{}
		for _, n := range line {
			if n == "" {
				continue
			}
			number, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			g = append(g, BingoSquare{
				Value: number,
			})
		}
		grid = append(grid, g)
	}

	for _, v := range drawn {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		for _, grid := range grids {
			grid.Mark(value)
			if grid.HasWon() {
				fmt.Printf("Part 1 result: %d\n", grid.UnmarkedTotal()*value)
				return
			}

		}
	}
}

func part2(lines []string) {
	drawn := strings.Split(lines[0], ",")
	lines = lines[2:]
	grid := [][]BingoSquare{}
	grids := []BingoGrid{grid}
	for _, v := range lines {
		if v == "" {
			grids = append(grids, grid)
			grid = [][]BingoSquare{}
			continue
		}
		line := strings.Split(v, " ")
		g := []BingoSquare{}
		for _, n := range line {
			if n == "" {
				continue
			}
			number, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			g = append(g, BingoSquare{
				Value: number,
			})
		}
		grid = append(grid, g)
	}

	for _, v := range drawn {
		temp := []BingoGrid{}
		fmt.Println(len(grids))
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		for _, grid := range grids {
			grid.Mark(value)
			if !grid.HasWon() && len(grids) > 1 {
				temp = append(temp, grid)
			}
			if grid.HasWon() && len(grids) == 2 {
				fmt.Printf("Part 2 result: %d\n", grids[1].UnmarkedTotal()*value)
				return
			}

		}
		grids = temp

	}
}
