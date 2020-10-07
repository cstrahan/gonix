package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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
var ex = []byte(`''abc ${ { "${ XYZ }" }} } "''' def''`)

func main() {

	lexer := NewLexer(ex)

	for {
		t, err := lexer.Lex()
		if t.TokenType == EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(printTok(t))
	}
}

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

type parser struct {
	pos     int
	tokens  []Token
	data    []byte
	symbols SymbolTable
}

func showAttrPath(attrPath []AttrName) string {
	out := ""
	first := true
	for _, i := range attrPath {
		if !first {
			out = out + "."
		} else {
			first = false
		}

		if i.Expr == nil {
			out = out + string(i.Symbol)
		} else {
			out = out + "\"${" + "<EXPR>" + "}" // TODO
			//out = out + "\"${" + i.Expr + "}" // TODO
		}
	}

	return out
}

func dupAttrPath(attrPath []AttrName, pos Pos, prevPos Pos) error {
	return fmt.Errorf("attribute ‘%s’ at %s already defined at %s", showAttrPath(attrPath), pos, prevPos)
}

func dupAttrSym(attr Symbol, pos Pos, prevPos Pos) error {
	return fmt.Errorf("attribute ‘%s’ at %s already defined at %s", attr, pos, prevPos)
}

func addAttr(attrs *ExprAttrs, attrPath []AttrName, expr Expr, pos Pos) error {
	for _, attr := range attrPath[:len(attrPath)-1] {
		if attr.Expr == nil {
			if existing, ok := attrs.AttrDefs[string(attr.Symbol)]; ok {
				if !existing.Inherited {
					if prevAttrs, ok := existing.Expr.(*ExprAttrs); ok {
						attrs = prevAttrs
					} else {
						return dupAttrPath(attrPath, pos, existing.Pos)
					}
				} else {
					return dupAttrPath(attrPath, pos, existing.Pos)
				}
			} else {
				nested := &ExprAttrs{
					AttrDefs:        map[string]AttrDef{},
					DynamicAttrDefs: map[string]DynamicAttrDef{},
				}
				attrs.AttrDefs[string(attr.Symbol)] = AttrDef{
					Expr:      nested,
					Inherited: false,
				}
				attrs = nested
			}
		} else {
			nested := &ExprAttrs{
				AttrDefs:        map[string]AttrDef{},
				DynamicAttrDefs: map[string]DynamicAttrDef{},
			}

			attrs.DynamicAttrDefs[string(attr.Symbol)] = DynamicAttrDef{
				NameExpr:  attr.Expr,
				ValueExpr: nested,
			}
			attrs = nested
		}
	}

	attr := attrPath[len(attrPath)-1]
	if attr.Expr == nil {
		if existing, ok := attrs.AttrDefs[string(attr.Symbol)]; ok {
			_ = existing
			return dupAttrPath(attrPath, pos, existing.Pos)
		} else {
			attrs.AttrDefs[string(attr.Symbol)] = AttrDef{
				Expr:      expr,
				Inherited: false,
			}
		}
	}

	return nil
}

func exprPtr(expr Expr) *Expr {
	return &expr
}

func addFormal(pos Pos, formals *Formals, formal Formal) error {
	if _, ok := formals.ArgNames[string(formal.Name)]; ok {
		return fmt.Errorf("duplicate formal function argument ‘%s’ at %s", formal.Name, pos)
	}

	formals.ArgNames[string(formal.Name)] = struct{}{}
	formals.Formals = append(formals.Formals, formal)
	return nil
}

func (self *parser) copyString(content []byte) *ExprString {
	sym := self.symbols.Create(string(content))
	var sval Value = &NixString{
		String:  []byte(sym),
		Context: nil,
	}

	return &ExprString{
		String: sym,
		Value:  &sval,
	}
}

