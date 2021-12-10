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

	scanner := bufio.NewScanner(f)

	data := make([]string, 0)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	powerConsumption(data)

	lifeSupportRating(data)
}

// Part 1
func powerConsumption(data []string) {
	var gamma string

	iter := len(data[0])
	for i := 0; i < iter; i++ {
		commonBit := countBits(data, i)
		switch commonBit {
		case "1":
			gamma += "1"
		case "0":
			gamma += "0"
		}
	}

	gammaRating, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilon := binaryCompliment(gamma)
	epsilonRating, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Power Consumption:%d\n", gammaRating*epsilonRating)
}

// Part 2
func lifeSupportRating(data []string) {
	o2Rating := o2Generator(data)
	co2Rating := co2Scrubber(data)
	fmt.Printf("Life Support:%d\n", o2Rating*co2Rating)
}

func o2Generator(data []string) (rating int64) {
	o2Data := make([]string, len(data))

	copy(o2Data, data)

	o2Gen := struct{ data []string }{data: o2Data}

	for i := 0; i < len(o2Data[0]); i++ {
		commonBit := countBits(o2Gen.data, i)

		if len(o2Gen.data) > 1 {
			o2Gen.data = filterData(o2Gen.data, commonBit, i)
		}
	}
	oxygen := o2Gen.data[0]

	rating, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		panic(err)
	}
	return
}

func co2Scrubber(data []string) (rating int64) {
	co2Data := make([]string, len(data))
	copy(co2Data, data)

	co2Scrub := struct{ data []string }{data: co2Data}
	for i := 0; i < len(co2Data[0]); i++ {
		uncommonBit := binaryCompliment(countBits(co2Scrub.data, i))
		if len(co2Scrub.data) > 1 {
			co2Scrub.data = filterData(co2Scrub.data, uncommonBit, i)
		}
	}
	co2 := co2Scrub.data[0]

	rating, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		panic(err)
	}
	return
}

func binaryCompliment(str string) (result string) {
	for char := range str {
		if str[char] == '0' {
			result += "1"
		} else {
			result += "0"
		}
	}
	return
}

func countBits(data []string, pos int) string {
	ones := 0
	zeros := 0

	for _, v := range data {
		if string(v[pos]) == "1" {
			ones++
			continue
		}
		zeros++
	}
	if ones >= zeros {
		return "1"
	}
	return "0"
}

func filterData(data []string, bitOfInterest string, pos int) []string {
	toFilter := make([]string, 0)

	if len(data) == 1 {
		return data
	}

	for i, v := range data {
		if string(v[pos]) != bitOfInterest {
			if i >= len(data) {
				break
			}
			toFilter = append(toFilter, v)
		} else {
			continue
		}
	}
	return difference(toFilter, data)
}

func difference(filter, data []string) (result []string) {
	filterMap := make(map[string]struct{}, len(filter))
	for _, v := range filter {
		filterMap[v] = struct{}{}
	}
	for _, v := range data {
		if _, ok := filterMap[v]; !ok {
			result = append(result, v)
		}
	}
	return
}
