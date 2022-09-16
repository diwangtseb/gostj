package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/spf13/cobra"
)

var gos string

var gosd string = "type Foo struct {	Bar string `json:\"bar\"` }"

func init() {
	structToJsonCmd.Flags().StringVarP(&gos, "gos", "j", gosd, "stj")
	rootCmd.AddCommand(structToJsonCmd)
}

const Instructions = "conversion go struct to json"

var structToJsonCmd = &cobra.Command{
	Use:   "stjc",
	Short: Instructions,
	Long:  Instructions,
	Run:   ToJson,
}

func ToJson(_ *cobra.Command, _ []string) {
	fmt.Println("toJson", gos)
	importStr := "package stj\n" + gos
	fmt.Println(importStr)
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", importStr, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		switch n.(type) {
		case *ast.StructType:
			// fmt.Println(n)
			return true
		default:
			return false
		}
	})
	// Print the AST.
	// ast.Print(fset, f)
}
