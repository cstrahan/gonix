//line lexer.rl:1
package main

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

func lex(data []byte) []Token {

//line lexer.rl:41

//line lexer-generated.go.in:46
	const scanner_start int = 60
	const scanner_first_final int = 60
	const scanner_error int = 0

	const scanner_en_string int = 106
	const scanner_en_ind_string int = 108
	const scanner_en_inside_dollar_curly int = 112
	const scanner_en_main int = 60

//line lexer.rl:42

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

	cs, p, pe, eof := 0, 0, len(data), len(data)
	ts, te, act := 0, 0, 0

	top := 0
	stack := make([]int, 4, 4)

	_, _, _, _, _, _ = top, ts, te, act, eof, stack

	tokens := make([]Token, 0, 0)
	_ = tokens

//line lexer.rl:230

//line lexer-generated.go.in:88
	{
		cs = scanner_start
		top = 0
		ts = 0
		te = 0
		act = 0
	}

//line lexer.rl:233

//line lexer-generated.go.in:99
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch cs {
		case 60:
			goto st60
		case 1:
			goto st1
		case 0:
			goto st0
		case 2:
			goto st2
		case 61:
			goto st61
		case 3:
			goto st3
		case 4:
			goto st4
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 66:
			goto st66
		case 8:
			goto st8
		case 67:
			goto st67
		case 68:
			goto st68
		case 9:
			goto st9
		case 10:
			goto st10
		case 69:
			goto st69
		case 70:
			goto st70
		case 71:
			goto st71
		case 11:
			goto st11
		case 72:
			goto st72
		case 12:
			goto st12
		case 13:
			goto st13
		case 73:
			goto st73
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 74:
			goto st74
		case 75:
			goto st75
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 76:
			goto st76
		case 77:
			goto st77
		case 22:
			goto st22
		case 23:
			goto st23
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
		case 104:
			goto st104
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 105:
			goto st105
		case 106:
			goto st106
		case 107:
			goto st107
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 108:
			goto st108
		case 109:
			goto st109
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 110:
			goto st110
		case 111:
			goto st111
		case 33:
			goto st33
		case 112:
			goto st112
		case 34:
			goto st34
		case 35:
			goto st35
		case 113:
			goto st113
		case 36:
			goto st36
		case 37:
			goto st37
		case 114:
			goto st114
		case 115:
			goto st115
		case 116:
			goto st116
		case 117:
			goto st117
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto st40
		case 118:
			goto st118
		case 41:
			goto st41
		case 119:
			goto st119
		case 120:
			goto st120
		case 42:
			goto st42
		case 43:
			goto st43
		case 121:
			goto st121
		case 122:
			goto st122
		case 123:
			goto st123
		case 44:
			goto st44
		case 124:
			goto st124
		case 45:
			goto st45
		case 46:
			goto st46
		case 125:
			goto st125
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 126:
			goto st126
		case 127:
			goto st127
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 128:
			goto st128
		case 129:
			goto st129
		case 55:
			goto st55
		case 56:
			goto st56
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
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 157:
			goto st157
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch cs {
		case 60:
			goto st_case_60
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 61:
			goto st_case_61
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 66:
			goto st_case_66
		case 8:
			goto st_case_8
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 11:
			goto st_case_11
		case 72:
			goto st_case_72
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 73:
			goto st_case_73
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
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
		case 104:
			goto st_case_104
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 33:
			goto st_case_33
		case 112:
			goto st_case_112
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 113:
			goto st_case_113
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 118:
			goto st_case_118
		case 41:
			goto st_case_41
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 44:
			goto st_case_44
		case 124:
			goto st_case_124
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 125:
			goto st_case_125
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
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
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 157:
			goto st_case_157
		}
		goto st_out
	tr3:
//line lexer.rl:198
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr6:
//line lexer.rl:200
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 60
					top++
					goto st112
				}
			}
		}
		goto st60
	tr7:
//line lexer.rl:190
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(AND), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr9:
//line lexer.rl:206
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 60
					top++
					goto st108
				}
			}
		}
		goto st60
	tr10:
//line lexer.rl:206
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 60
					top++
					goto st108
				}
			}
		}
		goto st60
	tr12:
//line NONE:1
		switch act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 54:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(IF), Pos: Pos{ts, te}, Text: nil})
			}
		case 55:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(THEN), Pos: Pos{ts, te}, Text: nil})
			}
		case 56:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ELSE), Pos: Pos{ts, te}, Text: nil})
			}
		case 57:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ASSERT), Pos: Pos{ts, te}, Text: nil})
			}
		case 58:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(WITH), Pos: Pos{ts, te}, Text: nil})
			}
		case 59:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(LET), Pos: Pos{ts, te}, Text: nil})
			}
		case 60:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(IN), Pos: Pos{ts, te}, Text: nil})
			}
		case 61:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(REC), Pos: Pos{ts, te}, Text: nil})
			}
		case 62:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(INHERIT), Pos: Pos{ts, te}, Text: nil})
			}
		case 63:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(OR_KW), Pos: Pos{ts, te}, Text: nil})
			}
		case 64:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ELLIPSIS), Pos: Pos{ts, te}, Text: nil})
			}
		case 68:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te}, Text: nil})
			}
		case 69:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te}, Text: nil})
			}
		case 79:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(CONCAT), Pos: Pos{ts, te}, Text: nil})
			}
		case 80:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ID), Pos: Pos{ts, te}, Text: nil})
			}
		case 81:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(INT), Pos: Pos{ts, te}, Text: nil})
			}
		case 82:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te}, Text: nil})
			}
		case 87:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(PATH), Pos: Pos{ts, te}, Text: nil})
			}
		case 88:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(HPATH), Pos: Pos{ts, te}, Text: nil})
			}
		case 94:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(DOT), Pos: Pos{ts, te}, Text: nil})
			}
		}

		goto st60
	tr16:
