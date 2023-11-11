package magnets

import (
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/spf13/cobra"
	"log"
)

// restartCmd represents the listMagnets command
var restartCmd = &cobra.Command{
	Use:   "restart id",
	Short: "Restart a failed magnet",
	Long:  `The restart command will restart a failed magnet.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var client *http.Client

		client = cmd.Context().Value("client").(*http.Client)

		deleteMagnetResponse, err := client.RestartMagnet(args[0])
		if err != nil {
			log.Fatal(err)
		}

		println(deleteMagnetResponse.Data.Message)

	},
}
