package magnets

import (
	"github.com/spf13/cobra"
)

// MagnetsCmd represents the magnets command
var MagnetsCmd = &cobra.Command{
	Use:   "magnets",
	Short: "Manage magnets",
	Long:  `The magnets command will allow you to manage your Alldebrid magnets.`,
}

func init() {
	MagnetsCmd.AddCommand(listCmd)
	MagnetsCmd.AddCommand(linksCmd)
	MagnetsCmd.AddCommand(uploadCmd)
	MagnetsCmd.AddCommand(deleteCmd)
	MagnetsCmd.AddCommand(restartCmd)
}
