package solutions

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Day08Part01() int {
	file, err := os.Open("inputs/day-08.txt")
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

	antennas := make(map[rune][][]int)
	hasAntennaMatrix := make([][]bool, 0)

	for li := range matrix {
		hasAntennaMatrix = append(hasAntennaMatrix, make([]bool, 0))
		for lj, c := range matrix[li] {
			hasAntennaMatrix[li] = append(hasAntennaMatrix[li], false)
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				if antennas[c] == nil {
					antennas[c] = make([][]int, 0)
				}
				antennas[c] = append(antennas[c], []int{li, lj})
			}
		}
	}

	matI := len(matrix)
	matJ := len(matrix[0])
	for _, cCoords := range antennas {
		for i, curCoord := range cCoords {
			xi, yi := curCoord[0], curCoord[1]
			for j, otherCoord := range cCoords {
				xj, yj := otherCoord[0], otherCoord[1]
				if i == j {
					continue
				}
				dx := xj - xi
				dy := yj - yi

				xn := xi - dx
				yn := yi - dy

				if xn >= 0 && xn < matI && yn >= 0 && yn < matJ {
					hasAntennaMatrix[xn][yn] = true
				}
			}
		}
	}

	count := 0
	for _, row := range hasAntennaMatrix {
		for _, hasAntenna := range row {
			if hasAntenna {
				count++
			}
		}
	}

	printMatrix(matrix, hasAntennaMatrix)

	return count
}

func printMatrix(ogMatrix []string, matrix [][]bool) {
	for i, row := range matrix {
		for j, hasAntenna := range row {
			if hasAntenna {
				fmt.Printf("#")
			} else {
				fmt.Printf("%c", ogMatrix[i][j])
			}
		}
		fmt.Printf("\n")
	}
}

