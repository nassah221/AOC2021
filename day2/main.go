package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	Direction string
	Distance  int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := make([]Movement, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		num, err := strconv.Atoi(input[1])
		if err != nil {
			panic(err)
		}

		data = append(data, Movement{input[0], num})
	}

	horizontalSum := 0
	verticalSum := 0
	aim := 0

	for _, m := range data {
		switch m.Direction {
		case "forward":
			horizontalSum += m.Distance
			verticalSum += m.Distance * aim
		case "up":
			aim += m.Distance
		case "down":
			aim -= m.Distance
		}
	}
	fmt.Println(horizontalSum, verticalSum, horizontalSum*verticalSum)
}
