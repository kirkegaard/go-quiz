package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Category struct {
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Text    string   `json:"question"`
	Answers []string `json:"answers"`
	Correct int      `json:"answer"`
}

// ANSI escape codes for text colors
const (
	whiteText     = "\033[97m" // White text
	blackText     = "\033[30m" // Black text
	redText       = "\033[31m" // Red text
	greenText     = "\033[32m" // Green text
	yellowText    = "\033[33m" // Yellow text
	blueText      = "\033[34m" // Blue text
	magentaText   = "\033[35m" // Magenta text
	cyanText      = "\033[36m" // Cyan text
	grayText      = "\033[37m" // Gray text
	lightGrayText = "\033[90m" // Light gray text

	// ANSI escape codes for background colors
	blackBackground     = "\033[40m" // Black background
	redBackground       = "\033[41m" // Red background
	greenBackground     = "\033[42m" // Green background
	yellowBackground    = "\033[43m" // Yellow background
	blueBackground      = "\033[44m" // Blue background
	magentaBackground   = "\033[45m" // Magenta background
	cyanBackground      = "\033[46m" // Cyan background
	whiteBackground     = "\033[47m" // White background
	lightGrayBackground = "\033[48m" // Light gray background

	reset = "\033[0m" // Reset color to default
)

func printColoredText(text, textColor, backgroundColor string) {
	fmt.Printf("%s%s%s%s\n", textColor, backgroundColor, text, reset)
}

func main() {
	// Load categories from JSON file
	data, err := os.ReadFile("questions.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var categories []Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var categoryIndex int = -1
	var score int = 0

	for {
		// Ask user to select a category if one is not chosen or if they chose to switch
		if categoryIndex == -1 {
			fmt.Println("Select a category:")
			for i, category := range categories {
				fmt.Printf("%d. %s\n", i+1, category.Title)
			}
			fmt.Print("Enter the category number or 'q' to quit: ")

			var input string
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Check if user wants to quit
			if input == "q" {
				break
			}

			// Convert input to int for category selection
			_, err = fmt.Sscan(input, &categoryIndex)
			if err != nil || categoryIndex < 1 || categoryIndex > len(categories) {
				fmt.Println("Invalid input. Please enter a valid category number or 'q' to quit.")
				// Reset to prompt selection again
				categoryIndex = -1
				continue
			}

			categoryIndex--
		}

		selectedCategory := categories[categoryIndex]

		// Create a new Rand object with a seed
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Select a random question from the chosen category
		question := selectedCategory.Questions[r.Intn(len(selectedCategory.Questions))]

		// Ask the question
		printColoredText(fmt.Sprintf("Score: %d", score), whiteText, blueBackground)
		fmt.Printf("Category: %s\n", selectedCategory.Title)
		fmt.Printf("Question: %s\n", question.Text)
		for i, opt := range question.Answers {
			fmt.Printf("%d. %s\n", i+1, opt)
		}
		fmt.Printf("Enter your answer (1-%d), 'c' to change category, or 'q' to quit: ", len(question.Answers))

		// Get the answer from the user
		var input string
		_, err = fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Check if user wants to quit
		if input == "q" {
			break
		}

		// Check if user wants to change category
		if input == "c" {
			categoryIndex = -1
			continue
		}

		var answer int
		_, err = fmt.Sscan(input, &answer)
		if err != nil || answer < 1 || answer > len(question.Answers) {
			fmt.Println("Invalid input. Please enter a valid answer number, 'c' to change category, or 'q' to quit.")
			continue
		}

		// Check the user's answer
		if answer == question.Correct+1 {
			score++
			printColoredText("Correct!", whiteText, greenBackground)
		} else {
			printColoredText(fmt.Sprintf("Sorry, the correct answer was %d.", question.Correct+1), whiteText, redBackground)
		}
	}
}