//line lexer.rl:219
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr21:
//line lexer.rl:193
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(UPDATE), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr23:
//line lexer.rl:217
		te = p + 1

		goto st60
	tr25:
//line lexer.rl:188
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(LEQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr27:
//line lexer.rl:211
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(SPATH), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr28:
//line lexer.rl:186
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(EQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr29:
//line lexer.rl:189
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(GEQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr33:
//line lexer.rl:191
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(OR), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr85:
//line lexer.rl:221
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 60
					top++
					goto st106
				}
			}
		}
		goto st60
	tr96:
//line lexer.rl:179
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(COLON), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr97:
//line lexer.rl:180
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(SCOLON), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr101:
//line lexer.rl:181
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(AT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr112:
//line lexer.rl:203
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(LCURLY), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr114:
//line lexer.rl:202
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(RCURLY), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr116:
//line lexer.rl:198
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr118:
//line lexer.rl:214
		te = p
		p--

		goto st60
	tr119:
//line lexer.rl:184
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(NOT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr120:
//line lexer.rl:187
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(NEQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr121:
//line lexer.rl:216
		te = p
		p--

		goto st60
	tr122:
//line lexer.rl:206
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 60
					top++
					goto st108
				}
			}
		}
		goto st60
	tr123:
//line lexer.rl:183
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr125:
//line lexer.rl:209
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(PATH), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr126:
//line lexer.rl:182
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr127:
//line lexer.rl:192
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IMPL), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr128:
//line lexer.rl:219
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr132:
//line lexer.rl:197
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(INT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr134:
//line lexer.rl:196
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(ID), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr135:
//line lexer.rl:212
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(URI), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr146:
//line lexer.rl:174
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(IN), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	tr163:
//line lexer.rl:210
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(HPATH), Pos: Pos{ts, te}, Text: nil})
		}
		goto st60
	st60:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line NONE:1
		ts = p

//line lexer-generated.go.in:1045
		switch data[p] {
		case 0:
			goto st1
		case 13:
			goto st63
		case 32:
			goto st63
		case 33:
			goto st64
		case 34:
			goto tr85
		case 35:
			goto st65
		case 36:
			goto st5
		case 38:
			goto st6
		case 39:
			goto st7
		case 43:
			goto tr90
		case 45:
			goto tr91
		case 46:
			goto tr92
		case 47:
			goto st14
		case 48:
			goto tr94
		case 58:
			goto tr96
		case 59:
			goto tr97
		case 60:
			goto st17
		case 61:
			goto st20
		case 62:
			goto st21
		case 64:
			goto tr101
		case 95:
			goto tr103
		case 97:
			goto tr104
		case 101:
			goto tr105
		case 105:
			goto tr106
		case 108:
			goto tr107
		case 111:
			goto tr108
		case 114:
			goto tr109
		case 116:
			goto tr110
		case 119:
			goto tr111
		case 123:
			goto tr112
		case 124:
			goto st24
		case 125:
			goto tr114
		case 126:
			goto st25
		}
		switch {
		case data[p] < 49:
			if 9 <= data[p] && data[p] <= 10 {
				goto st63
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr102
				}
			case data[p] >= 65:
				goto tr102
			}
		default:
			goto tr95
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

		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line lexer-generated.go.in:1164
		switch data[p] {
		case 69:
			goto st3
		case 101:
			goto st3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto tr116
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
			goto st62
		}
		goto tr3
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st62
		}
		goto tr3
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if 48 <= data[p] && data[p] <= 57 {
			goto st62
		}
		goto tr116
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 13:
			goto st63
		case 32:
			goto st63
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st63
		}
		goto tr118
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 61 {
			goto tr120
		}
		goto tr119
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 10:
			goto tr121
		case 13:
			goto tr121
		}
		goto st65
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

		goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lexer-generated.go.in:1281
		switch data[p] {
		case 10:
			goto tr10
		case 32:
			goto st8
		}
		goto tr122
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
	tr90:
//line NONE:1
		te = p + 1

