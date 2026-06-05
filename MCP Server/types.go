package main

import "encoding/json"

// Req and Res

type MCPError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *MCPErrorData `json:"data,omitempty"`
}

type MCPErrorData struct {
	Supported []string `json:"supported"`
	Requested string   `json:"requested"`
}

type MCPRequest struct {
	ID      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	JSONRPC string          `json:"jsonrpc"`
}

type MCPResponse struct {
	JSONRPC string           `json:"jsonrpc"`
	ID      int              `json:"id"`
	Result  *InitializeResult `json:"result,omitempty"`
	Error   *MCPError        `json:"error,omitempty"`
}

// Initialize Req

type InitializeParams struct {
	ProtocolVersion string       `json:"protocolVersion"`
	Capabilities    Capabilities `json:"capabilities"`
	ClientInfo      ClientInfo   `json:"clientInfo"`
}

type Capabilities struct {
	Roots       RootsCapability       `json:"roots"`
	Sampling    struct{}              `json:"sampling"`
	Elicitation ElicitationCapability `json:"elicitation"`
	Tasks       TasksCapability       `json:"tasks"`
}

type RootsCapability struct {
	ListChanged bool `json:"listChanged"`
}

type ElicitationCapability struct {
	Form struct{} `json:"form"`
	URL  struct{} `json:"url"`
}

type TasksCapability struct {
	Requests TaskRequests `json:"requests"`
}

type TaskRequests struct {
	Elicitation ElicitationRequests `json:"elicitation"`
	Sampling    SamplingRequests    `json:"sampling"`
}

type ElicitationRequests struct {
	Create struct{} `json:"create"`
}

type SamplingRequests struct {
	CreateMessage struct{} `json:"createMessage"`
}

type ClientInfo struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Icons       []Icon `json:"icons"`
	WebsiteURL  string `json:"websiteUrl"`
}

type Icon struct {
	Src      string   `json:"src"`
	MimeType string   `json:"mimeType"`
	Sizes    []string `json:"sizes"`
}

// Initialize Res

type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo         `json:"serverInfo"`
	Instructions    string             `json:"instructions,omitempty"`
}

type ServerCapabilities struct {
	Logging   struct{}                  `json:"logging"`
	Prompts   ServerPromptsCapability   `json:"prompts"`
	Resources ServerResourcesCapability `json:"resources"`
	Tools     ServerToolsCapability     `json:"tools"`
	Tasks     ServerTasksCapability     `json:"tasks"`
}

type ServerPromptsCapability struct {
	ListChanged bool `json:"listChanged"`
}

type ServerResourcesCapability struct {
	Subscribe   bool `json:"subscribe"`
	ListChanged bool `json:"listChanged"`
}

type ServerToolsCapability struct {
	ListChanged bool `json:"listChanged"`
}

type ServerTasksCapability struct {
	List     struct{}           `json:"list"`
	Cancel   struct{}           `json:"cancel"`
	Requests ServerTaskRequests `json:"requests"`
}

type ServerTaskRequests struct {
	Tools ServerTaskToolRequests `json:"tools"`
}

type ServerTaskToolRequests struct {
	Call struct{} `json:"call"`
}

type ServerInfo struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Icons       []Icon `json:"icons"`
	WebsiteURL  string `json:"websiteUrl"`
}
