package eval

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

// Exprは算術式
type Expr interface{
	// Evalは、環境env内でこのExprの値を返します。
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

func (v Var) Eval(env Env) float64 {
	_, ok := env[v]
	if !ok {
		fmt.Printf("value for %s: ", v)
		value, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		value = value[:len(value)-1]
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}
		env[v] = f
	}
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (l list) Eval(env Env) float64 {
	switch l.fn {
	case "min":
		min := l.args[0].Eval(env)
		for i := 1; i < len(l.args); i++ {
			v := l.args[i].Eval(env)
			if min > v {
				min = v
			}
		}
		return min
	case "max":
		max := l.args[0].Eval(env)
		for i := 1; i < len(l.args); i++ {
			v := l.args[i].Eval(env)
			if max < v {
				max = v
			}
		}
		return max
	case "avg":
		avg := float64(0.0)
		for _, expr := range l.args {
			avg += expr.Eval(env)
		}
		return avg / float64(len(l.args))
	case "mid":
		var temp []float64
		for _, expr := range l.args {
			temp = append(temp, expr.Eval(env))
		}
		sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
		return temp[len(temp) / 2]
	}
	panic(fmt.Sprintf("unsupported function call: %s", l.fn))
}

// Var は変数を特定します。例: x
type Var string

// literal は数値定数。例: 3.141
type literal float64

// unary は単項演算子気を表します。例:  -x
type unary struct {
	op rune // '+' か'-'のどちらか
	x Expr
}

// binary は二項演算式を表します。例:  x+y
type binary struct {
	op rune // '+' か'-'のどちらか
	x, y Expr
}

// call は関数呼び出し式を表します。例: sin(x)
type call struct {
	fn string // "pow", "sin", "sqrt"のどれか
	args []Expr
}

type list struct {
	fn   string // "min", "max", "avg", "mid"のどれか
	args []Expr
}

// 変数名を値へと対応付ける環境(environment)
type Env map[Var]float64

