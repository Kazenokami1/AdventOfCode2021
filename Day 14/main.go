package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("Day14Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	compositeFormula := scanner.Text()
	scanner.Scan()
	insertFormulae := make(map[string]string)
	formulaNumbers := make(map[string]int)
	elementNumbers := make(map[string]int)
	for scanner.Scan() {
		insertFormula := strings.Split(scanner.Text(), " -> ")
		insertFormulae[insertFormula[0]] = insertFormula[1]
		formulaNumbers[insertFormula[0]] = 0
		elementNumbers[insertFormula[1]] = 0
	}
	answer := partOne(compositeFormula, insertFormulae, 10)
	fmt.Printf("The answer is: %d \n", answer)
	for i := 0; i < len(compositeFormula)-1; i++ {
		formulaNumbers[string(compositeFormula[i])+string(compositeFormula[i+1])]++
	}
	answer = partTwo(formulaNumbers, elementNumbers, insertFormulae, compositeFormula, 40)
	fmt.Printf("The answer for Part Two is: %d \n", answer)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(composite string, inserts map[string]string, steps int) int {
	elements := make(map[string]int)
	for i := 0; i < steps; i++ {
		composite = findComposite(composite, inserts)
	}
	for _, v := range composite {
		if _, ok := elements[string(v)]; ok {
			elements[string(v)]++
		} else {
			elements[string(v)] = 1
		}
	}
	minVal := 9999999
	maxVal := 0
	for _, val := range elements {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal - minVal
}

func findComposite(composite string, inserts map[string]string) string {
	newComposite := string(composite[0])
	for i := 0; i < len(composite)-1; i++ {
		insert := inserts[string(composite[i])+string(composite[i+1])]
		newComposite += insert + string(composite[i+1])
	}
	return newComposite
}

func partTwo(numbers, elementNumbers map[string]int, inserts map[string]string, composite string, steps int) int {
	firstElement := string(composite[0]) + string(composite[1])
	for i := 0; i < steps; i++ {
		numbers = hashElements(numbers, inserts)
		firstElement = string(firstElement[0]) + inserts[string(firstElement[0])+string(firstElement[1])]
	}
	finalElements := elementNumbers
	for k, v := range numbers {
		if k == firstElement {
			finalElements[string(k[0])]++
		}
		finalElements[string(k[1])] += v
	}
	minVal := 0
	maxVal := 0
	for _, val := range finalElements {
		if val < minVal || minVal == 0 {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal - minVal
}

func hashElements(numbers map[string]int, inserts map[string]string) map[string]int {
	elements := make(map[string]int)
	for k, v := range numbers {
		elements[k] = v
	}
	for k, v := range numbers {
		elements[k] -= v
		elements[string(k[0])+inserts[k]] += v
		elements[inserts[k]+string(k[1])] += v
	}
	return elements
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
