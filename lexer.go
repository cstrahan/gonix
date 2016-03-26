package main

//go:generate ragel -G2 -Z lexer.rl -o lexer-generated.go.in
//go:generate echo $PATH > foo.txt
//go:generate cpp -traditional-cpp -P -undef -Wundef -std=c99 -nostdinc -Wtrigraphs -fdollars-in-identifiers -C lexer-generated.go.in > lexer-generated.go
//go:generate rm lexer-generated.go.in
