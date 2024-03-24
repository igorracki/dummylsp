package lsp

type InitializeRequest struct {
    Request
    Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
    ClientInfo *ClientInfo `json:"clientInfo"`
    // there's tons more that goes in here
}

type ClientInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

type InitializeResponse struct {
    Response
    Result InitializeResult `json:"result"`
}

type InitializeResult struct {
    Capabilities ServerCapabilities `json:"capabilities"`
    ServerInfo ServerInfo `json:"serverInfo"`
}

type ServerCapabilities struct {
    TextDocumentSync int `json:"textDocumentSync"`
    HoverProvider bool `json:"hoverProvider"`
    DefinitionProvider bool `json:"definitionProvider"`
}

type ServerInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

func NewInitializedResponse(id int) InitializeResponse {
    return InitializeResponse {
        Response: Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: InitializeResult{
            Capabilities: ServerCapabilities{
                TextDocumentSync: 1,
                HoverProvider: true,
                DefinitionProvider: true,
            },
            ServerInfo: ServerInfo{
                Name: "dummylsp",
                Version: "0.0.0-dummy",
            },
        },
    }
}
