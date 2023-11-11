package magnets

import (
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
)

// listCmd represents the listMagnets command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your magnets",
	Long: `The list command will list your magnets.
You can limit the number of results with the --limit flag. By default, it will last 10 magnets uploaded to your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		var client *http.Client
		var writer *table.Table

		client = cmd.Context().Value("client").(*http.Client)
		writer = cmd.Context().Value("writer").(*table.Table)

		magnetsResponse, err := client.ListMagnets()
		if err != nil {
			log.Fatal(err)
		}

		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			log.Fatal(err)
		}

		writer.AppendHeader(table.Row{"ID", "Filename", "Size", "Status"})
		for _, magnet := range magnetsResponse.Data.Magnets[:limit] {
			writer.AppendRow(table.Row{magnet.Id, magnet.Filename, magnet.Size, magnet.Status})
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

func init() {
	listCmd.Flags().IntP("limit", "l", 10, "Limit the number of results")
}
