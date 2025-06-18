package processor

import (
	"strings"
	"unicode"
	"math/big"
	"regexp"
)

func ProcessCommands(words []string) []string {
	words = MergeStrings(words)
	re := regexp.MustCompile(`\((cap|up|low)(?:,\s*(-?\d+))?\)`)

	for i := 0; i < len(words); i++ {
		word := words[i]

		if re.MatchString(word) {
			match := re.FindStringSubmatch(word)
			if len(match) > 0 {
				command := match[1]
				num := 1 // по умолчанию одно слово

				// если указано число — парсим его
				// if match[2] != "" {
				// 	bigNum := new(big.Int)
				// 	_, ok := bigNum.SetString(match[2], 10)
				// 	if ok && bigNum.Sign() > 0 {
				// 		maxRange := big.NewInt(int64(i))
				// 		if bigNum.Cmp(maxRange) == 1 {
				// 			bigNum = maxRange
				// 		}
				// 		num = int(bigNum.Int64())
				// 	}
				// }

				if match[2] != "" {
					bigNum := new(big.Int)
					_, ok := bigNum.SetString(match[2], 10)
					if ok {
						num = int(bigNum.Int64())
						if num <= 0 {
							words = append(words[:i], words[i+1:]...)
							i--
							continue
						}
						maxRange := i
						if num > maxRange {
							num = maxRange
						}
					}
				}


				start := i - num
				if start < 0 {
					start = 0
				}

				for j := start; j < i; j++ {
					switch command {
					case "cap":
						words[j] = Capitalize(words[j])
					case "up":
						words[j] = strings.ToUpper(words[j])
					case "low":
						words[j] = strings.ToLower(words[j])
					}
				}
			}
			// удаляем команду из массива
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}


func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	runes := []rune(word)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			for j := i + 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			break
		}
	}
	return string(runes)
}

func MergeStrings(slice []string) []string {
	var result []string
	reStart := regexp.MustCompile(`\((cap|up|low)`)
	reEnd := regexp.MustCompile(`(\d+)\)`)

	i := 0
	for i < len(slice) {
		if reStart.MatchString(slice[i]) {
			if i+1 < len(slice) && reEnd.MatchString(slice[i+1]) {
				result = append(result, slice[i]+" "+slice[i+1])
				i++
			} else {
				result = append(result, slice[i])
			}
		} else {
			result = append(result, slice[i])
		}
		i++
	}
	result = SplitAfterParenthesis(result)
	return result

}

func SplitAfterParenthesis(input []string) []string {
	var result []string

	for _, item := range input {
		index := strings.Index(item, ")")
		if index != -1 {
			result = append(result, item[:index+1])

			if index+1 < len(item) {
				result = append(result, item[index+1:])
			}
		} else {
			result = append(result, item)
		}
	}

	return result
}
