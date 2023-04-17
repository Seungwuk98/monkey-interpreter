package main

import (
	"fmt"
	"io/ioutil"
	"monkey/excute"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 2 {
		filename := os.Args[1]
		code, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		input := string(code)
		excute.ExcuteProgram(input)
	} else {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
