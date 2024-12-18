package main

import (
	"aoc-2024/solutions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	solution := solutions.Day07Part01()

	fmt.Printf("Solution: %d\n", solution)

	elapsed := time.Since(start)
	fmt.Printf("Running time: %s\n", elapsed)
}
