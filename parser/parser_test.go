package parser

import (
	"monkey-lang/ast"
	"monkey-lang/lexer"
	"testing"
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
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not return 3 statements, returned %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stm := program.Statements[i]
		if !testLetStatement(t, stm, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stm ast.Statement, name string) bool {
	if stm.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() not return 'let', returned %s", stm.TokenLiteral())
		return false
	}
	letStmt, ok := stm.(*ast.LetStatement)
	if !ok {
		t.Errorf("stm not *ast.LetStatement, returned %d", stm)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', returned '%s'", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("let.Stmt.Name.TokenLiteral() not return '%s', returned '%s'", name, letStmt.TokenLiteral())
	}

	return true
}
