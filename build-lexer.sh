#!/usr/bin/env bash

ragel -G2 -Z -L lexer.rl -o lexer-generated.go.in &&
sed '/^\/\/line/d' -i lexer-generated.go.in &&
cpp -traditional-cpp -P -C -nostdinc -fdollars-in-identifiers lexer-generated.go.in > lexer-generated.go &&
rm lexer-generated.go.in &&
gofmt -s -w lexer-generated.go
