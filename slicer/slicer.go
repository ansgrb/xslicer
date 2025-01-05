package slicer

import (
	"fmt"
	"strings"
)

// SplitTextIntoTweets splits a large text into tweetable chunks.
func SplitTextIntoTweets(text string, maxLength int) []string {
	var tweets []string
	words := strings.Fields(text) // Split text into words
	currentTweet := ""
	currentLength := 0

	for _, word := range words {
		// If adding the next word exceeds the max length, finalize the current tweet
		if currentLength+len(word)+1 > maxLength {
			tweets = append(tweets, currentTweet)
			currentTweet = ""
			currentLength = 0
		}

		// Add the word to the current tweet
		if currentLength > 0 {
			currentTweet += " "
			currentLength++
		}
		currentTweet += word
		currentLength += len(word)
	}

	// Add the last tweet if it's not empty
	if currentTweet != "" {
		tweets = append(tweets, currentTweet)
	}

	return tweets
}

// AddThreadNumbers adds thread numbering to the tweets (e.g., "1/5", "2/5").
func AddThreadNumbers(tweets []string) []string {
	totalTweets := len(tweets)
	if totalTweets <= 1 {
		return tweets
	}

	for i := range tweets {
		tweets[i] = fmt.Sprintf("%d/%d %s", i+1, totalTweets, tweets[i])
	}

	return tweets
}