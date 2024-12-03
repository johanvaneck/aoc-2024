package solutions

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Day03Part01() int {
	file, err := os.Open("inputs/day-03.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += parseLine01(line)
	}

	return sum
}

func Day03Part02() int {
	file, err := os.Open("inputs/day-03.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	isActivated := true

	for scanner.Scan() {
		line := scanner.Text()
		var lineSum int
		lineSum, isActivated = parseLine02(line, isActivated)
		sum += lineSum
	}

	return sum
}

func parseLine01(line string) int {
	sum := 0
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}

	return sum
}

func searchActivate(line string) bool {
	r := regexp.MustCompile(`do\(\)`)
	return r.MatchString(line)

}

func searchDeactivate(line string) bool {
	r := regexp.MustCompile(`don't\(\)`)
	return r.MatchString(line)

}

func parseLine02(line string, isActivated bool) (int, bool) {
	sum := 0

	// do()         - 4
	// don't()      - 7
	// mul(#,#)     - 8
	// mul(###,###) - 12
	maxLength := 12
	lineLength := len(line)
	i := 0
	for range line {
		if i+maxLength > lineLength {
			break
		}

		slice := line[i : i+maxLength]
		if searchDeactivate(slice) {
			isActivated = false
		}
		if searchActivate(slice) {
			isActivated = true
		}
		if isActivated {
			lineSum := parseLine01(slice)
			if lineSum > 0 {
				sum += lineSum
				i += maxLength
				continue
			}
		}
		i++
	}

	return sum, isActivated
}

