package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	xCoordinate int
	yCoordinate int
}

type line struct {
	points []point
}

var lines []line

func main() {
	start := time.Now()
	f, err := os.Open("day5Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var lineOne []point
		p := strings.Split(scanner.Text(), " -> ")
		pointOne := strings.Split(p[0], ",")
		xPoint, _ := strconv.Atoi(pointOne[0])
		yPoint, _ := strconv.Atoi(pointOne[1])
		pointTwo := strings.Split(p[1], ",")
		xPointTwo, _ := strconv.Atoi(pointTwo[0])
		yPointTwo, _ := strconv.Atoi(pointTwo[1])
		p1 := point{xPoint, yPoint}
		p2 := point{xPointTwo, yPointTwo}
		lineOne = append(lineOne, p1, p2)
		lines = append(lines, line{lineOne})
	}

	pointsCovered := partOne(lines)
	duplicatesCovered := 0
	for _, value := range pointsCovered {
		if value > 1 {
			duplicatesCovered += 1
		}
	}
	fmt.Println("Duplicate Points Covered in Part One: " + fmt.Sprint(duplicatesCovered))
	pointsCovered = partTwo(lines)
	duplicatesCovered = 0
	for _, value := range pointsCovered {
		if value > 1 {
			duplicatesCovered += 1
		}
	}
	fmt.Println("Duplicate Points Covered in Part Two: " + fmt.Sprint(duplicatesCovered))
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(lines []line) map[point]int {
	pointsCovered := make(map[point]int)
	for _, line := range lines {
		x1 := line.points[0].xCoordinate
		x2 := line.points[1].xCoordinate
		y1 := line.points[0].yCoordinate
		y2 := line.points[1].yCoordinate
		if x1 == x2 {
			if y1 > y2 {
				for i := 0; i < y1-y2+1; i++ {
					_, exist := pointsCovered[point{x1, y2 + i}]
					if exist {
						pointsCovered[point{x1, y2 + i}] += 1
					} else {
						pointsCovered[point{x1, y2 + i}] = 1
					}
				}
			} else {
				for i := 0; i < y2-y1+1; i++ {
					_, exist := pointsCovered[point{x1, y1 + i}]
					if exist {
						pointsCovered[point{x1, y1 + i}] += 1
					} else {
						pointsCovered[point{x1, y1 + i}] = 1
					}
				}
			}
		} else if y1 == y2 {
			if x1 > x2 {
				for i := 0; i < x1-x2+1; i++ {
					_, exist := pointsCovered[point{x2 + i, y1}]
					if exist {
						pointsCovered[point{x2 + i, y1}] += 1
					} else {
						pointsCovered[point{x2 + i, y1}] = 1
					}
				}
			} else {
				for i := 0; i < x2-x1+1; i++ {
					_, exist := pointsCovered[point{x1 + i, y1}]
					if exist {
						pointsCovered[point{x1 + i, y1}] += 1
					} else {
						pointsCovered[point{x1 + i, y1}] = 1
					}
				}
			}
		}
	}
	return pointsCovered
}

func partTwo(lines []line) map[point]int {
	pointsCovered := partOne(lines)
	for _, line := range lines {
		x1 := line.points[0].xCoordinate
		x2 := line.points[1].xCoordinate
		y1 := line.points[0].yCoordinate
		y2 := line.points[1].yCoordinate
		if x1 != x2 && y1 != y2 {
			if x1 > x2 && y1 > y2 {
				for i := 0; i < x1-x2+1; i++ {
					_, exist := pointsCovered[point{x2 + i, y2 + i}]
					if exist {
						pointsCovered[point{x2 + i, y2 + i}] += 1
					} else {
						pointsCovered[point{x2 + i, y2 + i}] = 1
					}
				}
			} else if x1 > x2 && y2 > y1 {
				for i := 0; i < x1-x2+1; i++ {
					_, exist := pointsCovered[point{x1 - i, y1 + i}]
					if exist {
						pointsCovered[point{x1 - i, y1 + i}] += 1
					} else {
						pointsCovered[point{x1 - i, y1 + i}] = 1
					}
				}
			} else if x2 > x1 && y1 > y2 {
				for i := 0; i < x2-x1+1; i++ {
					_, exist := pointsCovered[point{x1 + i, y1 - i}]
					if exist {
						pointsCovered[point{x1 + i, y1 - i}] += 1
					} else {
						pointsCovered[point{x1 + i, y1 - i}] = 1
					}
				}
			} else {
				for i := 0; i < x2-x1+1; i++ {
					_, exist := pointsCovered[point{x1 + i, y1 + i}]
					if exist {
						pointsCovered[point{x1 + i, y1 + i}] += 1
					} else {
						pointsCovered[point{x1 + i, y1 + i}] = 1
					}
				}
			}
		}
	}
	return pointsCovered
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
