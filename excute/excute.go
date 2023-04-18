package excute

import (
	"fmt"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

func ExcuteProgram(input string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) == 0 {
		env := object.NewEnvironment()
		evaluator.Eval(program, env)
		return
	}

	for _, msg := range p.Errors() {
		fmt.Printf("parser error: %q\n", msg)
	}
}
