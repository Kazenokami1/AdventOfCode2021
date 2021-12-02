package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type command struct {
	Direction string
	Units     int
}

func main() {
	start := time.Now()
	f, err := os.Open("day2Input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var commands []command
	for scanner.Scan() {
		commandSplit := strings.Fields(scanner.Text())
		var newCommand command
		newCommand.Direction = commandSplit[0]
		newCommand.Units, err = strconv.Atoi(commandSplit[1])
		check(err)
		commands = append(commands, newCommand)
	}
	forwardPosition, depth := partOne(commands)
	text := fmt.Sprintf("Forward Position: %d. Depth: %d. Multiplied: %d", forwardPosition, depth, forwardPosition*depth)
	fmt.Println(text)
	forwardPosition, depth = partTwo(commands)
	text = fmt.Sprintf("Forward Position: %d. Depth: %d. Multiplied: %d", forwardPosition, depth, forwardPosition*depth)
	fmt.Println(text)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func partOne(commands []command) (int, int) {
	forwardPosition := 0
	depth := 0
	for _, command := range commands {
		switch command.Direction {
		case "forward":
			forwardPosition += command.Units
		case "down":
			depth += command.Units
		case "up":
			depth -= command.Units
		}
	}
	return forwardPosition, depth
}

func partTwo(commands []command) (int, int) {
	forwardPosition := 0
	depth := 0
	aim := 0
	for _, command := range commands {
		switch command.Direction {
		case "forward":
			forwardPosition += command.Units
			depth += aim * command.Units
		case "down":
			aim += command.Units
		case "up":
			aim -= command.Units
		}
	}
	return forwardPosition, depth
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
