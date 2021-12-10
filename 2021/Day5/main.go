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

type Coords struct {
	X int
	Y int
}

func NewCoords(x, y int) Coords {
	return Coords{
		X: x,
		Y: y,
	}
}

type Vents struct {
	Start, End Coords
}

func NewVents(startX, startY, endX, endY int) Vents {
	return Vents{
		Start: Coords{
			X: startX,
			Y: startY,
		},
		End: Coords{
			X: endX,
			Y: endY,
		},
	}
}

func (v Vents) GetLine() []Coords {
	line := []Coords{}
	if v.Start.X == v.End.X {
		x := v.Start.X
		if v.Start.Y < v.End.Y {
			for y := v.Start.Y; y <= v.End.Y; y++ {
				line = append(line, NewCoords(x, y))
			}
		}
		if v.Start.Y > v.End.Y {
			for y := v.Start.Y; y >= v.End.Y; y-- {
				line = append(line, NewCoords(x, y))
			}
		}
	}

	if v.Start.Y == v.End.Y {
		y := v.Start.Y
		if v.Start.X < v.End.X {
			for x := v.Start.X; x <= v.End.X; x++ {
				line = append(line, NewCoords(x, y))
			}
		}
		if v.Start.X > v.End.X {
			for x := v.Start.X; x >= v.End.X; x-- {
				line = append(line, NewCoords(x, y))
			}
		}
	}

	return line
}

func AbsEqual(num1, num2 int) bool {
	if num1 < 0 {
		num1 = num1 * -1
	}
	if num2 < 0 {
		num2 = num2 * -1
	}
	return num1 == num2
}

func (v Vents) GetDiagonal() []Coords {
	line := []Coords{}
	xDiff := v.End.X - v.Start.X
	yDiff := v.End.Y - v.Start.Y

	if AbsEqual(xDiff, yDiff) {
		y := v.Start.Y
		if xDiff > 0 {
			if yDiff > 0 {
				for x := v.Start.X; x <= v.End.X; x++ {
					line = append(line, NewCoords(x, y))
					y++
				}
			}
			if yDiff < 0 {
				for x := v.Start.X; x <= v.End.X; x++ {
					line = append(line, NewCoords(x, y))
					y--
				}
			}
		}
		if xDiff < 0 {
			if yDiff > 0 {
				for x := v.Start.X; x >= v.End.X; x-- {
					line = append(line, NewCoords(x, y))
					y++
				}
			}
			if yDiff < 0 {
				for x := v.Start.X; x >= v.End.X; x-- {
					line = append(line, NewCoords(x, y))
					y--
				}
			}
		}
	}
	return line
}

type Diagram map[Coords]int

func (d Diagram) Mark(coords []Coords) {
	for _, coord := range coords {
		d[coord]++
	}
}

func (d Diagram) CountOverlaps() int {
	count := 0
	for _, v := range d {
		if v > 1 {
			count++
		}
	}
	return count
}

func MarkVents(lines []string, diagonals bool) int {
	diagram := Diagram{}

	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")

		startX, err := strconv.Atoi(start[0])
		if err != nil {
			log.Fatal(err)
		}
		startY, err := strconv.Atoi(start[1])
		if err != nil {
			log.Fatal(err)
		}
		endX, err := strconv.Atoi(end[0])
		if err != nil {
			log.Fatal(err)
		}
		endY, err := strconv.Atoi(end[1])
		if err != nil {
			log.Fatal(err)
		}

		vent := NewVents(startX, startY, endX, endY)
		diagram.Mark(vent.GetLine())
		if diagonals {
			diagram.Mark(vent.GetDiagonal())
		}
	}
	return diagram.CountOverlaps()
}

func part1(lines []string) {
	fmt.Printf("Part 1 result: %d\n", MarkVents(lines, false))
}

func part2(lines []string) {
	fmt.Printf("Part 2 result: %d\n", MarkVents(lines, true))
}
