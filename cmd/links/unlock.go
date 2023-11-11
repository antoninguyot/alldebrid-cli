package links

import (
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/spf13/cobra"
	"log"
)

// unlockCmd represents the listMagnets command
var unlockCmd = &cobra.Command{
	Use:   "unlock link",
	Short: "Unlock a link",
	Long: `The unlock command will unlock a link.
It will return the unlocked link as text.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var client *http.Client

		client = cmd.Context().Value("client").(*http.Client)

		linkUnlockResponse, err := client.LinkUnlock(args[0])

		if err != nil {
			log.Fatal(err)
		}

		println(linkUnlockResponse.Data.Link)

	},
}
