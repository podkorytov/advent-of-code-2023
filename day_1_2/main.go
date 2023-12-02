package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// Constants for numbers
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

var numbers = []string{one, two, three, four, five, six, seven, eight, nine}

var replacements = map[string]string{
	one:   "1",
	two:   "2",
	three: "3",
	four:  "4",
	five:  "5",
	six:   "6",
	seven: "7",
	eight: "8",
	nine:  "9",
}

func findMinNumber(myMap map[string]int) string {
	var minKey string
	minValue := 1000000

	for key, value := range myMap {
		if value < minValue {
			minValue = value
			minKey = key
		}
	}

	return replacements[minKey]
}

func findMaxNumber(myMap map[string]int) string {
	var maxKey string
	maxValue := -1

	for key, value := range myMap {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}

	return replacements[maxKey]
}

func calculateNumberFromString(input string) int {
	firstPositions := make(map[string]int)
	lastPositions := make(map[string]int)

	for _, number := range numbers {
		position := strings.Index(input, number)
		if position >= 0 {
			firstPositions[number] = position
		}
		position = strings.Index(input, replacements[number])
		if position >= 0 {
			if existingValue, ok := firstPositions[number]; !ok || position < existingValue {
				firstPositions[number] = position
			}
		}

		position = strings.LastIndex(input, number)
		if position >= 0 {
			lastPositions[number] = position
		}
		position = strings.LastIndex(input, replacements[number])
		if position >= 0 {
			if existingValue, ok := lastPositions[number]; !ok || position > existingValue {
				lastPositions[number] = position
			}
		}
	}

	outString := findMinNumber(firstPositions) + findMaxNumber(lastPositions)

	number, err := strconv.Atoi(outString)
	if err != nil {
		fmt.Println("Error converting string to number:", err)
	}

	return number
}

func main() {
	file, err := os.Open("day_1_2_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Close the file when the function exits
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	result := 0

	// Iterate through each line
	for scanner.Scan() {
		result += calculateNumberFromString(scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Result: %d\n", result)
}
