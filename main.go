package main

import (
	"fmt"
	processor "go-reloaded/processor"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	text := string(content)

	processedText := processor.ProcessText(text)

	err = os.WriteFile(outputFile, []byte(processedText), 0o644)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Job is done. You can check the file named", outputFile)
}
