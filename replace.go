package main

import (
    "fmt"
    "bufio"
    "os"
    "math/rand"
    "time"
    "strings"
    "log"
)

func main() {
    nouns := []string{}
    nouns, err := readLines("nounlist.txt")
    if err != nil {
        log.Fatalf("readLines", err)
    }

    fmt.Println("\n---------------------------------------------------------")
    fmt.Println("Type in a sentence with X (noun) and Y (noun).")
    fmt.Println("For example:", "The X is in a Y", "The X is hitting Y")
    reader := bufio.NewReader(os.Stdin)
    sentence, _ := reader.ReadString('\n')

    rand.Seed(time.Now().Unix()) 
    index1 := rand.Intn(len(nouns))
    new1 := nouns[index1]

    index2 := rand.Intn(len(nouns))
    new2 := nouns[index2]

    words := strings.Fields(sentence)

    xIndex, yIndex := findIndices(words)

    words[xIndex] = new1
    words[yIndex] = new2

    finalSentence := ""
    for _, word := range words {
        finalSentence += word + " "
    }

    fmt.Println(finalSentence, "\n")
}

func findIndices(words []string) (int, int) {
    indexX, indexY := -1, -1
    //words := strings.Fields(sentence)
    for i, word := range words {
        if word == "X" {
            indexX = i
        } 
        if word == "Y" {
            indexY = i
        }
    }
    return indexX, indexY
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























