package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type visitor int

type Node struct {
	val   string
	edges []Node
}

func (n *Node) addEdge(newNode Node) {
	n.edges = append(n.edges, newNode)
}

func (n *Node) walk() {
	fmt.Print(n.val)
	for _, node := range n.edges {
		node.walk()
	}
}

type Tree struct {
	root Node
}

func (t *Tree) walk() {
	t.root.walk()
}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {
	case *ast.AssignStmt:
		if d.Tok != token.DEFINE {
			return v
		}
		for _, name := range d.Lhs {
			printIdent(name)
		}
	case *ast.RangeStmt:
		printIdent(d.Key)
		printIdent(d.Value)
	case *ast.GenDecl:
		for _, spec := range d.Specs {
			if val, ok := spec.(*ast.ValueSpec); ok {
				for _, name := range val.Names {
					fmt.Printf("%v\n", name)
				}
			}
		}
	}
	return v
}

func printIdent(n ast.Node) {
	ident, ok := n.(*ast.Ident)
	if !ok {
		return
	}
	if ident.Name == "_" {
		return
	}
	fmt.Printf("%v\n", ident.Name)
}

func parse() {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", nil, 0)
	if err != nil {
		panic(err)
	}

	//spew.Dump(f)

	var v visitor
	ast.Walk(v, f)
}

func main() {
	parse()
}
