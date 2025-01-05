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

	// Print the tweets
	for _, tweet := range tweets {
		fmt.Println(tweet)
	}
}