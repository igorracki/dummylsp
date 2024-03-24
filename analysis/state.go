package analysis

import (
	"fmt"
	"igorracki/dummylsp/lsp"
)

type State struct {
    //  Map of f ile names to contents
    Documents map[string]string
}

func NewState() State {
    return State{ Documents: map[string]string{}}
}

func (s *State) OpenDocument(document, text string) {
    s.Documents[document] = text
}

func (s *State) UpdateDocument(uri, text string) {
    s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse  {
    // In reality, this would look up the type in our type analysis code.
    document := s.Documents[uri]
    return lsp.HoverResponse {
        Response: lsp.Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: lsp.HoverResult{
            Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
        },
    }
}

func (s *State) Definiton(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
    return lsp.DefinitionResponse {
        Response: lsp.Response{
            RPC: "2.0",
            ID: &id,
        },
        Result: lsp.Location{
            URI: uri,
            Range: lsp.Range{
                Start: lsp.Position{
                    Line: position.Line - 1,
                    Character: 0,
                },
                End: lsp.Position{
                    Line: position.Line - 1,
                    Character: 0,
                },
            },
        },
    }
}
