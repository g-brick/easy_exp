package easy_exp

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type expression interface {
	String() string
	eval(*ctx, []*ast) bool
}

type logicOperator interface {
	String() string
	eval(*ctx, []*ast) bool
	ArgNum() int
	IsRightAssociative() bool
	Priority() int
}

type identifier struct {
	name string
}

func newIdentifier(name string) *identifier {
	return &identifier{name: name}
}

func (i *identifier) String() string {
	return i.name
}

func (i *identifier) eval(c *ctx, _ []*ast) (ok bool) {
	return c.include(i.name)
}

type ast struct {
	value    expression
	children []*ast
}

func newAST(value expression, children []*ast) *ast {
	return &ast{
		value:    value,
		children: children,
	}
}

type property struct {
	name               string
	argNum             int
	isRightAssociative bool
	priority           int
}

func (p *property) String() string {
	return p.name
}

func (p *property) ArgNum() int {
	return p.argNum
}

func (p *property) IsRightAssociative() bool {
	return p.isRightAssociative
}

func (p *property) Priority() int {
	return p.priority
}

type and struct {
	*property
}

type ctx struct {
	root string
}

func (c *ctx) include(name string) (ok bool) {
	return strings.Contains(c.root, name)
}

func newAND() *and {
	return &and{
		&property{
			name:               "AND",
			argNum:             2,
			isRightAssociative: false,
			priority:           2,
		},
	}
}

func (a *and) eval(c *ctx, args []*ast) (ok bool) {
	for _, arg := range args {
		if !arg.value.eval(c, arg.children) {
			return
		}
	}
	ok = true
	return
}

type or struct {
	*property
}

func newOR() *or {
	return &or{
		&property{
			name:               "OR",
			argNum:             2,
			isRightAssociative: false,
			priority:           1,
		},
	}
}

func (a *or) eval(c *ctx, args []*ast) (ok bool) {
	for _, arg := range args {
		if arg.value.eval(c, arg.children) {
			ok = true
			return
		}
	}
	return
}

type not struct {
	*property
}

func newNOT() *not {
	return &not{
		&property{
			name:               "NOT",
			argNum:             1,
			isRightAssociative: false,
			priority:           3,
		},
	}
}

func (a *not) eval(c *ctx, args []*ast) (ok bool) {
	ok = !args[0].value.eval(c, args[0].children)
	return
}

type leftParenthesis struct{}

func (l *leftParenthesis) String() string {
	return "("
}
func (l *leftParenthesis) eval(*ctx, []*ast) (ok bool) {
	return
}

type rightParenthesis struct{}

func (r *rightParenthesis) String() string {
	return ")"
}
func (r *rightParenthesis) eval(*ctx, []*ast) (ok bool) {
	return
}

var logicOpSet = map[string]struct{}{
	"!": {},
	" ": {},
	"&": {},
	"|": {},
	"(": {},
	")": {},
}

