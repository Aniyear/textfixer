package processor

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

func ProcessText(input string) string {
    // 1. Обработка кавычек
    // 2. Обработка (hex), (bin)
    // 3. Обработка (up), (low), (cap)
    // 4. Обработка пунктуации
    // 5. Обработка "a"/"an"
    input = convertHexAndBin(input)
    input = applyTransformations(input)
    return input
}

//==================================================================================================
// convertHexAndBin заменяет `число (hex)` и `число (bin)` на десятичные значения
//==================================================================================================
func convertHexAndBin(text string) string {
    // Регулярное выражение: слово, пробел, (hex|bin)
    re := regexp.MustCompile(`(?i)\b([a-fA-F0-9]+) \((hex|bin)\)`)
    return re.ReplaceAllStringFunc(text, func(match string) string {
        parts := re.FindStringSubmatch(match)
        value := parts[1]
        mode := strings.ToLower(parts[2])

        var num int64
        var err error

        if mode == "hex" {
            num, err = strconv.ParseInt(value, 16, 64)
        } else {
            num, err = strconv.ParseInt(value, 2, 64)
        }

        if err != nil {
            return match // если ошибка, возвращаем как есть
        }
        return fmt.Sprintf("%d", num)
    })
}

//==================================================================================================
// applyTransformations применяет все трансформации к тексту
//==================================================================================================
func applyTransformations(text string) string {
    words := strings.Fields(text)
    var result []string

    for i := 0; i < len(words); i++ {
        word := words[i]

        if match := parseCommand(word); match != nil {
            n := match.count
            start := max(0, len(result)-n)
            for j := start; j < len(result); j++ {
                result[j] = applyCase(result[j], match.cmd)
            }
            continue // пропускаем команду, не добавляя её в результат
        }

        result = append(result, word)
    }

    return strings.Join(result, " ")
}


type command struct {
    cmd   string
    count int
}

// parseCommand парсит строку вида (up), (cap, 3) и возвращает команду
func parseCommand(s string) *command {
    re := regexp.MustCompile(`^\((up|low|cap)(?:,\s*(\d+))?\)$`)
    m := re.FindStringSubmatch(strings.ToLower(s))
    if m == nil {
        return nil
    }

    cmd := m[1]
    count := 1
    if m[2] != "" {
        parsed, err := strconv.Atoi(m[2])
        if err == nil {
            count = parsed
        }
    }

    return &command{cmd, count}
}

func applyCase(word, cmd string) string {
    switch cmd {
    case "up":
        return strings.ToUpper(word)
    case "low":
        return strings.ToLower(word)
    case "cap":
        return strings.Title(strings.ToLower(word)) // Title capitalizes only the first letter
    default:
        return word
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

