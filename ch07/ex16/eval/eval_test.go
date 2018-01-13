package eval

import (
	"fmt"
	"math"
	"testing"
)


func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		{"max[10, x, y]", Env{"x": 1, "y": 2}, "10"},
		{"max[10, x, sqrt(A / pi)]", Env{"x": 1, "A": 87616, "pi": math.Pi}, "167"},
		{"min[10, x, y]", Env{"x": 1, "y": 2}, "1"},
		{"avg[x, y, z]", Env{"x": 1, "y": 2, "z": 3}, "2"},
		{"mid[x, y, z]", Env{"x": 3, "y": 2, "z": 1}, "2"},
		{"mid[x, y, z, a]", Env{"x": 5, "y": 6, "z": 1, "a": 10}, "6"},
		{"mid[x, y, z, a, b]", Env{"x": 5, "y": 6, "z": 1, "a": 10, "b": 4}, "5"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
