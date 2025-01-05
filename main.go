package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"dev.ansgrb.xslicer/slicer"
)

func main() {
	// Define command-line flags
	textPtr := flag.String("text", "", "The text to split into tweets")
	filePtr := flag.String("file", "", "Path to a text file containing the input")
	outputPtr := flag.String("output", "", "Path to save the output tweets (e.g., output.txt)")
	maxLengthPtr := flag.Int("maxlength", 280, "Maximum length of each tweet")
	flag.Parse()

	// Validate input
	if *textPtr == "" && *filePtr == "" {
		fmt.Println("Please provide text using the -text flag or a file using the -file flag.")
		return
	}

	var inputText string

	// Prioritize file input over direct text input
	if *filePtr != "" {
		// Read text from file
		content, err := os.ReadFile(*filePtr)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		inputText = string(content)
	} else {
		inputText = *textPtr
	}

	// Split the text into tweets
	tweets := slicer.SplitTextIntoTweets(inputText, *maxLengthPtr)

	// Add thread numbering
	tweets = slicer.AddThreadNumbers(tweets)

	// Print the tweets to the console
	for _, tweet := range tweets {
		fmt.Println(tweet)
	}

	// Save tweets to output file if -output flag is provided
	if *outputPtr != "" {
		err := saveTweetsToFile(tweets, *outputPtr)
		if err != nil {
			log.Fatalf("Error saving tweets to file: %v", err)
		}
		fmt.Printf("Tweets saved to %s\n", *outputPtr)
	}
}

// saveTweetsToFile writes the tweets to a file.
func saveTweetsToFile(tweets []string, filePath string) error {
	// Create or truncate the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write each tweet to the file
	for _, tweet := range tweets {
		_, err := file.WriteString(tweet + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}