//line lexer.rl:183
		act = 69
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line lexer-generated.go.in:1313
		switch data[p] {
		case 43:
			goto tr124
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
		goto tr123
	tr17:
//line NONE:1
		te = p + 1

//line lexer.rl:178
		act = 64
		goto st68
	tr124:
//line NONE:1
		te = p + 1

//line lexer.rl:194
		act = 79
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line lexer-generated.go.in:1354
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

//line lexer.rl:209
		act = 87
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line lexer-generated.go.in:1443
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
		goto tr125
	tr91:
//line NONE:1
		te = p + 1

//line lexer.rl:182
		act = 68
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line lexer-generated.go.in:1477
		switch data[p] {
		case 43:
			goto st9
		case 47:
			goto st10
		case 62:
			goto tr127
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
		goto tr126
	tr92:
//line NONE:1
		te = p + 1

//line lexer.rl:219
		act = 94
		goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lexer-generated.go.in:1513
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
				goto tr130
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr128
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
	tr130:
//line NONE:1
		te = p + 1

//line lexer.rl:198
		act = 82
		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line lexer-generated.go.in:1579
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
			goto tr130
		}
		goto tr116
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

//line lexer.rl:198
		act = 82
		goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lexer-generated.go.in:1683
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
		goto tr116
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
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 42 {
			goto st16
		}
		goto st15
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 42:
			goto st16
		case 47:
			goto tr23
		}
		goto st15
	tr94:
//line NONE:1
		te = p + 1

//line lexer.rl:197
		act = 81
		goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line lexer-generated.go.in:1771
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
			goto tr94
		}
		goto tr132
	tr95:
//line NONE:1
		te = p + 1

//line lexer.rl:197
		act = 81
		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line lexer-generated.go.in:1810
		switch data[p] {
		case 43:
			goto st9
		case 45:
			goto st9
		case 46:
			goto tr130
		case 47:
			goto st10
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr95
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr132
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 43:
			goto st18
		case 61:
			goto tr25
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
			goto tr27
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
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 61 {
			goto tr28
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 61 {
			goto tr29
		}
		goto st0
	tr102:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st76
	tr140:
//line NONE:1
		te = p + 1

//line lexer.rl:171
		act = 57
		goto st76
	tr143:
//line NONE:1
		te = p + 1

//line lexer.rl:170
		act = 56
		goto st76
	tr144:
//line NONE:1
		te = p + 1

//line lexer.rl:168
		act = 54
		goto st76
	tr151:
//line NONE:1
		te = p + 1

//line lexer.rl:176
		act = 62
		goto st76
	tr153:
//line NONE:1
		te = p + 1

//line lexer.rl:173
		act = 59
		goto st76
	tr154:
//line NONE:1
		te = p + 1

//line lexer.rl:177
		act = 63
		goto st76
	tr156:
//line NONE:1
		te = p + 1

//line lexer.rl:175
		act = 61
		goto st76
	tr159:
//line NONE:1
		te = p + 1

//line lexer.rl:169
		act = 55
		goto st76
	tr162:
//line NONE:1
		te = p + 1

//line lexer.rl:172
		act = 58
		goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line lexer-generated.go.in:2017
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr12
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 39:
			goto st77
		case 45:
			goto st77
		case 95:
			goto st77
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st77
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st77
			}
		default:
			goto st77
		}
		goto tr134
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 43:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		default:
			goto st22
		}
		goto tr12
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 33:
			goto st78
		case 61:
			goto st78
		case 95:
			goto st78
		case 126:
			goto st78
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st78
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st78
				}
			case data[p] >= 63:
				goto st78
			}
		default:
			goto st78
		}
		goto tr12
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 33:
			goto st78
		case 61:
			goto st78
		case 95:
			goto st78
		case 126:
			goto st78
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st78
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st78
				}
			case data[p] >= 63:
				goto st78
			}
		default:
			goto st78
		}
		goto tr135
	tr103:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line lexer-generated.go.in:2177
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st9
		case 46:
			goto st9
		case 47:
			goto st10
		case 95:
			goto tr103
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr103
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr134
	tr104:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line lexer-generated.go.in:2215
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 115:
			goto tr136
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr136:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line lexer-generated.go.in:2257
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 115:
			goto tr137
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr137:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line lexer-generated.go.in:2299
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 101:
			goto tr138
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr138:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line lexer-generated.go.in:2341
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 114:
			goto tr139
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr139:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
//line lexer-generated.go.in:2383
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 116:
			goto tr140
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr105:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line lexer-generated.go.in:2425
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 108:
			goto tr141
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr141:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line lexer-generated.go.in:2467
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 115:
			goto tr142
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr142:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line lexer-generated.go.in:2509
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 101:
			goto tr143
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr106:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line lexer-generated.go.in:2551
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 102:
			goto tr144
		case 110:
			goto tr145
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr145:
//line NONE:1
		te = p + 1

//line lexer.rl:174
		act = 60
		goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line lexer-generated.go.in:2595
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 104:
			goto tr147
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr146
	tr147:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line lexer-generated.go.in:2637
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 101:
			goto tr148
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr148:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line lexer-generated.go.in:2679
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 114:
			goto tr149
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr149:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line lexer-generated.go.in:2721
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 105:
			goto tr150
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr150:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line lexer-generated.go.in:2763
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 116:
			goto tr151
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr107:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line lexer-generated.go.in:2805
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 101:
			goto tr152
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr152:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line lexer-generated.go.in:2847
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 116:
			goto tr153
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr108:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line lexer-generated.go.in:2889
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 114:
			goto tr154
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr109:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line lexer-generated.go.in:2931
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 101:
			goto tr155
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr155:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line lexer-generated.go.in:2973
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 99:
			goto tr156
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr110:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line lexer-generated.go.in:3015
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 104:
			goto tr157
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr157:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line lexer-generated.go.in:3057
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 101:
			goto tr158
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr158:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
//line lexer-generated.go.in:3099
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 110:
			goto tr159
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr111:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line lexer-generated.go.in:3141
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 105:
			goto tr160
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr160:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line lexer-generated.go.in:3183
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 116:
			goto tr161
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	tr161:
//line NONE:1
		te = p + 1

