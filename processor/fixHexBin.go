package processor

import (
	"math/big"
	"strconv"
	"strings"
)

func HexToDecimal(word string) string {
	n := new(big.Int)
	_, ok := n.SetString(word, 16)
	if !ok {
		return word 
	}
	return n.String() 
}

func BinToDecimal(word string) string {
	num, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(num))
}

func FixBinHex(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" {
			if i-1 >= 0 {
				prefix := ""
				suffix := ""
				word := words[i-1]

				// Если есть кавычки в начале или конце слова — запоминаем их
				if strings.HasPrefix(word, "\"") {
					prefix = "\""
					word = strings.TrimPrefix(word, "\"")
				}
				if strings.HasSuffix(word, "\"") {
					suffix = "\""
					word = strings.TrimSuffix(word, "\"")
				}

				// Конвертируем само значение
				converted := BinToDecimal(word)

				// Собираем обратно с кавычками
				words[i-1] = prefix + converted + suffix
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if words[i] == "(hex)" {
			if i-1 >= 0 {
				prefix := ""
				suffix := ""
				word := words[i-1]

				if strings.HasPrefix(word, "\"") {
					prefix = "\""
					word = strings.TrimPrefix(word, "\"")
				}
				if strings.HasSuffix(word, "\"") {
					suffix = "\""
					word = strings.TrimSuffix(word, "\"")
				}

				converted := HexToDecimal(word)

				words[i-1] = prefix + converted + suffix
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}
	}
	return words
}

