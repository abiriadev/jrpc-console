package main

import (
	"fmt"

	"github.com/abiriadev/jrpc-console/lib"
	"github.com/alecthomas/participle"
	"github.com/chzyer/readline"
)

func main() {
	parser := participle.MustBuild[lib.Invocation]()

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		ast, err := parser.ParseString("", line)
		if err != nil {
			panic(err)
		}

		fmt.Println(ast)
	}
}
