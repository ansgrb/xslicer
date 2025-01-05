package slicer

import (
	"testing"
)

func TestSplitTextIntoTweets(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		maxLength int
		expected  []string
	}{
		{
			name:      "Short text",
			text:      "Hello world",
			maxLength: 280,
			expected:  []string{"Hello world"},
		},
		{
			name:      "Long text",
			text:      "This is a long text that needs to be split into multiple tweets because it exceeds the character limit of a single tweet.",
			maxLength: 50,
			expected:  []string{"This is a long text that needs to be split", "into multiple tweets because it exceeds the", "character limit of a single tweet."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitTextIntoTweets(tt.text, tt.maxLength)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d tweets, got %d", len(tt.expected), len(result))
			}
			for i, tweet := range result {
				if tweet != tt.expected[i] {
					t.Errorf("Expected tweet %d to be '%s', got '%s'", i+1, tt.expected[i], tweet)
				}
			}
		})
	}
}

func TestAddThreadNumbers(t *testing.T) {
	tests := []struct {
		name     string
		tweets   []string
		expected []string
	}{
		{
			name:     "Single tweet",
			tweets:   []string{"Hello world"},
			expected: []string{"Hello world"},
		},
		{
			name:     "Multiple tweets",
			tweets:   []string{"Tweet one", "Tweet two", "Tweet three"},
			expected: []string{"1/3 Tweet one", "2/3 Tweet two", "3/3 Tweet three"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddThreadNumbers(tt.tweets)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d tweets, got %d", len(tt.expected), len(result))
			}
			for i, tweet := range result {
				if tweet != tt.expected[i] {
					t.Errorf("Expected tweet %d to be '%s', got '%s'", i+1, tt.expected[i], tweet)
				}
			}
		})
	}
}