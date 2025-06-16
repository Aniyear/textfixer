package processor

import (
	"regexp"
	"strings"
)

func FixPunctuation(text string) string {
	lines := strings.Split(text, "\n")

	for i := range lines {
		// what( is' it".    ->      what ( is' it ".
		AddSpaceBeforeBraces := regexp.MustCompile(`([^ \t\r\n])(["({[\]])`)
		lines[i] = AddSpaceBeforeBraces.ReplaceAllString(lines[i], "$1 $2 ")
		// word ( up) -> word (up)
		RemoveSpacesAfterBraces := regexp.MustCompile(`(["({[\]])\s*`)
		lines[i] = RemoveSpacesAfterBraces.ReplaceAllString(lines[i], "$1")
		// 1. what ,is it? -> what , is it?
		AddSpaceAfterPreps := regexp.MustCompile(`([,.:;!?)}])(\S)`)
		lines[i] = AddSpaceAfterPreps.ReplaceAllString(lines[i], "$1 $2")
		// '   herro ' toyota -> 'herro' toyota
		RemoveSpacesQuote := regexp.MustCompile(`'\s*(.*?)\s*'`)
		lines[i] = RemoveSpacesQuote.ReplaceAllString(lines[i], " '$1' ")
		// "  toyota " "" " mitsubishi "-> "toyota" "" "mitsubishi"
		RemoveSpaces2Quote := regexp.MustCompile(`"\s*(.*?)\s*"`)
		lines[i] = RemoveSpaces2Quote.ReplaceAllString(lines[i], "\"$1\" ")
		// 2. what , is it? -> what, is it?
		RemoveSpacesBeforePreps := regexp.MustCompile(`\s*([.,:;!?)}[\]-])`)
		lines[i] = RemoveSpacesBeforePreps.ReplaceAllString(lines[i], "$1")
		// убирает лишние пробелы
		RemoveSapces := regexp.MustCompile(`\s+`)
		lines[i] = RemoveSapces.ReplaceAllLiteralString(strings.TrimSpace(lines[i]), " ")

	}

	cleanedText := strings.Join(lines, "\n")

	return cleanedText
}
