package lsp

type TextDocumentItem struct {
    URI string `json:"uri"`
    LanguageID string `json:"languageId"`
    Version int `json:"version"`
    Text string `json:"text"`
}

type TextDocumentIdentifier struct {
    URI string `json:"uri"`
}

type VersionTextDocumentIdentifier struct {
    TextDocumentIdentifier
    Version int `json:"version"`
}

type TextDocumentContentChangeEvent struct {
    Text string `json:"text"`
}

type DidOpenTextDocumentNotification struct {
    Notification
    Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
    TextDocument TextDocumentItem `json:"textDocument"`
}

type DidChangeTextDocumentNotification struct {
    Notification
    Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
    TextDocument VersionTextDocumentIdentifier `json:"textDocument"`
    ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentPositionParams struct {
    TextDocument TextDocumentIdentifier `json:"textDocument"`
    Position Position `json:"position"`
}

type Position struct {
    Line int `json:"line"`
    Character int `json:"character"`
}

type HoverRequest struct {
    Request
    Params HoverParams `json:"params"`
}

type HoverParams struct {
    TextDocumentPositionParams
}

type HoverResponse struct {
    Response
    Result HoverResult `json:"result"`
}

type HoverResult struct {
    Contents string `json:"contents"`
}

type DefinitionRequest struct {
    Request
    Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
    TextDocumentPositionParams
}

type DefinitionResponse struct {
    Response
    Result Location `json:"result"`
}

type DefinitionResult struct {
    Contents string `json:"contents"`
}

type Location struct {
    URI string `json:"uri"`
    Range Range `json:"range"`
}

type Range struct {
    Start Position `json:"start"`
    End Position `json:"end"`
}
