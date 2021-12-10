package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type draw []int
type Numbers struct {
	draw
}

type unit map[int]bool
type arr [5][5]unit

type All struct {
	boards     []*Board
	firstToWin int
	lastToWin  int
}

type Board struct {
	id int
	arr
	winningSum  int
	winningDraw int
	won         bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers Numbers
	b := &Board{}
	var boards All

	id := 0
	entryCount := 0
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsAny(",", line) {
			inputStr := strings.Split(line, ",")
			for _, num := range inputStr {
				if num, err := strconv.Atoi(num); err == nil {
					numbers.draw = append(numbers.draw, num)
				}
			}
		} else if strings.ContainsAny(" ", line) {
			boardStr := strings.Split(line, " ")
			j := 0
			for _, num := range boardStr {
				if num == "" {
					continue
				}
				num, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				b.arr[i][j] = map[int]bool{num: false}
				j++
				entryCount++
			}
			i++
			if entryCount%25 == 0 {
				id++
				b.id = id
				boards.boards = append(boards.boards, b)
				i = 0
				newBoard := &Board{}
				b = newBoard
			}
		}
	}
	for _, draw := range numbers.draw {
		for _, b := range boards.boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					b.markUnit(i, j, draw)
					if j == 4 {
						if marked := b.isRowMarked(i); marked {
							if b.won {
								continue
							}
							b.winningDraw = draw
							b.winningSum = b.calculateSum()

							boards.lastToWin = b.id

							b.won = true
							if boards.firstToWin == 0 {
								boards.firstToWin = b.id
							}
						}
					}
				}
				if i == 4 {
					if marked := b.isColMarked(); marked {
						if b.won {
							continue
						}
						b.winningDraw = draw
						b.winningSum = b.calculateSum()
						boards.lastToWin = b.id

						b.won = true
						if boards.firstToWin == 0 {
							boards.firstToWin = b.id
						}
					}
				}
			}
		}
	}

	for _, b := range boards.boards {
		if b.id == boards.firstToWin {
			fmt.Printf("Board %d first to win. Score %d\n", b.id, b.winningDraw*b.winningSum)
		}
		if b.id == boards.lastToWin {
			fmt.Printf("Board %d last to win. Score %d\n", b.id, b.winningDraw*b.winningSum)
		}
	}
}

func (b *Board) isRowMarked(row int) bool {
	var markedCount int
	for col := 0; col < 5; col++ {
		for _, v := range b.arr[row][col] {
			if v {
				markedCount++
			}
		}
	}

	return markedCount == 5
}

func (b *Board) isColMarked() bool {
	for row := 0; row < 5; row++ {
		var markedCount int
		for col := 0; col < 5; col++ {
			for _, v := range b.arr[col][row] {
				if v {
					markedCount++
				}
			}
		}
		if markedCount == 5 {
			return true
		}
	}
	return false
}

func (b *Board) calculateSum() (sum int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k, v := range b.arr[i][j] {
				if !v {
					sum += k
				}
			}
		}
	}
	return
}

func (b *Board) markUnit(row, col, entry int) {
	if _, ok := b.arr[row][col][entry]; ok {
		b.arr[row][col][entry] = true
	}
}
