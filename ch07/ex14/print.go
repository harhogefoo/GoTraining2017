package eval

import (
	"bytes"
	"fmt"
)

func Format(e Expr) string {
	var buffer bytes.Buffer
	out(&buffer, e)
	return buffer.String()
}

func out(buffer *bytes.Buffer, expr Expr) {
	switch e := expr.(type) {
	case literal:
		fmt.Fprintf(buffer, "%g", e)

	case Var:
		fmt.Fprintf(buffer, "%s", e)

	case unary:
		fmt.Fprintf(buffer, "(%c", e.op)
		out(buffer, e.x)
		buffer.WriteByte(')')

	case binary:
		buffer.WriteByte('(')
		out(buffer, e.x)
		fmt.Fprintf(buffer, " %c ", e.op)
		out(buffer, e.y)
		buffer.WriteByte(')')

	case call:
		fmt.Fprintf(buffer, "%s(", e.fn)
		for i, arg := range e.args {
			if i > 0 {
				buffer.WriteString(", ")
			}
			out(buffer, arg)
		}
		buffer.WriteByte(')')

	default:
		panic(fmt.Sprintf("unknown Expr: %T", e))
	}
}
