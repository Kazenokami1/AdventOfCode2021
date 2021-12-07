package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("day7Input.txt")
	check(err)
	defer f.Close()
	var crabPositions []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		crabStrings := strings.Split(scanner.Text(), ",")
		for _, value := range crabStrings {
			crabX, _ := strconv.Atoi(value)
			crabPositions = append(crabPositions, crabX)
		}
	}
	totalFuel := partOne(crabPositions)
	text := fmt.Sprintf("Total Fuel for Part One: %d", totalFuel)
	fmt.Println(text)
	totalFuel = partTwo(crabPositions)
	text = fmt.Sprintf("Total Fuel for Part Two: %d", totalFuel)
	fmt.Println(text)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(crabPositions []int) int {
	crabMap := make(map[int]int)
	for _, value := range crabPositions {
		crabMap[value]++
	}
	minFuel := 9999999
	for i := range crabMap {
		currentFuel := 0
		for j, value := range crabMap {
			currentFuel += int(math.Abs(float64(i-j))) * value
		}
		if currentFuel < minFuel {
			minFuel = currentFuel
		}
	}
	return minFuel
}

func partTwo(crabPositions []int) int {
	crabMap := make(map[int]int)
	for _, value := range crabPositions {
		crabMap[value]++
	}
	for i := 0; i < len(crabMap); i++ {
		if _, ok := crabMap[i]; !ok {
			crabMap[i] = 0
		}
	}
	minFuel := -1
	for i := range crabMap {
		currentFuel := 0
		for j, value := range crabMap {
			movement := int(math.Abs(float64(i - j)))
			fuelCost := 0
			for i := 1; i < movement+1; i++ {
				fuelCost += i
			}
			currentFuel += fuelCost * value
		}
		if currentFuel < minFuel || minFuel == -1 {
			minFuel = currentFuel
		}
	}
	return minFuel
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