func (self *parser) stripIndentation(pos Pos, exprs []Expr) Expr {
	if len(exprs) == 0 {
		return self.copyString([]byte{})
	}

	atStartOfLine := true
	minIndent := 1000000
	curIndent := 0

	// Figure out minimum indentation.
	for _, expr := range exprs {
		e, ok := expr.(*exprIndStr)
		if !ok {
			// Anti-quotations end the current start-of-line whitespace.
			if atStartOfLine {
				atStartOfLine = false
				if curIndent < minIndent {
					minIndent = curIndent
				}
			}
			continue
		}
		for j := 0; j < len(e.String); j++ {
			if atStartOfLine {
				if e.String[j] == ' ' {
					curIndent++
				} else if e.String[j] == '\n' {
					curIndent = 0
				} else {
					atStartOfLine = false
					if curIndent < minIndent {
						minIndent = curIndent
					}
				}
			} else if e.String[j] == '\n' {
				atStartOfLine = true
				curIndent = 0
			}
		}
	}

	// Strip spaces from each line.
	exprs2 := []Expr{}
	atStartOfLine = true
	curDropped := 0
	for n, i := range exprs {
		e, ok := i.(*exprIndStr)
		if !ok {
			atStartOfLine = false
			curDropped = 0
			exprs2 = append(exprs2, i)
			continue
		}

		str2 := ""
		for j := 0; j < len(e.String); j++ {
			if atStartOfLine {
				if e.String[j] == ' ' {
					if curDropped >= minIndent {
						str2 += string(e.String[j])
					}
					curDropped++
				} else if e.String[j] == '\n' {
					curDropped = 0
					str2 += string(e.String[j])
				} else {
					atStartOfLine = false
					curDropped = 0
					str2 += string(e.String[j])
				}
			} else {
				str2 += string(e.String[j])
				if e.String[j] == '\n' {
					atStartOfLine = true
				}
			}
		}

		// Remove the last line if it is empty and consists only of spaces.
		if n == len(exprs)-1 {
			idx := strings.LastIndexByte(str2, '\n')
			if idx != -1 {
				for k := idx + 1; k < len(str2); k++ {
					if str2[k] != ' ' {
						goto Done
					}
				}

				str2 = str2[0:idx]
			}
		}

	Done:
		exprs2 = append(exprs2, &ExprString{String: self.symbols.Create(str2)})
	}

	// If this is a single string, then don't do a concatenation.
	if len(exprs2) == 1 {
		if e, ok := exprs2[0].(*ExprString); ok {
			return e
		}
	}

	return &ExprConcatStrings{Exprs: exprs2, ForceString: true}
}

func (self *parser) tok(offset int) Token {
	return self.tokens[self.pos+offset]
}

func (self *parser) advance(offset int) {
	self.pos += offset
}

func unexpected(tok Token) error {
	panic(fmt.Sprintf("unexpected '%v' token at %v,%v", tok.TokenType, tok.Pos.Row, tok.Pos.Col))
	return UnexpectedToken{}
}

type UnexpectedToken struct {
}

func (self UnexpectedToken) Error() string {
	return "Unexpected Token"
}

func (self *parser) mustMatch(typ TokenType, offset int) (Token, error) {
	tok := self.tok(offset)
	if tok.TokenType != typ {
		return tok, unexpected(tok)
	}
	return tok, nil
}

func (self *parser) Parse() (Expr, error) {
	expr, err := self.parseExpr()
	if err != nil {
		return nil, err
	}

	tok := self.tok(0)
	if tok.TokenType != EOF {
		return nil, fmt.Errorf("Unexpected token '%s'", printTok(tok))
	}

	return expr, nil
}

func (self *parser) parseExpr() (Expr, error) {
	return self.parseExprFunction()
}

