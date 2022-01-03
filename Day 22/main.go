package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, _ := os.Open("Day22Input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var instructions [][]string
	for scanner.Scan() {
		instructions = append(instructions, strings.Split(scanner.Text(), " "))
	}

	cubesOn := initializeReactor(instructions[0:10])
	fmt.Printf("Total Cubes Turned On: %d \n", cubesOn)

	cubesOn = fullReboot(instructions)
	fmt.Printf("Total Cubes Turn On Part 2: %d \n", cubesOn)

	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

type Cube struct {
	Position Vector
	Status   string
}

type Cuboid struct {
	XMin   int
	XMax   int
	YMin   int
	YMax   int
	ZMin   int
	ZMax   int
	Volume int
}

func checkOverlap(new Cuboid, old map[*Cuboid]struct{}, status string) map[*Cuboid]struct{} {
	overlaps := true
	for overlaps {
		overlaps = false
		for v := range old {
			if v.XMin >= new.XMin && v.XMax <= new.XMax && v.YMin >= new.YMin && v.YMax <= new.YMax && v.ZMin >= new.ZMin && v.ZMax <= new.ZMax {
				//Old Cuboid is Completely Contained in the new Cuboid
				delete(old, v)
			} else if new.XMin <= v.XMax && new.XMax >= v.XMin && new.YMin <= v.YMax && new.YMax >= v.YMin && new.ZMin <= v.ZMax && new.ZMax >= v.ZMin {
				//Cuboids overlap, split the old Cuboid based on the first condition it overlaps on, some overlap may still be present but we're going to loop through again until no overlaps occur
				var newCube Cuboid
				if new.XMin <= v.XMax && new.XMin > v.XMin && !(v.XMin >= new.XMin && v.XMax <= new.XMax) {
					newCube = Cuboid{XMin: new.XMin, XMax: v.XMax, YMin: v.YMin, YMax: v.YMax, ZMin: v.ZMin, ZMax: v.ZMax}
					newCube.getVolume()
					old[&newCube] = struct{}{}
					v.XMax = new.XMin - 1
				} else if new.XMax < v.XMax && new.XMax >= v.XMin && !(v.XMin >= new.XMin && v.XMax <= new.XMax) {
					newCube = Cuboid{XMin: v.XMin, XMax: new.XMax, YMin: v.YMin, YMax: v.YMax, ZMin: v.ZMin, ZMax: v.ZMax}
					newCube.getVolume()
					old[&newCube] = struct{}{}
					v.XMin = new.XMax + 1
				} else if new.YMin <= v.YMax && new.YMin > v.YMin && !(v.YMin >= new.YMin && v.YMax <= new.YMax) {
					newCube = Cuboid{XMin: v.XMin, XMax: v.XMax, YMin: new.YMin, YMax: v.YMax, ZMin: v.ZMin, ZMax: v.ZMax}
					newCube.getVolume()
					old[&newCube] = struct{}{}
					v.YMax = new.YMin - 1
				} else if new.YMax < v.YMax && new.YMax >= v.YMin && !(v.YMin >= new.YMin && v.YMax <= new.YMax) {
					newCube = Cuboid{XMin: v.XMin, XMax: v.XMax, YMin: v.YMin, YMax: new.YMax, ZMin: v.ZMin, ZMax: v.ZMax}
					newCube.getVolume()
					old[&newCube] = struct{}{}
					v.YMin = new.YMax + 1
				} else if new.ZMin <= v.ZMax && new.ZMin > v.ZMin && !(v.ZMin >= new.ZMin && v.ZMax <= new.ZMax) {
					newCube = Cuboid{XMin: v.XMin, XMax: v.XMax, YMin: v.YMin, YMax: v.YMax, ZMin: new.ZMin, ZMax: v.ZMax}
					newCube.getVolume()
					old[&newCube] = struct{}{}
					v.ZMax = new.ZMin - 1
				} else if new.ZMax < v.ZMax && new.ZMax >= v.ZMin && !(v.ZMin >= new.ZMin && v.ZMax <= new.ZMax) {
					newCube = Cuboid{XMin: v.XMin, XMax: v.XMax, YMin: v.YMin, YMax: v.YMax, ZMin: v.ZMin, ZMax: new.ZMax}
					newCube.getVolume()
					old[&newCube] = struct{}{}
					v.ZMin = new.ZMax + 1
				}
				v.getVolume()
				overlaps = true
			}
		}
	}
	//If the status is "on" add the new Cuboid to our list.  If it's "off", we've already deleted the overlap so no need to do anything more
	if status == "on" {
		new.getVolume()
		old[&new] = struct{}{}
	}
	return old
}

func getTotalCubesLit(cuboids map[*Cuboid]struct{}) int {
	var totalCubesOn int
	for k := range cuboids {
		totalCubesOn += k.Volume
	}
	return totalCubesOn
}

func (a *Cuboid) getVolume() {
	a.Volume = (a.XMax - a.XMin + 1) * (a.YMax - a.YMin + 1) * (a.ZMax - a.ZMin + 1)
}

type Vector struct {
	X int
	Y int
	Z int
}

func fullReboot(instructions [][]string) int {
	existingCuboids := make(map[*Cuboid]struct{})
	for _, v := range instructions {
		regSplit := regexp.MustCompile("[xyz=,.]")
		cubeValues := regSplit.Split(v[1], -1)
		x1, _ := strconv.Atoi(cubeValues[2])
		x2, _ := strconv.Atoi(cubeValues[4])
		y1, _ := strconv.Atoi(cubeValues[7])
		y2, _ := strconv.Atoi(cubeValues[9])
		z1, _ := strconv.Atoi(cubeValues[12])
		z2, _ := strconv.Atoi(cubeValues[14])
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		if z1 > z2 {
			z1, z2 = z2, z1
		}
		newCuboid := Cuboid{XMin: x1, XMax: x2, YMin: y1, YMax: y2, ZMin: z1, ZMax: z2}
		existingCuboids = checkOverlap(newCuboid, existingCuboids, v[0])
	}
	return getTotalCubesLit(existingCuboids)
}

func initializeReactor(instructions [][]string) int {
	var cubesOn int
	cubes := make(map[Vector]string)
	for _, v := range instructions {
		regSplit := regexp.MustCompile("[xyz=,.]")
		cubeValues := regSplit.Split(v[1], -1)
		x1, _ := strconv.Atoi(cubeValues[2])
		x2, _ := strconv.Atoi(cubeValues[4])
		y1, _ := strconv.Atoi(cubeValues[7])
		y2, _ := strconv.Atoi(cubeValues[9])
		z1, _ := strconv.Atoi(cubeValues[12])
		z2, _ := strconv.Atoi(cubeValues[14])
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		if z1 > z2 {
			z1, z2 = z2, z1
		}
		for x := x1; x < x2+1; x++ {
			for y := y1; y < y2+1; y++ {
				for z := z1; z < z2+1; z++ {
					cubes[Vector{X: x, Y: y, Z: z}] = v[0]
				}
			}
		}
	}
	for _, v := range cubes {
		if v == "on" {
			cubesOn++
		}
	}
	return cubesOn
}
