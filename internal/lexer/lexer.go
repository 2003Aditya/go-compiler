package lexer

import (
)



type Lexer struct {

    input         string
    position      int
    readposition  int
    Ch            byte
}


func (l *Lexer) ReadChar() {
    if l.readposition >= len(l.input) {
        l.Ch = 0
    } else {
        l.Ch  = l.input[l.readposition]
    }

    l.position = l.readposition

    l.readposition++
}

func NewLexer(input string) *Lexer {
    l := &Lexer{}
    l.input = input
    l.position = 0
    l.readposition = 0
    l.ReadChar()

    return l;

}
