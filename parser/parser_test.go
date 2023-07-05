package parser

import (
	"testing"

	"github.com/kasperbe/monkey/ast"
	"github.com/kasperbe/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
    input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    if program == nil {
        t.Fatalf("Parse program: nil")
    }
    checkParserErrors(t, p)

    if len(program.Statements) != 3 {
        t.Fatalf("Statement count: %d", len(program.Statements))
    }

    tests := []struct {
        expectedIdentifier string
    }{
        {"x"},
        {"y"},
        {"foobar"},
    }

    for i, tt := range tests {
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdentifier) {
            return
        }
    }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("Expected let, got %s", s.TokenLiteral())
        return false
    }

    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("Expected let, got %s", s.TokenLiteral())
        return false
    }
    
    if letStmt.Name.Value != name {
        t.Errorf("Expected %s, got %s", name, letStmt.Name) 
        return false
    }

    if letStmt.Name.TokenLiteral() != name {
        
        t.Errorf("Expected %s, got %s", name, letStmt.Name.TokenLiteral()) 
        return false
    }

    return true
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("Parser error: %q", msg)
    }

    t.FailNow()
}
