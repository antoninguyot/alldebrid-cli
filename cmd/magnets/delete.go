package magnets

import (
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/spf13/cobra"
	"log"
)

// deleteCmd represents the listMagnets command
var deleteCmd = &cobra.Command{
	Use:   "delete id",
	Short: "Delete a magnet",
	Long:  `The delete command will delete a magnet.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var client *http.Client

		client = cmd.Context().Value("client").(*http.Client)

		deleteMagnetResponse, err := client.DeleteMagnet(args[0])
		if err != nil {
			log.Fatal(err)
		}

		println(deleteMagnetResponse.Data.Message)

	},
}
