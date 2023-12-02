package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() string {
	input, err := os.Open("input.txt")
	if err != nil {
		return fmt.Sprintf("unable to read input file : %s", err)
	}
	defer input.Close()

	var total int

	// Better than doing Atoi as we know game ID are serial?
	var gameCounter int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		gameCounter++
		if gamePossible(scanner.Text(), 12, 13, 14) {
			total += gameCounter
		}
	}

	return fmt.Sprint("Answer to part one is ", total)
}

func partTwo() string {
	input, err := os.Open("input.txt")
	if err != nil {
		return fmt.Sprintf("unable to read input file : %s", err)
	}
	defer input.Close()

	var total int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		total += minSetPower(scanner.Text())
	}

	return fmt.Sprint("Answer to part one is ", total)
}

// gamePossible check the possiblity of the game based on the
// number of cubes vs cubes showed in a game.
func gamePossible(game string, maxRed int, maxGreen int, maxBlue int) bool {
	g := strings.Split(game, ":")
	rounds := strings.Split(g[1], ";")
	for _, round := range rounds {
		for _, cubes := range strings.Split(round, ",") {
			cube := strings.Fields(strings.TrimSpace(cubes))
			number, _ := strconv.Atoi(cube[0])
			switch cube[1] {
			case "red":
				if number > maxRed {
					return false
				}
			case "blue":
				if number > maxBlue {
					return false
				}
			case "green":
				if number > maxGreen {
					return false
				}
			}
		}
	}
	return true
}

// gamePossible check the possiblity of the game based on the
// number of cubes vs cubes showed in a game.
func minSetPower(game string) int {
	minRed := 0
	minGreen := 0
	minBlue := 0

	g := strings.Split(game, ":")
	rounds := strings.Split(g[1], ";")
	for _, round := range rounds {
		for _, cubes := range strings.Split(round, ",") {
			cube := strings.Fields(strings.TrimSpace(cubes))
			number, _ := strconv.Atoi(cube[0])
			switch cube[1] {
			case "red":
				if number > minRed {
					minRed = number
				}
			case "blue":
				if number > minBlue {
					minBlue = number
				}
			case "green":
				if number > minGreen {
					minGreen = number
				}
			}
		}
	}
	return minRed * minGreen * minBlue
}
