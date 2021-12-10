package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := make([]int, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		data = append(data, num)
	}

	result := 0

	window := []int{2: 0}
	prevSum := 0

	for i := range data {
		third := i + 2
		if third == len(data) {
			break
		}

		window[0], window[1], window[2] = data[i], data[third-1], data[third]

		curSum := 0
		for _, v := range window {
			curSum += v
		}

		if curSum > prevSum {
			if prevSum != 0 {
				result++
			}
		}

		prevSum = curSum
	}
	fmt.Println(result)
}
