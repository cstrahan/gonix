/*

TODO:

- get basic variable deref working
	- get let working
	- get both forms of `with` working
- sanity check blackholes
- get subtraction working
- get concatenation of strings working

*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Errors
// Error            : BaseError
// InvalidPathError : EvalError
// JSONParseError   : EvalError
// EvalError        : Error
// AssertionError   : EvalError
// Abort            : EvalError
// TypeError        : EvalError
type EvalError string

func (self EvalError) Error() string { return string(self) }

// values
type Value interface {
	isValue() // private
}

// Approximation of a Value union type
type NixInt int64
type NixBool bool
type NixString struct {
	String  []byte
	Context [][]byte
}
type NixPath string
type Null struct{}
type NixAttrs Bindings
type NixList []Value
type Thunk struct {
	Env  *Env
	Expr Expr
}
type App struct {
	Left  Expr
	Right Expr
}
type Lambda struct {
	Env *Env
	Fun ExprLambda
}
type Blackhole struct{}
type PrimOp struct{}
type PrimOpApp struct {
	Left  Value
	Right Value
}
type NixFloat float64

func (self NixInt) isValue()    {}
func (self NixBool) isValue()   {}
func (self NixString) isValue() {}
func (self NixPath) isValue()   {}
func (self Null) isValue()      {}
func (self NixAttrs) isValue()  {}
func (self NixList) isValue()   {}
func (self Thunk) isValue()     {}
func (self App) isValue()       {}
func (self Lambda) isValue()    {}
func (self Blackhole) isValue() {}
func (self PrimOp) isValue()    {}
func (self PrimOpApp) isValue() {}
func (self NixFloat) isValue()  {}

type Env struct {
	//size     uint16 // used by ‘valueSize’
	up       *Env
	prevWith uint16 // nr of levels up to next `with' environment
	kind     EnvKind
	values   []Value
}

type EnvKind byte

const (
	EnvPlain EnvKind = iota
	EnvHasWithExpr
	EnvHasWithAttrs
)

type StaticEnv struct {
	isWith bool
	up     *StaticEnv
	vars   map[Symbol]uint
}

type Attr struct {
	name  Symbol
	value *Value
	//pos *Pos
}

type Bindings []Attr

func (self *Bindings) Add(name Symbol, val *Value) {
	*self = append([]Attr(*self), Attr{name: name, value: val})
}

func (self Bindings) Find(name Symbol) *Value {
	// TODO: make this more efficient
	for _, kv := range self {
		if string(kv.name) == string(name) {
			return kv.value
		}
	}

	return nil
}

type SymbolTable struct {
	symbols map[string]string
}

func (self SymbolTable) Create(s string) Symbol {
	if found, ok := self.symbols[s]; ok {
		return Symbol(found)
	}

	self.symbols[s] = s
	return Symbol(s)
}

type EvalState struct {
	baseEnv       *Env
	staticBaseEnv *StaticEnv
	symbols       SymbolTable
	baseEnvDispl  uint
}

func (self *EvalState) createBaseEnv() {
	builtins := NixAttrs(make([]Attr, 128))
	self.addConstant("builtins", &builtins)

	n := NixInt(42)
	self.addConstant("__foo", Value(&n))
}

func (self *EvalState) coerceToString(pos Pos, val Value, context [][]byte, coerceMore bool, copyToStore bool) string {
	panic("not implemented")
}

func (self *EvalState) lookupVar(env *Env, variable *ExprVar, noEval bool) (*Value, error) {
	for l := variable.Level; l > 0; l, env = l-1, env.up {
	}

	if !variable.FromWith {
		v := env.values[variable.Displ]
		if v == nil {
			return nil, nil
		} else {
			return &env.values[variable.Displ], nil
		}
	}

	for {
		if env.kind == EnvHasWithExpr {
			if noEval {
				return nil, nil
			}

			var v Value = nil
			err := self.EvalAttrs(env.up, env.values[0].(Expr), &v)
			if err != nil {
				return nil, err
			}

			env.values[0] = v
			env.kind = EnvHasWithAttrs
		}

		val := Bindings(env.values[0].(NixAttrs)).Find(variable.Name)
		if val != nil {
			return val, nil
		}

		if env.prevWith == 0 {
			//errors.New("undefined variable '%1%' at %2%", var.name, var.pos)
			return nil, errors.New("undefined variable '" + string(variable.Name) + "'")
		}

		for l := env.prevWith; l > 0; l, env = l-1, env.up {
		}
	}
}

func (self *EvalState) EvalAttrs(env *Env, expr Expr, val *Value) error {
	err := expr.Eval(self, env, val)
	if err != nil {
		return err
	}

	if _, ok := (*val).(NixAttrs); !ok {
		return errors.New("value is ??? while a set was expected")
		//throwTypeError("value is %1% while a set was expected", v);
	}

	return nil
}

func (self *EvalState) Eval(expr Expr, val *Value) error {
	return expr.Eval(self, self.baseEnv, val)
}

func (self *EvalState) ForceAttrs(val *Value, pos Pos) (Bindings, error) {
	err := self.ForceValue(val, pos)
	if err != nil {
		return nil, err
	}

	switch v := (*val).(type) {
	case *NixAttrs:
		return (Bindings)(*v), nil
	default:
		if pos != noPos {
			panic("attrs was expected")
			// throwTypeError("value is %1% while a set was expected, at %2%", v, pos);
		} else {
			panic("attrs was expected")
			// throwTypeError("value is %1% while a set was expected", v);
		}
	}

	return nil, nil
}

func (self *EvalState) ForceString(val *Value, pos Pos) (string, error) {
	err := self.ForceValue(val, pos)
	if err != nil {
		return "", err
	}

	switch v := (*val).(type) {
	case *NixString:
		return string(v.String), nil
	default:
		if pos != noPos {
			panic("string was expected")
			// throwTypeError("value is %1% while a string was expected, at %2%", v, pos);
		} else {
			panic("string was expected")
			// throwTypeError("value is %1% while a string was expected", v);
		}
	}

	return "", nil
}

func copyContext(val Value, context *[][]byte) {
	from := val.(*NixString).Context
	if from != nil {
		for _, path := range from {
			*context = append(*context, path)
		}
	}
}

func (self *EvalState) ForceStringWithCtx(val *Value, context [][]byte, pos Pos) (string, error) {
	str, err := self.ForceString(val, pos)
	if err != nil {
		return "", err
	}
	copyContext(*val, &context)
	return str, nil
}

func (self *EvalState) ForceStringNoCtx(val *Value, pos Pos) (string, error) {
	str, err := self.ForceString(val, pos)
	if err != nil {
		return "", err
	}

	if (*val).(*NixString).Context != nil {
		if pos != noPos {
			panic("not allowed to refer to a store path")
			//throwEvalError("the string '%1%' is not allowed to refer to a store path (such as '%2%'), at %3%",
			//    v.string.s, v.string.context[0], pos);
		} else {
			panic("not allowed to refer to a store path")
			//throwEvalError("the string '%1%' is not allowed to refer to a store path (such as '%2%')",
			//    v.string.s, v.string.context[0]);
		}
	}

	return str, nil
}

func (self *EvalState) ForceValue(val *Value, pos Pos) error {
	if thunkPtr, ok := (*val).(*Thunk); ok {
		thunk := *thunkPtr
		env := thunk.Env
		expr := thunk.Expr
		*val = &Blackhole{}
		err := expr.Eval(self, env, val)
		if err != nil {
			*val = thunk
			return err
		}
	} else if app, ok := (*val).(*App); ok {
		return self.callFunction(app.Left, app.Right, val, noPos)
	} else if _, ok := (*val).(*Blackhole); ok {
		return EvalError(fmt.Sprintf("infinite recursion encountered, at %v", pos))
	}

	return nil
}

func (self *EvalState) callFunction(left Expr, right Expr, val *Value, pos Pos) error {
	panic("callFunction not impl.")
	return nil
}

func (self *EvalState) addConstant(name string, v Value) {
	self.staticBaseEnv.vars[self.symbols.Create(name)] = self.baseEnvDispl
	self.baseEnv.values = append(self.baseEnv.values, v)
	self.baseEnvDispl += 1

	name = strings.TrimPrefix(name, "__")
	bindings := (*Bindings)(self.baseEnv.values[0].(*NixAttrs))
	bindings.Add(self.symbols.Create(name), &v)
}
