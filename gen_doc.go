package main

import (
	"github.com/antoninguyot/alldebrid-cli/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	cobra.CheckErr(doc.GenMarkdownTree(cmd.RootCmd, "docs/"))
}
