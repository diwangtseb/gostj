package cmd

import (
	"fmt"

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

func ToJson(cmd *cobra.Command, args []string) {
	fmt.Println("toJson", gos)
}
