package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestXxx(t *testing.T) {
	var gos string = "type Foo struct {	Bar string `json:\"bar\"` }"
	importStr := "package stj\n" + gos
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", importStr, 0)
	if err != nil {
		panic(err)
	}
	for _, d := range f.Decls {
		if f, ok := d.(*ast.GenDecl); ok {
			for _, s := range f.Specs {
				t := s.(*ast.TypeSpec)
				st := t.Type.(*ast.StructType)
				for _, field := range st.Fields.List {
					fmt.Println(t.Name.Name)
					fmt.Println(field.Names[0])
					fmt.Println(field.Tag.Value)
					fmt.Println(field.Type)
				}
			}
		}
	}
}
