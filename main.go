package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"igorracki/dummylsp/analysis"
	"igorracki/dummylsp/lsp"
	"igorracki/dummylsp/rpc"
	"io"
	"log"
	"os"
)

func main() {
    logger := getLogger("./lsp.log")
    logger.Println("LSP started")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    state := analysis.NewState()
    writer := os.Stdout

    for scanner.Scan()  {
        message := scanner.Bytes()
        method, contents, err := rpc.DecodeMessage(message)
        if err != nil {
            logger.Printf("Got an error: %s", err)
        }
        handleMessage(logger, writer, state, method, contents)
    }
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
    logger.Printf("Received message with method: %s", method)
    switch method {
    case "initialize":
        var request lsp.InitializeRequest
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("[initialize] Could not parse the received message: %s", err)
        }
        logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
        message := lsp.NewInitializedResponse(request.ID)
        writeResponse(writer, message)
    case "textDocument/didOpen":
        var request lsp.DidOpenTextDocumentNotification
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("[textDocument/didOpen] Could not parse the received message: %s", err)
        }
        logger.Printf("Opened: %s, Content: %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
        state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
    case "textDocument/didChange":
        var request lsp.DidChangeTextDocumentNotification
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("[textDocument/didChange] Could not parse the received message: %s", err)
        }
        logger.Printf("Changed: %s", request.Params.TextDocument.URI)
        for _, change := range request.Params.ContentChanges {
            state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
        }
    case "textDocument/hover":
        var request lsp.HoverRequest
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("[textDocument/hover] Could not parse the received message: %s", err)
        }
        response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
        writeResponse(writer, response)
    case "textDocument/definition":
        var request lsp.DefinitionRequest
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("[textDocument/definition] Could not parse the received message: %s", err)
        }
        response := state.Definiton(request.ID, request.Params.TextDocument.URI, request.Params.Position)
        writeResponse(writer, response)
    }
}

func writeResponse(writer io.Writer, message any) {
    reply := rpc.EncodeMessage(message)
    writer.Write([]byte(reply))
}

func getLogger(filename string) * log.Logger {
    logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic(fmt.Sprintf("Could not open a file with name [%s]", filename))
    }
    return log.New(logFile, "[dummylsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
