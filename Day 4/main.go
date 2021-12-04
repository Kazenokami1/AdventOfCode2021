package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type bingoCardString struct {
	Rows    [][]string
	Columns [][]string
}

type bingoCardNumber struct {
	Rows     [][]int
	Columns  [][]int
	Winner   bool
	SetIndex int
}

var numberCards []bingoCardNumber

func main() {
	start := time.Now()
	f, err := os.Open("day4Input.txt")
	check(err)
	defer f.Close()
	var cards []bingoCardString
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	calledStrings := strings.Split(scanner.Text(), ",")
	var calledNumbers []int
	for _, value := range calledStrings {
		number, err := strconv.Atoi(value)
		check(err)
		calledNumbers = append(calledNumbers, number)
	}
	scanner.Scan()
	for scanner.Scan() {
		var card bingoCardString
		for i := 0; i < 5; i++ {
			row := strings.Trim(scanner.Text(), " ")
			rowTwo := strings.ReplaceAll(row, "  ", " ")
			rowThree := strings.Split(rowTwo, " ")
			card.Rows = append(card.Rows, rowThree)
			scanner.Scan()
		}
		for i := 0; i < 5; i++ {
			var column []string
			for j := 0; j < 5; j++ {
				column = append(column, card.Rows[j][i])
			}
			card.Columns = append(card.Columns, column)
		}
		cards = append(cards, card)
	}
	numberCards = turnCardsToNumbers(cards)
	partOne(calledNumbers, numberCards)
	partTwo(calledNumbers, numberCards)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(numbers []int, bingoCards []bingoCardNumber) int {
	var calledNumbers []int
	for index, value := range numbers {
		var bingo bool
		if index < 4 {
			calledNumbers = append(calledNumbers, value)
		} else {
			calledNumbers = append(calledNumbers, value)
			for _, card := range bingoCards {
				for _, row := range card.Rows {
					bingo = subset(calledNumbers, row)
					if bingo == true {
						sumUncalledNumbers := getUncalledNumbers(calledNumbers, card)
						fmt.Println(sumUncalledNumbers * value)
						return sumUncalledNumbers
					}
				}
				for _, column := range card.Columns {
					bingo = subset(calledNumbers, column)
					if bingo == true {
						sumUncalledNumbers := getUncalledNumbers(calledNumbers, card)
						fmt.Println(sumUncalledNumbers * value)
						return sumUncalledNumbers
					}
				}
			}
		}
	}
	return 0
}

func partTwo(numbers []int, bingoCards []bingoCardNumber) int {
	var calledNumbers []int
	var lastWinningCard bingoCardNumber
	var lastCalledNumber int
	set := make(map[int]bingoCardNumber)
	for index, card := range bingoCards {
		set[index] = card
	}
	for index, value := range numbers {
		var bingo bool
		if index < 4 {
			calledNumbers = append(calledNumbers, value)
		} else if len(set) > 0 {
			calledNumbers = append(calledNumbers, value)
			for _, card := range bingoCards {
				if _, ok := set[card.SetIndex]; ok {
					card := card
					for _, row := range card.Rows {
						bingo = subset(calledNumbers, row)
						if bingo == true {
							card.Winner = true
							lastWinningCard = card
							lastCalledNumber = value
							delete(set, card.SetIndex)
						}
					}
					for _, column := range card.Columns {
						bingo = subset(calledNumbers, column)
						if bingo == true {
							card.Winner = true
							lastWinningCard = card
							lastCalledNumber = value
							delete(set, card.SetIndex)
						}
					}
				}
			}
		}
	}
	sumUncalledNumbers := getUncalledNumbers(calledNumbers, lastWinningCard)
	fmt.Println(sumUncalledNumbers * lastCalledNumber)
	return 0
}

func getUncalledNumbers(calledNumbers []int, card bingoCardNumber) int {
	sum := 0
	for _, value := range card.Rows {
		for _, valueTwo := range value {
			if !subsetSingle(calledNumbers, valueTwo) {
				sum += valueTwo
			}
		}
	}
	return sum
}

func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range first {
		set[value] += 1
	}

	for _, value := range second {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}
	return true
}

func subsetSingle(first []int, second int) bool {
	set := make(map[int]int)
	for _, value := range first {
		set[value] += 1
	}

	if count, found := set[second]; !found {
		return false
	} else if count < 1 {
		return false
	} else {
		set[second] = count - 1
	}
	return true
}

func turnCardsToNumbers(cards []bingoCardString) []bingoCardNumber {
	var numberCards []bingoCardNumber
	for index, stringCard := range cards {
		var card bingoCardNumber
		for _, stringRow := range stringCard.Rows {
			var numbers []int
			for _, numberString := range stringRow {
				number, err := strconv.Atoi(numberString)
				check(err)
				numbers = append(numbers, number)
			}
			card.Rows = append(card.Rows, numbers)
		}
		for _, stringCol := range stringCard.Columns {
			var numbers []int
			for _, numberString := range stringCol {
				number, err := strconv.Atoi(numberString)
				check(err)
				numbers = append(numbers, number)
			}
			card.Columns = append(card.Columns, numbers)
			card.Winner = false
			card.SetIndex = index
		}
		numberCards = append(numberCards, card)
	}
	return numberCards
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
