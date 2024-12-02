package main

import (
	"aoc-2024/solutions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	solutions.Day01Part02()

	elapsed := time.Since(start)
	fmt.Printf("Running time: %s\n", elapsed)
}
