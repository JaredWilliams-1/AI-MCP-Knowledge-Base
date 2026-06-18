package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


// tools/write
func ToolsWriteHandler(MCPRequest){
	//Extract the text that is to be appended
	
	
	textToAppend

	// Append the text
	filename := "file.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed opening the file: %s", err)
	}
	defer file.Close()

	if _, err := file.WriteString(textToAppend); err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}

// tools/read
func ToolsReadHandler(){
	//Get the contents of the file
	file, ferr := os.Open("file.txt")
	
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	var sb strings.Builder

	for scanner.Scan() {
		line := scanner.Text() + "\n"
		sb.WriteString(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	fileString := sb.String()
	fmt.Println(fileString)

	//Write the contents of the file
}