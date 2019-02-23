package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := input("Enter file path: ")
	fmt.Println("file path read is ", filePath, "\n")
	numLines, lines := ReadFileLines(filePath, true)
	fmt.Print("number of lines is ", numLines)
	fmt.Printf(lines)
}

func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return "  "
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print(prompt)
	// stringRead, _ := reader.ReadString('\n')
	// return stringRead[0 : len(stringRead)-1]
}

// ReadFileLines read file lines with option to reverse these lines
func ReadFileLines(path string, reverse bool) (numLines int, lines string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // What does this line do?

	fileBytes, _ := ioutil.ReadAll(file)
	fileText := string(fileBytes)

	arrayLines := strings.Split(fileText, "\n")
	if reverse {
		arrayLines = reversArray(arrayLines)
	}

	lines = flattenArrayToStr(arrayLines)
	numLines = len(arrayLines)

	return
}
func reversArray(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reversArray(input[1:]), input[0])
}
func flattenArrayToStr(array []string) (text string) {
	text = ""
	for _, ele := range array {
		text += ele + "\n"
	}
	return
}
