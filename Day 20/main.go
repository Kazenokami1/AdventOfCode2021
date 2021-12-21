package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	f, _ := os.Open("Day20Input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var enhancementAlgorithm string
	var image []string
	for scanner.Scan() {
		enhancementAlgorithm = scanner.Text()
		for scanner.Scan() {
			if scanner.Text() != "" {
				image = append(image, scanner.Text())
			}
		}
	}
	lightPixels := partOne(image, enhancementAlgorithm, 50)
	fmt.Printf("Number of Light Pixels: %d \n", lightPixels)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(image []string, algorithm string, number int) int {
	var numberOfLights int
	for i := 0; i < number; i++ {
		for k, v := range image {
			if i%2 == 0 {
				image[k] = ".." + v + ".."
			} else {
				image[k] = "##" + v + "##"
			}
		}
		darks := getAddString(image, i)
		image = append(image, "0", "0")
		copy(image[2:], image[0:])
		for j := 0; j < 2; j++ {
			image[j] = darks
		}
		image = append(image, darks, darks)
		image = runEnhance(image, algorithm)
	}
	for _, v := range image {
		for _, w := range v {
			if string(w) == "#" {
				numberOfLights++
			}
		}
	}
	return numberOfLights
}

func runEnhance(image []string, algorithm string) []string {
	var enhancedImage []string
	for i := 1; i < len(image)-1; i++ {
		var enhancedLine string
		for j := 1; j < len(image[0])-1; j++ {
			var binaryString string
			for k := -1; k < 2; k++ {
				for l := -1; l < 2; l++ {
					if string(image[i+k][j+l]) == "." {
						binaryString += "0"
					} else {
						binaryString += "1"
					}
				}
			}
			binaryInt, _ := strconv.ParseInt(binaryString, 2, 0)
			enhancedLine += string(algorithm[binaryInt])
		}
		enhancedImage = append(enhancedImage, enhancedLine)
	}
	return enhancedImage
}

func getAddString(image []string, i int) string {
	var addString string
	for j := 0; j < len(image[0]); j++ {
		if i%2 == 0 {
			addString += "."
		} else {
			addString += "#"
		}
	}
	return addString
}

func partTwo() {

}
