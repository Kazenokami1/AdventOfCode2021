package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("Day18Input.txt")
	check(err)
	defer f.Close()
	var snailFishNumbers [][]string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		snailFishNumbers = append(snailFishNumbers, strings.Split(scanner.Text(), ""))
	}
	magnitude := partOne(snailFishNumbers)
	fmt.Printf("Magnitude of Part One: %d \n", magnitude)
	magnitude = partTwo(snailFishNumbers)
	fmt.Printf("Highest Magnitude of Sums: %d \n", magnitude)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(numbers [][]string) int {
	magnitude := 0
	firstNumber := []string{}
	for _, v := range numbers[0] {
		firstNumber = append(firstNumber, v)
	}
	for i := 0; i < len(numbers)-1; i++ {
		firstNumber = append(firstNumber, "0")
		copy(firstNumber[1:], firstNumber[0:])
		firstNumber[0] = "["
		firstNumber = append(firstNumber, ",")
		for _, v := range numbers[i+1] {
			firstNumber = append(firstNumber, v)
		}
		firstNumber = append(firstNumber, "]")
		firstNumber = reduceNumber(firstNumber)
	}
	fmt.Println(firstNumber)
	magnitude = findMagnitude(firstNumber)
	return magnitude
}

func reduceNumber(number []string) []string {
	var reducedNumber []string
	reduced := false
	for !reduced {
		reduced = true
		needToExplode := false
		needToSplit := false
		checkForExplode := 0
		var firstInt int
		var secondInt int
		leftIntFound := false
		rightIntFound := false
		var leftInt int
		var rightInt int
		var firstIntIndex int
		var secondIntIndex int
		leftIntIndex := 0
		rightIntIndex := len(number) - 1
		var splitIndex int
		var splitInt int
		//Check for Need to Explode and capture related integers
		for i, v := range number {
			if string(v) == "[" {
				checkForExplode++
				if checkForExplode > 4 {
					needToExplode = true
					firstInt, _ = strconv.Atoi(number[i+1])
					firstIntIndex = i + 1
					secondInt, _ = strconv.Atoi(number[i+3])
					secondIntIndex = i + 3
					reduced = false
					break
				}
			} else if string(v) == "]" {
				checkForExplode--
			} else if string(v) != "," {
				leftIntFound = true
				leftInt, _ = strconv.Atoi(string(v))
				leftIntIndex = i
				if leftInt > 9 && !needToSplit {
					reduced = false
					needToSplit = true
					splitIndex = i
					splitInt = leftInt
				}
			}
		}
		if needToExplode {
			reducedNumber = nil
			for i, v := range number[secondIntIndex+1:] {
				if v != "[" && v != "]" && v != "," {
					rightIntFound = true
					rightInt, _ = strconv.Atoi(v)
					rightIntIndex = i + secondIntIndex + 1
					break
				}
			}
			if leftIntFound {
				for _, v := range number[0:leftIntIndex] {
					reducedNumber = append(reducedNumber, v)
				}
				reducedNumber = append(reducedNumber, fmt.Sprint(leftInt+firstInt))
				for _, v := range number[leftIntIndex+1 : firstIntIndex-1] {
					reducedNumber = append(reducedNumber, v)
				}
			} else {
				for _, v := range number[0 : firstIntIndex-1] {
					reducedNumber = append(reducedNumber, v)
				}
			}
			reducedNumber = append(reducedNumber, "0")
			if rightIntFound {
				for _, v := range number[secondIntIndex+2 : rightIntIndex] {
					reducedNumber = append(reducedNumber, v)
				}
				reducedNumber = append(reducedNumber, fmt.Sprint(rightInt+secondInt))
				for _, v := range number[rightIntIndex+1:] {
					reducedNumber = append(reducedNumber, v)
				}
			} else {
				for _, v := range number[secondIntIndex+2:] {
					reducedNumber = append(reducedNumber, v)
				}
			}
		} else if needToSplit {
			reducedNumber = nil
			for _, v := range number[0:splitIndex] {
				reducedNumber = append(reducedNumber, v)
			}
			add := 0
			if splitInt%2 == 1 {
				add = 1
			}
			reducedNumber = append(reducedNumber, "[")
			reducedNumber = append(reducedNumber, fmt.Sprint(splitInt/2))
			reducedNumber = append(reducedNumber, ",")
			reducedNumber = append(reducedNumber, fmt.Sprint(splitInt/2+add))
			reducedNumber = append(reducedNumber, "]")
			if number[splitIndex+1] != "]" {
				reducedNumber = append(reducedNumber, ",")
			} else {
				reducedNumber = append(reducedNumber, "]")
			}
			for _, v := range number[splitIndex+2:] {
				reducedNumber = append(reducedNumber, v)
			}
		}
		number = nil
		for _, v := range reducedNumber {
			number = append(number, v)
		}
	}
	return reducedNumber
}

func findMagnitude(number []string) int {
	var magnitude int
	var magnitudeArray []string
	finished := false
	for !finished {
		for i, v := range number {
			magnitudeArray = nil
			if v == "[" && number[i+2] == "," && number[i+4] == "]" {
				for _, w := range number[0:i] {
					magnitudeArray = append(magnitudeArray, w)
				}
				leftValue, _ := strconv.Atoi(number[i+1])
				rightValue, _ := strconv.Atoi(number[i+3])
				magnitudeArray = append(magnitudeArray, fmt.Sprint(leftValue*3+rightValue*2))
				for _, w := range number[i+5:] {
					magnitudeArray = append(magnitudeArray, w)
				}
				break
			}
		}
		if len(magnitudeArray) == 1 {
			finished = true
			magnitude, _ = strconv.Atoi(magnitudeArray[0])
		} else if len(magnitudeArray) == 0 {
			return 0
		} else {
			number = nil
			for _, v := range magnitudeArray {
				number = append(number, v)
			}
		}
	}
	return magnitude
}
func partTwo(numbers [][]string) int {
	var number []string
	var maxMag int
	for _, v := range numbers {
		for _, w := range numbers {
			number = nil
			if !reflect.DeepEqual(v, w) {
				number = append(number, "[")
				for _, x := range v {
					number = append(number, x)
				}
				number = append(number, ",")
				for _, x := range w {
					number = append(number, x)
				}
				number = append(number, "]")
				number = reduceNumber(number)
				magnitude := findMagnitude(number)
				if magnitude > maxMag {
					maxMag = magnitude
				}
			}
		}
	}
	return maxMag
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
