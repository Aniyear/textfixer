package processor

import (
	"regexp"
	"strings"
)

func ProcessText(text string) string {

	lines := strings.Split(text, "\n")

	for i, line := range lines {
		line = FixPunctuation(line) // Фиксирует пунктуацию в строке изначально
		words := strings.Fields(line) // разбивает строку на слова   

		words = ProcessCommands(words)
		words = FixBinHex(words)
		words = FixArticles(words) 

		lines[i] = strings.Join(words, " ") // склеивает слова обратно в строку
		lines[i] = FixPunctuation(lines[i]) // Фиксирует пунктуацию еще раз после обработки 
	}

	result := strings.Join(lines, "\n") // склеивает строки обратно в текст
	finalresult := ControlCheck(result) // проверяет, остались ли необработанные конструкции, если да — запускает всё снова.

	return finalresult
}


func ControlCheck(text string) string {
	if strings.Contains(text, "(cap)") || strings.Contains(text, "(up)") || strings.Contains(text, "(low)") {
		return ProcessText(text)
	}

	re := regexp.MustCompile(`\((cap|up|low),\s*\d+\)`)
	if re.MatchString(text) {
		return ProcessText(text)
	}

	return text
}