//line lexer.rl:196
		act = 80
		goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line lexer-generated.go.in:3225
		switch data[p] {
		case 39:
			goto st77
		case 43:
			goto st22
		case 46:
			goto st22
		case 47:
			goto st10
		case 58:
			goto st23
		case 95:
			goto tr103
		case 104:
			goto tr162
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr102
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr102
			}
		default:
			goto tr102
		}
		goto tr134
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 124 {
			goto tr33
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 47 {
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
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

//line lexer.rl:210
		act = 88
		goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line lexer-generated.go.in:3314
		switch data[p] {
		case 43:
			goto tr35
		case 47:
			goto st26
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
		goto tr163
	tr36:
//line lexer.rl:86
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(STR), Pos: Pos{ts, te}, Text: data[ts:te]})
		}
		goto st106
	tr38:
//line lexer.rl:81
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(STR), Pos: Pos{ts, te - 1}, Text: data[ts : te-1]})
			tokens = append(tokens, Token{TokenType: TokenType(DQUOTE), Pos: Pos{te - 1, te}, Text: nil})
			{
				top--
				cs = stack[top]
				goto _again
			}
		}
		goto st106
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
				tokens = append(tokens, Token{TokenType: TokenType(STR), Pos: Pos{ts, te}, Text: data[ts:te]})
			}
		}

		goto st106
	tr41:
//line lexer.rl:88
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 106
					top++
					goto st112
				}
			}
		}
		goto st106
	tr164:
//line lexer.rl:90
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te}, Text: nil})
			{
				top--
				cs = stack[top]
				goto _again
			}
		}
		goto st106
	tr166:
//line lexer.rl:86
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(STR), Pos: Pos{ts, te}, Text: data[ts:te]})
		}
		goto st106
	st106:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line NONE:1
		ts = p

//line lexer-generated.go.in:3399
		switch data[p] {
		case 34:
			goto tr164
		case 36:
			goto st29
		case 92:
			goto st28
		}
		goto tr37
	tr37:
//line NONE:1
		te = p + 1

//line lexer.rl:86
		act = 2
		goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line lexer-generated.go.in:3421
		switch data[p] {
		case 34:
			goto tr166
		case 36:
			goto st27
		case 92:
			goto st28
		}
		goto tr37
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 34:
			goto tr38
		case 92:
			goto st28
		case 123:
			goto tr36
		}
		goto tr37
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		goto tr37
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 34:
			goto tr38
		case 92:
			goto st28
		case 123:
			goto tr41
		}
		goto tr37
	tr42:
//line lexer.rl:96
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te}, Text: data[ts:te]})
		}
		goto st108
	tr44:
//line lexer.rl:100
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 108
					top++
					goto st112
				}
			}
		}
		goto st108
	tr45:
//line lexer.rl:102
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_CLOSE), Pos: Pos{ts, te}, Text: nil})
			{
				top--
				cs = stack[top]
				goto _again
			}
		}
		goto st108
	tr46:
//line lexer.rl:99
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te}, Text: unescapeStr(data[ts+2 : te])})
		}
		goto st108
	tr170:
//line lexer.rl:96
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te}, Text: data[ts:te]})
		}
		goto st108
	tr173:
//line lexer.rl:104
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te}, Text: []byte("'")})
		}
		goto st108
	tr175:
//line lexer.rl:102
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_CLOSE), Pos: Pos{ts, te}, Text: nil})
			{
				top--
				cs = stack[top]
				goto _again
			}
		}
		goto st108
	tr176:
//line lexer.rl:97
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te}, Text: []byte("$")})
		}
		goto st108
	tr177:
//line lexer.rl:98
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STR), Pos: Pos{ts, te}, Text: []byte("''")})
		}
		goto st108
	st108:
//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line NONE:1
		ts = p

//line lexer-generated.go.in:3531
		switch data[p] {
		case 36:
			goto st32
		case 39:
			goto st110
		}
		goto tr43
	tr43:
//line NONE:1
		te = p + 1

		goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line lexer-generated.go.in:3549
		switch data[p] {
		case 36:
			goto st30
		case 39:
			goto st31
		}
		goto tr43
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 39:
			goto tr42
		case 123:
			goto tr42
		}
		goto tr43
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 36:
			goto tr42
		case 39:
			goto tr42
		}
		goto tr43
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 39:
			goto st0
		case 123:
			goto tr44
		}
		goto tr43
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 36:
			goto tr173
		case 39:
			goto tr174
		}
		goto tr43
	tr174:
//line NONE:1
		te = p + 1

		goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line lexer-generated.go.in:3615
		switch data[p] {
		case 36:
			goto tr176
		case 39:
			goto tr177
		case 92:
			goto st33
		}
		goto tr175
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		goto tr46
	tr49:
