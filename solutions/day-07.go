package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day07Part01() int {
	file, err := os.Open("inputs/day-07.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		stringValues := strings.Split(strings.TrimSpace(parts[1]), " ")
		values := make([]int, len(stringValues))
		for i, value := range stringValues {
			values[i], _ = strconv.Atoi(value)
		}

		if isValid(values, target, true) {
			sum += target
		}
	}
	return sum
}

func isValid(values []int, target int, isPart2 bool) bool {
	if len(values) == 1 {
		return values[0] == target
	}
	x := values[0]
	y := values[1]
	plus := x + y
	multiply := x * y
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", x, y))
	nextValues := values[2:]
	plusSlice := append([]int{plus}, nextValues...)
	multiplySlice := append([]int{multiply}, nextValues...)
	concatSlice := append([]int{concat}, nextValues...)

	if isValid(plusSlice, target, isPart2) {
		return true
	} else if isValid(multiplySlice, target, isPart2) {
		return true
	} else if isPart2 && isValid(concatSlice, target, isPart2) {
		return true
	}
	return false
}
