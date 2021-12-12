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

type octopus struct {
	Position  []int
	Energy    int
	Neighbors []*octopus
	Flash     bool
}

func main() {
	start := time.Now()
	f, err := os.Open("day11Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	x := 0
	var octoList []*octopus
	for scanner.Scan() {
		energyList := strings.Split(scanner.Text(), "")
		for i, v := range energyList {
			vstring, _ := strconv.Atoi(v)
			octo := octopus{Position: []int{i, x}, Energy: vstring, Flash: false}
			octoList = append(octoList, &octo)
		}
		x++
	}
	for _, v := range octoList {
		for _, w := range octoList {
			if math.Abs(float64(v.Position[0]-w.Position[0])) < 2 && math.Abs(float64(v.Position[1]-w.Position[1])) < 2 && !reflect.DeepEqual(v.Position, w.Position) {
				v.Neighbors = append(v.Neighbors, w)
			}
		}
	}
	//flashes := partOne(octoList, 100)
	//fmt.Printf("Total Number of Flashes: %d \n", flashes)
	simultaneousFlashDay := partTwo(octoList, 0)
	fmt.Printf("Day of first Simultaneous Flash: %d \n", simultaneousFlashDay)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(list []*octopus, days int) int {
	numberOfFlashes := 0
	for i := 0; i < days; i++ {
		for _, v := range list {
			v.Energy++
		}
		flash := true
		for flash {
			flash = false
			for _, v := range list {
				if v.Energy > 9 && !v.Flash {
					numberOfFlashes++
					v.Flash = true
					flash = true
					for _, w := range v.Neighbors {
						w.Energy++
					}
				}
			}
		}
		for _, v := range list {
			if v.Energy > 9 {
				v.Energy = 0
			}
			v.Flash = false
		}
	}
	return numberOfFlashes
}

func partTwo(list []*octopus, startDay int) int {
	simultaneousFlash := false
	simultaneousFlashDay := startDay
	for !simultaneousFlash {
		flashes := 0
		for _, v := range list {
			v.Energy++
		}
		flash := true
		for flash {
			flash = false
			for _, v := range list {
				if v.Energy > 9 && !v.Flash {
					v.Flash = true
					flash = true
					for _, w := range v.Neighbors {
						w.Energy++
					}
				}
			}
		}
		for _, v := range list {
			if v.Energy > 9 {
				v.Energy = 0
				flashes++
			}
			v.Flash = false
		}
		if flashes == len(list) {
			simultaneousFlash = true
		}
		simultaneousFlashDay++
	}
	return simultaneousFlashDay
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