//line lexer.rl:138
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr52:
//line lexer.rl:140
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOLLAR_CURLY), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 112
					top++
					goto st112
				}
			}
		}
		goto st112
	tr53:
//line lexer.rl:130
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(AND), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr55:
//line lexer.rl:148
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 112
					top++
					goto st108
				}
			}
		}
		goto st112
	tr56:
//line lexer.rl:148
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 112
					top++
					goto st108
				}
			}
		}
		goto st112
	tr58:
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
				tokens = append(tokens, Token{TokenType: TokenType(IF), Pos: Pos{ts, te}, Text: nil})
			}
		case 13:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(THEN), Pos: Pos{ts, te}, Text: nil})
			}
		case 14:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ELSE), Pos: Pos{ts, te}, Text: nil})
			}
		case 15:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ASSERT), Pos: Pos{ts, te}, Text: nil})
			}
		case 16:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(WITH), Pos: Pos{ts, te}, Text: nil})
			}
		case 17:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(LET), Pos: Pos{ts, te}, Text: nil})
			}
		case 18:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(IN), Pos: Pos{ts, te}, Text: nil})
			}
		case 19:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(REC), Pos: Pos{ts, te}, Text: nil})
			}
		case 20:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(INHERIT), Pos: Pos{ts, te}, Text: nil})
			}
		case 21:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(OR_KW), Pos: Pos{ts, te}, Text: nil})
			}
		case 22:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ELLIPSIS), Pos: Pos{ts, te}, Text: nil})
			}
		case 26:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te}, Text: nil})
			}
		case 27:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te}, Text: nil})
			}
		case 37:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(CONCAT), Pos: Pos{ts, te}, Text: nil})
			}
		case 38:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(ID), Pos: Pos{ts, te}, Text: nil})
			}
		case 39:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(INT), Pos: Pos{ts, te}, Text: nil})
			}
		case 40:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te}, Text: nil})
			}
		case 45:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(PATH), Pos: Pos{ts, te}, Text: nil})
			}
		case 46:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(HPATH), Pos: Pos{ts, te}, Text: nil})
			}
		case 52:
			{
				p = (te) - 1
				tokens = append(tokens, Token{TokenType: TokenType(DOT), Pos: Pos{ts, te}, Text: nil})
			}
		}

		goto st112
	tr62:
//line lexer.rl:161
		p = (te) - 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr67:
//line lexer.rl:133
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(UPDATE), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr69:
//line lexer.rl:159
		te = p + 1

		goto st112
	tr71:
//line lexer.rl:128
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(LEQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr73:
//line lexer.rl:153
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(SPATH), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr74:
//line lexer.rl:126
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(EQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr75:
//line lexer.rl:129
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(GEQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr79:
//line lexer.rl:131
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(OR), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr182:
//line lexer.rl:163
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(DQUOTE), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 112
					top++
					goto st106
				}
			}
		}
		goto st112
	tr193:
//line lexer.rl:119
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(COLON), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr194:
//line lexer.rl:120
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(SCOLON), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr198:
//line lexer.rl:121
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(AT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr209:
//line lexer.rl:144
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(LCURLY), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 112
					top++
					goto st112
				}
			}
		}
		goto st112
	tr211:
//line lexer.rl:142
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(RCURLY), Pos: Pos{ts, te}, Text: nil})
			{
				top--
				cs = stack[top]
				goto _again
			}
		}
		goto st112
	tr213:
//line lexer.rl:138
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(FLOAT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr215:
//line lexer.rl:156
		te = p
		p--

		goto st112
	tr216:
//line lexer.rl:124
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(NOT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr217:
//line lexer.rl:127
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(NEQ), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr218:
//line lexer.rl:158
		te = p
		p--

		goto st112
	tr219:
//line lexer.rl:148
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(IND_STRING_OPEN), Pos: Pos{ts, te}, Text: nil})
			{
				if len(stack) <= top {
					stack = append(stack, 0)
				}
				{
					stack[top] = 112
					top++
					goto st108
				}
			}
		}
		goto st112
	tr220:
