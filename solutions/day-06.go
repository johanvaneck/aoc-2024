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
		i, j, marker = move(i, j, marker, matrix[:])
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

func move(i int, j int, marker string, matrix [][]rune) (int, int, string) {
	if marker == "^" {
		if i-1 < 0 {
			return i - 1, j, marker
		} else if matrix[i-1][j] == '#' {
			return i, j + 1, ">"
		} else {
			return i - 1, j, "^"
		}

	} else if marker == ">" {
		if j+1 >= len(matrix[i]) {
			return i, j + 1, marker
		} else if matrix[i][j+1] == '#' {
			return i + 1, j, "v"
		} else {
			return i, j + 1, ">"
		}
	} else if marker == "<" {
		if j-1 < 0 {
			return i, j - 1, marker
		} else if matrix[i][j-1] == '#' {
			return i - 1, j, "^"
		} else {
			return i, j - 1, "<"
		}
	} else if marker == "v" {
		if i+1 >= len(matrix) {
			return i + 1, j, marker
		} else if matrix[i+1][j] == '#' {
			return i, j - 1, "<"
		} else {
			return i + 1, j, "v"
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

	originalMatrix := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		originalMatrix = append(originalMatrix, []rune(line))
	}

	width := len(originalMatrix)
	height := len(originalMatrix[0])
	startI, startJ, startMarker := findMarker(originalMatrix)

	count := 0
	for i := range originalMatrix {
		for j := range originalMatrix[i] {
			if (i == startI && j == startJ) || (string(originalMatrix[i][j]) == "#") {
				continue
			}
			// Copy matrix
			matrix := make([][]rune, len(originalMatrix))
			for i, line := range originalMatrix {
				matrix[i] = make([]rune, len(line))
				copy(matrix[i], line)
			}
			// Insert new obstacle
			matrix[i][j] = '#'
			if checkLoop(startI, startJ, string(startMarker), matrix, width, height) {
				count++
			}
		}
	}

	return count
}

func checkLoop(startI int, startJ int, startMarker string, matrix [][]rune, width int, height int) bool {
	i, j := startI, startJ
	marker := startMarker
	visited := make([][]string, width)
	for i := range visited {
		visited[i] = make([]string, height)
	}

	for true {
		if i < 0 ||
			j < 0 ||
			i >= width-1 ||
			j >= height-1 ||
			marker == "" {
			break
		}

		if visited[i][j] == marker {
			return true
		}

		visited[i][j] = marker

		i, j, marker = move(i, j, marker, matrix[:])
	}

	return false
}

// Tried:
// 1893 too low
