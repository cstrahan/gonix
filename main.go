package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	// "github.com/davecgh/go-spew/spew"
	"github.com/davecgh/go-spew/spew"
)

var noPos = Pos{}

// var ex = []byte(". } http://foo.com <foo> 123 ~/a in 123.1 inter inherit ++ -> ${ ./foo/bar /** **/ ./baz # asdf\n foobar")

//var ex = []byte("1 + 2 * 3")

//var ex = []byte("{ a, b, a, c }@y: x")
//var ex = []byte("let x = 123; y = 321; in x")
//var ex = []byte("x.a or x.b or x.c")

//var ex = []byte("rec { a = 123; }")

//var ex = []byte("foo.\"abc${bar}def\" ")

//var ex = []byte(`"abc ${foo.bar {}}  baz" ''foo ''$ baz''`)
//var ex = []byte(`! { a = 1; } ? a || true`)
//var ex = []byte(`true == true -> true == true`)
//var ex = []byte(`[ a b c ]`)

//var ex = []byte(`'' ''\n ''`)

// var ex = []byte(`'' ''$ ''`)

// var ex = []byte("foo/bar/baz")
// var ex = []byte("123.45 2.")

//var ex = []byte("123.456")
//var ex = []byte("12345")
//var ex = []byte("builtins")
//var ex = []byte("let x = 12345.0; in let y = x; in y")
//var ex = []byte("\"blah blah\"")
//var ex = []byte("1 + 2.2 + 3")
//var ex = []byte("_foo")
//var ex = []byte("builtins.${\"foo\"}")
//var ex = []byte("[ 1 2 2 ]")
//var ex = []byte("rec { a = 123; b = a+a; }")
//var ex = []byte("(rec { f = 777; a = 10; b = 20; c = 30; d = 40; e = c+1; __overrides = { c = 31; f = 42; }; }).f")

//var ex = []byte("./foo/bar")

//var ex = []byte("fooBarBaz")

//var ex = []byte(`foo: bar`)

func printTok(tok Token) string {
	return fmt.Sprintf("%v:%v (line %v col %v)  %s  [%s]", tok.Pos.Start, tok.Pos.End, tok.Pos.Row, tok.Pos.Col, tok.TokenType, string(tok.Text))
}

//func (self Pos) String() string {
//	return fmt.Sprintf("(%v:%v)", self.Start, self.End)
//}

// var ex = []byte("print \"asdf\"    ")
//var ex = []byte(`''abc ${ { "${ XYZ }" }} "''' def''`)
var ex = []byte(`a b c d e f`)

func main() {
	file, err := ioutil.ReadFile("all-packages.nix")
	_ = file
	if err != nil {
		panic(err)
	}

	start := time.Now()

	p := NewParser(ex)
	expr, err := p.Parse()
	_ = expr
	if err != nil {
		log.Fatalln("parse failed:", err)
	}

	spew.Dump(expr)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed.Milliseconds())
}

// func main() {
// 	file, err := ioutil.ReadFile("all-packages.nix")
// 	if err != nil {
// 		panic(err)
// 	}

// 	start := time.Now()

// 	for i := 0; i < 1000; i++ {
// 		lexer := NewLexer(file)

// 		for {
// 			t, err := lexer.Lex()
// 			if err != nil {
// 				panic(err)
// 			} else if t.TokenType == EOF {
// 				break
// 			}
// 			// fmt.Println(printTok(t))
// 		}
// 	}

// 	t := time.Now()
// 	elapsed := t.Sub(start)
// 	fmt.Println(elapsed.Milliseconds())
// }

// ----------------------

// func main() {
// 	//file, _ := ioutil.ReadFile("/home/cstrahan/src/nixpkgs/pkgs/top-level/all-packages.nix")
// 	//ex = file

// 	symbols := SymbolTable{symbols: map[string]string{}}

// 	data := ex

// 	start := time.Now()
// 	tokens, err := lex(ex)
// 	//for _, tok := range tokens {
// 	//	fmt.Println(printTok(tok))
// 	//}
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	elapsedLex := time.Since(start)

// 	//fmt.Printf("file size: %v\n", len(file))
// 	//fmt.Printf("tokens: %v\n", len(tokens))
// 	//fmt.Printf("lexed in %f\n", elapsedLex.Seconds())
// 	//os.Exit(0)

// 	// synthesize some EOF tokens for the sake of lookahead
// 	eof := Token{TokenType: EOF}
// 	tokens = append(tokens, eof, eof, eof, eof)

// 	start = time.Now()
// 	p := parser{pos: 0, tokens: tokens, data: data, symbols: symbols}
// 	expr, err := p.Parse()
// 	_ = expr
// 	if err != nil {
// 		log.Fatalln("parse failed:", err)
// 	}

// 	elapsedParse := time.Since(start)
// 	fmt.Printf("lexed in %f\n", elapsedLex.Seconds())
// 	fmt.Printf("parsed in %f\n", elapsedParse.Seconds())
// 	fmt.Printf("both in %f\n", (elapsedLex + elapsedParse).Seconds())

// 	//spew.Dump(expr)
// 	//log.Println(expr)

// 	state := &EvalState{
// 		baseEnvDispl: 0,
// 		baseEnv: &Env{
// 			up:       nil,
// 			prevWith: 0,
// 			kind:     EnvPlain,
// 			values:   []Value{},
// 		},
// 		staticBaseEnv: &StaticEnv{
// 			up:     nil,
// 			isWith: false,
// 			vars:   map[Symbol]uint{},
// 		},
// 		symbols: symbols,
// 	}
// 	state.createBaseEnv()

// 	err = expr.BindVars(state.staticBaseEnv)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	//spew.Dump(expr)
// 	//os.Exit(0)

// 	var val Value = nil

// 	err = state.Eval(expr, &val)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(">> UNFORCED:")
// 	//spew.Dump(val)

// 	fmt.Println(">> FORCED:")
// 	err = state.ForceValue(&val, noPos)
// 	if err != nil {
// 		panic(err)
// 	}
// 	spew.Dump(val)
// }