//line lexer.rl:123
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(PLUS), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr222:
//line lexer.rl:151
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(PATH), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr223:
//line lexer.rl:122
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(MINUS), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr224:
//line lexer.rl:132
		te = p + 1
		{
			tokens = append(tokens, Token{TokenType: TokenType(IMPL), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr225:
//line lexer.rl:161
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(DOT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr229:
//line lexer.rl:137
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(INT), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr231:
//line lexer.rl:136
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(ID), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr232:
//line lexer.rl:154
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(URI), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr243:
//line lexer.rl:114
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(IN), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	tr260:
//line lexer.rl:152
		te = p
		p--
		{
			tokens = append(tokens, Token{TokenType: TokenType(HPATH), Pos: Pos{ts, te}, Text: nil})
		}
		goto st112
	st112:
//line NONE:1
		ts = 0

//line NONE:1
		act = 0

		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line NONE:1
		ts = p

//line lexer-generated.go.in:3933
		switch data[p] {
		case 0:
			goto st34
		case 13:
			goto st115
		case 32:
			goto st115
		case 33:
			goto st116
		case 34:
			goto tr182
		case 35:
			goto st117
		case 36:
			goto st38
		case 38:
			goto st39
		case 39:
			goto st40
		case 43:
			goto tr187
		case 45:
			goto tr188
		case 46:
			goto tr189
		case 47:
			goto st47
		case 48:
			goto tr191
		case 58:
			goto tr193
		case 59:
			goto tr194
		case 60:
			goto st50
		case 61:
			goto st53
		case 62:
			goto st54
		case 64:
			goto tr198
		case 95:
			goto tr200
		case 97:
			goto tr201
		case 101:
			goto tr202
		case 105:
			goto tr203
		case 108:
			goto tr204
		case 111:
			goto tr205
		case 114:
			goto tr206
		case 116:
			goto tr207
		case 119:
			goto tr208
		case 123:
			goto tr209
		case 124:
			goto st57
		case 125:
			goto tr211
		case 126:
			goto st58
		}
		switch {
		case data[p] < 49:
			if 9 <= data[p] && data[p] <= 10 {
				goto st115
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 98 <= data[p] && data[p] <= 122 {
					goto tr199
				}
			case data[p] >= 65:
				goto tr199
			}
		default:
			goto tr192
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 46 {
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr48
		}
		goto st0
	tr48:
//line NONE:1
		te = p + 1

		goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line lexer-generated.go.in:4048
		switch data[p] {
		case 69:
			goto st36
		case 101:
			goto st36
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr48
		}
		goto tr213
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 43:
			goto st37
		case 45:
			goto st37
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st114
		}
		goto tr49
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if 48 <= data[p] && data[p] <= 57 {
			goto st114
		}
		goto tr49
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		if 48 <= data[p] && data[p] <= 57 {
			goto st114
		}
		goto tr213
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 13:
			goto st115
		case 32:
			goto st115
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st115
		}
		goto tr215
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		if data[p] == 61 {
			goto tr217
		}
		goto tr216
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 10:
			goto tr218
		case 13:
			goto tr218
		}
		goto st117
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 123 {
			goto tr52
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 38 {
			goto tr53
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 39 {
			goto tr54
		}
		goto st0
	tr54:
//line NONE:1
		te = p + 1

		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line lexer-generated.go.in:4165
		switch data[p] {
		case 10:
			goto tr56
		case 32:
			goto st41
		}
		goto tr219
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 10:
			goto tr56
		case 32:
			goto st41
		}
		goto tr55
	tr187:
//line NONE:1
		te = p + 1

//line lexer.rl:123
		act = 27
		goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line lexer-generated.go.in:4197
		switch data[p] {
		case 43:
			goto tr221
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr220
	tr63:
//line NONE:1
		te = p + 1

//line lexer.rl:118
		act = 22
		goto st120
	tr221:
//line NONE:1
		te = p + 1

//line lexer.rl:134
		act = 37
		goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line lexer-generated.go.in:4238
		switch data[p] {
		case 43:
			goto st42
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr58
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 43:
			goto st42
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr58
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 43:
			goto tr61
		case 95:
			goto tr61
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr61
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr61
				}
			case data[p] >= 65:
				goto tr61
			}
		default:
			goto tr61
		}
		goto tr58
	tr61:
//line NONE:1
		te = p + 1

//line lexer.rl:151
		act = 45
		goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line lexer-generated.go.in:4327
		switch data[p] {
		case 43:
			goto tr61
		case 47:
			goto st43
		case 95:
			goto tr61
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr61
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr61
			}
		default:
			goto tr61
		}
		goto tr222
	tr188:
//line NONE:1
		te = p + 1

//line lexer.rl:122
		act = 26
		goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line lexer-generated.go.in:4361
		switch data[p] {
		case 43:
			goto st42
		case 47:
			goto st43
		case 62:
			goto tr224
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr223
	tr189:
//line NONE:1
		te = p + 1

//line lexer.rl:161
		act = 52
		goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line lexer-generated.go.in:4397
		switch data[p] {
		case 43:
			goto st42
		case 45:
			goto st42
		case 46:
			goto st44
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr227
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr225
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 43:
			goto st42
		case 46:
			goto tr63
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr62
	tr227:
//line NONE:1
		te = p + 1

//line lexer.rl:138
		act = 40
		goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line lexer-generated.go.in:4463
		switch data[p] {
		case 43:
			goto st42
		case 47:
			goto st43
		case 69:
			goto st45
		case 95:
			goto st42
		case 101:
			goto st45
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st42
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st42
				}
			case data[p] >= 65:
				goto st42
			}
		default:
			goto tr227
		}
		goto tr213
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 43:
			goto st46
		case 45:
			goto st46
		case 46:
			goto st42
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr65
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr49
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 43:
			goto st42
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st42
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st42
				}
			case data[p] >= 65:
				goto st42
			}
		default:
			goto tr65
		}
		goto tr49
	tr65:
//line NONE:1
		te = p + 1

//line lexer.rl:138
		act = 40
		goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line lexer-generated.go.in:4567
		switch data[p] {
		case 43:
			goto st42
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st42
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st42
				}
			case data[p] >= 65:
				goto st42
			}
		default:
			goto tr65
		}
		goto tr213
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 42:
			goto st48
		case 43:
			goto tr61
		case 47:
			goto tr67
		case 95:
			goto tr61
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr61
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr61
			}
		default:
			goto tr61
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 42 {
			goto st49
		}
		goto st48
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 42:
			goto st49
		case 47:
			goto tr69
		}
		goto st48
	tr191:
