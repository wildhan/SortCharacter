package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input one line of words (S) :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	charArr := []rune(strings.ToLower(scanner.Text()))
	vowel := ""
	consonant := ""

	for i := 0; i < len(charArr); i++ {
		selectChar := charArr[i]
		colectChar := []rune{selectChar}
		if selectChar == ' ' {
			continue
		}
		for j := i + 1; j < len(charArr); j++ {
			if selectChar == charArr[j] {
				colectChar = append(colectChar, charArr[j])
				charArr[j] = ' '
			}
		}
		if selectChar == 'a' || selectChar == 'i' || selectChar == 'u' || selectChar == 'e' || selectChar == 'o' {
			vowel += string(colectChar)
		} else {
			consonant += string(colectChar)
		}
	}

	fmt.Print("Vowel Characters :\n", vowel, "\n")
	fmt.Print("Consonant Characters :\n", consonant, "\n")
}
