package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateFuel(d map[int]int, from int, constRate bool) int {
	fuel := 0
	for k, v := range d {
		if k == from {
			continue
		}
		if constRate {
			fuel += v * int(math.Abs(float64(k-from)))
			continue
		}
		n := int(math.Abs(float64(k - from)))
		fuel += v * ((n * (1 + n)) / 2)
	}
	return fuel
}

func findMinFuel(d map[int]int, max int) int {
	i := 0
	lastFuel := 0
	for {
		if i == max {
			return 0
		}

		fuel := calculateFuel(d, i, false)

		switch {
		case lastFuel == 0:
			lastFuel = fuel
			i++
		case fuel < lastFuel:
			lastFuel = fuel
			i++
		case fuel > lastFuel:
			return lastFuel
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	d := map[int]int{}
	var max int
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		for _, numStr := range s {
			num, err := strconv.Atoi(numStr)
			if num > max {
				max = num
			}
			if err != nil {
				panic(err)
			}
			d[num]++
		}
	}

	fmt.Println(findMinFuel(d, max))
}