//line NONE:1
		te = p + 1

//line lexer.rl:137
		act = 39
		goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line lexer-generated.go.in:4655
		switch data[p] {
		case 43:
			goto st42
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st42
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st42
				}
			case data[p] >= 65:
				goto st42
			}
		default:
			goto tr191
		}
		goto tr229
	tr192:
//line NONE:1
		te = p + 1

//line lexer.rl:137
		act = 39
		goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line lexer-generated.go.in:4694
		switch data[p] {
		case 43:
			goto st42
		case 45:
			goto st42
		case 46:
			goto tr227
		case 47:
			goto st43
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr192
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr229
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 43:
			goto st51
		case 61:
			goto tr71
		case 95:
			goto st51
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st51
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st51
				}
			case data[p] >= 65:
				goto st51
			}
		default:
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 43:
			goto st51
		case 47:
			goto st52
		case 62:
			goto tr73
		case 95:
			goto st51
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st51
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st51
			}
		default:
			goto st51
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 43:
			goto st51
		case 95:
			goto st51
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st51
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st51
				}
			case data[p] >= 65:
				goto st51
			}
		default:
			goto st51
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 61 {
			goto tr74
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 61 {
			goto tr75
		}
		goto st0
	tr199:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st128
	tr237:
//line NONE:1
		te = p + 1

//line lexer.rl:111
		act = 15
		goto st128
	tr240:
//line NONE:1
		te = p + 1

//line lexer.rl:110
		act = 14
		goto st128
	tr241:
//line NONE:1
		te = p + 1

//line lexer.rl:108
		act = 12
		goto st128
	tr248:
//line NONE:1
		te = p + 1

//line lexer.rl:116
		act = 20
		goto st128
	tr250:
//line NONE:1
		te = p + 1

//line lexer.rl:113
		act = 17
		goto st128
	tr251:
//line NONE:1
		te = p + 1

//line lexer.rl:117
		act = 21
		goto st128
	tr253:
//line NONE:1
		te = p + 1

//line lexer.rl:115
		act = 19
		goto st128
	tr256:
//line NONE:1
		te = p + 1

//line lexer.rl:109
		act = 13
		goto st128
	tr259:
//line NONE:1
		te = p + 1

//line lexer.rl:112
		act = 16
		goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line lexer-generated.go.in:4901
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr58
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
		goto tr231
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 43:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st55
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st55
			}
		default:
			goto st55
		}
		goto tr58
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
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
		goto tr58
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
		goto tr232
	tr200:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line lexer-generated.go.in:5061
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st42
		case 46:
			goto st42
		case 47:
			goto st43
		case 95:
			goto tr200
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr200
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr200
			}
		default:
			goto tr200
		}
		goto tr231
	tr201:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line lexer-generated.go.in:5099
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 115:
			goto tr233
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr233:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line lexer-generated.go.in:5141
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 115:
			goto tr234
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr234:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line lexer-generated.go.in:5183
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 101:
			goto tr235
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr235:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
//line lexer-generated.go.in:5225
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 114:
			goto tr236
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr236:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line lexer-generated.go.in:5267
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 116:
			goto tr237
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr202:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line lexer-generated.go.in:5309
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 108:
			goto tr238
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr238:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line lexer-generated.go.in:5351
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 115:
			goto tr239
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr239:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line lexer-generated.go.in:5393
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 101:
			goto tr240
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr203:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line lexer-generated.go.in:5435
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 102:
			goto tr241
		case 110:
			goto tr242
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr242:
//line NONE:1
		te = p + 1

//line lexer.rl:114
		act = 18
		goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line lexer-generated.go.in:5479
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 104:
			goto tr244
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr243
	tr244:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line lexer-generated.go.in:5521
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 101:
			goto tr245
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr245:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line lexer-generated.go.in:5563
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 114:
			goto tr246
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr246:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line lexer-generated.go.in:5605
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 105:
			goto tr247
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr247:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line lexer-generated.go.in:5647
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 116:
			goto tr248
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr204:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
//line lexer-generated.go.in:5689
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 101:
			goto tr249
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr249:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line lexer-generated.go.in:5731
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 116:
			goto tr250
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr205:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line lexer-generated.go.in:5773
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 114:
			goto tr251
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr206:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line lexer-generated.go.in:5815
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 101:
			goto tr252
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr252:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line lexer-generated.go.in:5857
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 99:
			goto tr253
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr207:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
//line lexer-generated.go.in:5899
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 104:
			goto tr254
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr254:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st152
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
//line lexer-generated.go.in:5941
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 101:
			goto tr255
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr255:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st153
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
//line lexer-generated.go.in:5983
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 110:
			goto tr256
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr208:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st154
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
//line lexer-generated.go.in:6025
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 105:
			goto tr257
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr257:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st155
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
//line lexer-generated.go.in:6067
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 116:
			goto tr258
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	tr258:
//line NONE:1
		te = p + 1

