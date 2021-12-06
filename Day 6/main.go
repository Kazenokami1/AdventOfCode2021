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
	f, err := os.Open("day6SampleInput.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var ageStrings []string
	var ages []int
	for scanner.Scan() {
		ageStrings = strings.Split(scanner.Text(), ",")
		check(err)
	}
	for _, value := range ageStrings {
		age, err := strconv.Atoi(value)
		check(err)
		ages = append(ages, age)
	}
	//numberOfFish := partOne(ages, 80)
	//text := fmt.Sprintf("Number of Fish Part One: %d", numberOfFish)
	//fmt.Println(text)
	//numberOfDays := 1
	for i := 0; i < 24; i++ {
		numberOfFish := firstWeekOfFish(ages, i)
		text := fmt.Sprintf("Number of Fish After %d Days: %d", i, numberOfFish)
		fmt.Println(text)
	}
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(ages []int, numberOfDays int) int {
	for i := 0; i < numberOfDays; i++ {
		minValue := minArray(ages)
		i += minValue - 1
		if i >= numberOfDays {
			minValue = i - numberOfDays
		}
		for i := range ages {
			ages[i] = ages[i] - minValue
			if ages[i] == -1 {
				ages[i] = 6
				ages = append(ages, 8)
			}
		}
	}
	return len(ages)
}

func minArray(ages []int) int {
	min := 9
	for _, value := range ages {
		if value < min {
			min = value
		}
	}
	return min + 1
}

func firstWeekOfFish(initialFish []int, numberOfDays int) int {
	totalNumberOfFish := len(initialFish)
	var passingFish []int
	if numberOfDays-7 > 0 {
		weeks := numberOfDays / 7
		remainingDays := numberOfDays % 7
		totalNumberOfFish *= weeks
		for i := 0; i < len(initialFish); i++ {
			passingFish = append(passingFish, initialFish[i])
		}
		totalNumberOfFish += setsOfNewFish(passingFish, numberOfDays-7)
		for i := range initialFish {
			if initialFish[i]-remainingDays < 0 {
				totalNumberOfFish++
			}
		}
	} else {
		remainingDays := numberOfDays
		for i := range initialFish {
			if initialFish[i]-remainingDays < 0 {
				totalNumberOfFish++
			}
		}
	}
	return totalNumberOfFish
}

func setsOfNewFish(initialFish []int, numberOfDays int) int {
	totalNumberOfFish := len(initialFish)
	var passingFish []int
	var remainingDays int
	weeks := (numberOfDays - 2) / 7
	if numberOfDays-9 > 0 {
		loop := 1
		for i := numberOfDays - 9; i > 0; i -= 7 {
			for i := 0; i < len(initialFish); i++ {
				passingFish = append(passingFish, initialFish[i])
			}
			totalNumberOfFish *= weeks
			totalNumberOfFish += setsOfNewFish(passingFish, i)
			if loop == 1 {
				remainingDays = i % 9
			} else {
				remainingDays = i % 7
			}
			loop++
		}
		for i := range initialFish {
			if initialFish[i]-remainingDays < 0 {
				totalNumberOfFish += loop
			}
		}
	} else {
		remainingDays := numberOfDays
		for _, age := range initialFish {
			if age+2-remainingDays < 0 {
				totalNumberOfFish++
			}
		}
	}
	return totalNumberOfFish
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
