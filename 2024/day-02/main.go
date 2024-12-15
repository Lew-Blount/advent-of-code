package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1Answer, p1Err := part1()
	if p1Err != nil {
		log.Fatalf("Error in part1: %v", p1Err)
	}

	fmt.Printf("Part 1: Safe reports: %d\n", p1Answer)

	p2Answer, p2Err := part2()
	if p2Err != nil {
		log.Fatalf("Error in part2: %v", p2Err)
	}

	fmt.Printf("Part 2: Safe reports: %d\n", p2Answer)
}

func readInput(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reports = append(reports, strings.Fields(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func part1() (int, error) {
	reports, err := readInput("input.txt")
	if err != nil {
		return 0, err
	}

	safe := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safe++
		}
	}

	return safe, nil
}

func part2() (int, error) {
	reports, err := readInput("input.txt")
	if err != nil {
		return 0, err
	}

	safe := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safe++
		} else {
			for i := 0; i < len(report); i++ {
				copyReport := append([]string(nil), report...)
				copyReport = append(copyReport[:i], copyReport[i+1:]...)
				if isSafeReport(copyReport) {
					safe++
					break
				}
			}
		}
	}

	return safe, nil
}

func isSafeReport(report []string) bool {
	increase, decrease := 0, 0

	for i := 0; i < len(report)-1; i++ {
		first, err := strconv.Atoi(report[i])
		if err != nil {
			log.Printf("Error parsing first number: %s, error: %v", report[i], err)
			return false
		}
		second, err := strconv.Atoi(report[i+1])
		if err != nil {
			log.Printf("Error parsing second number: %s, error: %v", report[i+1], err)
			return false
		}

		difference := 0
		if first == second {
			return false
		}

		if first > second {
			difference = first - second
			decrease++
		} else {
			difference = second - first
			increase++
		}

		if difference > 3 {
			return false
		}
	}
	return increase == len(report)-1 || decrease == len(report)-1
}
