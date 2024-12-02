package solutions

import (
	"aoc-2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day01Part01() {
	text := utils.FileToString("inputs/day-01.txt")

	lines := strings.Split(text, "\n")

	var firstColumn []int
	var secondColumn []int

	for _, line := range lines {
		lineParts := strings.Split(line, "   ")

		if len(lineParts) != 2 {
			continue
		}

		first, _ := strconv.Atoi(lineParts[0])
		second, _ := strconv.Atoi(lineParts[1])

		firstColumn = append(firstColumn, first)
		secondColumn = append(secondColumn, second)
	}

	sort.Ints(firstColumn)
	sort.Ints(secondColumn)

	sum := 0
	for i := 0; i < len(firstColumn); i++ {
		sum += int(math.Abs(float64(firstColumn[i] - secondColumn[i])))
	}

	fmt.Println(firstColumn)

	fmt.Printf("Day01PartA: %d\n", sum)
}

func Day01Part02() {
	text := utils.FileToString("inputs/day-01.txt")

	lines := strings.Split(text, "\n")

	var firstColumn []int
	var secondColumn []int

	for _, line := range lines {
		lineParts := strings.Split(line, "   ")

		if len(lineParts) != 2 {
			continue
		}

		first, _ := strconv.Atoi(lineParts[0])
		second, _ := strconv.Atoi(lineParts[1])

		firstColumn = append(firstColumn, first)
		secondColumn = append(secondColumn, second)
	}

	firstColumnSet := make(map[int]bool)

	for _, item := range firstColumn {
		firstColumnSet[item] = true
	}

	similaritySum := 0
	for key := range firstColumnSet {
		count := 0
		for _, item := range secondColumn {
			if item == key {
				count++
			}
		}
		similaritySum += similaririty(key, count)
	}

	fmt.Printf("Day01PartB: %d\n", similaritySum)
}

func similaririty(value int, count int) int {
	return value * count
}
