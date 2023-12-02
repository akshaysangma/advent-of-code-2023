package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var conversions = []string{
	"one", "o1e",
	"two", "t2o",
	"three", "t3e",
	"four", "f4r",
	"five", "f5e",
	"six", "s6x",
	"seven", "s7n",
	"eight", "e8t",
	"nine", "n9e",
	"zero", "z0o",
}

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

	re := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(input)
	var total int

	for scanner.Scan() {
		if digits := re.FindAllString(scanner.Text(), -1); digits != nil {
			twoDigit := digits[0] + digits[len(digits)-1]

			// Can ignore error as regex guarantees digits are ints
			num, _ := strconv.Atoi(twoDigit)
			total += num
		}
	}

	return fmt.Sprint("Answer to part one is ", total)
}

// sadly (?=(one|two|three|four|five|six|seven|eight|nine|zero|\d)) povitive lookahead is not supported
// in go. Hence unable to handle overlapping kinds like oneight (1,8). Refer https://github.com/google/re2/wiki/Syntax
func partTwo() string {
	input, err := os.Open("input.txt")
	if err != nil {
		return fmt.Sprintf("unable to read input file : %s", err)
	}
	defer input.Close()

	re := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(input)
	var total int

	for scanner.Scan() {
		parser := strings.NewReplacer(conversions...)

		// Since string.NewReplacer performs non-overlapping replacement
		// performing replacement twice will be able to handle overlapping
		// conversion from digit  word to digit as long as we maintain
		// first and last char of the digit word to enable secord replacement.
		parsed := parser.Replace(parser.Replace(scanner.Text()))

		if digits := re.FindAllString(parsed, -1); digits != nil {
			twoDigit := digits[0] + digits[len(digits)-1]
			// Can ignore error as regex guarantees digits are ints
			num, _ := strconv.Atoi(twoDigit)
			total += num
		}
	}

	return fmt.Sprint("Answer to part two is ", total)
}
