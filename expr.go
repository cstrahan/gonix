package main

import (
	"errors"
)

type Symbol string

type Expr interface {
	Eval(state *EvalState, env *Env, val *Value) error
	BindVars(staticEnv *StaticEnv) error
	SetName(name Symbol)
	MaybeThunk(state *EvalState, env *Env) (*Value, error)
}

func defaultMaybeThunk(expr Expr, state *EvalState, env *Env) (*Value, error) {
	var thunk Value = &Thunk{Expr: expr, Env: env}
	return &thunk, nil
}

type Formals struct {
	Formals  []Formal
	ArgNames map[string]struct{}
	Ellipsis bool
}

type Formal struct {
	Name Symbol
	Def  Expr
}

//------------------------------------------------------------------------------

type ExprConcatStrings struct {
	ForceString bool
	Exprs       []Expr
}

func (self *ExprConcatStrings) Eval(state *EvalState, env *Env, val *Value) error {
	var context [][]byte
	var n NixInt = 0
	var nf NixFloat = 0
	var first Value = nil
	var s string = ""

	for _, expr := range self.Exprs {
		var tmp Value = nil
		expr.Eval(state, env, &tmp)

		if first == nil {
			first = tmp
		}

		switch first.(type) {
		case *NixInt:
			switch v := tmp.(type) {
			case *NixInt:
				n += *v
			case *NixFloat:
				first = v
				nf = NixFloat(n)
				nf += NixFloat(*v)
			default:
				return errors.New("cannot add ??? to an integer, at ???")
				//throwEvalError("cannot add %1% to an integer, at %2%", showType(vTmp), pos);
			}
		case *NixFloat:
			switch v := tmp.(type) {
			case *NixInt:
				nf += NixFloat(*v)
			case *NixFloat:
				nf += *v
			default:
				return errors.New("cannot add ??? to an float, at ???")
				//throwEvalError("cannot add %1% to an float, at %2%", showType(vTmp), pos);
			}
		default:
			_, isString := first.(*NixString)
			s += state.coerceToString(noPos /*self.Pos*/, tmp, context, false, isString)
		}
	}

	switch first.(type) {
	case *NixInt:
		*val = &n
	case *NixFloat:
		*val = &nf
	default:
		// mkString(v, s.str(), context);
		panic("not implemented")
	}

	return nil
}

func (self *ExprConcatStrings) BindVars(staticEnv *StaticEnv) error {
	for _, expr := range self.Exprs {
		err := expr.BindVars(staticEnv)
		if err != nil {
			return nil
		}
	}

	return nil
}

func (self *ExprConcatStrings) SetName(name Symbol) {}

func (self *ExprConcatStrings) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprAttrs struct {
	Recursive       bool
	AttrDefs        map[string]AttrDef
	DynamicAttrDefs map[string]DynamicAttrDef
}

func (self *ExprAttrs) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprAttrs) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprAttrs) SetName(name Symbol) {}

func (self *ExprAttrs) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

type AttrDef struct {
	Pos       Pos
	Inherited bool
	Expr      Expr
	Displ     uint
}

type DynamicAttrDef struct {
	//Pos Pos
	NameExpr  Expr
	ValueExpr Expr
}

//------------------------------------------------------------------------------

type ExprLambda struct {
	Pos
	Name       Symbol
	MatchAttrs bool
	Formals    Formals
	Body       Expr
}

func (self *ExprLambda) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprLambda) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprLambda) SetName(name Symbol) {
	self.Name = name
}

func (self *ExprLambda) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprAssert struct {
	Pos
	Cond Expr
	Body Expr
}

func (self *ExprAssert) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprAssert) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprAssert) SetName(name Symbol) {}

func (self *ExprAssert) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprWith struct {
	Pos
	Attrs Expr
	Body  Expr
}

func (self *ExprWith) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprWith) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprWith) SetName(name Symbol) {}

func (self *ExprWith) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprLet struct {
	Pos
	Attrs *ExprAttrs
	Body  Expr
}

