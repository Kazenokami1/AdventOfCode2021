package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	//Sample Input "target area: x=20..30, y=-10..-5"
	//Input "target area: x=248..285, y=-85..-56"
	targetArea := make(map[string]int)
	targetArea["xMin"] = 248
	targetArea["xMax"] = 285
	targetArea["yMin"] = -85
	targetArea["yMax"] = -56
	yMax, velocities := partOne(targetArea)
	fmt.Printf("Highest y value reached: %d \n", yMax)
	fmt.Printf("Number of Distinct Velocities: %d \n", velocities)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(targetArea map[string]int) (int, int) {
	yMax := targetArea["yMin"]
	minXReached := false
	minXVelocity := 1
	xVelocity := minXVelocity
	velocities := 0
	for !minXReached {
		xPosition := 0
		for x := minXVelocity; x > 0; x-- {
			xPosition += x
			if xVelocity >= 0 && xPosition >= targetArea["xMin"] {
				minXReached = true
			}
		}
		if !minXReached {
			minXVelocity++
		}
	}
	maxYVelocity := -targetArea["yMin"]
	for x := minXVelocity; x < targetArea["xMax"]+1; x++ {
		for y := targetArea["yMin"]; y < maxYVelocity; y++ {
			yHit, inTarget := hitTarget(x, y, targetArea)
			if yHit > yMax && inTarget {
				yMax = yHit
				velocities++
			} else if inTarget {
				velocities++
			}
		}
	}
	return yMax, velocities
}

func hitTarget(x, y int, target map[string]int) (int, bool) {
	targetAvailable := true
	targetHit := false
	xCurrent := 0
	yCurrent := 0
	yMax := target["yMin"]
	for targetAvailable {
		xCurrent += x
		yCurrent += y
		if xCurrent <= target["xMax"] && xCurrent >= target["xMin"] && yCurrent <= target["yMax"] && yCurrent >= target["yMin"] {
			targetHit = true
			targetAvailable = false
		} else if xCurrent > target["xMax"] || yCurrent < target["yMin"] {
			targetAvailable = false
		} else {
			if yCurrent > yMax {
				yMax = yCurrent
			}
			y -= 1
			if x > 0 {
				x -= 1
			} else if x < 0 {
				x += 1
			}
		}
	}
	return yMax, targetHit
}

func partTwo(depths []int) {

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
