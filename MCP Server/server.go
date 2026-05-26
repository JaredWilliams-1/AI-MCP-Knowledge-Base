package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const (
        SERVER_HOST = "localhost"
        SERVER_PORT = "9988"
        SERVER_TYPE = "tcp"
)

func main() {
    fmt.Println("Server Running...")
    server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

    defer server.Close()
    fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
    fmt.Println("Waiting for client...")

    for {
        connection, err := server.Accept()

        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }

        fmt.Println("client connected")
        go processClient(connection)
    }
}

func processClient(connection net.Conn) {
    state := ""
    for {
        buffer := make([]byte, 1024)
        mLen, err := connection.Read(buffer)

        if err != nil {
            fmt.Println("Error reading:", err.Error())
        }

        input := (buffer[:mLen])
        fmt.Println("Received: ", input)

        // get alles in the buffer then parse to json 
        var jsonReq MCPRequest
        err = json.Unmarshal(input, &jsonReq)
        if err != nil {
            fmt.Println("Error parsing JSON:", err)
            return
        }
        
        method := jsonReq.Method
        var res *MCPResponse = nil

        if method == "initialize" && state == "" {
            state = "SYNACK" //dont even know what the equivalent of this is called in MCP
            res = initialize(jsonReq)
        } else if method == "notifications/initliazed" {
            state = "INITIALIZED"
        } else if state == "INITIALIZED" {
            // now we actually go through the list of functions
            switch method {
            case "write":
                res = write(jsonReq)
            case "read":
                res = read(jsonReq)
            default:
                fmt.Println("None of the above options"); 
            }
        }

        if res != nil {
            fmt.Println("Response is nil")
        }

        // convert res to json and then send it back
        var output []string
        output, err = json.Marshal(res)

        _, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
    }
    connection.Close()

}

// initialize method

func initialize(req MCPRequest) *MCPResponse {

    var initParamas InitializeParams
    err := json.Unmarshal(req.Params, &initParamas)
    if err != nil {
        fmt.Println("Error parsing params JSON in initialize:", err)
        return nil
    }

    fmt.Println("Ça march !")

    // JSONRPC string `json:"jsonrpc"`
    // ID string `json:"id"`
    // Result *json.RawMessage `json:"params,omitempty"`
    // Error *json.RawMessage `json:"error,omitempty"`

    result := &MCPResponse{JSONRPC: req.JSONRPC, ID: req.ID, }

    return result
}

// function to read the file
func read(req MCPRequest) *MCPResponse {
    return nil
}

// function to write to the file
func write(req MCPRequest) *MCPResponse {
    return nil
}
