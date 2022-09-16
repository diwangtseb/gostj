package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"

	"github.com/bytedance/sonic"
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
	t := astParse(f)
	fmt.Println(t.toJson())
}

type Tag []string

func (t *Tag) append(ele string) {
	*t = append(*t, ele)
}
func (t *Tag) toJson() string {
	m := make(map[string]interface{})
	for _, i := range *t {
		m[i] = i
	}
	r, err := sonic.Marshal(m)
	if err != nil {
		panic("struct convert json failed")
	}
	return string(r)
}

func astParse(f *ast.File) *Tag {
	tag := &Tag{}
	for _, d := range f.Decls {
		if f, ok := d.(*ast.GenDecl); ok {
			for _, s := range f.Specs {
				t := s.(*ast.TypeSpec)
				st := t.Type.(*ast.StructType)
				for _, field := range st.Fields.List {
					// log.Println(t.Name.Name)
					// log.Println(field.Names[0])
					// log.Println(field.Tag.Value)
					tag.append(tagParse(field.Tag.Value))
				}
			}
		}
	}
	return tag
}

func tagParse(t string) string {
	// express := `(?<=").+?(?=")`
	express := `".*"`
	r, err := regexp.Compile(express)
	if err != nil {
		return "go struct json tag parse error"
	}
	mt := r.FindStringSubmatch(t)
	mts := mt[0]
	mtsb := []byte(mts)
	mts = string(mtsb[1 : len(mtsb)-1])
	return mts
}
