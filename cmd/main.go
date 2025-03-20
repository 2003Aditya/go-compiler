package main

import (
	"fmt"

	"github.com/2003Aditya/go-compiler/internal/lexer"
	// "github.com/gobwas/glob/syntax/lexer"
)

func main() {
    fmt.Println("hello")
    input := "let x = 5"

    l := lexer.NewLexer(input)

    for l.Ch != 0 {


        fmt.Printf("first char: %q\n", l.Ch)

        l.ReadChar()
    }
}



