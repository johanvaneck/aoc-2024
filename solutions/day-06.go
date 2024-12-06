package solutions

import (
	"bufio"
	"fmt"
	"os"
)

func Day06Part01() int {
	file, err := os.Open("inputs/day-06.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	width := len(matrix)
	height := len(matrix[0])
	visited := make([][]bool, width)
	for i := range visited {
		visited[i] = make([]bool, height)
	}

	i, j, marker := findMarker(matrix)

	for true {
		visited[i][j] = true
		if i == -1 ||
			j == -1 ||
			i >= width-1 ||
			j >= height-1 ||
			marker == "" {
			break
		}
		i, j, marker = move(i, j, marker, matrix[:], -1, -1)
	}

	count := 0
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				matrix[i][j] = 'X'
				count++
			}
		}
	}

	for i := range matrix {
		fmt.Println(string(matrix[i]))
	}

	return count
}

func move(i int, j int, marker string, matrix [][]rune, obsI int, ObsJ int) (int, int, string) {
	if marker == "^" {
		nextI := i - 1
		if nextI < 0 {
			return nextI, j, marker
		} else if matrix[nextI][j] == '#' || (nextI == obsI && j == ObsJ) {
			return i, j + 1, ">"
		} else {
			return nextI, j, "^"
		}

	} else if marker == ">" {
		nextJ := j + 1
		if nextJ >= len(matrix[i]) {
			return i, nextJ, marker
		} else if matrix[i][nextJ] == '#' || (i == obsI && nextJ == ObsJ) {
			return i + 1, j, "v"
		} else {
			return i, j + 1, ">"
		}
	} else if marker == "<" {
		nextJ := j - 1
		if nextJ < 0 {
			return i, nextJ, marker
		} else if matrix[i][j-1] == '#' || (i == obsI && nextJ == ObsJ) {
			return i - 1, j, "^"
		} else {
			return i, nextJ, "<"
		}
	} else if marker == "v" {
		nextI := i + 1
		if nextI >= len(matrix) {
			return nextI, j, marker
		} else if matrix[nextI][j] == '#' || (nextI == obsI && j == ObsJ) {
			return i, j - 1, "<"
		} else {
			return nextI, j, "v"
		}
	}

	return -1, -1, marker
}

func findMarker(matrix [][]rune) (int, int, string) {
	for i, line := range matrix {
		for j, c := range line {
			if c == '^' ||
				c == '>' ||
				c == '<' ||
				c == 'v' {
				return i, j, string(c)
			}
		}
	}
	return -1, -1, ""
}

func Day06Part02() int {
	file, err := os.Open("inputs/day-06.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	h := len(matrix)
	w := len(matrix[0])
	startI, startJ, startMarker := findMarker(matrix)

	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			if (i == startI && j == startJ) || (matrix[i][j] == '#') {
				continue
			}
			if checkLoop(startI, startJ, string(startMarker), matrix, w, h, i, j) {
				count++
			}
		}
	}

	return count
}

func checkLoop(startI int, startJ int, startMarker string, matrix [][]rune, w int, h int, obsI int, ObsJ int) bool {
	i, j := startI, startJ
	marker := startMarker
	visited := make([][]string, h)
	for i := range visited {
		visited[i] = make([]string, w)
	}

	for true {
		if i < 0 ||
			j < 0 ||
			i >= h-1 ||
			j >= w-1 ||
			marker == "" {
			break
		}

		if visited[i][j] == marker {
			return true
		}

		visited[i][j] = marker

		i, j, marker = move(i, j, marker, matrix[:], obsI, ObsJ)
	}

	return false
}

// Tried:
// 1893 too low
