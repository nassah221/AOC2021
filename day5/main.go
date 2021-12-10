package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction uint8

const (
	Horizontal Direction = iota + 1
	Vertical
	Diagonal
)

type Point struct {
	x, y int
}

func ptFromStr(str []string) Point {
	x, _ := strconv.Atoi(str[0])
	if x > coordMax {
		coordMax = x
	}

	y, _ := strconv.Atoi(str[1])

	return Point{x, y}
}

type Line struct {
	p1  Point
	p2  Point
	dir Direction
}

var coordMax int

type Grid struct {
	grid  [][]int
	score int
}

func (d *Grid) initGrid(max int) {
	for i := 0; i < max+1; i++ {
		d.grid[i] = make([]int, max+1)
	}
}

func (d *Grid) drawLine(p1, p2 Point, dir Direction) {
	switch dir {
	case Horizontal:
		if p1.x > p2.x {
			p1, p2 = p2, p1
		}
		for i := p1.x; i <= p2.x; i++ {
			d.grid[p1.y][i]++
			if d.grid[p1.y][i] == 2 {
				d.score++
			}
		}
	case Vertical:
		if p1.y > p2.y {
			p1, p2 = p2, p1
		}
		for i := p1.y; i <= p2.y; i++ {
			d.grid[i][p1.x]++
			if d.grid[i][p1.x] == 2 {
				d.score++
			}
		}
	case Diagonal:
		posX := p1.x
		posY := p1.y
		for {
			d.grid[posY][posX]++
			if d.grid[posY][posX] == 2 {
				d.score++
			}
			switch {
			case posX > p2.x:
				posX--
			case posX < p2.x:
				posX++
			}
			switch {
			case posY > p2.y:
				posY--
			case posY < p2.y:
				posY++
			}
			if posX == p2.x && posY == p2.y {
				d.grid[posY][posX]++
				if d.grid[posY][posX] == 2 {
					d.score++
				}
				break
			}
		}
	}
}

func (l *Line) calculateDirection() {
	switch {
	case l.p1.x == l.p2.x:
		l.dir = Vertical
	case l.p1.y == l.p2.y:
		l.dir = Horizontal
	default:
		l.dir = Diagonal
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]*Line, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := &Line{}

		line := scanner.Text()
		points := strings.Split(line, " -> ")
		p1Str := strings.Split(points[0], ",")
		p2Str := strings.Split(points[1], ",")

		l.p1 = ptFromStr(p1Str)
		l.p2 = ptFromStr(p2Str)

		l.calculateDirection()

		lines = append(lines, l)
	}

	diagram := Grid{grid: make([][]int, coordMax+1)}
	diagram.initGrid(coordMax)

	for _, line := range lines {
		x1, x2 := line.p1, line.p2
		diagram.drawLine(x1, x2, line.dir)
	}
	fmt.Printf("Score: %d", diagram.score)
}
