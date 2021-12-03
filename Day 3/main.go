package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("day3Input.txt")
	check(err)
	defer f.Close()

	var binaries []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		binaries = append(binaries, scanner.Text())
	}
	gamma, epsilon := partOne(binaries)
	text := fmt.Sprintf("Part One Gamma Value: %d", gamma)
	fmt.Println(text)
	text = fmt.Sprintf("Part Two Epsilon Value: %d", epsilon)
	fmt.Println(text)
	fmt.Println(gamma * epsilon)
	oxygen, co := partTwo(binaries)
	text = fmt.Sprintf("Part Two Epsilon Value: %d", oxygen)
	fmt.Println(text)
	text = fmt.Sprintf("Part Two Epsilon Value: %d", co)
	fmt.Println(text)
	fmt.Println(oxygen * co)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(binaries []string) (int64, int64) {
	var gamma string
	var epsilon string
	for i := 0; i < len(binaries[0]); i++ {
		oneCount := 0
		zeroCount := 0
		for _, value := range binaries {
			if value[i] == 49 {
				oneCount++
			} else {
				zeroCount++
			}
		}
		if oneCount > zeroCount {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	check(err)
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	check(err)
	return gammaInt, epsilonInt
}

func partTwo(binaries []string) (int64, int64) {
	oxygenInt := findOxygen(binaries)
	coInt := findCo(binaries)
	return oxygenInt, coInt
}

func findOxygen(binaries []string) int64 {
	for i := 0; i < len(binaries[0]); i++ {
		if len(binaries) > 1 {
			var oxygenOnes []string
			var oxygenZeros []string
			for _, value := range binaries {
				if value[i] == 49 {
					oxygenOnes = append(oxygenOnes, value)
				} else {
					oxygenZeros = append(oxygenZeros, value)
				}
			}
			if len(oxygenOnes) >= len(oxygenZeros) {
				binaries = oxygenOnes
			} else {
				binaries = oxygenZeros
			}
		}
	}
	oxygenInt, err := strconv.ParseInt(binaries[0], 2, 64)
	check(err)
	return oxygenInt
}

func findCo(binaries []string) int64 {
	for i := 0; i < len(binaries[0]); i++ {
		if len(binaries) > 1 {
			var coOnes []string
			var coZeros []string
			for _, value := range binaries {
				if value[i] == 49 {
					coOnes = append(coOnes, value)
				} else {
					coZeros = append(coZeros, value)
				}
			}
			if len(coOnes) < len(coZeros) {
				binaries = coOnes
			} else {
				binaries = coZeros
			}
		}
	}
	coInt, err := strconv.ParseInt(binaries[0], 2, 64)
	check(err)
	return coInt
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