func (self *ExprLet) Eval(state *EvalState, env *Env, val *Value) error {
	var err error

	// Create a new environment that contains the attributes in this
	// `let'.
	env2 := &Env{
		up:       env,
		prevWith: 0,
		kind:     EnvPlain,
		values:   make([]Value, len(self.Attrs.AttrDefs)),
	}

	// The recursive attributes are evaluated in the new environment,
	// while the inherited attributes are evaluated in the original
	// environment.
	displ := uint(0)
	for _, attr := range self.Attrs.AttrDefs {
		var val *Value
		if attr.Inherited {
			val, err = attr.Expr.MaybeThunk(state, env)
		} else {
			val, err = attr.Expr.MaybeThunk(state, env2)
		}

		if err != nil {
			return err
		}

		env2.values[displ] = *val
		displ += 1
	}

	err = self.Body.Eval(state, env2, val)
	if err != nil {
		return err
	}

	return nil
}

func (self *ExprLet) BindVars(staticEnv *StaticEnv) error {
	newEnv := &StaticEnv{
		up:     staticEnv,
		isWith: false,
		vars:   map[Symbol]uint{},
	}

	displ := uint(0)
	for name, attr := range self.Attrs.AttrDefs {
		newEnv.vars[Symbol(name)] = displ
		attr.Displ = displ
		displ += 1
	}

	for _, attr := range self.Attrs.AttrDefs {
		if attr.Inherited {
			attr.Expr.BindVars(staticEnv)
		} else {
			attr.Expr.BindVars(newEnv)
		}
	}

	err := self.Body.BindVars(newEnv)
	if err != nil {
		return err
	}

	return nil
}

func (self *ExprLet) SetName(name Symbol) {}

func (self *ExprLet) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprIf struct {
	Pos
	Cond Expr
	Then Expr
	Else Expr
}

func (self *ExprIf) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprIf) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprIf) SetName(name Symbol) {}

func (self *ExprIf) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type UnOp byte

const (
	UnOpNegate UnOp = iota
	UnOpNot
)

type ExprUnOp struct {
	Pos
	Type UnOp
	Expr Expr
}

func (self *ExprUnOp) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprUnOp) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprUnOp) SetName(name Symbol) {}

func (self *ExprUnOp) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type BinOp byte

const (
	BinOpNEQ BinOp = iota
	BinOpEQ
	BinOpLE
	BinOpLEQ
	BinOpGE
	BinOpGEQ
	BinOpAnd
	BinOpOr
	BinOpImpl
	BinOpUpdate
	BinOpHasAttr
	BinOpSub
	BinOpMult
	BinOpDiv
	BinOpConcat
)

type ExprBinOp struct {
	Pos
	Type  BinOp
	Expr1 Expr
	Expr2 Expr
}

func (self *ExprBinOp) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprBinOp) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprBinOp) SetName(name Symbol) {}

func (self *ExprBinOp) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprApp struct {
	Pos
	Fun Expr
	Arg Expr
}

func (self *ExprApp) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprApp) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprApp) SetName(name Symbol) {}

func (self *ExprApp) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprVar struct {
	Pos
	Name Symbol

	/* Whether the variable comes from an environment (e.g. a rec, let
	   or function argument) or from a "with". */
	FromWith bool

	/* In the former case, the value is obtained by going `level'
	   levels up from the current environment and getting the
	   `displ'th value in that environment.  In the latter case, the
	   value is obtained by getting the attribute named `name' from
	   the set stored in the environment that is `level' levels up
	   from the current one.*/
	Level uint
	Displ uint
}

func (self *ExprVar) Eval(state *EvalState, env *Env, val *Value) error {
	v2, err := state.lookupVar(env, self, false)
	if err != nil {
		return err
	}

	err = state.ForceValue(v2, self.Pos)
	if err != nil {
		return err
	}

	*val = *v2
	return nil
}

func (self *ExprVar) BindVars(staticEnv *StaticEnv) error {
	// Check whether the variable appears in the environment.  If so,
	// set its level and displacement. */
	curEnv := staticEnv
	level := uint(0)
	withLevel := -1
	for ; curEnv != nil; curEnv, level = curEnv.up, level+1 {
		if curEnv.isWith {
			if withLevel == -1 {
				withLevel = int(level)
			}
		} else {
			if displ, ok := curEnv.vars[self.Name]; ok {
				self.FromWith = false
				self.Level = level
				self.Displ = displ
				return nil
			}
		}
	}

	// Otherwise, the variable must be obtained from the nearest
	// enclosing `with'.  If there is no `with', then we can issue an
	// "undefined variable" error now.
	if withLevel == -1 {
		return errors.New("undefined variable ??? at ???")
		//throw UndefinedVarError(format("undefined variable '%1%' at %2%") % name % pos);
	}

	self.FromWith = true
	self.Level = uint(withLevel)

	return nil
}

