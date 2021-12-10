package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Stack []string

var closingSets = map[string]string{
	"<": ">",
	"(": ")",
	"{": "}",
	"[": "]",
}

var errorScores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var autoScores = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func main() {
	start := time.Now()
	f, err := os.Open("day10Input.txt")
	check(err)
	defer f.Close()
	var lines [][]string
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
		i++
	}
	highScore, partTwoLines := partOne(lines)
	fmt.Printf("Total Syntax Error Score Part One: %d \n", highScore)
	autoCompleteScore := partTwo(partTwoLines)
	fmt.Printf("Total Auto Complete Score: %d \n", autoCompleteScore)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(lines [][]string) (int, [][]string) {
	var partTwoLines [][]string
	errorScore := 0
	for _, v := range lines {
		errorValue := checkForError(v)
		if errorValue == 0 {
			partTwoLines = append(partTwoLines, v)
		}
		errorScore += errorValue
	}
	return errorScore, partTwoLines
}

func checkForError(line []string) int {
	var stack Stack
	for i := 0; i < len(line); i++ {
		if len(stack) == 0 {
			stack = append(stack, line[i])
		} else if inOpener(line[i]) {
			stack = append(stack, line[i])
		} else if validCloser(line[i], stack) {
			stack = stack[:len(stack)-1]
		} else {
			return errorScores[line[i]]
		}
	}
	return 0
}

func inOpener(val string) bool {
	if _, ok := closingSets[val]; ok {
		return true
	}
	return false
}

func validCloser(val string, stack Stack) bool {
	if val == closingSets[stack[len(stack)-1]] {
		return true
	}
	return false
}

func partTwo(lines [][]string) int {
	var scores []int
	for _, v := range lines {
		autoScore := autoCompleteLines(v)
		scores = append(scores, autoScore)
	}
	sort.Ints(scores)
	scoreIndex := len(scores) / 2
	return scores[scoreIndex]
}

func autoCompleteLines(line []string) int {
	var stack Stack
	for i := 0; i < len(line); i++ {
		if len(stack) == 0 {
			stack = append(stack, line[i])
		} else if inOpener(line[i]) {
			stack = append(stack, line[i])
		} else if validCloser(line[i], stack) {
			stack = stack[:len(stack)-1]
		}
	}
	score := 0
	for i := len(stack) - 1; i > -1; i-- {
		score = score*5 + autoScores[closingSets[stack[i]]]
	}
	return score
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
