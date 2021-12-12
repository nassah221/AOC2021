package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Top         = 1
	TopLeft     = 2
	TopRight    = 3
	Middle      = 4
	BottomLeft  = 5
	BottomRight = 6
	Bottom      = 7
)

type digit map[int]string

type digitMap struct {
	seg digit
}

func invertMap(m map[int]string) map[string]int {
	newMap := make(map[string]int)
	for k, v := range m {
		newMap[v] = k
	}
	return newMap
}

func (d *digitMap) parseOutput(str string) string {
	m := invertMap(d.seg)
	on := make([]string, 0)

	for _, char := range str {
		on = append(on, strconv.Itoa(m[string(char)]))
	}
	sort.Strings(on)
	s := strings.Join(on, "")
	switch len(s) {
	case 2:
		return "1"
	case 3:
		return "7"
	case 4:
		return "4"
	case 5:
		if strings.EqualFold(s, "12467") {
			return "5"
		} else if strings.EqualFold(s, "13457") {
			return "2"
		} else {
			return "3"
		}
	case 6:
		if strings.EqualFold(s, "123467") {
			return "9"
		} else if strings.EqualFold(s, "124567") {
			return "6"
		} else {
			return "0"
		}
	case 7:
		return "8"
	default:
		return ""
	}
}

func (d *digitMap) parseDigitPattern(str []string) {
	sortedPatterns := make([]string, 0)
	for _, pattern := range str {
		v := strings.Split(pattern, "")
		sort.Strings(v)
		sortedPatterns = append(sortedPatterns, strings.Join(v, ""))
	}

	iter := 0

	for {
		switch iter {
		case 0:
			one, seven := sortedPatterns[iter], sortedPatterns[iter+1]
			for _, digit := range one {
				seven = strings.ReplaceAll(seven, string(digit), "")
			}
			d.seg[Top] = seven
		case 1:
			four, zeroSixNine := sortedPatterns[iter+1], sortedPatterns[6:len(sortedPatterns)-1]

			for _, digit := range zeroSixNine {
				reduceFour := digit
				for _, seg := range four {
					reduceFour = strings.ReplaceAll(reduceFour, string(seg), "")
				}
				reduced := strings.ReplaceAll(reduceFour, d.seg[Top], "")
				if len(reduced) == 1 {
					d.seg[Bottom] = reduced
					break
				}
			}
		case 2:
			four, eight := sortedPatterns[2], sortedPatterns[len(sortedPatterns)-1]

			for _, seg := range four {
				eight = strings.ReplaceAll(eight, string(seg), "")
			}
			eight = strings.ReplaceAll(eight, d.seg[Top], "")
			eight = strings.ReplaceAll(eight, d.seg[Bottom], "")

			if len(eight) == 1 {
				d.seg[BottomLeft] = eight
			}
		case 3:
			twoThreeFive := sortedPatterns[3:6]
			seven := sortedPatterns[1]
			four := sortedPatterns[2]
			one := sortedPatterns[0]
			for _, digit := range twoThreeFive {
				reduceSeven := digit
				for _, seg := range seven {
					reduceSeven = strings.ReplaceAll(reduceSeven, string(seg), "")
				}
				reducedSeven := strings.ReplaceAll(reduceSeven, d.seg[Bottom], "")
				if len(reducedSeven) == 1 {
					d.seg[Middle] = reducedSeven
					break
				}
			}
			for _, digit := range twoThreeFive {
				reduceTwo := digit
				reduceTwo = strings.ReplaceAll(reduceTwo, d.seg[Middle], "")
				reduceTwo = strings.ReplaceAll(reduceTwo, d.seg[Top], "")
				reduceTwo = strings.ReplaceAll(reduceTwo, d.seg[Bottom], "")
				reduceTwo = strings.ReplaceAll(reduceTwo, d.seg[BottomLeft], "")
				if len(reduceTwo) == 1 {
					d.seg[TopRight] = reduceTwo
					break
				}
			}
			d.seg[BottomRight] = strings.ReplaceAll(one, d.seg[TopRight], "")

			reduceFour := four
			for _, v := range d.seg {
				reduceFour = strings.ReplaceAll(reduceFour, v, "")
			}
			if len(reduceFour) == 1 {
				d.seg[TopLeft] = reduceFour
				return
			}
		}
		iter++
	}
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	patterns := []string{}
	outputs := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		str := strings.Split(line, "|")

		if str[0] != "\n" {
			patterns = append(patterns, strings.Trim(str[0], " "))
		}
		if str[1] != "\n" {
			outputs = append(outputs, strings.Trim(str[1], " "))
		}
	}

	sum := 0
	uniqueCount := 0
	for i, s := range patterns {
		sorted := make([]string, 0)
		dm := &digitMap{seg: make(digit, 7)}

		sorted = append(sorted, strings.Split(s, " ")...)
		sort.SliceStable(sorted, func(i, j int) bool {
			return len(sorted[i]) < len(sorted[j])
		})
		dm.parseDigitPattern(sorted)

		out := strings.Split(outputs[i], " ")
		numStr := ""
		for _, r := range out {
			for _, out := range strings.Split(r, " ") {
				// Part 1
				lenOut := len(out)
				if lenOut == 2 || lenOut == 3 || lenOut == 4 || lenOut == 7 {
					uniqueCount++
				}
				// Part 2
				numStr += dm.parseOutput(out)
			}
		}
		result, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		sum += result
	}
	end := time.Now()

	fmt.Printf("Part 1: %d\nPart 2: %d\nTime elapsed: %v", uniqueCount, sum, end.Sub(start))
}
