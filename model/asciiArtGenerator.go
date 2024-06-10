package model

import (
	"errors"
	"os"
	"strings"
    "bufio"
    "fmt"
)
const (
	GetCharacterHeight = 8
	GetNumOfAsciiXters = 95
)
func CheckBannerLineCount(filename string) error {
    const BannerLineCount = 855

    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineCount := 0
    for scanner.Scan() {
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    if lineCount != BannerLineCount {
        return fmt.Errorf("\x1b[38;5;9m[Error] Expected %d but got %d characters in the alphabet. Ensure you have the correct ascii art file\x1b[38;5;9m", BannerLineCount, lineCount)
    }

    return nil
}
func LoadAsciiArtFromFile(filename string) (bool, map[rune][]string) {
    // Check the line count of the banner file
    err := CheckBannerLineCount(filename)
    if err != nil {
        fmt.Println(err)
        return false, nil
    }

    file, err := os.Open(filename)
    if err != nil {
        return false, nil
    }
    defer file.Close()

    // Create a new scanner to read the file
    scanner := bufio.NewScanner(file)
    scanner.Scan()

    // Initialize the map to hold the ASCII art
    AsciiSymbol := make(map[rune][]string)
	// Initialize the current letter and count
	var letter rune = ' '
	var count int
	// Scan through the file line by line
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()
		// If the count is not equal to the character height, append the line to the current letter's ASCII art
		if count != GetCharacterHeight {
			AsciiSymbol[letter] = append(AsciiSymbol[letter], line)
			count++
		} else {
			// If the count is equal to the character height, move on to the next letter and reset the count
			letter++
			count = 0
		}
		// If there's an error scanning the file, print an error message and exit
		if err := scanner.Err(); err != nil {
			fmt.Printf("\x1b[38;5;9m[Internal Error] Error reading file: %v\x1b[38;5;9m", err)
			os.Exit(1)
		}
	}
	// If the number of characters in the ASCII art does not match the expected number, print an error message and exit
	if len(AsciiSymbol) != GetNumOfAsciiXters {
		fmt.Printf("\x1b[38;5;9m[Error] Expected %d but got %d characters in the alphabet. Ensure you have the correct ascii art file\n\x1b[38;5;9m", GetNumOfAsciiXters, len(AsciiSymbol))
		os.Exit(1)
	}
	// If everything went well, return true and the ASCII art
	return true, AsciiSymbol
}
// GenerateAsciiArt generates ASCII art from the given text and style.
func GenerateAsciiArt(text string, style string) (string, error) {
    text = strings.ReplaceAll(text, "\r\n", "\n")
    text = strings.ReplaceAll(text, "\t", "    ")

    lines := strings.Split(text, "\n")

    var asciiArtLines [][][]string
    asciiFile := "model/" + style + ".txt"

    // Load the ASCII art from the file
    _, AsciiSymbol := LoadAsciiArtFromFile(asciiFile)
    if AsciiSymbol == nil {
        return "", errors.New("error loading ASCII art from file")
    }

    for _, line := range lines {
        var asciiArtLine [][]string
        for _, char := range line {
            asciiArtLine = append(asciiArtLine, AsciiSymbol[char])
        }
        asciiArtLines = append(asciiArtLines, asciiArtLine)
    }

    return strings.Join(FormatAsciiArtLines(asciiArtLines), ""), nil
}

// FormatAsciiArtLines formats the given 2D slice of ASCII art lines into a slice of strings.
func FormatAsciiArtLines(asciiArtLines [][][]string) []string {
	var formattedLines []string

	for _, line := range asciiArtLines {
		formattedLines = append(formattedLines, FormatAsciiArt(line))
	}

	return formattedLines
}

// FormatAsciiArt formats the given 2D slice of ASCII art lines into a string.
func FormatAsciiArt(asciiArtLines [][]string) string {
	var formattedArt []string
	var line string

	// For each row of the ASCII art, append the corresponding part of the ASCII art for each character to line.
	for j := 0; j < 8; j++ {
		for _, art := range asciiArtLines {
			if art[0] == "" {
				art[0] = "      "
			}
			line += art[j]
		}
		line += "\n"
		formattedArt = append(formattedArt, line)
		line = ""
	}

	// Join the formatted ASCII art into a single string and return it.
	return strings.Join(formattedArt, "")
}
