package functionfrequency

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func FunctionFrequency(gocode []byte) []string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src", gocode, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	fnCounter := make(map[string]int, 0)
	ast.Inspect(f, func(n ast.Node) bool {
		ce, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		var fnName string
		switch x := ce.Fun.(type) {
		case *ast.SelectorExpr:
			fnName = getIdentName(x.X) + "." + getIdentName(x.Sel)
		case *ast.Ident:
			fnName = getIdentName(x)
		default:
			return true
		}

		fnCounter[fnName] += 1

		return true
	})

	res := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		max, mk := 0, ""
		for k, v := range fnCounter {
			if v > max {
				max = v
				mk = k
			}
		}

		if max != 0 {
			res = append(res, mk)
			delete(fnCounter, mk)
		}
	}

	return res
}

func getIdentName(n ast.Node) string {
	ident := n.(*ast.Ident)

	return ident.Name
}
