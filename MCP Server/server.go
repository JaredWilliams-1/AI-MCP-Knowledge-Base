package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("HI")

    reader := bufio.NewReader(os.Stdin)
    
    for {
        //Getting the content length 

        fmt.Println("1")

        header, _ := reader.ReadString('\n')
        length := parseContentLength(header)

        fmt.Println("2")

        reader.ReadString('\n')

        body := make([]byte, length)
        io.ReadFull(reader, body)
        
        fmt.Println("3")

        input := parseInput(body) // input is parsed into one of the classes
        input = input // remove this later it was just to get rid of errors

        fmt.Println(body)
        break;
    }

    // General flow:
    /*
        Receive input
        Parse into one of the classes
        determine type of input
            - req
            - notif
            - i cant remember the other ones lol XD
        
        if req
            process
            get return value
            put it into class
            parse into json-formatted string
            send it out
    */
}

func parseContentLength(header string) int {

    start := strings.Index(header, " ")
    end := strings.Index(header, "\r")
    // end := strings.Index(header, "a")

    conLen := header[start+1:end]
    num, err := strconv.Atoi(conLen)

    if err != nil {
        fmt.Printf("Error", err)
        return -1
    }

    return num
}

func parseInput(a []byte) []byte {
    inputStr := (string)(a)

    return ([]byte)(inputStr) // temporary
}