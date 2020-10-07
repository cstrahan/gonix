//line lexer.rl:1

//line lexer.rl:2

package main

import "fmt"

//line lexer-generated.go.in:12
const scanner_start int = 58
const scanner_first_final int = 58
const scanner_error int = 0

const scanner_en_string int = 105
const scanner_en_ind_string int = 107
const scanner_en_inside_dollar_curly int = 111
const scanner_en_main int = 58

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

	// make go stop complaining about unused variables
	_ = scanner_start
	_ = scanner_first_final
	_ = scanner_error
	_ = scanner_en_main
	//_ = _scanner_nfa_targs
	//_ = _scanner_nfa_offsets
	//_ = _scanner_nfa_push_actions
	//_ = _scanner_nfa_pop_trans
	_ = scanner_en_inside_dollar_curly
	_ = scanner_en_string
	_ = scanner_en_ind_string

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

	_, _, _, _, _, _ = top, ts, te, act, eof, stack

	// Try to reserve enough capacity so we can avoid copying.
	// This factor is based on lexing all-packages.nix, where we reduce our lexing
	// time to one third compared to the naive empty slice.
	token := Token{}
	_ = token

//line lexer-generated.go.in:99
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch cs {
		case 58:
			goto st58
		case 1:
			goto st1
		case 0:
			goto st0
		case 2:
			goto st2
		case 59:
			goto st59
		case 3:
			goto st3
		case 4:
			goto st4
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 64:
			goto st64
		case 8:
			goto st8
		case 65:
			goto st65
		case 66:
			goto st66
		case 9:
			goto st9
		case 10:
			goto st10
		case 67:
			goto st67
		case 68:
			goto st68
		case 69:
			goto st69
		case 11:
			goto st11
		case 70:
			goto st70
		case 12:
			goto st12
		case 13:
			goto st13
		case 71:
			goto st71
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 72:
			goto st72
		case 73:
			goto st73
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 74:
			goto st74
		case 20:
			goto st20
		case 75:
			goto st75
		case 76:
			goto st76
		case 21:
			goto st21
		case 22:
			goto st22
		case 77:
			goto st77
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 82:
			goto st82
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 86:
			goto st86
		case 87:
			goto st87
		case 88:
			goto st88
		case 89:
			goto st89
		case 90:
			goto st90
		case 91:
			goto st91
		case 92:
			goto st92
		case 93:
			goto st93
		case 94:
			goto st94
		case 95:
			goto st95
		case 96:
			goto st96
		case 97:
			goto st97
		case 98:
			goto st98
		case 99:
			goto st99
		case 100:
			goto st100
		case 101:
			goto st101
		case 102:
			goto st102
		case 103:
			goto st103
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 104:
			goto st104
		case 105:
			goto st105
		case 106:
			goto st106
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 107:
			goto st107
		case 108:
			goto st108
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 109:
			goto st109
		case 110:
			goto st110
		case 32:
			goto st32
		case 111:
			goto st111
		case 33:
			goto st33
		case 34:
			goto st34
		case 112:
			goto st112
		case 35:
			goto st35
		case 36:
			goto st36
		case 113:
			goto st113
		case 114:
			goto st114
		case 115:
			goto st115
		case 116:
			goto st116
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 117:
			goto st117
		case 40:
			goto st40
		case 118:
			goto st118
		case 119:
			goto st119
		case 41:
			goto st41
		case 42:
			goto st42
		case 120:
			goto st120
		case 121:
			goto st121
		case 122:
			goto st122
		case 43:
			goto st43
		case 123:
			goto st123
		case 44:
			goto st44
		case 45:
			goto st45
		case 124:
			goto st124
		case 46:
			goto st46
		case 47:
			goto st47
		case 48:
			goto st48
		case 125:
			goto st125
		case 126:
			goto st126
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 127:
			goto st127
		case 52:
			goto st52
		case 128:
			goto st128
		case 129:
			goto st129
		case 53:
			goto st53
		case 54:
			goto st54
		case 130:
			goto st130
		case 131:
			goto st131
		case 132:
			goto st132
		case 133:
			goto st133
		case 134:
			goto st134
		case 135:
			goto st135
		case 136:
			goto st136
		case 137:
			goto st137
		case 138:
			goto st138
		case 139:
			goto st139
		case 140:
			goto st140
		case 141:
			goto st141
		case 142:
			goto st142
		case 143:
			goto st143
		case 144:
			goto st144
		case 145:
			goto st145
		case 146:
			goto st146
		case 147:
			goto st147
		case 148:
			goto st148
		case 149:
			goto st149
		case 150:
			goto st150
		case 151:
			goto st151
		case 152:
			goto st152
		case 153:
			goto st153
		case 154:
			goto st154
		case 155:
			goto st155
		case 156:
			goto st156
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 157:
			goto st157
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch cs {
		case 58:
			goto st_case_58
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 59:
			goto st_case_59
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 64:
			goto st_case_64
		case 8:
			goto st_case_8
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 11:
			goto st_case_11
		case 70:
			goto st_case_70
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 71:
			goto st_case_71
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 74:
			goto st_case_74
		case 20:
			goto st_case_20
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 32:
			goto st_case_32
		case 111:
			goto st_case_111
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 112:
			goto st_case_112
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 117:
			goto st_case_117
		case 40:
			goto st_case_40
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 43:
			goto st_case_43
		case 123:
			goto st_case_123
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 124:
			goto st_case_124
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 127:
			goto st_case_127
		case 52:
			goto st_case_52
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 157:
			goto st_case_157
		}
		goto st_out
	tr3:
//line lexer.rl:374
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr6:
		cs = 58
//line lexer.rl:377
		te = p + 1
		{
			token = Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 58
			top++
			cs = 111
			{
				p++
				goto _out
			}
		}
		goto _again
	tr7:
//line lexer.rl:359
		te = p + 1
		{
			token = Token{TokenType: TokenType(AND), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr9:
		cs = 58
//line lexer.rl:391
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 58
			top++
			cs = 107
			{
				p++
				goto _out
			}
		}
		goto _again
	tr10:
		cs = 58
//line lexer.rl:109
		lineStart = p + 1
		lineCount++
//line lexer.rl:391
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 58
			top++
			cs = 107
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
		case 62:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(IF), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 63:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(THEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 64:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ELSE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 65:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ASSERT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 66:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(WITH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 67:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(LET), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 68:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(IN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 69:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(REC), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 70:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(INHERIT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 71:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(OR_KW), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 72:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ELLIPSIS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 76:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 77:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 88:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(CONCAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 89:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ID), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 90:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(INT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 91:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 96:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 97:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 58
					goto _out
				}
			}
		case 103:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 58
					goto _out
				}
			}
		}

		goto st58
	tr16:
//line lexer.rl:414
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr21:
//line lexer.rl:365
		te = p + 1
		{
			token = Token{TokenType: TokenType(UPDATE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr24:
//line lexer.rl:412
		te = p + 1

		goto st58
	tr26:
//line lexer.rl:355
		te = p + 1
		{
			token = Token{TokenType: TokenType(LEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr28:
//line lexer.rl:404
		te = p + 1
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr29:
//line lexer.rl:357
		te = p + 1
		{
			token = Token{TokenType: TokenType(GEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr33:
//line lexer.rl:361
		te = p + 1
		{
			token = Token{TokenType: TokenType(OR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr89:
		cs = 58
//line lexer.rl:417
		te = p + 1
		{
			token = Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 58
			top++
			cs = 105
			{
				p++
				goto _out
			}
		}
		goto _again
	tr94:
//line lexer.rl:432
		te = p + 1
		{
			token = Token{TokenType: TokenType(LPAREN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr95:
//line lexer.rl:434
		te = p + 1
		{
			token = Token{TokenType: TokenType(RPAREN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr96:
//line lexer.rl:348
		te = p + 1
		{
			token = Token{TokenType: TokenType(STAR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr98:
//line lexer.rl:430
		te = p + 1
		{
			token = Token{TokenType: TokenType(COMMA), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr104:
//line lexer.rl:336
		te = p + 1
		{
			token = Token{TokenType: TokenType(COLON), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr105:
//line lexer.rl:338
		te = p + 1
		{
			token = Token{TokenType: TokenType(SCOLON), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr109:
//line lexer.rl:428
		te = p + 1
		{
			token = Token{TokenType: TokenType(QMARK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr110:
//line lexer.rl:340
		te = p + 1
		{
			token = Token{TokenType: TokenType(AT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr112:
//line lexer.rl:436
		te = p + 1
		{
			token = Token{TokenType: TokenType(LBRACK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr113:
//line lexer.rl:438
		te = p + 1
		{
			token = Token{TokenType: TokenType(RBRACK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr123:
//line lexer.rl:387
		te = p + 1
		{
			token = Token{TokenType: TokenType(LCURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr125:
//line lexer.rl:385
		te = p + 1
		{
			token = Token{TokenType: TokenType(RCURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr127:
//line lexer.rl:374
		te = p
		p--
		{
			token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr129:
//line lexer.rl:409
		te = p
		p--

		goto st58
	tr130:
//line lexer.rl:346
		te = p
		p--
		{
			token = Token{TokenType: TokenType(NOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr131:
//line lexer.rl:353
		te = p + 1
		{
			token = Token{TokenType: TokenType(NEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr132:
//line lexer.rl:411
		te = p
		p--

		goto st58
	tr133:
		cs = 58
//line lexer.rl:391
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 58
			top++
			cs = 107
			{
				p++
				goto _out
			}
		}
		goto _again
	tr134:
//line lexer.rl:344
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr136:
//line lexer.rl:400
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr137:
//line lexer.rl:342
		te = p
		p--
		{
			token = Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr138:
//line lexer.rl:363
		te = p + 1
		{
			token = Token{TokenType: TokenType(IMPL), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr139:
//line lexer.rl:414
		te = p
		p--
		{
			token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr143:
//line lexer.rl:372
		te = p
		p--
		{
			token = Token{TokenType: TokenType(INT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr144:
//line lexer.rl:426
		te = p
		p--
		{
			token = Token{TokenType: TokenType(EQS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr145:
//line lexer.rl:351
		te = p + 1
		{
			token = Token{TokenType: TokenType(EQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr147:
//line lexer.rl:370
		te = p
		p--
		{
			token = Token{TokenType: TokenType(ID), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr148:
//line lexer.rl:406
		te = p
		p--
		{
			token = Token{TokenType: TokenType(URI), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr159:
//line lexer.rl:326
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	tr176:
//line lexer.rl:402
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 58
				goto _out
			}
		}
		goto st58
	st58:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line NONE:1
		ts = p

//line lexer-generated.go.in:1158
		switch data[p] {
		case 0:
			goto st1
		case 9:
			goto st61
		case 10:
			goto tr87
		case 13:
			goto st61
		case 32:
			goto st61
		case 33:
			goto st62
		case 34:
			goto tr89
		case 35:
			goto st63
		case 36:
			goto st5
		case 38:
			goto st6
		case 39:
			goto st7
		case 40:
			goto tr94
		case 41:
			goto tr95
		case 42:
			goto tr96
		case 43:
			goto tr97
		case 44:
			goto tr98
		case 45:
			goto tr99
		case 46:
			goto tr100
		case 47:
			goto st14
		case 48:
			goto tr102
		case 58:
			goto tr104
		case 59:
			goto tr105
		case 60:
			goto st17
		case 61:
			goto st74
		case 62:
			goto st20
		case 63:
			goto tr109
		case 64:
			goto tr110
		case 91:
			goto tr112
		case 93:
			goto tr113
		case 95:
			goto tr114
		case 97:
			goto tr115
		case 101:
			goto tr116
		case 105:
			goto tr117
		case 108:
			goto tr118
		case 111:
			goto tr119
		case 114:
			goto tr120
		case 116:
			goto tr121
		case 119:
			goto tr122
		case 123:
			goto tr123
		case 124:
			goto st23
		case 125:
			goto tr125
		case 126:
			goto st24
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr103
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
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

		goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line lexer-generated.go.in:1290
		switch data[p] {
		case 69:
			goto st3
		case 101:
			goto st3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto tr127
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
			goto st60
		}
		goto tr3
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st60
		}
		goto tr3
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		if 48 <= data[p] && data[p] <= 57 {
			goto st60
		}
		goto tr127
	tr87:
//line lexer.rl:109
		lineStart = p + 1
		lineCount++
		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line lexer-generated.go.in:1343
		switch data[p] {
		case 9:
			goto st61
		case 10:
			goto tr87
		case 13:
			goto st61
		case 32:
			goto st61
		}
		goto tr129
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 61 {
			goto tr131
		}
		goto tr130
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 10:
			goto tr132
		case 13:
			goto tr132
		}
		goto st63
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

		goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line lexer-generated.go.in:1413
		switch data[p] {
		case 10:
			goto tr10
		case 32:
			goto st8
		}
		goto tr133
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
	tr97:
//line NONE:1
		te = p + 1

//line lexer.rl:344
		act = 77
		goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line lexer-generated.go.in:1445
		switch data[p] {
		case 43:
			goto tr135
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
		goto tr134
	tr17:
//line NONE:1
		te = p + 1

//line lexer.rl:334
		act = 72
		goto st66
	tr135:
//line NONE:1
		te = p + 1

//line lexer.rl:367
		act = 88
		goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lexer-generated.go.in:1486
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

//line lexer.rl:400
		act = 96
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line lexer-generated.go.in:1575
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
		goto tr136
	tr99:
//line NONE:1
		te = p + 1

//line lexer.rl:342
		act = 76
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line lexer-generated.go.in:1609
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 62:
			goto tr138
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
		goto tr137
	tr100:
//line NONE:1
		te = p + 1

//line lexer.rl:414
		act = 103
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line lexer-generated.go.in:1645
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
				goto tr141
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr139
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
	tr141:
//line NONE:1
		te = p + 1

//line lexer.rl:374
		act = 91
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line lexer-generated.go.in:1711
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
			goto tr141
		}
		goto tr127
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

//line lexer.rl:374
		act = 91
		goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lexer-generated.go.in:1815
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
		goto tr127
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
//line lexer.rl:109
		lineStart = p + 1
		lineCount++
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lexer-generated.go.in:1879
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
	tr102:
//line NONE:1
		te = p + 1

//line lexer.rl:372
		act = 90
		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line lexer-generated.go.in:1913
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
			goto tr102
		}
		goto tr143
	tr103:
//line NONE:1
		te = p + 1

//line lexer.rl:372
		act = 90
		goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lexer-generated.go.in:1952
		switch data[p] {
		case 43:
			goto st9
		case 45:
			goto st9
		case 46:
			goto tr141
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr103
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr143
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
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 61 {
			goto tr145
		}
		goto tr144
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 61 {
			goto tr29
		}
		goto st0
	tr111:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st75
	tr153:
//line NONE:1
		te = p + 1

//line lexer.rl:320
		act = 65
		goto st75
	tr156:
//line NONE:1
		te = p + 1

//line lexer.rl:318
		act = 64
		goto st75
	tr157:
//line NONE:1
		te = p + 1

//line lexer.rl:314
		act = 62
		goto st75
	tr164:
//line NONE:1
		te = p + 1

//line lexer.rl:330
		act = 70
		goto st75
	tr166:
//line NONE:1
		te = p + 1

//line lexer.rl:324
		act = 67
		goto st75
	tr167:
//line NONE:1
		te = p + 1

//line lexer.rl:332
		act = 71
		goto st75
	tr169:
//line NONE:1
		te = p + 1

//line lexer.rl:328
		act = 69
		goto st75
	tr172:
//line NONE:1
		te = p + 1

//line lexer.rl:316
		act = 63
		goto st75
	tr175:
//line NONE:1
		te = p + 1

//line lexer.rl:322
		act = 66
		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line lexer-generated.go.in:2159
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr12
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 39:
			goto st76
		case 45:
			goto st76
		case 95:
			goto st76
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st76
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st76
			}
		default:
			goto st76
		}
		goto tr147
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
			goto st77
		case 61:
			goto st77
		case 95:
			goto st77
		case 126:
			goto st77
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st77
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st77
				}
			case data[p] >= 63:
				goto st77
			}
		default:
			goto st77
		}
		goto tr12
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 33:
			goto st77
		case 61:
			goto st77
		case 95:
			goto st77
		case 126:
			goto st77
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st77
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st77
				}
			case data[p] >= 63:
				goto st77
			}
		default:
			goto st77
		}
		goto tr148
	tr114:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line lexer-generated.go.in:2319
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st9
		case 46:
			goto st9
		case 47:
			goto st10
		case 95:
			goto tr114
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr114
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr114
			}
		default:
			goto tr114
		}
		goto tr147
	tr115:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line lexer-generated.go.in:2357
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 115:
			goto tr149
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr149:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line lexer-generated.go.in:2399
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 115:
			goto tr150
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr150:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line lexer-generated.go.in:2441
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 101:
			goto tr151
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr151:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line lexer-generated.go.in:2483
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 114:
			goto tr152
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr152:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line lexer-generated.go.in:2525
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 116:
			goto tr153
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr116:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line lexer-generated.go.in:2567
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 108:
			goto tr154
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr154:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line lexer-generated.go.in:2609
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 115:
			goto tr155
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr155:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line lexer-generated.go.in:2651
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 101:
			goto tr156
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr117:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line lexer-generated.go.in:2693
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 102:
			goto tr157
		case 110:
			goto tr158
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr158:
//line NONE:1
		te = p + 1

//line lexer.rl:326
		act = 68
		goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line lexer-generated.go.in:2737
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 104:
			goto tr160
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr159
	tr160:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line lexer-generated.go.in:2779
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 101:
			goto tr161
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr161:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line lexer-generated.go.in:2821
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 114:
			goto tr162
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr162:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line lexer-generated.go.in:2863
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 105:
			goto tr163
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr163:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line lexer-generated.go.in:2905
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 116:
			goto tr164
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr118:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line lexer-generated.go.in:2947
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 101:
			goto tr165
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr165:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line lexer-generated.go.in:2989
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 116:
			goto tr166
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr119:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line lexer-generated.go.in:3031
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 114:
			goto tr167
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr120:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line lexer-generated.go.in:3073
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 101:
			goto tr168
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr168:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line lexer-generated.go.in:3115
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 99:
			goto tr169
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr121:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line lexer-generated.go.in:3157
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 104:
			goto tr170
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr170:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line lexer-generated.go.in:3199
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 101:
			goto tr171
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr171:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line lexer-generated.go.in:3241
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 110:
			goto tr172
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr122:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
//line lexer-generated.go.in:3283
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 105:
			goto tr173
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr173:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line lexer-generated.go.in:3325
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 116:
			goto tr174
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
	tr174:
//line NONE:1
		te = p + 1

//line lexer.rl:370
		act = 89
		goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line lexer-generated.go.in:3367
		switch data[p] {
		case 39:
			goto st76
		case 43:
			goto st21
		case 46:
			goto st21
		case 47:
			goto st10
		case 58:
			goto st22
		case 95:
			goto tr114
		case 104:
			goto tr175
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr147
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

//line lexer.rl:402
		act = 97
		goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line lexer-generated.go.in:3456
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
		goto tr176
	tr36:
//line lexer.rl:130
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]} // ?
			{
				p++
				cs = 105
				goto _out
			}
		}
		goto st105
	tr38:
		cs = 105
//line lexer.rl:123
		te = p + 1
		{
			token = Token{TokenType: TokenType(STR), Pos: Pos{ts, te - 1, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts : te-1]}
			token = Token{TokenType: TokenType(DQUOTE), Pos: Pos{te - 1, te, lineCount + 1, (te - 1 - lineStart) + 1}, Text: nil}
			top--
			cs = (stack[top])
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
					cs = 105
					goto _out
				}
			}
		}

		goto st105
	tr42:
		cs = 105
//line lexer.rl:132
		te = p + 1
		{
			token = Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 105
			top++
			cs = 111
			{
				p++
				goto _out
			}
		}
		goto _again
	tr177:
		cs = 105
//line lexer.rl:141
		te = p + 1
		{
			token = Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			top--
			cs = (stack[top])
			{
				p++
				goto _out
			}

		}
		goto _again
	tr179:
//line lexer.rl:130
		te = p
		p--
		{
			token = Token{TokenType: TokenType(STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]} // ?
			{
				p++
				cs = 105
				goto _out
			}
		}
		goto st105
	st105:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line NONE:1
		ts = p

//line lexer-generated.go.in:3553
		switch data[p] {
		case 34:
			goto tr177
		case 36:
			goto st28
		case 92:
			goto st27
		}
		goto tr37
	tr37:
//line NONE:1
		te = p + 1

//line lexer.rl:130
		act = 2
		goto st106
	tr41:
//line NONE:1
		te = p + 1

//line lexer.rl:109
		lineStart = p + 1
		lineCount++
//line lexer.rl:130
		act = 2
		goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line lexer-generated.go.in:3584
		switch data[p] {
		case 34:
			goto tr179
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
//line lexer.rl:151
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 107
				goto _out
			}
		}
		goto st107
	tr46:
		cs = 107
//line lexer.rl:159
		te = p + 1
		{
			token = Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 107
			top++
			cs = 111
			{
				p++
				goto _out
			}
		}
		goto _again
	tr47:
		cs = 107
//line lexer.rl:167
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(IND_STRING_CLOSE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			top--
			cs = (stack[top])
			{
				p++
				goto _out
			}
		}
		goto _again
	tr48:
//line lexer.rl:157
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: unescapeStr(data[ts+2 : te])}
			{
				p++
				cs = 107
				goto _out
			}
		}
		goto st107
	tr49:
//line lexer.rl:109
		lineStart = p + 1
		lineCount++
//line lexer.rl:157
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: unescapeStr(data[ts+2 : te])}
			{
				p++
				cs = 107
				goto _out
			}
		}
		goto st107
	tr183:
//line lexer.rl:151
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 107
				goto _out
			}
		}
		goto st107
	tr186:
//line lexer.rl:171
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: []byte("'")}
			{
				p++
				cs = 107
				goto _out
			}
		}
		goto st107
	tr188:
		cs = 107
//line lexer.rl:167
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STRING_CLOSE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			top--
			cs = (stack[top])
			{
				p++
				goto _out
			}
		}
		goto _again
	tr189:
//line lexer.rl:153
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: []byte("$")}
			{
				p++
				cs = 107
				goto _out
			}
		}
		goto st107
	tr190:
//line lexer.rl:155
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: []byte("'")}
			{
				p++
				cs = 107
				goto _out
			}
		}
		goto st107
	st107:
//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line NONE:1
		ts = p

//line lexer-generated.go.in:3720
		switch data[p] {
		case 10:
			goto tr45
		case 36:
			goto st31
		case 39:
			goto st109
		}
		goto tr44
	tr44:
//line NONE:1
		te = p + 1

		goto st108
	tr45:
//line NONE:1
		te = p + 1

//line lexer.rl:109
		lineStart = p + 1
		lineCount++
		goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line lexer-generated.go.in:3747
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
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 10:
			goto tr45
		case 36:
			goto tr186
		case 39:
			goto tr187
		}
		goto tr44
	tr187:
//line NONE:1
		te = p + 1

		goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line lexer-generated.go.in:3823
		switch data[p] {
		case 36:
			goto tr189
		case 39:
			goto tr190
		case 92:
			goto st32
		}
		goto tr188
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 10 {
			goto tr49
		}
		goto tr48
	tr52:
//line lexer.rl:237
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr55:
		cs = 111
//line lexer.rl:240
		te = p + 1
		{
			token = Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 111
			top++
			cs = 111
			{
				p++
				goto _out
			}
		}
		goto _again
	tr56:
//line lexer.rl:221
		te = p + 1
		{
			token = Token{TokenType: TokenType(AND), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr58:
		cs = 111
//line lexer.rl:262
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 111
			top++
			cs = 107
			{
				p++
				goto _out
			}
		}
		goto _again
	tr59:
		cs = 111
//line lexer.rl:109
		lineStart = p + 1
		lineCount++
//line lexer.rl:262
		te = p + 1
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 111
			top++
			cs = 107
			{
				p++
				goto _out
			}
		}
		goto _again
	tr61:
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
					cs = 111
					goto _out
				}
			}
		case 13:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(THEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 14:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ELSE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 15:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ASSERT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 16:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(WITH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 17:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(LET), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 18:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(IN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 19:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(REC), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 20:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(INHERIT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 21:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(OR_KW), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 22:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ELLIPSIS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 26:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 27:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 38:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(CONCAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 39:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(ID), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 40:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(INT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 41:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 46:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 47:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
				{
					p++
					cs = 111
					goto _out
				}
			}
		case 53:
			{
				p = (te) - 1
				token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
				{
					p++
					cs = 111
					goto _out
				}
			}
		}

		goto st111
	tr65:
//line lexer.rl:285
		p = (te) - 1
		{
			token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr70:
//line lexer.rl:227
		te = p + 1
		{
			token = Token{TokenType: TokenType(UPDATE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr73:
//line lexer.rl:283
		te = p + 1

		goto st111
	tr75:
//line lexer.rl:217
		te = p + 1
		{
			token = Token{TokenType: TokenType(LEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr77:
//line lexer.rl:275
		te = p + 1
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr78:
//line lexer.rl:219
		te = p + 1
		{
			token = Token{TokenType: TokenType(GEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr82:
//line lexer.rl:223
		te = p + 1
		{
			token = Token{TokenType: TokenType(OR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr196:
		cs = 111
//line lexer.rl:288
		te = p + 1
		{
			token = Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 111
			top++
			cs = 105
			{
				p++
				goto _out
			}
		}
		goto _again
	tr201:
//line lexer.rl:303
		te = p + 1
		{
			token = Token{TokenType: TokenType(LPAREN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr202:
//line lexer.rl:305
		te = p + 1
		{
			token = Token{TokenType: TokenType(RPAREN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr203:
//line lexer.rl:210
		te = p + 1
		{
			token = Token{TokenType: TokenType(STAR), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr205:
//line lexer.rl:301
		te = p + 1
		{
			token = Token{TokenType: TokenType(COMMA), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr211:
//line lexer.rl:198
		te = p + 1
		{
			token = Token{TokenType: TokenType(COLON), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr212:
//line lexer.rl:200
		te = p + 1
		{
			token = Token{TokenType: TokenType(SCOLON), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr216:
//line lexer.rl:299
		te = p + 1
		{
			token = Token{TokenType: TokenType(QMARK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr217:
//line lexer.rl:202
		te = p + 1
		{
			token = Token{TokenType: TokenType(AT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr219:
//line lexer.rl:307
		te = p + 1
		{
			token = Token{TokenType: TokenType(LBRACK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr220:
//line lexer.rl:309
		te = p + 1
		{
			token = Token{TokenType: TokenType(RBRACK), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr230:
		cs = 111
//line lexer.rl:252
		te = p + 1
		{
			token = Token{TokenType: TokenType(LCURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 111
			top++
			cs = 111
			{
				p++
				goto _out
			}
		}
		goto _again
	tr232:
		cs = 111
//line lexer.rl:248
		te = p + 1
		{
			token = Token{TokenType: TokenType(RCURLY), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			top--
			cs = (stack[top])
			{
				p++
				goto _out
			}
		}
		goto _again
	tr234:
//line lexer.rl:237
		te = p
		p--
		{
			token = Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr236:
//line lexer.rl:280
		te = p
		p--

		goto st111
	tr237:
//line lexer.rl:208
		te = p
		p--
		{
			token = Token{TokenType: TokenType(NOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr238:
//line lexer.rl:215
		te = p + 1
		{
			token = Token{TokenType: TokenType(NEQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr239:
//line lexer.rl:282
		te = p
		p--

		goto st111
	tr240:
		cs = 111
//line lexer.rl:262
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			if len(stack) <= top {
				stack = append(stack, 0)
			}
			stack[top] = 111
			top++
			cs = 107
			{
				p++
				goto _out
			}
		}
		goto _again
	tr241:
//line lexer.rl:206
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr243:
//line lexer.rl:271
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr244:
//line lexer.rl:204
		te = p
		p--
		{
			token = Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr245:
//line lexer.rl:225
		te = p + 1
		{
			token = Token{TokenType: TokenType(IMPL), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr246:
//line lexer.rl:285
		te = p
		p--
		{
			token = Token{TokenType: TokenType(DOT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr250:
//line lexer.rl:235
		te = p
		p--
		{
			token = Token{TokenType: TokenType(INT), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr251:
//line lexer.rl:297
		te = p
		p--
		{
			token = Token{TokenType: TokenType(EQS), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr252:
//line lexer.rl:213
		te = p + 1
		{
			token = Token{TokenType: TokenType(EQ), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr254:
//line lexer.rl:233
		te = p
		p--
		{
			token = Token{TokenType: TokenType(ID), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr255:
//line lexer.rl:277
		te = p
		p--
		{
			token = Token{TokenType: TokenType(URI), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr266:
//line lexer.rl:188
		te = p
		p--
		{
			token = Token{TokenType: TokenType(IN), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: nil}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	tr283:
//line lexer.rl:273
		te = p
		p--
		{
			token = Token{TokenType: TokenType(PATH), Pos: Pos{ts, te, lineCount + 1, (ts - lineStart) + 1}, Text: data[ts:te]}
			{
				p++
				cs = 111
				goto _out
			}
		}
		goto st111
	st111:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line NONE:1
		ts = p

//line lexer-generated.go.in:4261
		switch data[p] {
		case 0:
			goto st33
		case 9:
			goto st114
		case 10:
			goto tr194
		case 13:
			goto st114
		case 32:
			goto st114
		case 33:
			goto st115
		case 34:
			goto tr196
		case 35:
			goto st116
		case 36:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto tr201
		case 41:
			goto tr202
		case 42:
			goto tr203
		case 43:
			goto tr204
		case 44:
			goto tr205
		case 45:
			goto tr206
		case 46:
			goto tr207
		case 47:
			goto st46
		case 48:
			goto tr209
		case 58:
			goto tr211
		case 59:
			goto tr212
		case 60:
			goto st49
		case 61:
			goto st127
		case 62:
			goto st52
		case 63:
			goto tr216
		case 64:
			goto tr217
		case 91:
			goto tr219
		case 93:
			goto tr220
		case 95:
			goto tr221
		case 97:
			goto tr222
		case 101:
			goto tr223
		case 105:
			goto tr224
		case 108:
			goto tr225
		case 111:
			goto tr226
		case 114:
			goto tr227
		case 116:
			goto tr228
		case 119:
			goto tr229
		case 123:
			goto tr230
		case 124:
			goto st55
		case 125:
			goto tr232
		case 126:
			goto st56
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr210
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 46 {
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr51
		}
		goto st0
	tr51:
//line NONE:1
		te = p + 1

		goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line lexer-generated.go.in:4389
		switch data[p] {
		case 69:
			goto st35
		case 101:
			goto st35
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr51
		}
		goto tr234
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 43:
			goto st36
		case 45:
			goto st36
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st113
		}
		goto tr52
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if 48 <= data[p] && data[p] <= 57 {
			goto st113
		}
		goto tr52
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		if 48 <= data[p] && data[p] <= 57 {
			goto st113
		}
		goto tr234
	tr194:
//line lexer.rl:109
		lineStart = p + 1
		lineCount++
		goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line lexer-generated.go.in:4442
		switch data[p] {
		case 9:
			goto st114
		case 10:
			goto tr194
		case 13:
			goto st114
		case 32:
			goto st114
		}
		goto tr236
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		if data[p] == 61 {
			goto tr238
		}
		goto tr237
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 10:
			goto tr239
		case 13:
			goto tr239
		}
		goto st116
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 123 {
			goto tr55
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 38 {
			goto tr56
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 39 {
			goto tr57
		}
		goto st0
	tr57:
//line NONE:1
		te = p + 1

		goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
//line lexer-generated.go.in:4512
		switch data[p] {
		case 10:
			goto tr59
		case 32:
			goto st40
		}
		goto tr240
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 10:
			goto tr59
		case 32:
			goto st40
		}
		goto tr58
	tr204:
//line NONE:1
		te = p + 1

//line lexer.rl:206
		act = 27
		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line lexer-generated.go.in:4544
		switch data[p] {
		case 43:
			goto tr242
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr241
	tr66:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 22
		goto st119
	tr242:
//line NONE:1
		te = p + 1

//line lexer.rl:229
		act = 38
		goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line lexer-generated.go.in:4585
		switch data[p] {
		case 43:
			goto st41
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr61
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 43:
			goto st41
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr61
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 43:
			goto tr64
		case 95:
			goto tr64
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr64
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr64
				}
			case data[p] >= 65:
				goto tr64
			}
		default:
			goto tr64
		}
		goto tr61
	tr64:
//line NONE:1
		te = p + 1

//line lexer.rl:271
		act = 46
		goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line lexer-generated.go.in:4674
		switch data[p] {
		case 43:
			goto tr64
		case 47:
			goto st42
		case 95:
			goto tr64
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr64
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr64
			}
		default:
			goto tr64
		}
		goto tr243
	tr206:
//line NONE:1
		te = p + 1

//line lexer.rl:204
		act = 26
		goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line lexer-generated.go.in:4708
		switch data[p] {
		case 43:
			goto st41
		case 47:
			goto st42
		case 62:
			goto tr245
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr244
	tr207:
//line NONE:1
		te = p + 1

//line lexer.rl:285
		act = 53
		goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line lexer-generated.go.in:4744
		switch data[p] {
		case 43:
			goto st41
		case 45:
			goto st41
		case 46:
			goto st43
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr248
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr246
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 43:
			goto st41
		case 46:
			goto tr66
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr65
	tr248:
//line NONE:1
		te = p + 1

//line lexer.rl:237
		act = 41
		goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line lexer-generated.go.in:4810
		switch data[p] {
		case 43:
			goto st41
		case 47:
			goto st42
		case 69:
			goto st44
		case 95:
			goto st41
		case 101:
			goto st44
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st41
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st41
				}
			case data[p] >= 65:
				goto st41
			}
		default:
			goto tr248
		}
		goto tr234
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 43:
			goto st45
		case 45:
			goto st45
		case 46:
			goto st41
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr68
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr52
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 43:
			goto st41
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st41
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st41
				}
			case data[p] >= 65:
				goto st41
			}
		default:
			goto tr68
		}
		goto tr52
	tr68:
//line NONE:1
		te = p + 1

//line lexer.rl:237
		act = 41
		goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line lexer-generated.go.in:4914
		switch data[p] {
		case 43:
			goto st41
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st41
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st41
				}
			case data[p] >= 65:
				goto st41
			}
		default:
			goto tr68
		}
		goto tr234
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 42:
			goto st47
		case 43:
			goto tr64
		case 47:
			goto tr70
		case 95:
			goto tr64
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr64
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr64
			}
		default:
			goto tr64
		}
		goto st0
	tr71:
//line lexer.rl:109
		lineStart = p + 1
		lineCount++
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lexer-generated.go.in:4978
		switch data[p] {
		case 10:
			goto tr71
		case 42:
			goto st48
		}
		goto st47
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 10:
			goto tr71
		case 42:
			goto st48
		case 47:
			goto tr73
		}
		goto st47
	tr209:
//line NONE:1
		te = p + 1

//line lexer.rl:235
		act = 40
		goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line lexer-generated.go.in:5012
		switch data[p] {
		case 43:
			goto st41
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st41
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st41
				}
			case data[p] >= 65:
				goto st41
			}
		default:
			goto tr209
		}
		goto tr250
	tr210:
//line NONE:1
		te = p + 1

//line lexer.rl:235
		act = 40
		goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line lexer-generated.go.in:5051
		switch data[p] {
		case 43:
			goto st41
		case 45:
			goto st41
		case 46:
			goto tr248
		case 47:
			goto st42
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr210
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto tr250
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 43:
			goto st50
		case 61:
			goto tr75
		case 95:
			goto st50
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st50
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st50
				}
			case data[p] >= 65:
				goto st50
			}
		default:
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 43:
			goto st50
		case 47:
			goto st51
		case 62:
			goto tr77
		case 95:
			goto st50
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st50
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st50
			}
		default:
			goto st50
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 43:
			goto st50
		case 95:
			goto st50
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st50
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st50
				}
			case data[p] >= 65:
				goto st50
			}
		default:
			goto st50
		}
		goto st0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		if data[p] == 61 {
			goto tr252
		}
		goto tr251
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 61 {
			goto tr78
		}
		goto st0
	tr218:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st128
	tr260:
//line NONE:1
		te = p + 1

//line lexer.rl:182
		act = 15
		goto st128
	tr263:
//line NONE:1
		te = p + 1

//line lexer.rl:180
		act = 14
		goto st128
	tr264:
//line NONE:1
		te = p + 1

//line lexer.rl:176
		act = 12
		goto st128
	tr271:
//line NONE:1
		te = p + 1

//line lexer.rl:192
		act = 20
		goto st128
	tr273:
//line NONE:1
		te = p + 1

//line lexer.rl:186
		act = 17
		goto st128
	tr274:
//line NONE:1
		te = p + 1

//line lexer.rl:194
		act = 21
		goto st128
	tr276:
//line NONE:1
		te = p + 1

//line lexer.rl:190
		act = 19
		goto st128
	tr279:
//line NONE:1
		te = p + 1

//line lexer.rl:178
		act = 13
		goto st128
	tr282:
//line NONE:1
		te = p + 1

//line lexer.rl:184
		act = 16
		goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line lexer-generated.go.in:5258
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr61
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 39:
			goto st129
		case 45:
			goto st129
		case 95:
			goto st129
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st129
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st129
			}
		default:
			goto st129
		}
		goto tr254
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 43:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st53
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st53
			}
		default:
			goto st53
		}
		goto tr61
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 33:
			goto st130
		case 61:
			goto st130
		case 95:
			goto st130
		case 126:
			goto st130
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st130
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st130
				}
			case data[p] >= 63:
				goto st130
			}
		default:
			goto st130
		}
		goto tr61
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 33:
			goto st130
		case 61:
			goto st130
		case 95:
			goto st130
		case 126:
			goto st130
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st130
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st130
				}
			case data[p] >= 63:
				goto st130
			}
		default:
			goto st130
		}
		goto tr255
	tr221:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line lexer-generated.go.in:5418
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st41
		case 46:
			goto st41
		case 47:
			goto st42
		case 95:
			goto tr221
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr221
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr221
			}
		default:
			goto tr221
		}
		goto tr254
	tr222:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line lexer-generated.go.in:5456
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 115:
			goto tr256
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr256:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line lexer-generated.go.in:5498
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 115:
			goto tr257
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr257:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line lexer-generated.go.in:5540
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 101:
			goto tr258
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr258:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
//line lexer-generated.go.in:5582
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 114:
			goto tr259
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr259:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line lexer-generated.go.in:5624
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 116:
			goto tr260
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr223:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line lexer-generated.go.in:5666
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 108:
			goto tr261
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr261:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line lexer-generated.go.in:5708
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 115:
			goto tr262
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr262:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line lexer-generated.go.in:5750
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 101:
			goto tr263
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr224:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line lexer-generated.go.in:5792
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 102:
			goto tr264
		case 110:
			goto tr265
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr265:
//line NONE:1
		te = p + 1

//line lexer.rl:188
		act = 18
		goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line lexer-generated.go.in:5836
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 104:
			goto tr267
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr266
	tr267:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line lexer-generated.go.in:5878
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 101:
			goto tr268
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr268:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line lexer-generated.go.in:5920
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 114:
			goto tr269
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr269:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line lexer-generated.go.in:5962
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 105:
			goto tr270
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr270:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line lexer-generated.go.in:6004
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 116:
			goto tr271
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr225:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
//line lexer-generated.go.in:6046
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 101:
			goto tr272
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr272:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line lexer-generated.go.in:6088
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 116:
			goto tr273
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr226:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line lexer-generated.go.in:6130
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 114:
			goto tr274
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr227:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line lexer-generated.go.in:6172
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 101:
			goto tr275
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr275:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line lexer-generated.go.in:6214
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 99:
			goto tr276
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr228:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
//line lexer-generated.go.in:6256
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 104:
			goto tr277
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr277:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st152
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
//line lexer-generated.go.in:6298
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 101:
			goto tr278
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr278:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st153
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
//line lexer-generated.go.in:6340
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 110:
			goto tr279
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr229:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st154
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
//line lexer-generated.go.in:6382
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 105:
			goto tr280
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr280:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st155
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
//line lexer-generated.go.in:6424
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 116:
			goto tr281
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	tr281:
//line NONE:1
		te = p + 1

//line lexer.rl:233
		act = 39
		goto st156
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
//line lexer-generated.go.in:6466
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st53
		case 46:
			goto st53
		case 47:
			goto st42
		case 58:
			goto st54
		case 95:
			goto tr221
		case 104:
			goto tr282
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr218
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr218
			}
		default:
			goto tr218
		}
		goto tr254
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 124 {
			goto tr82
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 47 {
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 43:
			goto tr84
		case 95:
			goto tr84
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr84
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr84
				}
			case data[p] >= 65:
				goto tr84
			}
		default:
			goto tr84
		}
		goto tr61
	tr84:
//line NONE:1
		te = p + 1

//line lexer.rl:273
		act = 47
		goto st157
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
//line lexer-generated.go.in:6555
		switch data[p] {
		case 43:
			goto tr84
		case 47:
			goto st57
		case 95:
			goto tr84
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr84
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr84
			}
		default:
			goto tr84
		}
		goto tr283
	st_out:
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
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
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
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
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof71:
		cs = 71
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
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
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
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
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
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof86:
		cs = 86
		goto _test_eof
	_test_eof87:
		cs = 87
		goto _test_eof
	_test_eof88:
		cs = 88
		goto _test_eof
	_test_eof89:
		cs = 89
		goto _test_eof
	_test_eof90:
		cs = 90
		goto _test_eof
	_test_eof91:
		cs = 91
		goto _test_eof
	_test_eof92:
		cs = 92
		goto _test_eof
	_test_eof93:
		cs = 93
		goto _test_eof
	_test_eof94:
		cs = 94
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof101:
		cs = 101
		goto _test_eof
	_test_eof102:
		cs = 102
		goto _test_eof
	_test_eof103:
		cs = 103
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
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof106:
		cs = 106
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
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof108:
		cs = 108
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
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof114:
		cs = 114
		goto _test_eof
	_test_eof115:
		cs = 115
		goto _test_eof
	_test_eof116:
		cs = 116
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof130:
		cs = 130
		goto _test_eof
	_test_eof131:
		cs = 131
		goto _test_eof
	_test_eof132:
		cs = 132
		goto _test_eof
	_test_eof133:
		cs = 133
		goto _test_eof
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof135:
		cs = 135
		goto _test_eof
	_test_eof136:
		cs = 136
		goto _test_eof
	_test_eof137:
		cs = 137
		goto _test_eof
	_test_eof138:
		cs = 138
		goto _test_eof
	_test_eof139:
		cs = 139
		goto _test_eof
	_test_eof140:
		cs = 140
		goto _test_eof
	_test_eof141:
		cs = 141
		goto _test_eof
	_test_eof142:
		cs = 142
		goto _test_eof
	_test_eof143:
		cs = 143
		goto _test_eof
	_test_eof144:
		cs = 144
		goto _test_eof
	_test_eof145:
		cs = 145
		goto _test_eof
	_test_eof146:
		cs = 146
		goto _test_eof
	_test_eof147:
		cs = 147
		goto _test_eof
	_test_eof148:
		cs = 148
		goto _test_eof
	_test_eof149:
		cs = 149
		goto _test_eof
	_test_eof150:
		cs = 150
		goto _test_eof
	_test_eof151:
		cs = 151
		goto _test_eof
	_test_eof152:
		cs = 152
		goto _test_eof
	_test_eof153:
		cs = 153
		goto _test_eof
	_test_eof154:
		cs = 154
		goto _test_eof
	_test_eof155:
		cs = 155
		goto _test_eof
	_test_eof156:
		cs = 156
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
	_test_eof157:
		cs = 157
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 59:
				goto tr127
			case 3:
				goto tr3
			case 4:
				goto tr3
			case 60:
				goto tr127
			case 61:
				goto tr129
			case 62:
				goto tr130
			case 63:
				goto tr132
			case 64:
				goto tr133
			case 8:
				goto tr9
			case 65:
				goto tr134
			case 66:
				goto tr12
			case 9:
				goto tr12
			case 10:
				goto tr12
			case 67:
				goto tr136
			case 68:
				goto tr137
			case 69:
				goto tr139
			case 11:
				goto tr16
			case 70:
				goto tr127
			case 12:
				goto tr3
			case 13:
				goto tr3
			case 71:
				goto tr127
			case 72:
				goto tr143
			case 73:
				goto tr143
			case 74:
				goto tr144
			case 75:
				goto tr12
			case 76:
				goto tr147
			case 21:
				goto tr12
			case 22:
				goto tr12
			case 77:
				goto tr148
			case 78:
				goto tr147
			case 79:
				goto tr147
			case 80:
				goto tr147
			case 81:
				goto tr147
			case 82:
				goto tr147
			case 83:
				goto tr147
			case 84:
				goto tr147
			case 85:
				goto tr147
			case 86:
				goto tr147
			case 87:
				goto tr147
			case 88:
				goto tr159
			case 89:
				goto tr147
			case 90:
				goto tr147
			case 91:
				goto tr147
			case 92:
				goto tr147
			case 93:
				goto tr147
			case 94:
				goto tr147
			case 95:
				goto tr147
			case 96:
				goto tr147
			case 97:
				goto tr147
			case 98:
				goto tr147
			case 99:
				goto tr147
			case 100:
				goto tr147
			case 101:
				goto tr147
			case 102:
				goto tr147
			case 103:
				goto tr147
			case 25:
				goto tr12
			case 104:
				goto tr176
			case 106:
				goto tr179
			case 26:
				goto tr36
			case 27:
				goto tr40
			case 108:
				goto tr183
			case 29:
				goto tr43
			case 30:
				goto tr43
			case 109:
				goto tr186
			case 110:
				goto tr188
			case 32:
				goto tr47
			case 112:
				goto tr234
			case 35:
				goto tr52
			case 36:
				goto tr52
			case 113:
				goto tr234
			case 114:
				goto tr236
			case 115:
				goto tr237
			case 116:
				goto tr239
			case 117:
				goto tr240
			case 40:
				goto tr58
			case 118:
				goto tr241
			case 119:
				goto tr61
			case 41:
				goto tr61
			case 42:
				goto tr61
			case 120:
				goto tr243
			case 121:
				goto tr244
			case 122:
				goto tr246
			case 43:
				goto tr65
			case 123:
				goto tr234
			case 44:
				goto tr52
			case 45:
				goto tr52
			case 124:
				goto tr234
			case 125:
				goto tr250
			case 126:
				goto tr250
			case 127:
				goto tr251
			case 128:
				goto tr61
			case 129:
				goto tr254
			case 53:
				goto tr61
			case 54:
				goto tr61
			case 130:
				goto tr255
			case 131:
				goto tr254
			case 132:
				goto tr254
			case 133:
				goto tr254
			case 134:
				goto tr254
			case 135:
				goto tr254
			case 136:
				goto tr254
			case 137:
				goto tr254
			case 138:
				goto tr254
			case 139:
				goto tr254
			case 140:
				goto tr254
			case 141:
				goto tr266
			case 142:
				goto tr254
			case 143:
				goto tr254
			case 144:
				goto tr254
			case 145:
				goto tr254
			case 146:
				goto tr254
			case 147:
				goto tr254
			case 148:
				goto tr254
			case 149:
				goto tr254
			case 150:
				goto tr254
			case 151:
				goto tr254
			case 152:
				goto tr254
			case 153:
				goto tr254
			case 154:
				goto tr254
			case 155:
				goto tr254
			case 156:
				goto tr254
			case 57:
				goto tr61
			case 157:
				goto tr283
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:83

	// store state in l
	l.cs = cs
	l.p = p
	l.top = top
	l.lineStart = lineStart
	l.lineCount = lineCount
	l.act = act

	var err error

	if cs == 0 {
		// TODO: handle error
		err = fmt.Errorf("Unexpected token at %v %v", lineCount+1, (p-lineStart)+1)
	}

	//if cs < scanner_first_final {
	//  err = fmt.Errorf("Unexpected token at %v %v", lineCount+1, (p-lineStart)+1)
	//}

	return token, err
}

//line lexer.rl:447
