package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluted := evaluator.Eval(program, env)
		if evaluted != nil {
			io.WriteString(out, evaluted.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const neko = `
       \'-._           __
        \\  '-..____,.'  '.
         :'.         /    \'.
         :  )       :      : \
          ;'        '   ;  |  :
          )..      .. .:.'.;  :
         /::...  .:::...   ' ;
         ; _ '    __        /:\
         ':o>   /\o_>      ;:. '.
        '-'.__ ;   __..--- /:.   \
        === \_/   ;=====_.':.     ;
         ,/''--'...'--....        ;
              ;                    ;
            .'                      ;
          .'                        ;
        .'     ..     ,      .       ;
       :       ::..  /      ;::.     |
      /      '.;::.  |       ;:..    ;
     :         |:.   :       ;:.    ;
     :         ::     ;:..   |.    ;
      :       :;      :::....|     |
      /\     ,/ \      ;:::::;     ;
    .:. \:..|    :     ; '.--|     ;
   ::.  :''  '-.,,;     ;'   ;     ;
.-'. _.'\      / ';      \,__:      \
'---'    '----'   ;      /    \,.,,,/
                   '----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, neko)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
