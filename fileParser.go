package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "github.com/common-nighthawk/go-figure"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// Read through the file and search for specific string
func readFile(filename, keyword string) {

    file, err := os.Open(filename)

    csvfile, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer csvfile.Close()

    writer := csv.NewWriter(csvfile)
    defer writer.Flush()

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        username := scanner.Text()
        var s = strings.Split(username, ";")
        if strings.Contains(s[0], keyword) {
            r := []string{s[0]}
            err := writer.Write(r)
            checkError("Cannot write to file", err)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}

// CSV checkError
func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

// Take two arguments, the first being a directory read from and the second a keyword to search
func main() {

    argsWithProg := os.Args[1:]
    directory := argsWithProg[0]
    keyword := argsWithProg[1]

    myFigure := figure.NewFigure("File Parser", "", true)
    myFigure.Print()

    start := time.Now()

    searchDir := directory

    fileList := []string{}
    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileList = append(fileList, path)
        return nil
    })

    if err != nil {
        fmt.Println(err)
    }

    for _, filename := range fileList {
        if filename != searchDir {
            readFile(filename, keyword)
        }
    }

    elapsed := time.Since(start)
    fmt.Println("\nElapsed time:", elapsed.Seconds())
}
