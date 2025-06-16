package processor

import (
	"strings"
	"unicode"
)

func IsArticle(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "a" || lowerWord == "an"
}

func IsConjunction(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "or" || lowerWord == "and"
}

func StartsWithVowel(word string) bool {
	for _, r := range word {
		if r == '\'' {
			continue
		}
		r = unicode.ToLower(r)
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'y' || r == 'h' && word[1] == 'o'
	}
	return false
}

func StartsWithUpperCase(word string) bool {
	if len(word) > 0 && unicode.IsUpper(rune(word[0])) {
		return true
	}
	return false
}

func FixArticles(words []string) []string {

	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]

		if IsArticle(currentWord) {
			if IsConjunction(nextWord) || IsArticle(nextWord) {
				continue
			}

			var correctedArticle string
			if StartsWithVowel(nextWord) {
				correctedArticle = "an"
			} else {
				correctedArticle = "a"
			}

			if strings.ToUpper(nextWord) == nextWord && (currentWord == "A" || currentWord == "AN") {
				correctedArticle = strings.ToUpper(correctedArticle)
			} else if unicode.IsUpper(rune(currentWord[0])) {

				correctedArticle = Capitalize(correctedArticle)
			}

			words[i] = correctedArticle
		}
	}
	return words
}
