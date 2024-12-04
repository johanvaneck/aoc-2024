package solutions

import (
	"bufio"
	"os"
)

func Day04Part01() int {
	file, err := os.Open("inputs/day-04.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	count := 0
	word := "XMAS"
	wordLastIndex := len(word) - 1
	rowsLastIndex := len(matrix) - 1
	columnsLastIndex := len(matrix[0]) - 1

	for i := range matrix {
		for j := range matrix[i] {
			if i <= rowsLastIndex-wordLastIndex && j <= columnsLastIndex-wordLastIndex {
				count += countDiagonal(i, j, matrix, word)
			}
			if i <= rowsLastIndex-wordLastIndex {
				count += countColumn(i, j, matrix, word)
			}
			if j <= columnsLastIndex-wordLastIndex {
				count += countRow(i, j, matrix, word)
			}
		}

	}

	return count
}

func countDiagonal(x int, y int, matrix []string, word string) int {
	wordLastIndex := len(word) - 1
	count := 0
	/*
		X###
		#M##
		##A#
		###S
	*/
	for i, c := range word {
		if matrix[x+i][y+i] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}
	/*
		###X
		##M#
		#A##
		S###
	*/
	for i, c := range word {
		if matrix[x+i][y+wordLastIndex-i] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}
	/*
		S###
		#A##
		##M#
		###X
	*/
	for i, c := range word {
		if matrix[x+wordLastIndex-i][y+i] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}
	/*
		###S
		##A#
		#M##
		X###
	*/
	for i, c := range word {
		if matrix[x+wordLastIndex-i][y+wordLastIndex-i] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}

	return count
}

func countRow(x int, y int, matrix []string, word string) int {
	wordLastIndex := len(word) - 1
	count := 0

	// XMAS
	for i, c := range word {
		if matrix[x][y+i] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}
	// SAMX
	for i, c := range word {
		if matrix[x][y+wordLastIndex-i] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}
	return count
}

func countColumn(x int, y int, matrix []string, word string) int {
	wordLastIndex := len(word) - 1
	count := 0
	/*
		X###
		M###
		A###
		S###
	*/
	for i, c := range word {
		if matrix[x+i][y] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}
	/*
		S###
		A###
		M###
		X###
	*/
	for i, c := range word {
		if matrix[x+wordLastIndex-i][y] != byte(c) {
			break
		} else if i == wordLastIndex {
			count++
		}
	}
	return count
}

func Day04Part02() int {
	file, err := os.Open("inputs/day-04.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	count := 0
	word := "MAS"
	wordLastIndex := len(word) - 1
	rowsLastIndex := len(matrix) - 1
	columnsLastIndex := len(matrix[0]) - 1

	for i := range matrix {
		for j := range matrix[i] {
			if i <= rowsLastIndex-wordLastIndex && j <= columnsLastIndex-wordLastIndex {
				count += countCrossPairs(i, j, matrix, word)
			}
		}

	}

	return count
}

func countCrossPairs(x int, y int, matrix []string, word string) int {
	wordLastIndex := len(word) - 1
	/*
		M#M
		#A#
		S#S
	*/
	for i, c := range word {
		if matrix[x+i][y+i] != byte(c) || matrix[x+wordLastIndex-i][y+i] != byte(c) {
			break
		} else if i == wordLastIndex {
			return 1
		}
	}
	/*
		S#M
		#A#
		S#M
	*/
	for i, c := range word {
		if matrix[x+wordLastIndex-i][y+wordLastIndex-i] != byte(c) || matrix[x+wordLastIndex-i][y+i] != byte(c) {
			break
		} else if i == wordLastIndex {
			return 1
		}
	}
	/*
		S#S
		#A#
		M#M
	*/
	for i, c := range word {
		if matrix[x+wordLastIndex-i][y+wordLastIndex-i] != byte(c) || matrix[x+i][y+wordLastIndex-i] != byte(c) {
			break
		} else if i == wordLastIndex {
			return 1
		}
	}
	/*
		M#S
		#A#
		M#S
	*/
	for i, c := range word {
		if matrix[x+i][y+i] != byte(c) || matrix[x+i][y+wordLastIndex-i] != byte(c) {
			break
		} else if i == wordLastIndex {
			return 1
		}
	}

	return 0
}
