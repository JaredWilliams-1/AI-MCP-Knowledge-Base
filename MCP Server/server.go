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

        header, _ := reader.ReadString('\n')
        length := parseContentLength(header)

        reader.ReadString('\n')

        body := make([]byte, length)
        io.ReadFull(reader, body)

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

    MCPRes := MCPResponse{
        JSONRPC: "2.0",
        ID: 1,
        Result: &initRes,
    }

    body, err := json.Marshal(MCPRes)

    fmt.Fprintf(os.Stdout, "Content-Length: %d\r\n\r\n", len(body))
    os.Stdout.Write(body)
}

func initializedHandler(MCPReq MCPRequest) {
    fmt.Print("Initialized.")
}