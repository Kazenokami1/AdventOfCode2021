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
	f, err := os.Open("Day9Input.txt")
	check(err)
	defer f.Close()
	var heights [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var heightInts []int
		caveHeights := strings.Split(scanner.Text(), "")
		for _, v := range caveHeights {
			for _, w := range v {
				height, _ := strconv.Atoi(string(w))
				heightInts = append(heightInts, height)
			}
		}
		heights = append(heights, heightInts)
	}
	sum := partOne(heights)
	fmt.Printf("Risk Factor of %d \n", sum)
	multiplier := partTwo(heights)
	fmt.Printf("Risk Factor of %d \n", multiplier)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}
