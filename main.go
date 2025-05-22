package main

import (
    "fmt"
    "os"
    "textfixer/processor"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . input.txt output.txt")
        return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    inputData, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Printf("Failed to read input file: %v\n", err)
        return
    }

    outputData := processor.ProcessText(string(inputData))

    err = os.WriteFile(outputFile, []byte(outputData), 0644)
    if err != nil {
        fmt.Printf("Failed to write output file: %v\n", err)
    }
}
