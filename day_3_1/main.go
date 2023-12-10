package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func readFile(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		matrix = append(matrix, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

func findNumberAtCoordinates(matrix [][]rune, i, j int) (startCoord [2]int, value int, found bool) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		// Coordinates out of bounds
		return
	}

	if !unicode.IsDigit(matrix[i][j]) {
		// Not a number at the specified coordinates
		return
	}

	// Initialize start coordinates to current coordinates
	startCoord = [2]int{i, j}

	// Read the entire number backward
	for startCoord[1] >= 0 && unicode.IsDigit(matrix[i][startCoord[1]]) {
		startCoord[1]--
	}

	// Read the entire number forward
	endCoord := j + 1
	for endCoord < len(matrix[i]) && unicode.IsDigit(matrix[i][endCoord]) {
		endCoord++
	}

	// Parse the number
	num, err := strconv.Atoi(string(matrix[i][startCoord[1]+1 : endCoord]))
	if err != nil {
		// Failed to convert to an integer
		return
	}

	return startCoord, num, true
}

func processCoordinates(matrix [][]rune, i, j int, uniqueStartCoords map[[2]int]struct{}, result int) int {
	var directions = [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		startCoord, num, found := findNumberAtCoordinates(matrix, i+dir[0], j+dir[1])
		if found {
			coordKey := [2]int{startCoord[0], startCoord[1]}
			if _, exists := uniqueStartCoords[coordKey]; !exists {
				uniqueStartCoords[coordKey] = struct{}{}
				result += num
			}
		}
	}

	return result
}

func main() {
	filePath := "input.txt"
	matrix, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	uniqueStartCoords := make(map[[2]int]struct{})
	result := 0

	for i, row := range matrix {
		for j, cell := range row {
			if cell != '.' && (cell < '0' || cell > '9') {
				result = processCoordinates(matrix, i, j, uniqueStartCoords, result)
			}
		}
	}

	fmt.Printf("Result: %d\n", result)
}
