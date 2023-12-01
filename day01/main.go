package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(partOne())
}

func partOne() string {
	input1, err := os.Open("input1.txt")
	if err != nil {
		return fmt.Sprintf("unable to read input file : %s", err)
	}

	re := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(input1)
	var total int

	for scanner.Scan() {
		if digits := re.FindAllString(scanner.Text(), -1); digits != nil {
			twoDigit := digits[0] + digits[len(digits)-1]

			// Can ignore error as regex guarantees digits are ints
			num, _ := strconv.Atoi(twoDigit)
			total += num
		}
	}

	return fmt.Sprintln("Answer to part one is ", total)
}