func buildAST(str []rune) (head *ast, err error) {
	var (
		exec = new(executor)
		strL = len(str)
		i    int
	)
	for i < strL {
		s := string(str[i])
		switch s {
		case "(":

			exec.updateOpStack(&leftParenthesis{})
			i++

		case ")":

			exec.updateOpStack(&rightParenthesis{})
			i++

		case "&":

			if i < strL-1 && string(str[i+1]) == "&" {
				exec.updateOpStack(newAND())
			} else {
				err = throwSyntaxErr(str, i)
				return
			}
			i += 2

		case "|":

			if i < strL-1 && string(str[i+1]) == "|" {
				exec.updateOpStack(newOR())
			} else {
				err = throwSyntaxErr(str, i)
				return
			}
			i += 2

		case "!":

			exec.updateOpStack(newNOT())
			i++

		case " ":

			i++

		default:

			isQuoted := s == `"`
			if isQuoted {
				i++
			}
			token := ""
			for {
				if i >= len(str) {
					break
				}
				if isQuoted && string(str[i]) == `"` {
					break
				}
				_, isLogicOp := logicOpSet[string(str[i])]
				if !isQuoted && isLogicOp {
					break
				}
				if string(str[i]) == `\` {
					i++
				}
				token += string(str[i])
				i++
			}
			if isQuoted {
				i++
			}
			isQuoted = false
			exec.updateOpStack(newIdentifier(token))

		}
	}

	for len(exec.operatorStack) > 0 {
		if !exec.updateOutputAST() {
			err = errors.New(" Invalid Syntax")
			return
		}
	}

	if len(exec.outputAST) > 1 {
		err = errors.New(" Invalid Syntax")
		return
	}
	head = exec.outputAST[0]
	return
}

type executor struct {
	outputAST     []*ast
	operatorStack []interface{}
}

func (exec *executor) updateOpStack(unit interface{}) {
	switch exp := unit.(type) {
	case *identifier:
		exec.outputAST = append(exec.outputAST, newAST(exp, nil))
		exec.checkUnary()
	case *and, *or, *not:
		exec.adjustByLogicOperator(exp.(logicOperator))
		exec.operatorStack = append(exec.operatorStack, exp)
	case *leftParenthesis:
		exec.operatorStack = append(exec.operatorStack, exp)
	case *rightParenthesis:
		exec.adjustByParenthesis()
	}
}

func (exec *executor) checkUnary() (isOk bool) {
	stackL := len(exec.operatorStack)
	if stackL == 0 {
		return
	}
	topOpStack := exec.operatorStack[stackL-1]
	lp, ok := topOpStack.(logicOperator)
	if ok && lp.ArgNum() == 1 { // !
		exec.updateOutputAST()
	}
	return
}

func (exec *executor) adjustByLogicOperator(unit logicOperator) {
	i := len(exec.operatorStack) - 1
	for i > 0 {
		topOpStack := exec.operatorStack[i]
		lp, ok := topOpStack.(logicOperator)
		if !ok || lp == nil {
			return
		}
		if !lp.IsRightAssociative() && lp.Priority() == unit.Priority() || lp.Priority() > unit.Priority() {
			exec.updateOutputAST()
		}
		i--
	}
	return
}

func (exec *executor) adjustByParenthesis() {
	i := len(exec.operatorStack) - 1
	for i > 0 {
		topOpStack := exec.operatorStack[i]
		_, ok := topOpStack.(leftParenthesis)
		if ok {
			break
		}
		exec.updateOutputAST()
		i--
	}
	exec.operatorStack = exec.operatorStack[:len(exec.operatorStack)-1] // pop left parenthesis
	exec.checkUnary()
	return
}

func (exec *executor) updateOutputAST() (ok bool) {
	stackL := len(exec.operatorStack)
	if stackL == 0 {
		return
	}
	topOpStack := exec.operatorStack[stackL-1]
	lp, ok := topOpStack.(logicOperator)
	if !ok {
		return
	}
	exec.operatorStack = exec.operatorStack[:stackL-1] // pop
	var (
		args []*ast
		num  = lp.ArgNum()
		astL = len(exec.outputAST)
	)

	if num > 0 {
		args = append(args, exec.outputAST[astL-num:astL]...)
		exec.outputAST = exec.outputAST[:astL-num]
		exec.outputAST = append(exec.outputAST, newAST(lp, args))
		return
	}
	exec.outputAST = append(exec.outputAST, newAST(lp, nil))
	return
}

func throwSyntaxErr(str []rune, i int) error {
	return fmt.Errorf(" Syntax Error at column %d: %s >>>%s<<<", i, string(str[int(math.Max(0, float64(i)-5)):i]), string(str[i]))
}

// Compile is main func in this package.
func Compile(expRule string) (check func(string) bool, err error) {
	headAst, err := buildAST(([]rune)(expRule))
	if err != nil {
		return
	}
	return func(target string) (ok bool) {
		var strCtx = &ctx{
			root: target,
		}
		if headAst == nil {
			ok = true
			return
		}
		ok = headAst.value.eval(strCtx, headAst.children)
		return
	}, nil
}
