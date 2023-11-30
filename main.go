package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/chzyer/readline"
)

func getTransliteration(word, lang string) string {
	transDict := make(map[string]string)
	for _, t := range transTups[lang] {
		transDict[t[0]] = t[1]
	}

	transliteratedWord := ""
	for _, letter := range word {
		transliteratedWord += transDict[string(letter)]
	}
	return transliteratedWord
}

func colorizeDiff(original, userInput, lang string) string {
	result := ""
	userIndex := 0

	transDict := make(map[string]string)
	for _, t := range transTups[lang] {
		transDict[t[0]] = t[1]
	}

	for _, origLetter := range original {
		expectedTrans := transDict[string(origLetter)]

		if userIndex >= len(userInput) {
			result += fmt.Sprintf("\033[91m%c\033[0m", origLetter)
			continue
		}

		endIndex := userIndex + len(expectedTrans)
		if endIndex > len(userInput) {
			endIndex = len(userInput)
		}
		userTrans := userInput[userIndex:endIndex]

		if strings.ToLower(expectedTrans) != strings.ToLower(userTrans) {
			result += fmt.Sprintf("\033[91m%c\033[0m", origLetter)
		} else {
			result += string(origLetter)
		}

		userIndex += len(userTrans)
	}
	return result
}

func victory(word string) string {
	greenColor := "\033[92m"
	resetColor := "\033[0m"
	return "*" + greenColor + "*" + resetColor + "*" + " " + greenColor + word + resetColor
}

func getSolution(word, lang string) string {
	// Initialize the transliteration dictionary
	transDict := make(map[rune]string)
	for _, t := range transTups[lang] {
		// Convert the first character of the string to a rune
		origRune, _ := utf8.DecodeRuneInString(t[0])
		transDict[origRune] = t[1]
	}

	var result []string
	uniqueLetters := make(map[rune]bool)

	for _, origLetter := range word {
		// Skip if the letter is already processed
		if _, exists := uniqueLetters[origLetter]; !exists {
			uniqueLetters[origLetter] = true
			transliteratedLetter := transDict[origLetter]
			result = append(result, fmt.Sprintf("\033[92m%c\033[0m=%s", origLetter, transliteratedLetter))
		}
	}

	return strings.Join(result, ", ")
}

func main() {
	armenian := flag.Bool("a", false, "Armenian (default)")
	georgian := flag.Bool("g", false, "Georgian")
	amharic := flag.Bool("e", false, "Amharic")
	flag.Parse()

	selectedLanguages := 0

	lang := "hy"
	if *georgian {
		lang = "ka"
		selectedLanguages++
	}
	if *amharic {
		lang = "am"
		selectedLanguages++
	}
	if *armenian {
		lang = "hy"
		selectedLanguages++
	}

	if selectedLanguages > 1 {
		fmt.Println("Please select exactly one language option.")
		os.Exit(1)
	} else if selectedLanguages == 0 {
		lang = "hy"
	}

	// Setup readline
	rl, err := readline.New(">>> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		// Choose a random word
		word := wordList[lang][rand.Intn(len(wordList[lang]))]
		displayWord := word
		correctTransliteration := getTransliteration(word, lang)

		if correctTransliteration == "" {
			fmt.Printf("No transliteration found for: %s\n", word)
			continue
		}

		for {
			fmt.Println("\033[92m>>>\033[0m " + displayWord)
			fmt.Printf(">>> ")
			userInput, err := rl.Readline()
			if err != nil { // Handle Ctrl-D
				fmt.Println("\n\033[92m***\033[0m 👋")
				os.Exit(0)
			}

			userInput = strings.TrimSpace(userInput)

			if userInput == "?" {
				fmt.Println(getSolution(word, lang))
				continue
			}

			if strings.ToLower(userInput) != strings.ToLower(correctTransliteration) {
				coloredWord := colorizeDiff(word, userInput, lang)
				displayWord = coloredWord
			} else {
				fmt.Println(victory(word))
				break // Exit the inner loop for the next word
			}
		}
	}
}
