package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	f, _ := os.Open("Day25Input.txt")
	defer f.Close()

	var cucumbers []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cucumbers = append(cucumbers, scanner.Text())
	}

	numberOfMoves := findFinalMap(cucumbers)
	fmt.Printf("Cucumbers Stopped Moving at %d \n", numberOfMoves)

	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func findFinalMap(cucumbers []string) int {
	var totalSteps int
	cucumberMoved := true
	for cucumberMoved {
		cucumberMoved = false
		var newMap []string
		for _, row := range cucumbers {
			var newRow string
			var skipOne bool
			for j, v := range row {
				if v == '>' && j == len(row)-1 {
					if row[0] == '.' {
						newRow = ">" + newRow[1:] + "."
						cucumberMoved = true
					} else {
						newRow += ">"
					}
				} else if v == '>' && row[j+1] == '.' {
					newRow += ".>"
					skipOne = true
					cucumberMoved = true
				} else if skipOne {
					skipOne = false
				} else {
					newRow += string(v)
				}
			}
			newMap = append(newMap, newRow)
		}
		var newNewMap []string
		moved := make(map[[2]int]string)
		for j := 0; j < len(newMap); j++ {
			var newRow string
			if j == 0 {
				for k, v := range newMap[len(newMap)-1] {
					if v == 'v' && newMap[j][k] == '.' {
						newRow += "v"
						moved[[2]int{len(newMap) - 1, k}] = "."
						cucumberMoved = true
					} else {
						newRow += string(newMap[j][k])
					}
				}
				newNewMap = append(newNewMap, newRow)
			} else {
				for k, v := range newMap[j-1] {
					if v == 'v' && newMap[j][k] == '.' {
						newRow += "v"
						newNewMap[j-1] = newNewMap[j-1][:k] + "." + newNewMap[j-1][k+1:]
						cucumberMoved = true
					} else if moved[[2]int{j, k}] != "." {
						newRow += string(newMap[j][k])
					} else {
						newRow += "."
						delete(moved, [2]int{j, k})
					}
				}
				newNewMap = append(newNewMap, newRow)
			}
		}
		cucumbers = newNewMap
		totalSteps++
	}
	return totalSteps
}
