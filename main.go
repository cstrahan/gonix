package main

import "log"

// var ex = []byte(". } http://foo.com <foo> 123 ~/a in 123.1 inter inherit ++ -> ${ ./foo/bar /** **/ ./baz # asdf\n foobar")
var ex = []byte(`"abc ${foo.bar} baz" ''foo''`)

// var ex = []byte("foo/bar/baz")
// var ex = []byte("123.45 2.")

// var ex = []byte("./foo/bar ./quux")

func main() {
	log.Println("LEXING")
	lex(ex)
}
