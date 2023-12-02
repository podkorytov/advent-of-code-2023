package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type GameAttempt struct {
	Red   int
	Green int
	Blue  int
}

type GameData struct {
	GameId   int
	Attempts []GameAttempt
}

func parseGameAttemptData(gameDataString string) (GameData, error) {
	attemptGroups := strings.Split(gameDataString, ";")
	var attempts []GameAttempt

	for _, attemptGroup := range attemptGroups {
		attemptStrings := strings.Split(strings.TrimSpace(attemptGroup), ",")
		var attempt GameAttempt

		for _, attemptString := range attemptStrings {
			parts := strings.Split(strings.TrimSpace(attemptString), " ")
			if len(parts) != 2 {
				return GameData{}, fmt.Errorf("invalid attempt format: %s", attemptString)
			}

			num, err := strconv.Atoi(parts[0])
			if err != nil {
				return GameData{}, fmt.Errorf("invalid number format: %s", parts[0])
			}

			switch parts[1] {
			case "red":
				attempt.Red = num
			case "green":
				attempt.Green = num
			case "blue":
				attempt.Blue = num
			default:
				return GameData{}, fmt.Errorf("invalid color: %s", parts[1])
			}
		}

		attempts = append(attempts, attempt)
	}

	return GameData{Attempts: attempts}, nil
}

func parseGameString(gameString string) (GameData, error) {
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	match := re.FindStringSubmatch(gameString)

	if len(match) != 3 {
		return GameData{}, fmt.Errorf("invalid game format: %s", gameString)
	}

	gameId, err := strconv.Atoi(match[1])
	if err != nil {
		return GameData{}, fmt.Errorf("invalid game ID format: %s", match[1])
	}

	gameData, err := parseGameAttemptData(match[2])
	if err != nil {
		return GameData{}, err
	}

	gameData.GameId = gameId
	return gameData, nil
}

func checkGame(game GameData, restriction GameAttempt) bool {
	for _, attempt := range game.Attempts {
		if attempt.Red > restriction.Red || attempt.Green > restriction.Green || attempt.Blue > restriction.Blue {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Close the file when the function exits
	defer file.Close()

	restriction := GameAttempt{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	result := 0

	// Iterate through each line
	for scanner.Scan() {
		gameData, err := parseGameString(scanner.Text())

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if checkGame(gameData, restriction) {
			result += gameData.GameId
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Result: %d\n", result)
}
