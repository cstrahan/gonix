#!/usr/bin/env bash

ragel -G2 -Z lexer.rl -o lexer-generated.go.in &&
cpp -traditional-cpp -P -undef -Wundef -std=c99 -nostdinc -Wtrigraphs -fdollars-in-identifiers -C lexer-generated.go.in > lexer-generated.go &&
rm lexer-generated.go.in &&
gofmt -s -w lexer-generated.go