func (self *parser) parseExprFunction() (Expr, error) {
	t1 := self.tok(0)

	switch t1.TokenType {
	case ID:
		t2 := self.tok(1)

		switch t2.TokenType {
		case COLON:
			self.advance(2)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			lambda := &ExprLambda{
				Pos:        t1.Pos,
				Arg:        self.symbols.Create(string(t1.Text)),
				MatchAttrs: false,
				Formals:    Formals{Formals: []Formal{}, ArgNames: map[string]struct{}{}},
				Body:       body,
			}

			return lambda, nil
		case AT:
			_, err := self.mustMatch(LCURLY, 2)
			if err != nil {
				return nil, err
			}

			self.advance(3)

			formals, err := self.parseFormals()
			if err != nil {
				return nil, err
			}

			_, err = self.mustMatch(RCURLY, 0)
			if err != nil {
				return nil, err
			}

			_, err = self.mustMatch(COLON, 1)
			if err != nil {
				return nil, err
			}

			self.advance(2)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			lambda := &ExprLambda{
				Pos:        t1.Pos,
				Arg:        self.symbols.Create(string(t1.Text)),
				MatchAttrs: true,
				Formals:    formals,
				Body:       body,
			}

			return lambda, nil
		}
	case ASSERT:
		self.advance(1)

		cond, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(SCOLON, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		body, err := self.parseExprFunction()
		if err != nil {
			return nil, err
		}

		return &ExprAssert{t1.Pos, cond, body}, nil
	case WITH:
		self.advance(1)

		attrs, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(SCOLON, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		body, err := self.parseExprFunction()
		if err != nil {
			return nil, err
		}

		return &ExprWith{t1.Pos, attrs, body}, nil
	case LET:
		if self.tok(1).TokenType == '{' {
			goto Fallthrough
		}

		self.advance(1)

		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		if len(binds.DynamicAttrDefs) != 0 {
			return nil, fmt.Errorf("dynamic attributes not allowed in let at %s", t1)
		}

		_, err = self.mustMatch(IN, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		body, err := self.parseExprFunction()
		if err != nil {
			return nil, err
		}

		return &ExprLet{t1.Pos, binds, body}, nil
	case LCURLY:
		if !self.isFunctionWithFormals() {
			goto Fallthrough
		}

		self.advance(1)

		formals, err := self.parseFormals()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)

		t2 := self.tok(1)
		switch t2.TokenType {
		case AT:
			nameTok, err := self.mustMatch(ID, 2)
			if err != nil {
				return nil, err
			}

			_, err = self.mustMatch(':', 3)
			if err != nil {
				return nil, err
			}

			self.advance(4)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			lambda := &ExprLambda{
				Pos:        t1.Pos,
				Arg:        self.symbols.Create(string(nameTok.Text)),
				MatchAttrs: true,
				Formals:    formals,
				Body:       body,
			}

			return lambda, nil
		case COLON:
			self.advance(2)

			body, err := self.parseExprFunction()
			if err != nil {
				return nil, err
			}

			lambda := &ExprLambda{
				Pos:        t1.Pos,
				Arg:        "",
				MatchAttrs: true,
				Formals:    formals,
				Body:       body,
			}

			return lambda, nil
		}
	default:
		goto Fallthrough
	}

Fallthrough:
	return self.parseExprIf()
}

func (self *parser) isTok(offset int, typ TokenType) bool {
	return self.tok(offset).TokenType == typ
}

// Use this to disambiguate a function with formal arguments from an attr-set.
//
// The subsequent tokens are part of a function with formals when they are one
// of the following sequences:
//
//     '{' ID '}'
//     '{' ID ','
//     '{' ID '?'
//     '{' '}' ':'
//     '{' '}' '@'
//     '{' ELLIPSIS
//
// Otherwise, the tokens begin an attr-set.
func (self *parser) isFunctionWithFormals() bool {
	if !self.isTok(0, '{') {
		return false
	}

	return (self.isTok(1, ID) && self.isTok(2, '}')) ||
		(self.isTok(1, ID) && self.isTok(2, ',')) ||
		(self.isTok(1, ID) && self.isTok(2, '?')) ||
		(self.isTok(1, '}') && self.isTok(2, ':')) ||
		(self.isTok(1, '}') && self.isTok(2, '@')) ||
		self.isTok(1, ELLIPSIS)
}

func undefined() Expr {
	return nil
}

func (self *parser) parseExprIf() (Expr, error) {
	t1 := self.tok(0)
	if t1.TokenType == IF {
		self.advance(1)

		cond, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(THEN, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		then, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(ELSE, 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)

		els, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		return &ExprIf{t1.Pos, cond, then, els}, nil
	} else {
		return self.parseExprOp()
	}
}

var tokToBinOp = [...]BinOp{
	EQ:     BinOpEQ,
	NEQ:    BinOpNEQ,
	'<':    BinOpLE,
	LEQ:    BinOpLEQ,
	'>':    BinOpGE,
	GEQ:    BinOpGEQ,
	AND:    BinOpAnd,
	OR:     BinOpOr,
	IMPL:   BinOpImpl,
	UPDATE: BinOpUpdate,
	'?':    BinOpHasAttr,
	'-':    BinOpSub,
	'*':    BinOpMult,
	'/':    BinOpDiv,
	CONCAT: BinOpConcat,
}

var tokToUnOp = [...]UnOp{
	'!':    UnOpNot,
	NEGATE: UnOpNegate,
}

func (self *parser) parseExprOp() (Expr, error) {
	ops := []Token{}
	exprs := []Expr{}

	// This is, more or less, the shunting-yard algorithm, though extended to
	// support unary operators and non-associative binary operators.
	for {
		// push all unary operators
		for {
			tok := self.tok(0)
			if tok.TokenType == NOT {
				ops = append(ops, tok)
				self.advance(1)
			} else if tok.TokenType == MINUS {
				tok.TokenType = NEGATE
				ops = append(ops, tok)
				self.advance(1)
			} else {
				break
			}
		}

		// parse/push expr
		expr, err := self.parseExprApp()
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)

		// next op
		op1 := self.tok(0)
		if isOperator(op1.TokenType) {
			self.advance(1)
			op1_ := operators[op1.TokenType]

			for len(ops) > 0 {
				op2 := ops[len(ops)-1]
				op2_ := operators[op2.TokenType]
				if op1_.Assoc == ASSOC_NONE && op2_.Assoc == ASSOC_NONE && op1.TokenType == op2.TokenType {
					return nil, unexpected(op1)
				} else if (op1_.Assoc == ASSOC_LEFT || op1_.Assoc == ASSOC_NONE) && op1_.Prec <= op2_.Prec ||
					op1_.Assoc == ASSOC_RIGHT && op1_.Prec < op2_.Prec {
					ops = ops[:len(ops)-1] // pop op2 off stack
					switch op2.TokenType {
					case '!', NEGATE:
						exprs[len(exprs)-1] = &ExprUnOp{op2.Pos, tokToUnOp[op2.TokenType], exprs[len(exprs)-1]}
					case '+':
						exprs[len(exprs)-2] = &ExprConcatStrings{
							Exprs: []Expr{
								exprs[len(exprs)-2],
								exprs[len(exprs)-1],
							},
							ForceString: false,
						}
						exprs = exprs[:len(exprs)-1]
					case EQ, NEQ, '<', LEQ, '>', GEQ, AND, OR, IMPL, UPDATE, '?', '-', '*', '/', CONCAT:
						exprs[len(exprs)-2] = &ExprBinOp{op2.Pos, tokToBinOp[op2.TokenType], exprs[len(exprs)-2], exprs[len(exprs)-1]}
						exprs = exprs[:len(exprs)-1]
					}
				} else {
					break
				}
			}

			// push the new op to the stack
			ops = append(ops, op1)
		} else {
			// no more operators/exprs; pop remaining ops from stack
			// TODO: find good way to dedupe this
			for len(ops) > 0 {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1] // pop op off stack
				switch op.TokenType {
				case '!', NEGATE:
					exprs[len(exprs)-1] = &ExprUnOp{op.Pos, tokToUnOp[op.TokenType], exprs[len(exprs)-1]}
				case '+':
					exprs[len(exprs)-2] = &ExprConcatStrings{
						Exprs: []Expr{
							exprs[len(exprs)-2],
							exprs[len(exprs)-1],
						},
						ForceString: false,
					}
					exprs = exprs[:len(exprs)-1]
				case EQ, NEQ, '<', LEQ, '>', GEQ, AND, OR, IMPL, UPDATE, '?', '-', '*', '/', CONCAT:
					exprs[len(exprs)-2] = &ExprBinOp{op.Pos, tokToBinOp[op.TokenType], exprs[len(exprs)-2], exprs[len(exprs)-1]}
					exprs = exprs[:len(exprs)-1]
				}
			}

			return exprs[0], nil
		}

	}

	return exprs[0], nil
}

const (
	ASSOC_LEFT  = 0
	ASSOC_RIGHT = 1
	ASSOC_NONE  = 2
)

type Op struct {
	Assoc int
	Prec  int
}

var operators = [...]Op{
	IMPL:   Op{Assoc: ASSOC_NONE, Prec: 0},
	OR:     Op{Assoc: ASSOC_LEFT, Prec: 1},
	AND:    Op{Assoc: ASSOC_LEFT, Prec: 2},
	EQ:     Op{Assoc: ASSOC_NONE, Prec: 3},
	NEQ:    Op{Assoc: ASSOC_NONE, Prec: 3},
	'<':    Op{Assoc: ASSOC_LEFT, Prec: 4},
	'>':    Op{Assoc: ASSOC_LEFT, Prec: 4},
	LEQ:    Op{Assoc: ASSOC_LEFT, Prec: 4},
	GEQ:    Op{Assoc: ASSOC_LEFT, Prec: 4},
	UPDATE: Op{Assoc: ASSOC_RIGHT, Prec: 5},
	NOT:    Op{Assoc: ASSOC_LEFT, Prec: 6},
	'+':    Op{Assoc: ASSOC_LEFT, Prec: 7},
	'-':    Op{Assoc: ASSOC_LEFT, Prec: 7},
	'*':    Op{Assoc: ASSOC_LEFT, Prec: 8},
	'/':    Op{Assoc: ASSOC_LEFT, Prec: 8},
	CONCAT: Op{Assoc: ASSOC_RIGHT, Prec: 9},
	'?':    Op{Assoc: ASSOC_NONE, Prec: 10},
	NEGATE: Op{Assoc: ASSOC_NONE, Prec: 11},
}

func isOperator(ty TokenType) bool {
	return ty == IMPL ||
		ty == OR ||
		ty == AND ||
		ty == EQ ||
		ty == NEQ ||
		ty == '<' ||
		ty == '>' ||
		ty == LEQ ||
		ty == GEQ ||
		ty == UPDATE ||
		ty == NOT ||
		ty == '+' ||
		ty == '-' ||
		ty == '*' ||
		ty == '/' ||
		ty == CONCAT ||
		ty == '?' ||
		ty == NEGATE
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (self *parser) parseExprApp() (Expr, error) {
	tok := self.tok(0)
	pos := tok.Pos

	expr, err := self.parseExprSelect()
	if err != nil {
		return nil, err
	}

	for {
		tok := self.tok(0)
		switch tok.TokenType {
		case ID, INT, FLOAT, '"', IND_STRING_OPEN, PATH, HPATH, SPATH, URI, '(', LET, REC, '{', '[':

			arg, err := self.parseExprSelect()
			if err != nil {
				return nil, err
			}

			expr = &ExprBinOp{
				Type:  BinOpApp,
				Pos:   pos,
				Expr1: expr,
				Expr2: arg,
			}
		default:
			return expr, nil
		}
	}

	return expr, nil
}

func (self *parser) parseExprSelect() (Expr, error) {
	simple, err := self.parseExprSimple()
	if err != nil {
		return nil, err
	}

	tok1 := self.tok(0)
	switch tok1.TokenType {
	case DOT:
		self.advance(1)
		path, err := self.parseAttrPath()
		if err != nil {
			return nil, err
		}

		var def Expr
		tok2 := self.tok(0)
		if tok2.TokenType == OR_KW {
			self.advance(1)
			def, err = self.parseExprSelect()
			if err != nil {
				return nil, err
			}
		}

		return &ExprSelect{tok1.Pos, simple, path, def}, nil
	case OR_KW:
		self.advance(1)
		return &ExprBinOp{
			Type:  BinOpApp,
			Pos:   tok1.Pos,
			Expr1: simple,
			Expr2: &ExprVar{tok1.Pos, self.symbols.Create("or"), false, 0, 0},
		}, nil
	}

	return simple, nil
}

func (self *parser) parseExprSimple() (Expr, error) {
	t1 := self.tok(0)
	switch t1.TokenType {
	case ID:
		self.advance(1)
		return &ExprVar{t1.Pos, self.symbols.Create(string(t1.Text)), false, 0, 0}, nil
	case INT:
		self.advance(1)

		n, err := strconv.ParseInt(string(t1.Text), 10, 32)
		if err != nil {
			return nil, err
		}

		var val Value = (*NixInt)(&n)
		return &ExprInt{&val, n}, nil
	case FLOAT:
		self.advance(1)

		n, _ := strconv.ParseFloat(string(t1.Text), 64)
		var val Value = (*NixFloat)(&n)
		return &ExprFloat{&val, n}, nil
	case '"':
		self.advance(1)

		stringParts, err := self.parseStringParts()
		if err != nil {
			return nil, err
		}
		_, err = self.mustMatch('"', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)
		return stringParts, nil
	case IND_STRING_OPEN:
		self.advance(1)

		stringParts, err := self.parseIndStringParts()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(IND_STRING_CLOSE, 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return self.stripIndentation(t1.Pos, stringParts), nil
	case PATH:
		self.advance(1)
		path := NixPath(t1.Text)
		var val Value = &path
		return &ExprPath{&val, string(t1.Text)}, nil
	case HPATH:
		self.advance(1)
		path := NixPath(t1.Text)
		var val Value = &path
		return &ExprPath{&val, string(t1.Text)}, nil
	case SPATH:
		self.advance(1)
		path := NixPath(t1.Text)
		var val Value = &path
		return &ExprPath{&val, string(t1.Text)}, nil
	case URI:
		self.advance(1)
		return self.copyString(t1.Text), nil
	case '(':
		self.advance(1)

		expr, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(')', 0)
		if err != nil {
			return nil, err
		}

		self.advance(1)
		return expr, nil
	case LET:
		self.advance(1)

		_, err := self.mustMatch('{', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		binds.Recursive = true
		attrPath := []AttrName{
			AttrName{Symbol: self.symbols.Create("body")},
		}
		return &ExprSelect{noPos, binds, attrPath, nil}, nil
	case REC:
		self.advance(1)

		_, err := self.mustMatch('{', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		binds.Recursive = true
		return binds, nil
	case '{':
		self.advance(1)
		binds, err := self.parseBinds()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return binds, err
	case '[':
		self.advance(1)
		list, err := self.parseExprList()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch(']', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return list, err
	}

	return nil, unexpected(t1)
}

func (self *parser) parseStringParts() (Expr, error) {
	parts := []Expr{}

	tok := self.tok(0)

	// Exit early if this is just a simple string literal.
	// We'll make use of the fact that we have an ExprString in this case
	// elsewhere.
	if t2 := self.tok(1); tok.TokenType == STR && t2.TokenType == '"' {
		self.advance(1)
		return self.copyString(tok.Text), nil
	}

	for {
		switch tok.TokenType {
		case STR:
			self.advance(1)

			str := self.copyString(tok.Text)
			parts = append(parts, str)
		case DOLLAR_CURLY:
			self.advance(1)

			expr, err := self.parseExpr()
			if err != nil {
				return nil, err
			}
			parts = append(parts, expr)

			_, err = self.mustMatch(RCURLY, 0)
			if err != nil {
				return nil, err
			}
			self.advance(1)
		default:
			goto Break
		}

		tok = self.tok(0)
	}

Break:
	concatExpr := &ExprConcatStrings{Exprs: parts, ForceString: true}
	return concatExpr, nil
}

func (self *parser) parseIndStringParts() ([]Expr, error) {
	parts := []Expr{}

	tok := self.tok(0)

	for {
		switch tok.TokenType {
		case IND_STR:
			self.advance(1)

			str := &exprIndStr{String: tok.Text}
			parts = append(parts, str)
		case DOLLAR_CURLY:
			self.advance(1)

			expr, err := self.parseExpr()
			if err != nil {
				return nil, err
			}
			parts = append(parts, expr)

			_, err = self.mustMatch(RCURLY, 0)
			if err != nil {
				return nil, err
			}
			self.advance(1)
		default:
			goto Break
		}

		tok = self.tok(0)
	}

Break:
	return parts, nil
}

func (self *parser) parseBinds() (*ExprAttrs, error) {
	binds := &ExprAttrs{
		AttrDefs:        map[string]AttrDef{},
		DynamicAttrDefs: map[string]DynamicAttrDef{},
	}

	for {
		tok := self.tok(0)

		switch tok.TokenType {
		case '}', IN:
			return binds, nil
		}

		switch tok.TokenType {
		case INHERIT:
			self.advance(1)
			tok = self.tok(0)

			switch tok.TokenType {
			case '(':
				self.advance(1)
				tok = self.tok(0)

				expr, err := self.parseExpr()
				if err != nil {
					return binds, err
				}

				_, err = self.mustMatch(')', 0)
				if err != nil {
					return binds, err
				}
				self.advance(1)

				attrs, err := self.parseAttrs()
				if err != nil {
					return binds, err
				}

				_, err = self.mustMatch(';', 0)
				if err != nil {
					return binds, err
				}
				self.advance(1)

				for _, attr := range attrs {
					if _, ok := binds.AttrDefs[string(attr.Symbol)]; ok {
						return binds, fmt.Errorf("attribute ‘%s’ at %s already defined at %s", attr.Symbol, "???", "???")
					}

					binds.AttrDefs[string(attr.Symbol)] = AttrDef{
						Expr: &ExprSelect{
							Expr:     expr,
							AttrPath: []AttrName{AttrName{Symbol: attr.Symbol}},
						},
						Inherited: false,
					}
				}
			default:
				attrs, err := self.parseAttrs()
				if err != nil {
					return binds, err
				}

				_, err = self.mustMatch(';', 0)
				if err != nil {
					return binds, err
				}
				self.advance(1)

				for _, attr := range attrs {
					if _, ok := binds.AttrDefs[string(attr.Symbol)]; ok {
						return binds, fmt.Errorf("attribute ‘%s’ at %s already defined at %s", attr.Symbol, "???", "???")
					}

					binds.AttrDefs[string(attr.Symbol)] = AttrDef{
						Expr: &ExprVar{
							Name: attr.Symbol,
						},
						Inherited: true,
					}
				}
			}
		default:
			attrpath, err := self.parseAttrPath()
			if err != nil {
				return binds, err
			}

			_, err = self.mustMatch('=', 0)
			if err != nil {
				return binds, err
			}
			self.advance(1)

			expr, err := self.parseExpr()
			if err != nil {
				return binds, err
			}

			_, err = self.mustMatch(';', 0)
			if err != nil {
				return binds, err
			}
			self.advance(1)

			err = addAttr(binds, attrpath, expr, tok.Pos)
			if err != nil {
				return binds, err
			}
		}
	}

	return binds, nil
}

func (self *parser) parseAttrs() ([]AttrName, error) {
	attrs := []AttrName{}

	for {
		tok := self.tok(0)

		switch tok.TokenType {
		case ID, OR_KW:
			attr, err := self.parseAttr()
			if err != nil {
				return nil, err
			}

			attrs = append(attrs, AttrName{Symbol: Symbol(attr)})
		case '"', DOLLAR_CURLY:
			strExpr, err := self.parseStringAttr()
			if err != nil {
				return nil, err
			}

			if str, ok := strExpr.(*ExprString); ok {
				attrs = append(attrs, AttrName{Symbol: Symbol(str.String)})
			} else {
				return nil, fmt.Errorf("dynamic attributes not allowed in inherit at %s", "???")
			}
		default:
			return attrs, nil
		}
	}

	return attrs, nil
}

func (self *parser) parseAttrPath() ([]AttrName, error) {
	attrs := []AttrName{}

	for {
		tok := self.tok(0)

		switch tok.TokenType {
		case ID, OR_KW:
			attr, err := self.parseAttr()
			if err != nil {
				return nil, err
			}

			attrs = append(attrs, AttrName{Symbol: Symbol(attr)})
		case '"', DOLLAR_CURLY:
			strExpr, err := self.parseStringAttr()
			if err != nil {
				return nil, err
			}

			if str, ok := strExpr.(*ExprString); ok {
				attrs = append(attrs, AttrName{Symbol: Symbol(str.String)})
			} else {
				attrs = append(attrs, AttrName{Expr: strExpr})
			}
		}

		tok = self.tok(0)
		if tok.TokenType != '.' {
			break
		}
		self.advance(1)
	}

	return attrs, nil
}

func (self *parser) parseAttr() (string, error) {
	tok := self.tok(0)
	self.advance(1)
	switch tok.TokenType {
	case ID:
		return string(tok.Text), nil
	case OR_KW:
		return "or", nil
	default:
		return "", unexpected(tok)
	}
}

func (self *parser) parseStringAttr() (Expr, error) {
	tok := self.tok(0)
	self.advance(1)
	switch tok.TokenType {
	case '"':
		expr, err := self.parseStringParts()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('"', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return expr, err
	case DOLLAR_CURLY:
		expr, err := self.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = self.mustMatch('}', 0)
		if err != nil {
			return nil, err
		}
		self.advance(1)

		return expr, err
	default:
		return nil, unexpected(tok)
	}
}

func (self *parser) parseExprList() (*ExprList, error) {
	elems := []Expr{}
	tok := self.tok(0)

	for {
		if tok.TokenType == ']' {
			break
		}

		elem, err := self.parseExprSelect()
		if err != nil {
			return &ExprList{}, err
		}

		elems = append(elems, elem)

		tok = self.tok(0)
	}

	return &ExprList{Elems: elems}, nil
}

func (self *parser) parseFormals() (Formals, error) {
	formals := Formals{
		ArgNames: map[string]struct{}{},
		Formals:  []Formal{},
		Ellipsis: false,
	}

	tok := self.tok(0)

	for {
		switch tok.TokenType {
		case ELLIPSIS:
			self.advance(1)
			formals.Ellipsis = true
			return formals, nil
		case ',':
			self.advance(1)
		case ID:
			formal, err := self.parseFormal()
			if err != nil {
				return Formals{}, err
			}

			err = addFormal(tok.Pos, &formals, formal)
			if err != nil {
				return Formals{}, err
			}
		default:
			return formals, nil
		}

		tok = self.tok(0)
	}

	return formals, nil
}

func (self *parser) parseFormal() (Formal, error) {
	tokId, err := self.mustMatch(ID, 0)
	if err != nil {
		return Formal{}, err
	}
	self.advance(1)

	if self.tok(0).TokenType == '?' {
		self.advance(1)
		expr, err := self.parseExpr()
		if err != nil {
			return Formal{}, err
		}

		return Formal{Name: self.symbols.Create(string(tokId.Text)), Def: expr}, nil
	} else {
		return Formal{Name: self.symbols.Create(string(tokId.Text)), Def: nil}, nil
	}
}
