package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	newBornTimer = 8
	Timer        = 6
)

var (
	numDays = flag.Int("days", 80, "Number of days to run the simlation for")
)

func main() {
	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	d := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), ",")
		for _, v := range str {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			if _, ok := d[n]; ok {
				d[n]++
			} else {
				d[n] = 1
			}
		}
	}

	for day := 1; day <= *numDays; day++ {
		opMap := make(map[int]int)
		for k, v := range d {
			if k > 0 {
				opMap[k-1] += v
				delete(d, k)
				continue
			}
			if k == 0 {
				opMap[newBornTimer] += v
				opMap[Timer] += v
				delete(d, k)
				continue
			}
		}
		d = opMap
	}

	sum := 0
	for _, v := range d {
		sum += v
	}
	fmt.Printf("Number of fish:%d", sum)
}
