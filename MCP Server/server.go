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
    reader := bufio.NewReader(os.Stdin)
    
    for {
        //Getting the content length 

        header, _ := reader.ReadString('\n')
        length := parseContentLength(header)

        reader.ReadString('\n')

        body := make([]byte, length)
        io.ReadFull(reader, body)
        
        parseInput(body)

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

    conLen := header[start:end]
    num, err := strconv.Atoi(conLen)

    if err != nil {
        fmt.Printf("Error", err)
        return -1
    }

    return num
}

func parseInput(a []byte) []byte {
    return a
}