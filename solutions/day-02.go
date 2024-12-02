package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day02Part01() int {
	file, err := os.Open("inputs/day-02.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	countSafe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		if isSafe01(fields) {
			countSafe++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return countSafe
}

func Day02Part02() int {
	file, err := os.Open("inputs/day-02.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	countSafe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")

		isSafe := isSafe02(fields)

		if isSafe {
			countSafe++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return countSafe
}

func isSafe01(fields []string) bool {
	var prev int
	isAscending := true
	for i, field := range fields {
		if i == 0 {
			prev, _ = strconv.Atoi(field)
			continue
		}

		v, _ := strconv.Atoi(field)

		if i == 1 {
			diff := v - prev
			if diff == 0 || (math.Abs(float64(diff)) > 3) {
				return false
			}
			if diff > 0 {
				isAscending = true
			} else {
				isAscending = false
			}
			prev = v
			continue
		}

		diff := v - prev

		if diff == 0 ||
			(math.Abs(float64(diff)) > 3) ||
			(isAscending && diff < 0) ||
			(!isAscending && diff > 0) {
			return false

		}

		prev = v
	}

	return true
}

func isSafe02(fields []string) bool {
	if isSafe01(fields) {
		return true
	}

	for i := range fields {
		first := append([]string{}, fields[:i]...)
		second := fields[i+1:]
		slice := append(first, second...)

		if isSafe01(slice) {
			return true
		}
	}

	return false
}
