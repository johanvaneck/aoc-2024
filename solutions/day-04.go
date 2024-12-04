package solutions

import (
	"bufio"
	"fmt"
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
		fmt.Println(line)
		matrix = append(matrix, line)
	}

	count := 0
	word := "XMAS"
	wordLastIndex := len(word) - 1
	rowsLastIndex := len(matrix) - 1
	columnsLastIndex := len(matrix[0]) - 1

	diagonalCount := 0
	colCount := 0
	rowCount := 0

	for i := range matrix {
		for j := range matrix[i] {
			if i <= rowsLastIndex-wordLastIndex && j <= columnsLastIndex-wordLastIndex {
				fmt.Println(i, j)
				count += countDiagonal(i, j, matrix, word)
				diagonalCount += countDiagonal(i, j, matrix, word)
			}
			if i <= rowsLastIndex-wordLastIndex {
				count += countColumn(i, j, matrix, word)
				colCount += countColumn(i, j, matrix, word)
			}
			if j <= columnsLastIndex-wordLastIndex {
				count += countRow(i, j, matrix, word)
				rowCount += countRow(i, j, matrix, word)
			}
		}

	}

	fmt.Printf("diagonal: %d\ncolumn: %d\nrow: %d\n", diagonalCount, colCount, rowCount)

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
