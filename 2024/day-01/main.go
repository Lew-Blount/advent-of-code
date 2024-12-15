package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	p1Answer, p1Err := part1()
	if p1Err != nil {
		log.Fatalf("Error in part1: %v", p1Err)
	}

	fmt.Printf("Part 1: Distance is: %d\n", p1Answer)

	p2Answer, p2Err := part2()
	if p2Err != nil {
		log.Fatalf("Error in part2: %v", p2Err)
	}

	fmt.Printf("Part 2: Similarity score is: %d\n", p2Answer)
}

func part1() (int, error) {
	left, right, err := createLists("input.txt")
	if err != nil {
		return 0, err
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			totalDistance += left[i] - right[i]
		} else {
			totalDistance += right[i] - left[i]
		}
	}

	return totalDistance, nil
}

func part2() (int, error) {
	left, right, err := createLists("input.txt")
	if err != nil {
		return 0, err
	}

	similarityScore := 0
	countMap := make(map[int]int)

	for _, value := range right {
		countMap[value]++
	}

	for _, value := range left {
		similarityScore += value * countMap[value]
	}

	return similarityScore, nil
}

func createLists(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var left, right []int

	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		if len(input) != 2 {
			log.Printf("Skipping invalid line: %s\n", scanner.Text())
			continue
		}

		first, err := strconv.Atoi(input[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in line: %w", err)
		}

		second, err := strconv.Atoi(input[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in line: %w", err)
		}

		left = append(left, first)
		right = append(right, second)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	return left, right, nil
}
