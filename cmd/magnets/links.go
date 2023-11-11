package magnets

import (
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
)

// linksCmd represents the listMagnets command
var linksCmd = &cobra.Command{
	Use:   "links id",
	Short: "Get links from a magnet",
	Long: `The links command will get links from a magnet.
Once you have the links, you can unlock them with the links unlock command.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var client *http.Client
		var writer *table.Table

		client = cmd.Context().Value("client").(*http.Client)
		writer = cmd.Context().Value("writer").(*table.Table)

		magnetResponse, err := client.ShowMagnet(args[0])
		if err != nil {
			log.Fatal(err)
		}

		writer.AppendHeader(table.Row{"Filename", "Size", "Link"})
		for _, link := range magnetResponse.Data.Magnets.Links {
			writer.AppendRow(table.Row{link.Filenme, link.Size, link.Link})
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
