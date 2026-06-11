package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

        MCPReq := parseInput(body) // input is parsed into one of the classes
        
        if MCPReq.Method == "initialize" {
            initializeHandler(MCPReq)
        } else if MCPReq.Method == "initialized" {
            initializedHandler(MCPReq)
        }

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


// ======== Functions for parsing input ========

func parseContentLength(header string) int {

    start := strings.Index(header, " ")
    end := strings.Index(header, "\r")
    // end := strings.Index(header, "a")

    conLen := header[start+1:end]
    num, err := strconv.Atoi(conLen)

    if err != nil {
        fmt.Printf("Error: %v", err)
        return -1
    }

    return num
}

func parseInput(input []byte) MCPRequest {
    var req MCPRequest
    err := json.Unmarshal(input, &req)

    if err != nil {
        log.Fatalf("Error parsing JSON: %v", input)
    }
    return req
}

// ======== Handlers ========

func initializeHandler(MCPReq MCPRequest) {
    //parse the intialize parameters here and respond
    var initParams InitializeParams

    err := json.Unmarshal(MCPReq.Params, &initParams)
    if err != nil {
        log.Fatalf("Error parsing JSON: %v", MCPReq.Params)
    }

    // For now I am not actually taking the negotiation into account
    initRes := InitializeResult{
        ProtocolVersion: "2024-11-05",
        Capabilities: ServerCapabilities{
            Logging: struct{}{},
            Prompts: ServerPromptsCapability{
                ListChanged: true,
            },
            Resources: ServerResourcesCapability{
                Subscribe: true,
                ListChanged: true,
            },
            Tools: ServerToolsCapability{
                ListChanged: true,
            },
            Tasks: ServerTasksCapability{
                List:   struct{}{},
                Cancel: struct{}{},
                Requests: ServerTaskRequests{
                    Tools: ServerTaskToolRequests{
                        Call: struct{}{},
                    },
                },
            },
        },
        ServerInfo: ServerInfo{
            Name: "Knowledge Server",
            Title: "Title",
            Version: "1.0.0",
            Description: "Store and retrieve information in text format",
            WebsiteURL: "www.example.com",
        },
    }

    // TODO: put everything in the proper response body and Send the response
}

func initializedHandler(MCPReq MCPRequest) {
    //TODO: Implement this. It should not return anything
}