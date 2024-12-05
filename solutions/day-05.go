package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day05Part01() int {
	file, err := os.Open("inputs/day-05.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ruleMap := make(map[int][]int)
	pageGroups := make([][]int, 0)

	isRulesSection := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isRulesSection = false
			continue
		}

		if isRulesSection {
			ruleStrings := strings.Split(line, "|")
			a, _ := strconv.Atoi(ruleStrings[0])
			b, _ := strconv.Atoi(ruleStrings[1])
			if ruleMap[a] == nil {
				ruleMap[a] = make([]int, 0)
			}
			ruleMap[a] = append(ruleMap[a], b)
		} else {
			groupStrings := strings.Split(line, ",")
			group := make([]int, 0)
			for _, s := range groupStrings {
				value, _ := strconv.Atoi(s)
				group = append(group, value)
			}
			pageGroups = append(pageGroups, group)
		}
	}

	sum := 0

	for _, group := range pageGroups {
		sum += groupValuePart01(group, ruleMap)
	}

	return sum
}

func groupValuePart01(group []int, ruleMap map[int][]int) int {
	for i, value := range group {
		beforeSlice := group[:i]
		for _, ruleValue := range ruleMap[value] {
			for _, beforeValue := range beforeSlice {
				if ruleValue == beforeValue {
					return 0
				}
			}
		}
	}
	return group[len(group)/2]
}

func Day05Part02() int {
	file, err := os.Open("inputs/day-05.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ruleMap := make(map[int][]int)
	pageGroups := make([][]int, 0)

	isRulesSection := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isRulesSection = false
			continue
		}

		if isRulesSection {
			ruleStrings := strings.Split(line, "|")
			a, _ := strconv.Atoi(ruleStrings[0])
			b, _ := strconv.Atoi(ruleStrings[1])
			if ruleMap[a] == nil {
				ruleMap[a] = make([]int, 0)
			}
			ruleMap[a] = append(ruleMap[a], b)
		} else {
			groupStrings := strings.Split(line, ",")
			group := make([]int, 0)
			for _, s := range groupStrings {
				value, _ := strconv.Atoi(s)
				group = append(group, value)
			}
			pageGroups = append(pageGroups, group)
		}
	}

	sum := 0

	for _, group := range pageGroups {
		if groupValuePart01(group, ruleMap) == 0 {
			sum += groupValuePart02(group, ruleMap)
		}
	}

	return sum
}

func groupValuePart02(group []int, ruleMap map[int][]int) int {
	for i, value := range group {
		beforeSlice := group[:i]
		for _, ruleValue := range ruleMap[value] {
			for j, beforeValue := range beforeSlice {
				if ruleValue == beforeValue {
					newGroup := []int{}
					newGroup = append(newGroup, group[:j]...)
					newGroup = append(newGroup, group[i])
					newGroup = append(newGroup, group[j:i]...)
					newGroup = append(newGroup, group[i+1:]...)
					return groupValuePart02(newGroup, ruleMap)
				}
			}
		}
	}
	return group[len(group)/2]
}
