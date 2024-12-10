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

func Day09Part02() int {
	file, err := os.Open("inputs/day-09.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	id := 0
	memory := make([]int, 0)
	for i, c := range line {
		num, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				memory = append(memory, id)
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				memory = append(memory, -1)
			}
		}
	}

	tail := len(memory) - 1
	for tail > 0 {
		if memory[tail] == -1 {
			tail--
			continue
		}
		req := 0
		for true {
			req++
			if tail == 0 || memory[tail-1] != memory[tail] {
				break
			}
			tail--
		}

		head := 0
		space := 0
		for head <= tail {
			for memory[head] == -1 {
				space++
				head++
			}

			if space >= req {
				for i := 0; i < req; i++ {
					swap := memory[head-space+i]
					memory[head-space+i] = memory[tail+i]
					memory[tail+i] = swap
				}
				break
			} else {
				space = 0
			}
			head++
		}
		tail--
	}

	checksum := 0
	for i, num := range memory {
		if num == -1 {
			continue
		} else {
			checksum += num * i
		}
	}

	return checksum
}
