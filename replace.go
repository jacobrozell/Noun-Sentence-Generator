package main

import (
    "fmt"
    "bufio"
    "os"
    "math/rand"
    "time"
    "strings"
    "log"
    "errors"
)

func main() {
Start:
    nouns := []string{}
    nouns, err := readLines("nounlist.txt")
    if err != nil {
        log.Fatalf("readLines", err)
    }

    words := strings.Fields(getUserInput())
    xIndex, yIndex, err := findIndices(words)
    if err != nil {
        fmt.Println("Please enter a valid response.")
        goto Start
    }
    words[xIndex], words[yIndex] = generateRandomWords(nouns)
    finalSentence := calcFinalSentence(words)

    fmt.Println(finalSentence, "\n---------------------------------------------------------\n")
}

func getUserInput() string {
    fmt.Println("\n---------------------------------------------------------")
    fmt.Println("Type in a sentence with X (noun) and Y (noun).")
    fmt.Println("For example:", "'The X is in a Y'", "'The X is hitting Y'")
    reader := bufio.NewReader(os.Stdin)
    sentence, _ := reader.ReadString('\n')
    return sentence
}

func generateRandomWords(nouns []string) (string, string) {
    rand.Seed(time.Now().Unix())
    return nouns[rand.Intn(len(nouns))], nouns[rand.Intn(len(nouns))]
}

func calcFinalSentence(words []string) string {
    finalSentence := ""
    for _, word := range words {
        finalSentence += word + " "
    }
    return finalSentence
}

func findIndices(words []string) (int, int, error) {
    indexX, indexY := -1, -1
    for i, word := range words {
        if word == "X" || word == "x" {
            indexX = i
        } 
        if word == "Y" || word == "y"{
            indexY = i
        }
    }
    if indexX == -1 || indexY == -1 {
        return indexX, indexY, errors.New("Enter a valid input.")
    } else {
        return indexX, indexY, nil
    }
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

