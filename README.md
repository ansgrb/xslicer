# xslicer ✂️

_xslicer_ is a Go-based command-line tool that helps you split long text into tweetable chunks (280 characters or less) and save them as a thread. It’s perfect for turning long paragraphs into Twitter-friendly threads!

## Features ✨

- Split Long Text: Easily split long text into tweet-sized chunks.
- Thread Numbering: Automatically adds thread numbering (e.g., "1/5", "2/5") to each tweet.
- File Support: Accepts input from both direct text and text files.
- Save Output: Save the sliced tweets to an output file.
- Append Mode: Optionally append tweets to an existing file.
- Verbose Mode: Get detailed information about the process (e.g., number of tweets generated).
- Customizable: Set a custom character limit for tweets.

## Installation 🛠️

### Prerequisites

- **Go**: Make sure you have Go installed. You can download it from [here](https://golang.org/dl/).

## Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/ansgrb/xslicer.git
   cd xslicer
   ```
2. Build the project:
   ```bash
    go build -o tweet-slicer
   ```
3. Run the app:
   ```bash
    ./xslicer -file input.txt -output output.txt -verbose
   ```

## Usage 🚀

1. Split Text from a File:
   ```bash
    ./xslicer -file input.txt -output output.txt
   ```
2. Split Direct Text Input:
   ```bash
    ./xslicer -text "Your long text here..." -output output.txt
   ```
3. Append Tweets to an Existing File::
   ```bash
    ./xslicer -file input.txt -output output.txt -append
   ```
4. Set Custom Tweet Length:
   ```bash
   ./xslicer -file input.txt -maxlength 100 -output output.txt
   ```
5. Enable Verbose Mode:
   ```bash
    ./xslicer -file input.txt -output output.txt -verbose
   ```

# Example 📝

### Input File `(input.txt)`

```txt
This is a long text that needs to be split into multiple tweets because it exceeds the character limit of a single tweet.
The goal is to ensure that no word is split in the middle and that each tweet is concise and readable.
This app will handle the splitting and add thread numbering to make it easy to follow.
```

### Command

```bash
./tweet-slicer -file input.txt -output output.txt -verbose
```

### Output File `(output.txt)`

```txt
1/3 This is a long text that needs to be split into multiple tweets because it exceeds the character limit of a single tweet. The goal is to ensure that no word is split in the middle and that each tweet is concise and readable.
---
2/3 This app will handle the splitting and add thread numbering to make it easy to follow.
---
```

### Verbose Output

```txt
Tweets saved to output.txt
Generated 2 tweets.
Max tweet length: 280 characters.
Tweets were written to the output file (overwritten if it existed).
```

# Command-Line Flags 🎯

| Flag         | Description                                                      |
| ------------ | ---------------------------------------------------------------- |
| `-text`      | Direct text input to split into tweets.                          |
| `-file`      | Path to a text file containing the input.                        |
| `-output`    | Path to save the output tweets (e.g., `output.txt`).             |
| `-append`    | Append tweets to the output file instead of overwriting it.      |
| `-maxlength` | Maximum length of each tweet (default: 280).                     |
| `-verbose`   | Print additional information (e.g., number of tweets generated). |

# Contributing 🤝

Contributions are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

### Happy Slicing!
