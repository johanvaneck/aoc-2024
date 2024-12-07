package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day07Part01() int64 {
	file, err := os.Open("inputs/day-07.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[int64][]int64)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		a, _ := strconv.Atoi(parts[0])
		b := strings.Split(parts[1], " ")
		b = b[1:]
		m[int64(a)] = make([]int64, int64(len(b)))
		for i, c := range b {
			c, _ := strconv.Atoi(c)
			m[int64(a)][i] = int64(c)
		}
	}

	sum := int64(0)
	for k, v := range m {
		sum += getSum(k, v)
	}

	return sum
}

func getSum(k int64, values []int64) int64 {
	operators := []string{"+", "*"}
	combinations := getCombinations(operators, len(values)-1)

	for _, c := range combinations {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Target: ", k)
		fmt.Println(equationToString(values, c))
		res := calculateCombination(values, c)
		if res == k {
			return k
		}
	}
	return 0
}

func getCombinations(operators []string, size int) [][]string {
	if size == 0 {
		return [][]string{}
	} else if size == 1 {
		res := make([][]string, len(operators))
		for i, o := range operators {
			res[i] = []string{o}
		}
		return res
	} else {
		res := make([][]string, 0)
		prev := getCombinations(operators, size-1)
		for _, o := range operators {
			for _, p := range prev {
				row := make([]string, 0)
				row = append([]string{o}, p...)
				res = append(res, row)
			}
		}
		return res
	}
}

func equationToString(values []int64, combination []string) string {
	str := ""
	for i, c := range combination {
		str += fmt.Sprintf("%d%s", values[i], c)
	}
	str += fmt.Sprintf("%d", values[len(values)-1])
	return str
}

func printMatrix(m [][]string) {
	for i := range m {
		for j := range m[i] {
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}

// (((v1 op v2) op v3) op v4)
func calculateCombination(values []int64, combination []string) int64 {
	if len(values) == 2 {
		x := values[0]
		y := values[1]
		op := combination[0]
		res := calculate(x, y, op)
		fmt.Println(x, op, y, "=", res)
		return res
	}
	lvi := len(values) - 1
	lci := len(combination) - 1
	x := calculateCombination(values[:lvi], combination[:lci])
	y := values[lvi]
	op := combination[lci]
	res := calculate(x, y, op)
	fmt.Println(x, op, y, "=", res)
	return res
}

func calculate(x int64, y int64, op string) int64 {
	if op == "+" {
		return x + y
	} else if op == "-" {
		return x - y
	} else if op == "*" {
		return x * y
	} else if op == "/" && y != 0 {
		return x / y
	}
	fmt.Println("Error: ", x, y, op)
	return 0
}

func Day07Part02() int64 {
	file, err := os.Open("samples/day-07.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[int64][]int64)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		a, _ := strconv.Atoi(parts[0])
		b := strings.Split(parts[1], " ")
		b = b[1:]
		m[int64(a)] = make([]int64, len(b))
		for i, c := range b {
			c, _ := strconv.Atoi(c)
			m[int64(a)][i] = int64(c)
		}
	}

	sum := int64(0)

	return sum
}

// Tried:
// 3351425134686 too high
// 3351424676480 too low
// 3351425131488

// +
// -

// ++
// +-
// -+
// --

// +++
// ++-
// +-+
// +--
// -++
// -+-
// --+
// ---

// ++++
// +++-
// ++-+
// ++--
// +-++
// +-+-
// +--+
// +---
// -+++
// -++-
// -+-+
// -+--
// --++
// --+-
// ---+
// ----

// +
// -
// *
// /

// ++
// +-
// +*
// +/
// -+
// --
// -*
// -/
// *+
// *-
// **
// */
// /+
// /-
// /*
// //

// +++
// ++-
// ++*
// ++/
// +-+
// +--
// +-*
// +-/
// +*+
// +*-
// +**
// +*/
// +/+
// +/-
// +/*
// +//

// -++
// -+-
// -+*
// -+/
// --+
// ---
// --*
// --/
// -*+
// -*-
// -**
// -*/
// -/+
// -/-
// -/*
// -//
// *++
// *+-
// *+*
// *+/
// *-+
// *--
// *-*
// *-/
// **+
// **-
// ***
// **/
// */+
// */-
// */*
// *//
// *++
// *+-
// *+*
// *+/
// *-+
// *--
// *-*
// *-/
// **+
// **-
// ***
// **/
// */+
// */-
// */*
// *//
