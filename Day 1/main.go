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
	f, err := os.Open("day1Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var depths []int
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)
		depths = append(depths, depth)
	}
	increasesPartOne := partOne(depths)
	increasesPartTwo := partTwo(depths)

	text := fmt.Sprintf("Number of Increases Part One: %d", increasesPartOne)
	fmt.Println(text)
	text = fmt.Sprintf("Number of Increases Part Two: %d", increasesPartTwo)
	fmt.Println(text)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(depths []int) int {
	increases := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	return increases
}

func partTwo(depths []int) int {
	increases := 0
	for i := 3; i < len(depths); i++ {
		if depths[i]+depths[i-1]+depths[i-2] > depths[i-1]+depths[i-2]+depths[i-3] {
			increases++
		}
	}
	return increases
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
