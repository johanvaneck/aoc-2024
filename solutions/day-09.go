package solutions

import (
	"bufio"
	"os"
	"strconv"
)

func Day09Part01() int {
	file, err := os.Open("inputs/day-09.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	id := 0
	unzipped := make([]int, 0)
	for i, c := range line {
		num, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				unzipped = append(unzipped, id)
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				unzipped = append(unzipped, -1)
			}
		}
	}

	head := 0
	tail := len(unzipped) - 1
	for head < tail {
		for unzipped[tail] == -1 {
			tail--
		}
		if unzipped[head] == -1 {
			tmp := unzipped[head]
			unzipped[head] = unzipped[tail]
			unzipped[tail] = tmp
			tail--
		}
		head++
	}

	checksum := 0
	for i, num := range unzipped {
		if num == -1 {
			break
		} else {
			checksum += num * i
		}
	}

	return checksum
}
