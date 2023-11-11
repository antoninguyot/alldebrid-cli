package magnets

import (
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// uploadCmd represents the listMagnets command
var uploadCmd = &cobra.Command{
	Use:   "upload <magnet or file>",
	Short: "Upload a magnet or a file to Alldebrid",
	Long: `The upload command will upload a magnet or a file to Alldebrid.
You can either provide a magnet link or a file path.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var client *http.Client
		var writer *table.Table

		client = cmd.Context().Value("client").(*http.Client)
		writer = cmd.Context().Value("writer").(*table.Table)

		var magnets []http.UploadedMagnet

		_, err := os.Stat(args[0])
		if err != nil {
			uploadMagnetResponse, err := client.UploadMagnet(args[0])
			if err != nil {
				log.Fatal(err)
			}
			magnets = uploadMagnetResponse.Data.Magnets
		} else {
			uploadFileResponse, err := client.UploadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
			magnets = uploadFileResponse.Data.Files
		}

		writer.AppendHeader(table.Row{"ID", "Name", "Ready", "Size"})
		for _, magnet := range magnets {
			writer.AppendRow(table.Row{magnet.Id, magnet.Name, magnet.Ready, magnet.Size})
		}

		outputFormat, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatal(err)
		}

		if outputFormat == "csv" {
			writer.RenderCSV()
			return
		}

		writer.Render()
	},
}
