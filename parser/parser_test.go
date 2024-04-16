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
	checkParseErrros(t, p)
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

func checkParseErrros(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error %q", msg)
	}
	t.FailNow()
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrros(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not return 3 statements, returned %d", len(program.Statements))
	}
	for _, stm := range program.Statements {
		returnStm, ok := stm.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stm not *ast.ReturnStatement, returned %d", stm)
			continue
		}
		if returnStm.TokenLiteral() != "return" {
			t.Errorf("returnStmt.Name.Value not '%s', returned return", returnStm.TokenLiteral())
			continue
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrros(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not return 1 statements, returned %d", len(program.Statements))
	}
	stm, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stm not *ast.ExpressionStatement, returned %T", stm)
	}
	iden, ok := stm.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stm.Expression)
	}
	if iden.Value != "foobar" {
		t.Errorf("iden.Value not %s, got=%s", "foobar", iden.Value)
	}
	if iden.TokenLiteral() != "foobar" {
		t.Errorf("iden.TokenLiteral() not %s, got=%s", "foobar", iden.TokenLiteral())
	}

}

func TestIntLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrros(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not return 1 statements, returned %d", len(program.Statements))
	}
	stm, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stm not *ast.ExpressionStatement, returned %T", stm)
	}
	literal, ok := stm.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stm.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("iden.Value not %d, got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("iden.TokenLiteral() not %s, got=%s", "5", literal.TokenLiteral())
	}
}