func (self *ExprVar) SetName(name Symbol) {}

func (self *ExprVar) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	v, err := state.lookupVar(env, self, true)
	if err != nil {
		return nil, err
	}

	if v != nil {
		return v, nil
	}

	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprSelect struct {
	Pos
	Expr     Expr
	AttrPath []AttrName
	Def      Expr
}

func (self *ExprSelect) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprSelect) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprSelect) SetName(name Symbol) {}

func (self *ExprSelect) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

type AttrName struct {
	Symbol Symbol
	Expr   Expr
}

//------------------------------------------------------------------------------

type ExprInt struct {
	Value *Value
	Int   int64
}

func (self *ExprInt) Eval(state *EvalState, env *Env, val *Value) error {
	*val = *self.Value
	return nil
}

func (self *ExprInt) BindVars(staticEnv *StaticEnv) error { return nil }

func (self *ExprInt) SetName(name Symbol) {}

func (self *ExprInt) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return self.Value, nil
}

//------------------------------------------------------------------------------

type ExprPath struct {
	Value *Value
	Path  string
}

func (self *ExprPath) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprPath) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprPath) SetName(name Symbol) {}

func (self *ExprPath) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return self.Value, nil
}

//------------------------------------------------------------------------------

type ExprFloat struct {
	Value *Value
	Float float64
}

func (self *ExprFloat) Eval(state *EvalState, env *Env, val *Value) error {
	*val = *self.Value
	return nil
}

func (self *ExprFloat) BindVars(staticEnv *StaticEnv) error { return nil }

func (self *ExprFloat) SetName(name Symbol) {}

func (self *ExprFloat) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return self.Value, nil
}

//------------------------------------------------------------------------------

type ExprString struct {
	Value  *Value
	String Symbol
}

func (self *ExprString) Eval(state *EvalState, env *Env, val *Value) error {
	*val = *self.Value
	return nil
}

func (self *ExprString) BindVars(staticEnv *StaticEnv) error { return nil }

func (self *ExprString) SetName(name Symbol) {}

func (self *ExprString) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return self.Value, nil
}

//------------------------------------------------------------------------------

// Temporary structure during parse.
type exprIndStr struct {
	String []byte
}

func (self *exprIndStr) Eval(state *EvalState, env *Env, val *Value) error {
	panic("exprIndStr: Eval should never be called")
}

func (self *exprIndStr) BindVars(staticEnv *StaticEnv) error {
	panic("exprIndStr: BindVars should never be called")
}

func (self *exprIndStr) SetName(name Symbol) {
	panic("exprIndStr: SetName should never be called")
}

func (self *exprIndStr) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	panic("exprIndStr: MaybeThunk should never be called")
}

//------------------------------------------------------------------------------

type ExprList struct {
	Elems []Expr
}

func (self *ExprList) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprList) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprList) SetName(name Symbol) {}

func (self *ExprList) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//------------------------------------------------------------------------------

type ExprPos struct {
	Pos Pos
}

func (self *ExprPos) Eval(state *EvalState, env *Env, val *Value) error {
	panic("not implemented")
}

func (self *ExprPos) BindVars(staticEnv *StaticEnv) error {
	panic("not implemented")
}

func (self *ExprPos) SetName(name Symbol) {}

func (self *ExprPos) MaybeThunk(state *EvalState, env *Env) (*Value, error) {
	return defaultMaybeThunk(self, state, env)
}

//--------------------------------

//type Position struct {
//	Filename string // filename, if any
//	Offset   int    // offset, starting at 0
//	Line     int    // line number, starting at 1
//	Column   int    // column number, starting at 1 (byte count)
//}
// IsValid reports whether the position is valid.
//func (pos *Position) IsValid() bool { return pos.Line > 0 }

//type Token struct {
//	TokenType TokenType
//	Start     int
//	End       int
//	Text      []byte
//}

//--------------------------------
// EXPRS
