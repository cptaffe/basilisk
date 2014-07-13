package main

import (
	"bufio"
	"fmt"
	"github.com/cptaffe/lang/optim"
	"github.com/cptaffe/lang/parser"
	"io"
	"log"
	"os"
)

type Program struct {
	Str string // string of input
	Len int    // length of tree last time
}

func Compute(s *Program) string {
	t := optim.Eval(parser.Parse(s.Str, "input"))
	if t == nil {
		return "error..."
	}
	var str string
	app := ", "
	for i := 0; i < (len(t.Sub) - s.Len); i++ {
		str += t.Sub[s.Len+i].String()
		if i != (len(t.Sub)-s.Len)-1 {
			str += app
		}
	}
	str = fmt.Sprintf("result: {%s}", str)
	s.Len = len(t.Sub) // set new len
	return str
}

// Read input from stdin & output result to stdout
func readFile(file io.Reader) string {
	r := bufio.NewReader(file)
	var str string
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				return str
			} else {
				log.Print(err)
			}
		}
		if string(b) == "exit" {
			os.Exit(0)
		} else if string(b) == "exec" {
			return str
		} else {
			str += string(b)
		}
	}
}

func main() {
	p := new(Program)
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		p.Str = readFile(file)
	} else {
		p.Str = readFile(os.Stdin)
	}
	Compute(p);
}
