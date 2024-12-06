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
		if matrix[i-1][j] == '#' {
			return i, j + 1, ">"
		} else {
			return i - 1, j, "^"
		}

	} else if marker == ">" {
		if matrix[i][j+1] == '#' {
			return i + 1, j, "v"
		} else {
			return i, j + 1, ">"
		}
	} else if marker == "<" {
		if matrix[i][j-1] == '#' {
			return i - 1, j, "^"
		} else {
			return i, j - 1, "<"
		}
	} else if marker == "v" {
		if matrix[i+1][j] == '#' {
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
