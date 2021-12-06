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
	f, err := os.Open("day6Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var ageStrings []string
	var ages []int
	fishMap := make(map[int]int)
	for scanner.Scan() {
		ageStrings = strings.Split(scanner.Text(), ",")
		check(err)
	}
	for _, value := range ageStrings {
		age, err := strconv.Atoi(value)
		check(err)
		ages = append(ages, age)
	}
	for i := 0; i < 9; i++ {
		fishMap[i] = 0
	}
	for _, value := range ages {
		if _, ok := fishMap[value]; ok {
			fishMap[value]++
		} else {
			fishMap[value] = 1
		}
	}
	for i := 256; i < 257; i++ {
		totalFish := 0
		_, totalFish = recursiveFish(fishMap, i)
		text := fmt.Sprintf("Number of Fish After %d Days: %d", i, totalFish)
		fmt.Println(text)
	}
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func recursiveFish(initialFish map[int]int, numberOfDays int) (map[int]int, int) {
	totalNumberOfFish := make(map[int]int)
	for key, value := range initialFish {
		totalNumberOfFish[key] = value
	}
	totalFish := 0
	if numberOfDays > 7 {
		for key, value := range initialFish {
			if key-7 >= 0 {
				totalNumberOfFish[key-7] += value
				totalNumberOfFish[key] -= value
			} else {
				totalNumberOfFish[key+2] += value
			}
		}
		totalFishReturned := 0
		totalNumberOfFish, totalFishReturned = recursiveFish(totalNumberOfFish, numberOfDays-7)
		totalFish += totalFishReturned
	} else if numberOfDays == 7 {
		for _, value := range initialFish {
			if value-numberOfDays < 0 {
				totalFish += value
			}
			totalFish += value
		}
	} else {
		for key, value := range initialFish {
			if key-numberOfDays < 0 {
				totalFish += value
			}
			totalFish += value
		}
	}
	return totalNumberOfFish, totalFish
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
