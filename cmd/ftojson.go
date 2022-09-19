package cmd

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var fos string

var fosd string = "xxx.go"

func init() {
	FileToJsonCmd.Flags().StringVarP(&fos, "fos", "f", fosd, "ftj")
	rootCmd.AddCommand(FileToJsonCmd)
}

const FileInstructions = "conversion go struct file to json"

var FileToJsonCmd = &cobra.Command{
	Use:   "fstjc",
	Short: FileInstructions,
	Long:  FileInstructions,
	Run:   FToJson,
}

func FToJson(_ *cobra.Command, _ []string) {
	fset := token.NewFileSet() // positions are relative to fset
	fmt.Println(fos)
	dst := parseFile(fos)
	f, err := parser.ParseFile(fset, "", dst, 0)
	if err != nil {
		panic(err)
	}
	t := astParse(f)
	fmt.Println(t.toJson())
}

func parseFile(src string) string {
	if src == "" {
		return ""
	}
	dst, err := ioutil.ReadFile(src)
	if err != nil {
		return ""
	}

	return string(dst)
}
