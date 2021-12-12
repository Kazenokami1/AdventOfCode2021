package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type cave struct {
	Name           string
	Size           string
	ConnectedCaves []*cave
	NumberOfPaths  int
}

func main() {
	start := time.Now()
	f, err := os.Open("day12Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var caves []*cave
	for scanner.Scan() {
		var caveOne *cave
		var caveTwo *cave
		cavePath := strings.Split(scanner.Text(), "-")
		caveExists, caveOne := checkCaveExists(cavePath[0], caves)
		if !caveExists {
			caveOne = &cave{Name: cavePath[0]}
			if strings.ToUpper(cavePath[0]) == cavePath[0] {
				caveOne.Size = "Big"
			} else {
				caveOne.Size = "Small"
			}
			caves = append(caves, caveOne)
		}
		caveExists, caveTwo = checkCaveExists(cavePath[1], caves)
		if !caveExists {
			caveTwo = &cave{Name: cavePath[1]}
			if strings.ToUpper(cavePath[1]) == cavePath[1] {
				caveTwo.Size = "Big"
			} else {
				caveTwo.Size = "Small"
			}
			caves = append(caves, caveTwo)
		}
		pathExists := cavesConnected(caveOne, caveTwo)
		if !pathExists {
			caveOne.ConnectedCaves = append(caveOne.ConnectedCaves, caveTwo)
			caveTwo.ConnectedCaves = append(caveTwo.ConnectedCaves, caveOne)
		}
	}
	var startingCave *cave
	for _, v := range caves {
		if v.Name == "start" {
			startingCave = v
		}
	}
	partOne(startingCave)
	for _, v := range caves {
		if v.Name == "end" {
			fmt.Printf("Total Number of Paths Found: %d \n", v.NumberOfPaths)
			v.NumberOfPaths = 0
		}
	}
	partTwo(startingCave)
	for _, v := range caves {
		if v.Name == "end" {
			fmt.Printf("Total Number of Paths Found: %d \n", v.NumberOfPaths)
			v.NumberOfPaths = 0
		}
	}
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func checkCaveExists(name string, caves []*cave) (bool, *cave) {
	var cave *cave
	for _, v := range caves {
		if v.Name == name {
			return true, v
		}
	}
	return false, cave
}

func cavesConnected(caveOne *cave, caveTwo *cave) bool {
	for _, v := range caveOne.ConnectedCaves {
		if v.Name == caveTwo.Name {
			return true
		}
	}
	return false
}

func partOne(cave *cave) {
	visitedSmall := []string{"start"}
	findCavePaths(cave, visitedSmall, "start")
}

func findCavePaths(cave *cave, visitedSmall []string, cavePath string) {
	for _, a := range cave.ConnectedCaves {
		currentVisitedSmall := visitedSmall
		currentCavePath := cavePath
		if a.Name == "end" {
			//We reached the end. Yay!
			currentCavePath += "->" + a.Name
			//fmt.Println(currentCavePath)
			a.NumberOfPaths++
		} else if a.Size == "Small" && !ifVisitedSmall(a, visitedSmall) {
			//Entering a small cave we haven't been to before
			currentVisitedSmall = append(currentVisitedSmall, a.Name)
			currentCavePath += "->" + a.Name
			findCavePaths(a, currentVisitedSmall, currentCavePath)
		} else if a.Size == "Small" && ifVisitedSmall(a, visitedSmall) {
			//Who says you can't go back... oh the puzzle did
		} else {
			//We're entering a large cave!
			currentCavePath += "->" + a.Name
			findCavePaths(a, currentVisitedSmall, currentCavePath)
		}
	}
}

func ifVisitedSmall(cave *cave, visited []string) bool {
	for _, v := range visited {
		if cave.Name == v {
			return true
		}
	}
	return false
}

func partTwo(cave *cave) {
	visitedSmall := []string{"start"}
	findCavePathsTwo(cave, visitedSmall, "start", false)
}

func findCavePathsTwo(cave *cave, visitedSmall []string, cavePath string, visitedTwice bool) {
	for _, a := range cave.ConnectedCaves {
		currentVisitedSmall := visitedSmall
		currentCavePath := cavePath
		if a.Name == "end" {
			//We reached the end. Yay!
			currentCavePath += "->" + a.Name
			//fmt.Println(currentCavePath)
			a.NumberOfPaths++
		} else if a.Size == "Small" && !ifVisitedSmall(a, visitedSmall) {
			//Entering a small cave we haven't been to before
			currentVisitedSmall = append(currentVisitedSmall, a.Name)
			currentCavePath += "->" + a.Name
			findCavePathsTwo(a, currentVisitedSmall, currentCavePath, visitedTwice)
		} else if a.Size == "Small" && ifVisitedSmall(a, visitedSmall) && !visitedTwice {
			if a.Name != "start" && a.Name != "end" {
				//We haven't visited a small cave twice... yet
				currentCavePath += "->" + a.Name
				findCavePathsTwo(a, currentVisitedSmall, currentCavePath, true)
			}
		} else if a.Size == "Small" && ifVisitedSmall(a, visitedSmall) {
			//Who says you can't go back... oh the puzzle did
		} else {
			//We're entering a large cave!
			currentCavePath += "->" + a.Name
			findCavePathsTwo(a, currentVisitedSmall, currentCavePath, visitedTwice)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
