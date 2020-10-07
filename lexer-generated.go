//line lexer.rl:1

//line lexer.rl:2

package main

import "fmt"

//line lexer-generated.go.in:12
const scanner_start int = 33
const scanner_first_final int = 33
const scanner_error int = 0

const scanner_en_string int = 80
const scanner_en_ind_string int = 82
const scanner_en_main int = 33

//line lexer.rl:8

func unescapeStr(s []byte) []byte {
	t := []byte{}

	idx := 0
	var c byte
	for idx < len(s) {
		c = s[idx]
		idx++
		switch c {
		case '\\':
			{
				c = s[idx]
				idx++
				switch c {
				case 'n':
					t = append(t, byte('\n'))
				case 'r':
					t = append(t, byte('\r'))
				case 't':
					t = append(t, byte('\t'))
				default:
					t = append(t, byte(c))
				}
			}
		case '\r':
			{
				t = append(t, byte('\n'))
				if s[idx] == '\n' {
					idx++
				}
			}
		default:
			t = append(t, byte(c))
		}
	}

	return t
}

func (l *Lexer) Lex() (Token, error) {
	if l.p == l.pe {
		return Token{TokenType: TokenType(EOF), Pos: Pos{l.pe, l.pe, l.lineCount + 1, (l.pe - l.lineStart) + 1}, Text: nil}, nil
	}

	// load state from l
	data := l.data
	pe := l.pe
	cs := l.cs
	p := l.p
	top := l.top
	stack := l.stack
	lineStart := l.lineStart
	lineCount := l.lineCount
	act := l.act

	eof := pe
	ts, te, act := 0, 0, 0

	token := Token{}

	var err error

//line lexer-generated.go.in:84
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch cs {
		case 33:
			goto st33
		case 1:
			goto st1
		case 0:
			goto st0
		case 2:
			goto st2
		case 34:
			goto st34
		case 3:
			goto st3
		case 4:
			goto st4
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		case 38:
			goto st38
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 39:
			goto st39
		case 8:
			goto st8
		case 40:
			goto st40
		case 41:
			goto st41
		case 9:
			goto st9
		case 10:
			goto st10
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 11:
			goto st11
		case 45:
			goto st45
		case 12:
			goto st12
		case 13:
			goto st13
		case 46:
			goto st46
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 47:
			goto st47
		case 48:
			goto st48
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 49:
			goto st49
		case 20:
			goto st20
		case 50:
			goto st50
		case 51:
			goto st51
		case 21:
			goto st21
		case 22:
			goto st22
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 66:
			goto st66
		case 67:
			goto st67
		case 68:
			goto st68
		case 69:
			goto st69
		case 70:
			goto st70
		case 71:
			goto st71
		case 72:
			goto st72
		case 73:
			goto st73
		case 74:
			goto st74
		case 75:
			goto st75
		case 76:
			goto st76
		case 77:
			goto st77
		case 78:
			goto st78
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 82:
			goto st82
		case 83:
			goto st83
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 84:
			goto st84
		case 85:
			goto st85
		case 32:
			goto st32
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch cs {
		case 33:
			goto st_case_33
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 34:
			goto st_case_34
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 39:
			goto st_case_39
		case 8:
			goto st_case_8
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 11:
			goto st_case_11
		case 45:
			goto st_case_45
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 46:
			goto st_case_46
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 49:
			goto st_case_49
		case 20:
			goto st_case_20
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 32:
			goto st_case_32
		}
		goto st_out
	tr3:
//line lexer.rl:217
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr6:
//line lexer.rl:220
		te = p + 1
		{
			token = Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			stack[top-1].curlyCount += 1
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr7:
//line lexer.rl:202
		te = p + 1
		{
			token = Token{TokenType: TokenType(AND), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr9:
		cs = 33
//line lexer.rl:235
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			cs = 82
			{
				p++
				goto _out
			}
		}
		goto _again
	tr10:
		cs = 33
//line lexer.rl:89
		lineStart = p + 1
		lineCount++
//line lexer.rl:235
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			cs = 82
			{
				p++
				goto _out
			}
		}
		goto _again
	tr12:
//line NONE:1
		switch act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 12:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(IF), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 13:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(THEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 14:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ELSE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 15:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ASSERT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 16:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(WITH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 17:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(LET), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 18:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(IN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 19:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(REC), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 20:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(INHERIT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 21:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(OR_KW), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 22:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ELLIPSIS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 26:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 27:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 38:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(CONCAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 39:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ID), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 40:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(INT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 41:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 46:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 47:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 33
					goto _out
				}
			}
		case 53:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 33
					goto _out
				}
			}
		}

		goto st33
	tr16:
//line lexer.rl:253
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr21:
//line lexer.rl:208
		te = p + 1
		{
			token = Token{TokenType: TokenType(UPDATE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr24:
//line lexer.rl:251
		te = p + 1

		goto st33
	tr26:
//line lexer.rl:198
		te = p + 1
		{
			token = Token{TokenType: TokenType(LEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr28:
//line lexer.rl:243
		te = p + 1
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr29:
//line lexer.rl:200
		te = p + 1
		{
			token = Token{TokenType: TokenType(GEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr33:
//line lexer.rl:204
		te = p + 1
		{
			token = Token{TokenType: TokenType(OR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr54:
		cs = 33
//line lexer.rl:256
		te = p + 1
		{
			token = Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			cs = 80
			{
				p++
				goto _out
			}
		}
		goto _again
	tr59:
//line lexer.rl:266
		te = p + 1
		{
			token = Token{TokenType: TokenType(LPAREN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr60:
//line lexer.rl:268
		te = p + 1
		{
			token = Token{TokenType: TokenType(RPAREN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr61:
//line lexer.rl:191
		te = p + 1
		{
			token = Token{TokenType: TokenType(STAR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr63:
//line lexer.rl:264
		te = p + 1
		{
			token = Token{TokenType: TokenType(COMMA), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr69:
//line lexer.rl:179
		te = p + 1
		{
			token = Token{TokenType: TokenType(COLON), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr70:
//line lexer.rl:181
		te = p + 1
		{
			token = Token{TokenType: TokenType(SCOLON), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr74:
//line lexer.rl:262
		te = p + 1
		{
			token = Token{TokenType: TokenType(QMARK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr75:
//line lexer.rl:183
		te = p + 1
		{
			token = Token{TokenType: TokenType(AT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr77:
//line lexer.rl:270
		te = p + 1
		{
			token = Token{TokenType: TokenType(LBRACK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr78:
//line lexer.rl:272
		te = p + 1
		{
			token = Token{TokenType: TokenType(RBRACK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr88:
//line lexer.rl:230
		te = p + 1
		{
			token = Token{TokenType: TokenType(LCURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			stack[top-1].curlyCount += 1
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr90:
		cs = 33
//line lexer.rl:223
		te = p + 1
		{
			token = Token{TokenType: TokenType(RCURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			stack[top-1].curlyCount -= 1
			if stack[top-1].curlyCount == 0 {
				cs = (stack[top-1].resumeState)
				top--
			}
			{
				p++
				goto _out
			}
		}
		goto _again
	tr92:
//line lexer.rl:217
		te = p
		p--
		{
			token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr94:
//line lexer.rl:248
		te = p
		p--

		goto st33
	tr95:
//line lexer.rl:189
		te = p
		p--
		{
			token = Token{TokenType: TokenType(NOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr96:
//line lexer.rl:196
		te = p + 1
		{
			token = Token{TokenType: TokenType(NEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr97:
//line lexer.rl:250
		te = p
		p--

		goto st33
	tr98:
		cs = 33
//line lexer.rl:235
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			cs = 82
			{
				p++
				goto _out
			}
		}
		goto _again
	tr99:
//line lexer.rl:187
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr101:
//line lexer.rl:239
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr102:
//line lexer.rl:185
		te = p
		p--
		{
			token = Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr103:
//line lexer.rl:206
		te = p + 1
		{
			token = Token{TokenType: TokenType(IMPL), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr104:
//line lexer.rl:253
		te = p
		p--
		{
			token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr108:
//line lexer.rl:215
		te = p
		p--
		{
			token = Token{TokenType: TokenType(INT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr109:
//line lexer.rl:260
		te = p
		p--
		{
			token = Token{TokenType: TokenType(EQS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr110:
//line lexer.rl:194
		te = p + 1
		{
			token = Token{TokenType: TokenType(EQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr112:
//line lexer.rl:213
		te = p
		p--
		{
			token = Token{TokenType: TokenType(ID), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr113:
//line lexer.rl:245
		te = p
		p--
		{
			token = Token{TokenType: TokenType(URI), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr124:
//line lexer.rl:169
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	tr141:
//line lexer.rl:241
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 33
				goto _out
			}
		}
		goto st33
	st33:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line NONE:1
		ts = p

//line lexer-generated.go.in:836
		switch data[p] {
		case 0:
			goto st1
		case 9:
			goto st36
		case 10:
			goto tr52
		case 13:
			goto st36
		case 32:
			goto st36
		case 33:
			goto st37
		case 34:
			goto tr54
		case 35:
			goto st38
		case 36:
			goto st5
		case 38:
			goto st6
		case 39:
			goto st7
		case 40:
			goto tr59
		case 41:
			goto tr60
		case 42:
			goto tr61
		case 43:
			goto tr62
		case 44:
			goto tr63
		case 45:
			goto tr64
		case 46:
			goto tr65
		case 47:
			goto st14
		case 48:
			goto tr67
		case 58:
			goto tr69
		case 59:
			goto tr70
		case 60:
			goto st17
		case 61:
			goto st49
		case 62:
			goto st20
		case 63:
			goto tr74
		case 64:
			goto tr75
		case 91:
			goto tr77
		case 93:
			goto tr78
		case 95:
			goto tr79
		case 97:
			goto tr80
		case 101:
			goto tr81
		case 105:
			goto tr82
		case 108:
			goto tr83
		case 111:
			goto tr84
		case 114:
			goto tr85
		case 116:
			goto tr86
		case 119:
			goto tr87
		case 123:
			goto tr88
		case 124:
			goto st23
		case 125:
			goto tr90
		case 126:
			goto st24
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr68
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 46 {
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
	tr2:
//line NONE:1
		te = p + 1

		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line lexer-generated.go.in:968
		switch data[p] {
		case 69:
			goto st3
		case 101:
			goto st3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto tr92
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 43:
			goto st4
		case 45:
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st35
		}
		goto tr3
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st35
		}
		goto tr3
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if 48 <= data[p] && data[p] <= 57 {
			goto st35
		}
		goto tr92
	tr52:
//line lexer.rl:89
		lineStart = p + 1
		lineCount++
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line lexer-generated.go.in:1021
		switch data[p] {
		case 9:
			goto st36
		case 10:
			goto tr52
		case 13:
			goto st36
		case 32:
			goto st36
		}
		goto tr94
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 61 {
			goto tr96
		}
		goto tr95
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 10:
			goto tr97
		case 13:
			goto tr97
		}
		goto st38
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 123 {
			goto tr6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 38 {
			goto tr7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 39 {
			goto tr8
		}
		goto st0
	tr8:
//line NONE:1
		te = p + 1

		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line lexer-generated.go.in:1091
		switch data[p] {
		case 10:
			goto tr10
		case 32:
			goto st8
		}
		goto tr98
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 10:
			goto tr10
		case 32:
			goto st8
		}
		goto tr9
	tr62:
//line NONE:1
		te = p + 1

//line lexer.rl:187
		act = 27
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line lexer-generated.go.in:1123
		switch data[p] {
		case 43:
			goto tr100
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr99
	tr17:
//line NONE:1
		te = p + 1

//line lexer.rl:177
		act = 22
		goto st41
	tr100:
//line NONE:1
		te = p + 1

//line lexer.rl:210
		act = 38
		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line lexer-generated.go.in:1164
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr12
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr12
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 43:
			goto tr15
		case 95:
			goto tr15
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr15
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr15
				}
			case data[p] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr12
	tr15:
//line NONE:1
		te = p + 1

//line lexer.rl:239
		act = 46
		goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line lexer-generated.go.in:1253
		switch data[p] {
		case 43:
			goto tr15
		case 47:
			goto st10
		case 95:
			goto tr15
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr101
	tr64:
//line NONE:1
		te = p + 1

//line lexer.rl:185
		act = 26
		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line lexer-generated.go.in:1287
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 62:
			goto tr103
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr102
	tr65:
//line NONE:1
		te = p + 1

//line lexer.rl:253
		act = 53
		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line lexer-generated.go.in:1323
		switch data[p] {
		case 43:
			goto st9
		case 45:
			goto st9
		case 46:
			goto st11
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr106
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr104
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 43:
			goto st9
		case 46:
			goto tr17
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr16
	tr106:
//line NONE:1
		te = p + 1

//line lexer.rl:217
		act = 41
		goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line lexer-generated.go.in:1389
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 69:
			goto st12
		case 95:
			goto st9
		case 101:
			goto st12
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto tr106
		}
		goto tr92
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 43:
			goto st13
		case 45:
			goto st13
		case 46:
			goto st9
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr19
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr3
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto tr19
		}
		goto tr3
	tr19:
//line NONE:1
		te = p + 1

//line lexer.rl:217
		act = 41
		goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line lexer-generated.go.in:1493
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto tr19
		}
		goto tr92
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 42:
			goto st15
		case 43:
			goto tr15
		case 47:
			goto tr21
		case 95:
			goto tr15
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
	tr22:
//line lexer.rl:89
		lineStart = p + 1
		lineCount++
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lexer-generated.go.in:1557
		switch data[p] {
		case 10:
			goto tr22
		case 42:
			goto st16
		}
		goto st15
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 10:
			goto tr22
		case 42:
			goto st16
		case 47:
			goto tr24
		}
		goto st15
	tr67:
//line NONE:1
		te = p + 1

//line lexer.rl:215
		act = 40
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lexer-generated.go.in:1591
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st9
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto tr67
		}
		goto tr108
	tr68:
//line NONE:1
		te = p + 1

//line lexer.rl:215
		act = 40
		goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line lexer-generated.go.in:1630
		switch data[p] {
		case 43:
			goto st9
		case 45:
			goto st9
		case 46:
			goto tr106
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr68
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr108
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 43:
			goto st18
		case 61:
			goto tr26
		case 95:
			goto st18
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st18
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st18
				}
			case data[p] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 43:
			goto st18
		case 47:
			goto st19
		case 62:
			goto tr28
		case 95:
			goto st18
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st18
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 43:
			goto st18
		case 95:
			goto st18
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st18
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st18
				}
			case data[p] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 61 {
			goto tr110
		}
		goto tr109
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 61 {
			goto tr29
		}
		goto st0
	tr76:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st50
	tr118:
//line NONE:1
		te = p + 1

//line lexer.rl:163
		act = 15
		goto st50
	tr121:
//line NONE:1
		te = p + 1

//line lexer.rl:161
		act = 14
		goto st50
	tr122:
//line NONE:1
		te = p + 1

//line lexer.rl:157
		act = 12
		goto st50
	tr129:
//line NONE:1
		te = p + 1

//line lexer.rl:173
		act = 20
		goto st50
	tr131:
//line NONE:1
		te = p + 1

//line lexer.rl:167
		act = 17
		goto st50
	tr132:
//line NONE:1
		te = p + 1

//line lexer.rl:175
		act = 21
		goto st50
	tr134:
//line NONE:1
		te = p + 1

//line lexer.rl:171
		act = 19
		goto st50
	tr137:
//line NONE:1
		te = p + 1

//line lexer.rl:159
		act = 13
		goto st50
	tr140:
//line NONE:1
		te = p + 1

//line lexer.rl:165
		act = 16
		goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line lexer-generated.go.in:1837
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr12
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 39:
			goto st51
		case 45:
			goto st51
		case 95:
			goto st51
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st51
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st51
			}
		default:
			goto st51
		}
		goto tr112
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 43:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st21
			}
		default:
			goto st21
		}
		goto tr12
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 33:
			goto st52
		case 61:
			goto st52
		case 95:
			goto st52
		case 126:
			goto st52
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st52
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st52
				}
			case data[p] >= 63:
				goto st52
			}
		default:
			goto st52
		}
		goto tr12
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 33:
			goto st52
		case 61:
			goto st52
		case 95:
			goto st52
		case 126:
			goto st52
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st52
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st52
				}
			case data[p] >= 63:
				goto st52
			}
		default:
			goto st52
		}
		goto tr113
	tr79:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line lexer-generated.go.in:1997
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st9
		case 46:
			goto st9
		case 47:
			goto st10
		case 95:
			goto tr79
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr79
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr79
			}
		default:
			goto tr79
		}
		goto tr112
	tr80:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line lexer-generated.go.in:2035
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 115:
			goto tr114
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr114:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line lexer-generated.go.in:2077
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 115:
			goto tr115
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr115:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line lexer-generated.go.in:2119
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 101:
			goto tr116
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr116:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line lexer-generated.go.in:2161
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 114:
			goto tr117
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr117:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line lexer-generated.go.in:2203
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 116:
			goto tr118
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr81:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line lexer-generated.go.in:2245
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 108:
			goto tr119
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr119:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line lexer-generated.go.in:2287
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 115:
			goto tr120
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr120:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line lexer-generated.go.in:2329
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 101:
			goto tr121
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr82:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line lexer-generated.go.in:2371
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 102:
			goto tr122
		case 110:
			goto tr123
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr123:
//line NONE:1
		te = p + 1

//line lexer.rl:169
		act = 18
		goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line lexer-generated.go.in:2415
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 104:
			goto tr125
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr124
	tr125:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line lexer-generated.go.in:2457
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 101:
			goto tr126
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr126:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line lexer-generated.go.in:2499
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 114:
			goto tr127
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr127:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lexer-generated.go.in:2541
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 105:
			goto tr128
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr128:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line lexer-generated.go.in:2583
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 116:
			goto tr129
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr83:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line lexer-generated.go.in:2625
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 101:
			goto tr130
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr130:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line lexer-generated.go.in:2667
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 116:
			goto tr131
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr84:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line lexer-generated.go.in:2709
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 114:
			goto tr132
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr85:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lexer-generated.go.in:2751
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 101:
			goto tr133
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr133:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line lexer-generated.go.in:2793
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 99:
			goto tr134
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr86:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lexer-generated.go.in:2835
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 104:
			goto tr135
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr135:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line lexer-generated.go.in:2877
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 101:
			goto tr136
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr136:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line lexer-generated.go.in:2919
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 110:
			goto tr137
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr87:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line lexer-generated.go.in:2961
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 105:
			goto tr138
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr138:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line lexer-generated.go.in:3003
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 116:
			goto tr139
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	tr139:
//line NONE:1
		te = p + 1

//line lexer.rl:213
		act = 39
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line lexer-generated.go.in:3045
		switch data[p] {
		case 39:
			goto st51
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr79
		case 104:
			goto tr140
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto tr112
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 124 {
			goto tr33
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 47 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 43:
			goto tr35
		case 95:
			goto tr35
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr35
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr35
				}
			case data[p] >= 65:
				goto tr35
			}
		default:
			goto tr35
		}
		goto tr12
	tr35:
//line NONE:1
		te = p + 1

//line lexer.rl:241
		act = 47
		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line lexer-generated.go.in:3134
		switch data[p] {
		case 43:
			goto tr35
		case 47:
			goto st25
		case 95:
			goto tr35
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr35
			}
		default:
			goto tr35
		}
		goto tr141
	tr36:
//line lexer.rl:114
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]} // ?
			{
				p++
				cs = 80
				goto _out
			}
		}
		goto st80
	tr38:
		cs = 80
//line lexer.rl:103
		te = p + 1
		{
			// don't consume the final " char.
			// we'll set the next state back to string,
			// which will handle emitting the final " char
			// and popping the stack.
			p--

			token = Token{TokenType: TokenType(STR), Pos: Pos{ts, te - 1, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts : te-1]}
			cs = 80
			{
				p++
				goto _out
			}
		}
		goto _again
	tr40:
//line NONE:1
		switch act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 2:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]} // ?
				{
					p++
					cs = 80
					goto _out
				}
			}
		}

		goto st80
	tr42:
		cs = 80
//line lexer.rl:116
		te = p + 1
		{
			token = Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, lexContext{})
			}
			stack[top] = lexContext{1, 80}
			top++
			cs = 33
			{
				p++
				goto _out
			}
		}
		goto _again
	tr142:
		cs = 80
//line lexer.rl:125
		te = p + 1
		{
			token = Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			cs = 33
			{
				p++
				goto _out
			}
		}
		goto _again
	tr144:
//line lexer.rl:114
		te = p
		p--
		{
			token = Token{TokenType: TokenType(STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]} // ?
			{
				p++
				cs = 80
				goto _out
			}
		}
		goto st80
	st80:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line NONE:1
		ts = p

//line lexer-generated.go.in:3234
		switch data[p] {
		case 34:
			goto tr142
		case 36:
			goto st28
		case 92:
			goto st27
		}
		goto tr37
	tr37:
//line NONE:1
		te = p + 1

//line lexer.rl:114
		act = 2
		goto st81
	tr41:
//line NONE:1
		te = p + 1

//line lexer.rl:89
		lineStart = p + 1
		lineCount++
//line lexer.rl:114
		act = 2
		goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line lexer-generated.go.in:3265
		switch data[p] {
		case 34:
			goto tr144
		case 36:
			goto st26
		case 92:
			goto st27
		}
		goto tr37
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 34:
			goto tr38
		case 92:
			goto st27
		case 123:
			goto tr36
		}
		goto tr37
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 10 {
			goto tr41
		}
		goto tr37
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 34:
			goto tr38
		case 92:
			goto st27
		case 123:
			goto tr42
		}
		goto tr37
	tr43:
//line lexer.rl:133
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 82
				goto _out
			}
		}
		goto st82
	tr46:
		cs = 82
//line lexer.rl:141
		te = p + 1
		{
			token = Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, lexContext{})
			}
			stack[top] = lexContext{1, 82}
			top++
			cs = 33
			{
				p++
				goto _out
			}
		}
		goto _again
	tr47:
		cs = 82
//line lexer.rl:149
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(IND_STRING_CLOSE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			cs = 33
			{
				p++
				goto _out
			}
		}
		goto _again
	tr48:
//line lexer.rl:139
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: unescapeStr(data[ts+2 : te])}
			{
				p++
				cs = 82
				goto _out
			}
		}
		goto st82
	tr49:
//line lexer.rl:89
		lineStart = p + 1
		lineCount++
//line lexer.rl:139
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: unescapeStr(data[ts+2 : te])}
			{
				p++
				cs = 82
				goto _out
			}
		}
		goto st82
	tr148:
//line lexer.rl:133
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 82
				goto _out
			}
		}
		goto st82
	tr151:
//line lexer.rl:152
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: []byte("'")}
			{
				p++
				cs = 82
				goto _out
			}
		}
		goto st82
	tr153:
		cs = 82
//line lexer.rl:149
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STRING_CLOSE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			cs = 33
			{
				p++
				goto _out
			}
		}
		goto _again
	tr154:
//line lexer.rl:135
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: []byte("$")}
			{
				p++
				cs = 82
				goto _out
			}
		}
		goto st82
	tr155:
//line lexer.rl:137
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: []byte("'")}
			{
				p++
				cs = 82
				goto _out
			}
		}
		goto st82
	st82:
//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line NONE:1
		ts = p

//line lexer-generated.go.in:3399
		switch data[p] {
		case 10:
			goto tr45
		case 36:
			goto st31
		case 39:
			goto st84
		}
		goto tr44
	tr44:
//line NONE:1
		te = p + 1

		goto st83
	tr45:
//line NONE:1
		te = p + 1

//line lexer.rl:89
		lineStart = p + 1
		lineCount++
		goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line lexer-generated.go.in:3426
		switch data[p] {
		case 10:
			goto tr45
		case 36:
			goto st29
		case 39:
			goto st30
		}
		goto tr44
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 10:
			goto tr45
		case 39:
			goto tr43
		case 123:
			goto tr43
		}
		goto tr44
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 10:
			goto tr45
		case 36:
			goto tr43
		case 39:
			goto tr43
		}
		goto tr44
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 10:
			goto tr45
		case 39:
			goto st0
		case 123:
			goto tr46
		}
		goto tr44
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 10:
			goto tr45
		case 36:
			goto tr151
		case 39:
			goto tr152
		}
		goto tr44
	tr152:
//line NONE:1
		te = p + 1

		goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line lexer-generated.go.in:3502
		switch data[p] {
		case 36:
			goto tr154
		case 39:
			goto tr155
		case 92:
			goto st32
		}
		goto tr153
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 10 {
			goto tr49
		}
		goto tr48
	st_out:
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof79:
		cs = 79
		goto _test_eof
	_test_eof80:
		cs = 80
		goto _test_eof
	_test_eof81:
		cs = 81
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 34:
				goto tr92
			case 3:
				goto tr3
			case 4:
				goto tr3
			case 35:
				goto tr92
			case 36:
				goto tr94
			case 37:
				goto tr95
			case 38:
				goto tr97
			case 39:
				goto tr98
			case 8:
				goto tr9
			case 40:
				goto tr99
			case 41:
				goto tr12
			case 9:
				goto tr12
			case 10:
				goto tr12
			case 42:
				goto tr101
			case 43:
				goto tr102
			case 44:
				goto tr104
			case 11:
				goto tr16
			case 45:
				goto tr92
			case 12:
				goto tr3
			case 13:
				goto tr3
			case 46:
				goto tr92
			case 47:
				goto tr108
			case 48:
				goto tr108
			case 49:
				goto tr109
			case 50:
				goto tr12
			case 51:
				goto tr112
			case 21:
				goto tr12
			case 22:
				goto tr12
			case 52:
				goto tr113
			case 53:
				goto tr112
			case 54:
				goto tr112
			case 55:
				goto tr112
			case 56:
				goto tr112
			case 57:
				goto tr112
			case 58:
				goto tr112
			case 59:
				goto tr112
			case 60:
				goto tr112
			case 61:
				goto tr112
			case 62:
				goto tr112
			case 63:
				goto tr124
			case 64:
				goto tr112
			case 65:
				goto tr112
			case 66:
				goto tr112
			case 67:
				goto tr112
			case 68:
				goto tr112
			case 69:
				goto tr112
			case 70:
				goto tr112
			case 71:
				goto tr112
			case 72:
				goto tr112
			case 73:
				goto tr112
			case 74:
				goto tr112
			case 75:
				goto tr112
			case 76:
				goto tr112
			case 77:
				goto tr112
			case 78:
				goto tr112
			case 25:
				goto tr12
			case 79:
				goto tr141
			case 81:
				goto tr144
			case 26:
				goto tr36
			case 27:
				goto tr40
			case 83:
				goto tr148
			case 29:
				goto tr43
			case 30:
				goto tr43
			case 84:
				goto tr151
			case 85:
				goto tr153
			case 32:
				goto tr47
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:69

	// store state in l
	l.cs = cs
	l.p = p
	l.top = top
	l.lineStart = lineStart
	l.lineCount = lineCount
	l.act = act

	if cs == 0 {
		// TODO: handle error
		err = fmt.Errorf("Unexpected token at %v %v", lineCount+1, (p-lineStart)+1)
	}

	return token, err
}

//line lexer.rl:281
