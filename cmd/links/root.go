package links

import (
	"github.com/spf13/cobra"
)

// LinksCmd represents the magnets command
var LinksCmd = &cobra.Command{
	Use:   "links",
	Short: "Manage links",
	Long:  `The links command will allow you to manage your Alldebrid links.`,
}

func init() {
	LinksCmd.AddCommand(unlockCmd)
	LinksCmd.AddCommand(streamCmd)
}
