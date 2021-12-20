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

type Probe struct {
	ProbeNo  int
	Position Vector
	Beacons  []*Beacon
	Solved   bool
}

type Beacon struct {
	Position        Vector
	VectorToBeacons []Vector
}

type Vector struct {
	X float64
	Y float64
	Z float64
}

func main() {
	start := time.Now()
	f, _ := os.Open("Day19Input.txt")
	defer f.Close()
	var probeNo int
	scanner := bufio.NewScanner(f)
	var probes []*Probe
	for scanner.Scan() {
		var probe Probe
		if scanner.Text()[0] == '-' {
			probe = Probe{ProbeNo: probeNo}
			for scanner.Scan() {
				if scanner.Text() != "" {
					vector := strings.Split(scanner.Text(), ",")
					x, _ := strconv.Atoi(vector[0])
					y, _ := strconv.Atoi(vector[1])
					z, _ := strconv.Atoi(vector[2])
					probe.Beacons = append(probe.Beacons, &Beacon{Position: Vector{X: float64(x), Y: float64(y), Z: float64(z)}})
					if probeNo == 0 {
						probe.Solved = true
					}
				} else {
					probeNo++
					break
				}
			}
			probes = append(probes, &probe)
			for _, v := range probe.Beacons {
				for _, w := range probe.Beacons {
					if !reflect.DeepEqual(v.Position, w.Position) {
						v.VectorToBeacons = append(v.VectorToBeacons, Vector{v.Position.X - w.Position.X, v.Position.Y - w.Position.Y, v.Position.Z - w.Position.Z})
					}
				}
			}
		}
	}
	totalBeacons := partOne(probes)
	fmt.Printf("Number of Beacons: %d \n", totalBeacons)
	manhattanDistance := partTwo(probes)
	fmt.Printf("Furthest Manhattan Distance: %f \n", manhattanDistance)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(probes []*Probe) int {
	uniqueBeacons := make(map[Vector]struct{})
	allSolved := false
	for !allSolved {
		allSolved = true
		for _, v := range probes {
			for _, w := range probes {
				if v.Solved && !w.Solved {
					w.Solved = checkOverlap(v, w)
				}
			}
		}
		for _, v := range probes {
			if !v.Solved {
				allSolved = false
			}
		}
	}
	for _, v := range probes {
		for _, w := range v.Beacons {
			if _, ok := uniqueBeacons[w.Position]; !ok {
				uniqueBeacons[w.Position] = struct{}{}
			}
		}
	}
	return len(uniqueBeacons)
}

func checkOverlap(one, two *Probe) bool {
	var sharedBeacons int
	var beaconOne *Beacon
	var beaconTwo *Beacon
	for _, v := range one.Beacons {
		for _, w := range two.Beacons {
			var shared int
			for _, x := range v.VectorToBeacons {
				for _, y := range w.VectorToBeacons {
					if vectorMagnitude(x) == vectorMagnitude(y) {
						shared++
						break
					}
				}
			}
			if shared >= 11 {
				beaconOne = v
				beaconTwo = w
				sharedBeacons++
			}
		}
	}
	if sharedBeacons >= 11 {
		solveProbe(one, two, beaconOne, beaconTwo)
		return true
	}
	return false
}

func solveProbe(solved, unsolved *Probe, beaconOne, beaconTwo *Beacon) {
	for _, x := range beaconOne.VectorToBeacons {
		for _, y := range beaconTwo.VectorToBeacons {
			if vectorMagnitude(x) == vectorMagnitude(y) {
				//Calculate change in position for Scanner/Beacons
				//Possible Orientations:
				var newVector Vector
				if x.X-y.X == 0 && x.Y-y.Y == 0 && x.Z-y.Z == 0 {
					//Orientation is xyz
					differenceVector := beaconOne.Position.Subtract(beaconTwo.Position)
					for _, b := range unsolved.Beacons {
						b.Position = b.Position.Add(differenceVector)
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Y == 0 && x.Y+y.X == 0 && x.Z-y.Z == 0 {
					//Orientation is y-xz
					newVector.X = beaconTwo.Position.Y
					newVector.Y = -beaconTwo.Position.X
					newVector.Z = beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Y + differenceVector.X
						newVector.Y = -b.Position.X + differenceVector.Y
						newVector.Z = b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Y
							newVector.Y = -magVector.X
							newVector.Z = magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.X == 0 && x.Y+y.Y == 0 && x.Z-y.Z == 0 {
					//Orientation is -x-yz
					newVector.X = -beaconTwo.Position.X
					newVector.Y = -beaconTwo.Position.Y
					newVector.Z = beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.X + differenceVector.X
						newVector.Y = -b.Position.Y + differenceVector.Y
						newVector.Z = b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.X
							newVector.Y = -magVector.Y
							newVector.Z = magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Y == 0 && x.Y-y.X == 0 && x.Z-y.Z == 0 {
					//Orientation is -yxz
					newVector.X = -beaconTwo.Position.Y
					newVector.Y = beaconTwo.Position.X
					newVector.Z = beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Y + differenceVector.X
						newVector.Y = b.Position.X + differenceVector.Y
						newVector.Z = b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Y
							newVector.Y = magVector.X
							newVector.Z = magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.X == 0 && x.Y+y.Y == 0 && x.Z+y.Z == 0 {
					//Orientation is x-y-z
					newVector.X = beaconTwo.Position.X
					newVector.Y = -beaconTwo.Position.Y
					newVector.Z = -beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.X + differenceVector.X
						newVector.Y = -b.Position.Y + differenceVector.Y
						newVector.Z = -b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.X
							newVector.Y = -magVector.Y
							newVector.Z = -magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Y == 0 && x.Y-y.X == 0 && x.Z+y.Z == 0 {
					//Orientation is yx-z
					newVector.X = beaconTwo.Position.Y
					newVector.Y = beaconTwo.Position.X
					newVector.Z = -beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Y + differenceVector.X
						newVector.Y = b.Position.X + differenceVector.Y
						newVector.Z = -b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Y
							newVector.Y = magVector.X
							newVector.Z = -magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.X == 0 && x.Y-y.Y == 0 && x.Z+y.Z == 0 {
					//Orientation is -xy-z
					newVector.X = -beaconTwo.Position.X
					newVector.Y = beaconTwo.Position.Y
					newVector.Z = -beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.X + differenceVector.X
						newVector.Y = b.Position.Y + differenceVector.Y
						newVector.Z = -b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.X
							newVector.Y = magVector.Y
							newVector.Z = -magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Y == 0 && x.Y+y.X == 0 && x.Z+y.Z == 0 {
					//Orientation is y-x-z
					newVector.X = beaconTwo.Position.Y
					newVector.Y = -beaconTwo.Position.X
					newVector.Z = -beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Y + differenceVector.X
						newVector.Y = -b.Position.X + differenceVector.Y
						newVector.Z = -b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Y
							newVector.Y = -magVector.X
							newVector.Z = -magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Z == 0 && x.Y-y.Y == 0 && x.Z-y.X == 0 {
					//Orientation is -zyx
					newVector.X = -beaconTwo.Position.Z
					newVector.Y = beaconTwo.Position.Y
					newVector.Z = beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Z + differenceVector.X
						newVector.Y = b.Position.Y + differenceVector.Y
						newVector.Z = b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Z
							newVector.Y = magVector.Y
							newVector.Z = magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Z == 0 && x.Y-y.X == 0 && x.Z+y.Y == 0 {
					//Orientation is -zx-y
					newVector.X = -beaconTwo.Position.Z
					newVector.Y = beaconTwo.Position.X
					newVector.Z = -beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Z + differenceVector.X
						newVector.Y = b.Position.X + differenceVector.Y
						newVector.Z = -b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Z
							newVector.Y = magVector.X
							newVector.Z = -magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Z == 0 && x.Y+y.Y == 0 && x.Z+y.X == 0 {
					//Orientation is -z-y-x
					newVector.X = -beaconTwo.Position.Z
					newVector.Y = -beaconTwo.Position.Y
					newVector.Z = -beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Z + differenceVector.X
						newVector.Y = -b.Position.Y + differenceVector.Y
						newVector.Z = -b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Z
							newVector.Y = -magVector.Y
							newVector.Z = -magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Z == 0 && x.Y+y.X == 0 && x.Z-y.Y == 0 {
					//Orientation is -z-xy
					newVector.X = -beaconTwo.Position.Z
					newVector.Y = -beaconTwo.Position.X
					newVector.Z = beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Z + differenceVector.X
						newVector.Y = -b.Position.X + differenceVector.Y
						newVector.Z = b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Z
							newVector.Y = -magVector.X
							newVector.Z = magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Z == 0 && x.Y+y.Y == 0 && x.Z-y.X == 0 {
					//Orientation is z-yx
					newVector.X = beaconTwo.Position.Z
					newVector.Y = -beaconTwo.Position.Y
					newVector.Z = beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Z + differenceVector.X
						newVector.Y = -b.Position.Y + differenceVector.Y
						newVector.Z = b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Z
							newVector.Y = -magVector.Y
							newVector.Z = magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Z == 0 && x.Y-y.X == 0 && x.Z-y.Y == 0 {
					//Orientation is zxy
					newVector.X = beaconTwo.Position.Z
					newVector.Y = beaconTwo.Position.X
					newVector.Z = beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Z + differenceVector.X
						newVector.Y = b.Position.X + differenceVector.Y
						newVector.Z = b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Z
							newVector.Y = magVector.X
							newVector.Z = magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Z == 0 && x.Y-y.Y == 0 && x.Z+y.X == 0 {
					//Orientation is zy-x
					newVector.X = beaconTwo.Position.Z
					newVector.Y = beaconTwo.Position.Y
					newVector.Z = -beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Z + differenceVector.X
						newVector.Y = b.Position.Y + differenceVector.Y
						newVector.Z = -b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Z
							newVector.Y = magVector.Y
							newVector.Z = -magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Z == 0 && x.Y+y.X == 0 && x.Z+y.Y == 0 {
					//Orientation is z-x-y
					newVector.X = beaconTwo.Position.Z
					newVector.Y = -beaconTwo.Position.X
					newVector.Z = -beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Z + differenceVector.X
						newVector.Y = -b.Position.X + differenceVector.Y
						newVector.Z = -b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Z
							newVector.Y = -magVector.X
							newVector.Z = -magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Y == 0 && x.Y+y.Z == 0 && x.Z-y.X == 0 {
					//Orientation is -y-zx
					newVector.X = -beaconTwo.Position.Y
					newVector.Y = -beaconTwo.Position.Z
					newVector.Z = beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Y + differenceVector.X
						newVector.Y = -b.Position.Z + differenceVector.Y
						newVector.Z = b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Y
							newVector.Y = -magVector.Z
							newVector.Z = magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.X == 0 && x.Y+y.Z == 0 && x.Z+y.Y == 0 {
					//Orientation is -x-z-y
					newVector.X = -beaconTwo.Position.X
					newVector.Y = -beaconTwo.Position.Z
					newVector.Z = -beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.X + differenceVector.X
						newVector.Y = -b.Position.Z + differenceVector.Y
						newVector.Z = -b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.X
							newVector.Y = -magVector.Z
							newVector.Z = -magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Y == 0 && x.Y+y.Z == 0 && x.Z+y.X == 0 {
					//Orientation is y-z-x
					newVector.X = beaconTwo.Position.Y
					newVector.Y = -beaconTwo.Position.Z
					newVector.Z = -beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Y + differenceVector.X
						newVector.Y = -b.Position.Z + differenceVector.Y
						newVector.Z = -b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Y
							newVector.Y = -magVector.Z
							newVector.Z = -magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.X == 0 && x.Y+y.Z == 0 && x.Z-y.Y == 0 {
					//Orientation is x-zy
					newVector.X = beaconTwo.Position.X
					newVector.Y = -beaconTwo.Position.Z
					newVector.Z = beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.X + differenceVector.X
						newVector.Y = -b.Position.Z + differenceVector.Y
						newVector.Z = b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.X
							newVector.Y = -magVector.Z
							newVector.Z = magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.Y == 0 && x.Y-y.Z == 0 && x.Z-y.X == 0 {
					//Orientation is yzx
					newVector.X = beaconTwo.Position.Y
					newVector.Y = beaconTwo.Position.Z
					newVector.Z = beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.Y + differenceVector.X
						newVector.Y = b.Position.Z + differenceVector.Y
						newVector.Z = b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.Y
							newVector.Y = magVector.Z
							newVector.Z = magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.X == 0 && x.Y-y.Z == 0 && x.Z-y.Y == 0 {
					//Orientation is -xzy
					newVector.X = -beaconTwo.Position.X
					newVector.Y = beaconTwo.Position.Z
					newVector.Z = beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.X + differenceVector.X
						newVector.Y = b.Position.Z + differenceVector.Y
						newVector.Z = b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.X
							newVector.Y = magVector.Z
							newVector.Z = magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Y == 0 && x.Y-y.Z == 0 && x.Z+y.X == 0 {
					//Orientation is -yz-x
					newVector.X = -beaconTwo.Position.Y
					newVector.Y = beaconTwo.Position.Z
					newVector.Z = -beaconTwo.Position.X
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Y + differenceVector.X
						newVector.Y = b.Position.Z + differenceVector.Y
						newVector.Z = -b.Position.X + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Y
							newVector.Y = magVector.Z
							newVector.Z = -magVector.X
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X-y.X == 0 && x.Y-y.Z == 0 && x.Z+y.Y == 0 {
					//Orientation is xz-y
					newVector.X = beaconTwo.Position.X
					newVector.Y = beaconTwo.Position.Z
					newVector.Z = -beaconTwo.Position.Y
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = b.Position.X + differenceVector.X
						newVector.Y = b.Position.Z + differenceVector.Y
						newVector.Z = -b.Position.Y + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = magVector.X
							newVector.Y = magVector.Z
							newVector.Z = -magVector.Y
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				} else if x.X+y.Y == 0 && x.Y+y.X == 0 && x.Z+y.Z == 0 {
					//Orientation is -y-x-z
					newVector.X = -beaconTwo.Position.Y
					newVector.Y = -beaconTwo.Position.X
					newVector.Z = -beaconTwo.Position.Z
					differenceVector := beaconOne.Position.Subtract(newVector)
					for _, b := range unsolved.Beacons {
						newVector.X = -b.Position.Y + differenceVector.X
						newVector.Y = -b.Position.X + differenceVector.Y
						newVector.Z = -b.Position.Z + differenceVector.Z
						b.Position = newVector
						for i, magVector := range b.VectorToBeacons {
							newVector.X = -magVector.Y
							newVector.Y = -magVector.X
							newVector.Z = -magVector.Z
							b.VectorToBeacons[i] = newVector
						}
					}
					unsolved.Position = differenceVector
					return
				}
			}
		}
	}
}

func partTwo(probes []*Probe) float64 {
	var maxManhattanDistance float64
	for _, v := range probes {
		for _, w := range probes {
			manhattanDistance := math.Abs(v.Position.X-w.Position.X) + math.Abs(v.Position.Y-w.Position.Y) + math.Abs(v.Position.Z-w.Position.Z)
			if manhattanDistance > maxManhattanDistance {
				maxManhattanDistance = manhattanDistance
			}
		}
	}
	return maxManhattanDistance
}

func vectorMagnitude(v Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (a Vector) Subtract(b Vector) Vector {
	return Vector{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}
