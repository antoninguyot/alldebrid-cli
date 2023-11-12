package main

import (
	"github.com/antoninguyot/alldebrid-cli/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

func main() {
	cobra.CheckErr(doc.GenMarkdownTree(cmd.RootCmd, os.Args[1]))
}
