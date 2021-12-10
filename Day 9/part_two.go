package main

import "sort"

type cavePoint struct {
	xCoordinate    int
	yCoordinate    int
	height         int
	connectedBasin []cavePoint
}

func partTwo(heights [][]int) int {
	var allCavePoints []cavePoint
	basins := make(map[int][]cavePoint)
	for i := range heights {
		for j, v := range heights[i] {
			cPoint := cavePoint{yCoordinate: i, xCoordinate: j, height: v}
			if i == 0 && v != 9 {
				if j == 0 {
					if heights[i][j+1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j + 1, height: heights[i][j+1]})
					}
					if heights[i+1][j] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i + 1, xCoordinate: j, height: heights[i+1][j]})
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j-1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j - 1, height: heights[i][j-1]})
					}
					if heights[i+1][j] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i + 1, xCoordinate: j, height: heights[i+1][j]})
					}
				} else {
					if heights[i][j-1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j - 1, height: heights[i][j-1]})
					}
					if heights[i][j+1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j + 1, height: heights[i][j+1]})
					}
					if heights[i+1][j] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i + 1, xCoordinate: j, height: heights[i+1][j]})
					}
				}
			} else if i == len(heights)-1 && v != 9 {
				if j == 0 {
					if heights[i-1][j] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i - 1, xCoordinate: j, height: heights[i-1][j]})
					}
					if heights[i][j+1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j + 1, height: heights[i][j+1]})
					}
				} else if j == len(heights[i])-1 {
					if heights[i][j-1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j - 1, height: heights[i][j-1]})
					}
					if heights[i-1][j] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i - 1, xCoordinate: j, height: heights[i-1][j]})
					}
				} else {
					if heights[i-1][j] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i - 1, xCoordinate: j, height: heights[i-1][j]})
					}
					if heights[i][j+1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j + 1, height: heights[i][j+1]})
					}
					if heights[i][j-1] != 9 {
						cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j - 1, height: heights[i][j-1]})
					}
				}
			} else if j == 0 && v != 9 {
				if heights[i-1][j] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i - 1, xCoordinate: j, height: heights[i-1][j]})
				}
				if heights[i+1][j] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i + 1, xCoordinate: j, height: heights[i+1][j]})
				}
				if heights[i][j+1] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j + 1, height: heights[i][j+1]})
				}
			} else if j == len(heights[i])-1 && v != 9 {
				if heights[i-1][j] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i - 1, xCoordinate: j, height: heights[i-1][j]})
				}
				if heights[i+1][j] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i + 1, xCoordinate: j, height: heights[i+1][j]})
				}
				if heights[i][j-1] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j - 1, height: heights[i][j-1]})
				}
			} else if v != 9 {
				if heights[i-1][j] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i - 1, xCoordinate: j, height: heights[i-1][j]})
				}
				if heights[i+1][j] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i + 1, xCoordinate: j, height: heights[i+1][j]})
				}
				if heights[i][j-1] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j - 1, height: heights[i][j-1]})
				}
				if heights[i][j+1] != 9 {
					cPoint.connectedBasin = append(cPoint.connectedBasin, cavePoint{yCoordinate: i, xCoordinate: j + 1, height: heights[i][j+1]})
				}
			}
			if cPoint.height != 9 {
				allCavePoints = append(allCavePoints, cPoint)
			}
		}
	}
	numberOfBasins := 0
	for _, v := range allCavePoints {
		basinFound := false
		var foundBasin int
		for i := range basins {
			for _, w := range v.connectedBasin {
				for _, x := range basins[i] {
					if w.xCoordinate == x.xCoordinate && w.yCoordinate == x.yCoordinate {
						if !basinFound {
							basins[i] = append(basins[i], v)
							basinFound = true
							foundBasin = i
						} else if i != foundBasin {
							for _, y := range basins[i] {
								basins[foundBasin] = append(basins[foundBasin], y)
							}
							delete(basins, i)
						}
					}
				}
			}
		}
		if !basinFound {
			basins[numberOfBasins] = append(basins[numberOfBasins], v)
			numberOfBasins++
		}
	}
	var unsortedBasins []int
	for i := range basins {
		unsortedBasins = append(unsortedBasins, len(basins[i]))
	}
	sort.Ints(unsortedBasins)
	returnValue := 1
	for i := len(unsortedBasins) - 1; i > len(unsortedBasins)-4; i-- {
		returnValue *= unsortedBasins[i]
	}
	return returnValue
}