//line lexer.rl:136
		act = 38
		goto st156
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
//line lexer-generated.go.in:6109
		switch data[p] {
		case 39:
			goto st129
		case 43:
			goto st55
		case 46:
			goto st55
		case 47:
			goto st43
		case 58:
			goto st56
		case 95:
			goto tr200
		case 104:
			goto tr259
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr199
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr199
			}
		default:
			goto tr199
		}
		goto tr231
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 124 {
			goto tr79
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 47 {
			goto st59
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 43:
			goto tr81
		case 95:
			goto tr81
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr81
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr81
				}
			case data[p] >= 65:
				goto tr81
			}
		default:
			goto tr81
		}
		goto tr58
	tr81:
//line NONE:1
		te = p + 1

//line lexer.rl:152
		act = 46
		goto st157
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
//line lexer-generated.go.in:6198
		switch data[p] {
		case 43:
			goto tr81
		case 47:
			goto st59
		case 95:
			goto tr81
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr81
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr81
			}
		default:
			goto tr81
		}
		goto tr260
	st_out:
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
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
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
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
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof73:
		cs = 73
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
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
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
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
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
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof106:
		cs = 106
		goto _test_eof
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof108:
		cs = 108
		goto _test_eof
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
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
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
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
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
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
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof157:
		cs = 157
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 61:
				goto tr116
			case 3:
				goto tr3
			case 4:
				goto tr3
			case 62:
				goto tr116
			case 63:
				goto tr118
			case 64:
				goto tr119
			case 65:
				goto tr121
			case 66:
				goto tr122
			case 8:
				goto tr9
			case 67:
				goto tr123
			case 68:
				goto tr12
			case 9:
				goto tr12
			case 10:
				goto tr12
			case 69:
				goto tr125
			case 70:
				goto tr126
			case 71:
				goto tr128
			case 11:
				goto tr16
			case 72:
				goto tr116
			case 12:
				goto tr3
			case 13:
				goto tr3
			case 73:
				goto tr116
			case 74:
				goto tr132
			case 75:
				goto tr132
			case 76:
				goto tr12
			case 77:
				goto tr134
			case 22:
				goto tr12
			case 23:
				goto tr12
			case 78:
				goto tr135
			case 79:
				goto tr134
			case 80:
				goto tr134
			case 81:
				goto tr134
			case 82:
				goto tr134
			case 83:
				goto tr134
			case 84:
				goto tr134
			case 85:
				goto tr134
			case 86:
				goto tr134
			case 87:
				goto tr134
			case 88:
				goto tr134
			case 89:
				goto tr146
			case 90:
				goto tr134
			case 91:
				goto tr134
			case 92:
				goto tr134
			case 93:
				goto tr134
			case 94:
				goto tr134
			case 95:
				goto tr134
			case 96:
				goto tr134
			case 97:
				goto tr134
			case 98:
				goto tr134
			case 99:
				goto tr134
			case 100:
				goto tr134
			case 101:
				goto tr134
			case 102:
				goto tr134
			case 103:
				goto tr134
			case 104:
				goto tr134
			case 26:
				goto tr12
			case 105:
				goto tr163
			case 107:
				goto tr166
			case 27:
				goto tr36
			case 28:
				goto tr40
			case 109:
				goto tr170
			case 30:
				goto tr42
			case 31:
				goto tr42
			case 110:
				goto tr173
			case 111:
				goto tr175
			case 33:
				goto tr45
			case 113:
				goto tr213
			case 36:
				goto tr49
			case 37:
				goto tr49
			case 114:
				goto tr213
			case 115:
				goto tr215
			case 116:
				goto tr216
			case 117:
				goto tr218
			case 118:
				goto tr219
			case 41:
				goto tr55
			case 119:
				goto tr220
			case 120:
				goto tr58
			case 42:
				goto tr58
			case 43:
				goto tr58
			case 121:
				goto tr222
			case 122:
				goto tr223
			case 123:
				goto tr225
			case 44:
				goto tr62
			case 124:
				goto tr213
			case 45:
				goto tr49
			case 46:
				goto tr49
			case 125:
				goto tr213
			case 126:
				goto tr229
			case 127:
				goto tr229
			case 128:
				goto tr58
			case 129:
				goto tr231
			case 55:
				goto tr58
			case 56:
				goto tr58
			case 130:
				goto tr232
			case 131:
				goto tr231
			case 132:
				goto tr231
			case 133:
				goto tr231
			case 134:
				goto tr231
			case 135:
				goto tr231
			case 136:
				goto tr231
			case 137:
				goto tr231
			case 138:
				goto tr231
			case 139:
				goto tr231
			case 140:
				goto tr231
			case 141:
				goto tr243
			case 142:
				goto tr231
			case 143:
				goto tr231
			case 144:
				goto tr231
			case 145:
				goto tr231
			case 146:
				goto tr231
			case 147:
				goto tr231
			case 148:
				goto tr231
			case 149:
				goto tr231
			case 150:
				goto tr231
			case 151:
				goto tr231
			case 152:
				goto tr231
			case 153:
				goto tr231
			case 154:
				goto tr231
			case 155:
				goto tr231
			case 156:
				goto tr231
			case 59:
				goto tr58
			case 157:
				goto tr260
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:234
	return tokens
}
