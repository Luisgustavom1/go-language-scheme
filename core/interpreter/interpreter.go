package interpreter

import (
	"fmt"
	"strconv"

	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
	"github.com/Luisgustavom1/go-language-scheme/core/parser"
)

type Context map[string]any

var builtIns = map[string]func(_ parser.Ast, ctx Context) any{}

func InitializeBuiltIns() {
	builtIns["if"] = func(args parser.Ast, ctx Context) any {
		condition := AstEvaluator(args[0], ctx)
		then := args[1]
		_else := args[2]

		if condition == true {
			return AstEvaluator(then, ctx)
		}

		return AstEvaluator(_else, ctx)
	}

	builtIns["+"] = func(args parser.Ast, ctx Context) any {
		var value int64
		for _, arg := range args {
			value += AstEvaluator(arg, ctx).(int64)
		}
		return value
	}

	builtIns["-"] = func(args parser.Ast, ctx Context) any {
		var value = AstEvaluator(args[0], ctx).(int64)
		for _, arg := range args[1:] {
			value -= AstEvaluator(arg, ctx).(int64)
		}
		return value
	}

	builtIns["begin"] = func(args parser.Ast, ctx Context) any {
		var last any
		for _, arg := range args {
			last = AstEvaluator(arg, ctx)
		}

		return last
	} 

	builtIns["func"] = func(args parser.Ast, ctx Context) any {
		functionName := (*args[0].Literal).Value
		params := *args[1].List
		body := *args[2].List
			
		ctx[functionName] = func(args []any, ctx Context) any {
			if len(params) != len(args) {
				panic(fmt.Sprintf("Expected %d args to `%s`, got %d", len(params), functionName, len(args)))
			}
			// not to change the global context
			functionCtx := copyContext(ctx)

			for i, param := range params {
				functionCtx[(*param.Literal).Value] = args[i] 
			}

			return InitAstWalk(body, functionCtx)
		}

		return ctx[functionName]
	}
}

func AstEvaluator(value parser.AstValue, ctx Context) any {
	if value.Kind == parser.LiteralValue {
		token := *value.Literal

		switch token.Kind {
		case lexer.IntegerToken:
			integer, err := strconv.ParseInt(token.Value, 10, 64)
			if err != nil {
				fmt.Println("Expected an integer, got: ", token.Value)
				panic(err)
			}

			return integer
		case lexer.IdentifierToken:
			return ctx[token.Value]
		}
	}

	return InitAstWalk(*value.List, ctx)
}

func copyContext(original Context) Context {
	copy := Context{}

	for key, value := range original {
		copy[key] = value
	}

	return copy
}

func InitAstWalk(ast parser.Ast, ctx Context) interface{} {
	functionName := (*ast[0].Literal).Value

	if function, ok := builtIns[functionName]; ok {
		return function(ast[1:], ctx)
	}

	userDefinedFunction := ctx[functionName].(func([]any, Context) any)

	var args []any 
	for _, unevaluatedArg := range ast[1:] {
		args = append(args, AstEvaluator(unevaluatedArg, ctx))
	}

	return userDefinedFunction(args, ctx)
}