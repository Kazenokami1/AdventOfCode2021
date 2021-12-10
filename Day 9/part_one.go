package main

func partOne(heights [][]int) int {
	var lowPoints []int
	risk := 0
	for i := range heights {
		for j, v := range heights[i] {
			if i == 0 && v != 9 {
				if j == 0 {
					if v < heights[i][j+1] && v < heights[i+1][j] {
						lowPoints = append(lowPoints, v)
					}
				} else if j == len(heights[i])-1 {
					if v < heights[i][j-1] && v < heights[i+1][j] {
						lowPoints = append(lowPoints, v)
					}
				} else {
					if v < heights[i][j-1] && v < heights[i][j+1] && v < heights[i+1][j] {
						lowPoints = append(lowPoints, v)
					}
				}
			} else if i == len(heights)-1 && v != 9 {
				if j == 0 {
					if v < heights[i-1][j] && v < heights[i][j+1] {
						lowPoints = append(lowPoints, v)
					}
				} else if j == len(heights[i])-1 {
					if v < heights[i][j-1] && v < heights[i-1][j] {
						lowPoints = append(lowPoints, v)
					}
				} else {
					if v < heights[i-1][j] && v < heights[i][j+1] && v < heights[i][j-1] {
						lowPoints = append(lowPoints, v)
					}
				}
			} else if j == 0 && v != 9 {
				if v < heights[i-1][j] && v < heights[i+1][j] && v < heights[i][j+1] {
					lowPoints = append(lowPoints, v)
				}
			} else if j == len(heights[i])-1 && v != 9 {
				if v < heights[i-1][j] && v < heights[i+1][j] && v < heights[i][j-1] {
					lowPoints = append(lowPoints, v)
				}
			} else if v != 9 {
				if v < heights[i-1][j] && v < heights[i+1][j] && v < heights[i][j-1] && v < heights[i][j+1] {
					lowPoints = append(lowPoints, v)
				}
			}
		}
	}
	for _, v := range lowPoints {
		risk += v + 1
	}
	return risk
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
