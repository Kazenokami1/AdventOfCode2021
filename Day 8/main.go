package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type placementArray struct {
	Top         string
	Middle      string
	Bottom      string
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
}

func main() {
	start := time.Now()
	f, err := os.Open("day8Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var signals [][]string
	var output [][]string
	for scanner.Scan() {
		notes := strings.Split(scanner.Text(), " | ")
		signals = append(signals, strings.Fields(notes[0]))
		output = append(output, strings.Fields(notes[1]))
	}
	numbers := partOne(output)
	text := fmt.Sprintf("Total Numbers of 1, 4, 7, or 8 for Part One: %d", numbers)
	fmt.Println(text)
	numbers = partTwo(signals, output)
	text = fmt.Sprintf("Total Sum for Part Two: %d", numbers)
	fmt.Println(text)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(output [][]string) int {
	total := 0
	for i := range output {
		for _, v := range output[i] {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				total++
			}
		}
	}
	return total
}

func partTwo(signals, output [][]string) int {
	var letterPlacements placementArray
	signalMap := make(map[int]string)
	totalSum := 0
	for i := range signals {
		for _, v := range signals[i] {
			if len(v) == 2 {
				signalMap[1] = v
			} else if len(v) == 3 {
				signalMap[7] = v
			} else if len(v) == 4 {
				signalMap[4] = v
			} else if len(v) == 7 {
				signalMap[8] = v
			}
		}
		letterPlacements, signalMap = determineNumberArray(signals[i], signalMap)
		signalMap[5] = letterPlacements.Top + letterPlacements.TopLeft + letterPlacements.Middle + letterPlacements.BottomRight + letterPlacements.Bottom
		signalMap[6] = signalMap[5] + letterPlacements.BottomLeft
		signalMap[0] = signalMap[7] + letterPlacements.Bottom + letterPlacements.TopLeft + letterPlacements.BottomLeft
		outputString := ""
		for _, v := range output[i] {
			output := string(findDisplayValue(v, signalMap))
			outputString += output
		}
		displayValue, _ := strconv.Atoi(outputString)
		totalSum += displayValue
	}
	return totalSum
}

func determineNumberArray(signal []string, signalMap map[int]string) (placementArray, map[int]string) {
	var placement placementArray
	placement.Top, _ = findLetterNotFound(signalMap[7], signalMap[1], 1)
	placement.Bottom, signalMap = findBottom(signal, placement.Top, signalMap)
	placement.BottomLeft, _ = findLetterNotFound(signalMap[8], signalMap[9], 1)
	placement.Middle, signalMap = findMiddle(signal, placement.Bottom, signalMap)
	placement.TopRight, signalMap = findTopRight(signal, placement, signalMap)
	placement.BottomRight, _ = findLetterNotFound(signalMap[1], placement.TopRight, 1)
	placement.TopLeft, _ = findLetterNotFound(signalMap[8], signalMap[2]+placement.BottomRight, 1)
	return placement, signalMap
}

func findDisplayValue(numberString string, signalMap map[int]string) string {
	for i := range signalMap {
		if len(signalMap[i]) == len(numberString) {
			_, finished := findLetterNotFound(signalMap[i], numberString, 0)
			if finished {
				return fmt.Sprint(i)
			}
		}
	}
	return "z"
}

func findBottom(signal []string, top string, signalMap map[int]string) (string, map[int]string) {
	four := signalMap[4]
	nineMinusBottom := four + top
	for i := range signal {
		if len(signal[i]) == 6 {
			letter, finished := findLetterNotFound(signal[i], nineMinusBottom, 1)
			if finished {
				signalMap[9] = signal[i]
				return letter, signalMap
			}
		}
	}
	return "z", signalMap
}

func findMiddle(signal []string, bottom string, signalMap map[int]string) (string, map[int]string) {
	threeMinusMiddle := signalMap[7] + bottom
	for i := range signal {
		if len(signal[i]) == 5 {
			letter, finished := findLetterNotFound(signal[i], threeMinusMiddle, 1)
			if finished {
				signalMap[3] = signal[i]
				return letter, signalMap
			}
		}
	}
	return "z", signalMap
}

func findTopRight(signal []string, pArray placementArray, signalMap map[int]string) (string, map[int]string) {
	twoMinusTopRight := pArray.Top + pArray.Middle + pArray.BottomLeft + pArray.Bottom
	for i := range signal {
		if len(signal[i]) == 5 {
			letter, finished := findLetterNotFound(signal[i], twoMinusTopRight, 1)
			if finished {
				signalMap[2] = signal[i]
				return letter, signalMap
			}
		}
	}
	return "z", signalMap
}

func findLetterNotFound(first, second string, limit int) (string, bool) {
	var lettersNotFoundArray []string
	lettersNotFound := 0
	for _, v := range first {
		letterFound := false
		for _, w := range second {
			if v == w {
				letterFound = true
			}
		}
		if !letterFound {
			lettersNotFound++
			lettersNotFoundArray = append(lettersNotFoundArray, string(v))
		}
	}
	if lettersNotFound == limit && limit > 0 {
		return lettersNotFoundArray[0], true
	} else if lettersNotFound == limit && limit == 0 {
		return "z", true
	}
	return "z", false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
