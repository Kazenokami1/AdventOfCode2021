package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type cave struct {
	Position       []int
	Entered        bool
	RiskLevel      float64
	Neighbors      []*cave
	End            bool
	TotalRiskToEnd float64
}

func main() {
	start := time.Now()
	f, err := os.Open("Day15Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	endCave, startCave := partOneParseInput(scanner)
	totalRisk := partOne(endCave, startCave)
	fmt.Printf("Total Risk of Part One: %f \n", totalRisk)
	f, err = os.Open("Day15Input.txt")
	check(err)
	defer f.Close()
	scannerTwo := bufio.NewScanner(f)
	endCave, startCave = partTwoParseInput(scannerTwo)
	totalRisk = partTwo(endCave, startCave)
	fmt.Printf("Total Risk of Part Two: %f \n", totalRisk)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOneParseInput(scanner *bufio.Scanner) (*cave, *cave) {
	x := 0
	var caveList []*cave
	var rowLength int
	for scanner.Scan() {
		chitonList := strings.Split(scanner.Text(), "")
		rowLength = len(chitonList)
		for i, v := range chitonList {
			vstring, _ := strconv.Atoi(v)
			aCave := cave{Position: []int{i, x}, RiskLevel: float64(vstring), Entered: false, TotalRiskToEnd: math.Inf(1), End: false}
			caveList = append(caveList, &aCave)
		}
		x++
	}
	var endCave *cave
	var startCave *cave
	for _, v := range caveList {
		for _, w := range caveList {
			if math.Abs(float64(v.Position[0]-w.Position[0])) < 2 && math.Abs(float64(v.Position[1]-w.Position[1])) < 2 && !reflect.DeepEqual(v.Position, w.Position) && (v.Position[1]-w.Position[1] == 0 || v.Position[0]-w.Position[0] == 0) {
				v.Neighbors = append(v.Neighbors, w)
			}
		}
		if v.Position[0] == rowLength-1 && v.Position[1] == x-1 {
			v.End = true
			v.TotalRiskToEnd = 0
			endCave = v
		} else if v.Position[0] == 0 && v.Position[1] == 0 {
			startCave = v
		}
	}
	return endCave, startCave
}

func partOne(end, start *cave) float64 {
	assignRisk(end, start, end.TotalRiskToEnd)
	return start.TotalRiskToEnd
}

func assignRisk(end, start *cave, neighborRisk float64) {
	loopAgain := false
	for _, v := range end.Neighbors {
		if v.TotalRiskToEnd > neighborRisk+end.RiskLevel {
			v.TotalRiskToEnd = neighborRisk + end.RiskLevel
			loopAgain = true
		}
		if reflect.DeepEqual(v.Position, start.Position) {
			return
		} else if loopAgain && neighborRisk < start.TotalRiskToEnd {
			assignRisk(v, start, v.TotalRiskToEnd)
		}
	}
}

func partTwoParseInput(scanner *bufio.Scanner) (*cave, *cave) {
	var endCave *cave
	var startCave *cave
	var chitonListArray [][]string
	for scanner.Scan() {
		chitonList := strings.Split(scanner.Text(), "")
		listLength := len(chitonList)
		for i := 1; i < 5; i++ {
			for j := 0; j < listLength; j++ {
				appendIntString, _ := strconv.Atoi(chitonList[j])
				appendInt := appendIntString + i
				if appendInt > 9 {
					appendInt -= 9
				}
				chitonList = append(chitonList, fmt.Sprint(appendInt))
			}
		}
		chitonListArray = append(chitonListArray, chitonList)
	}
	listLength := len(chitonListArray)
	for i := 1; i < 5; i++ {
		for j := 0; j < listLength; j++ {
			var chitonStringArrayAppend []string
			for _, v := range chitonListArray[j+(i-1)*listLength] {
				appendIntString, _ := strconv.Atoi(v)
				appendInt := appendIntString + 1
				if appendInt > 9 {
					appendInt -= 9
				}
				chitonStringArrayAppend = append(chitonStringArrayAppend, fmt.Sprint(appendInt))
			}
			chitonListArray = append(chitonListArray, chitonStringArrayAppend)
		}
	}
	var caveListTwo []*cave
	for i, v := range chitonListArray {
		for j, w := range v {
			vstring, _ := strconv.Atoi(w)
			aCave := cave{Position: []int{j, i}, RiskLevel: float64(vstring), Entered: false, TotalRiskToEnd: math.Inf(1), End: false}
			caveListTwo = append(caveListTwo, &aCave)
		}
	}
	for _, v := range caveListTwo {
		for _, w := range caveListTwo {
			if math.Abs(float64(v.Position[0]-w.Position[0])) < 2 && math.Abs(float64(v.Position[1]-w.Position[1])) < 2 && !reflect.DeepEqual(v.Position, w.Position) && (v.Position[1]-w.Position[1] == 0 || v.Position[0]-w.Position[0] == 0) {
				v.Neighbors = append(v.Neighbors, w)
			}
		}
		if v.Position[0] == len(chitonListArray[0])-1 && v.Position[1] == len(chitonListArray)-1 {
			v.End = true
			v.TotalRiskToEnd = 0
			endCave = v
		} else if v.Position[0] == 0 && v.Position[1] == 0 {
			startCave = v
		}
	}
	return endCave, startCave
}

func partTwo(end, start *cave) float64 {
	assignRisk(end, start, end.TotalRiskToEnd)
	return start.TotalRiskToEnd
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
