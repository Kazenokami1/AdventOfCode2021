package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("day13Input.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	maxXValue := 0
	maxYValue := 0
	var instructionList [][]int
	var foldList []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			for scanner.Scan() {
				foldList = append(foldList, scanner.Text())
			}
			break
		}
		line := strings.Split(scanner.Text(), ",")
		xValue, _ := strconv.Atoi(line[0])
		yValue, _ := strconv.Atoi(line[1])
		instructionList = append(instructionList, []int{xValue, yValue})
		if xValue > maxXValue {
			maxXValue = xValue
		}
		if yValue > maxYValue {
			maxYValue = yValue
		}
	}
	instructions := make([][]string, maxYValue+1)
	for i := 0; i < maxXValue+1; i++ {
		for j := 0; j < maxYValue+1; j++ {
			instructions[j] = append(instructions[j], ".")
		}
	}
	for _, v := range instructionList {
		instructions[v[1]][v[0]] = "#"
	}
	if len(instructions[0])%2 == 0 {
		for i := range instructions {
			instructions[i] = append(instructions[i], ".")
		}
	}
	numberOfDots := partOne(instructions, foldList, 0)
	fmt.Printf("Total Number of Dots: %d \n", numberOfDots)
	partTwo(instructions, foldList)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(instructions [][]string, foldList []string, numberOfInstructions int) int {
	numberOfDots := 0
	for w := 0; w < numberOfInstructions; w++ {
		foldLine := strings.SplitAfter(foldList[w], "fold along ")
		foldLine = strings.Split(foldLine[1], "=")
		foldLineInt, _ := strconv.Atoi(foldLine[1])
		if foldLine[0] == "y" {
			for i := 0; i < foldLineInt; i++ {
				for j := 0; j < len(instructions[0]); j++ {
					if instructions[i][j] == "#" || instructions[len(instructions)-i-1][j] == "#" {
						instructions[i][j] = "#"
						numberOfDots++
					}
				}
			}
			instructions = instructions[0:foldLineInt]
		} else {
			for i := 0; i < foldLineInt; i++ {
				for j := 0; j < len(instructions); j++ {
					if instructions[j][i] == "#" || instructions[j][len(instructions[0])-i-1] == "#" {
						instructions[j][len(instructions[0])-i-1] = "#"
						numberOfDots++
					}
				}
			}
			for i, v := range instructions {
				instructions[i] = v[foldLineInt+1 : len(instructions[i])]
			}
		}
	}
	return numberOfDots
}

func partTwo(instructions [][]string, foldList []string) {
	for _, w := range foldList {
		foldLine := strings.SplitAfter(w, "fold along ")
		foldLine = strings.Split(foldLine[1], "=")
		foldLineInt, _ := strconv.Atoi(foldLine[1])
		if foldLine[0] == "y" {
			if len(instructions)-foldLineInt < foldLineInt {
				var dotsToAdd []string
				for i := 0; i < len(instructions[0]); i++ {
					dotsToAdd = append(dotsToAdd, ".")
				}
				instructions = append(instructions, dotsToAdd, dotsToAdd)
			}
			for i := 0; i < foldLineInt; i++ {
				for j := 0; j < len(instructions[0]); j++ {
					if instructions[i][j] == "#" || instructions[len(instructions)-i-1][j] == "#" {
						instructions[i][j] = "#"
					}
				}
			}
			instructions = instructions[0:foldLineInt]
		} else {
			for i := 0; i < foldLineInt; i++ {
				for j := 0; j < len(instructions); j++ {
					if instructions[j][i] == "#" || instructions[j][len(instructions[0])-i-1] == "#" {
						instructions[j][len(instructions[0])-i-1] = "#"
					}
				}
			}
			for i, v := range instructions {
				instructions[i] = v[foldLineInt+1 : len(instructions[i])]
			}
		}
	}
	for _, v := range instructions {
		fmt.Println(v)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
