package links

import (
	"fmt"
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

// streamCmd represents the listMagnets command
var streamCmd = &cobra.Command{
	Use:   "stream [-id id] [--best] link",
	Short: "Get streaming links for a link",
	Long: `The stream command will get the streaming links for a link.
It will return the list of available streams and their corresponding links. 
Once listed, you can use the same command with the --id flag to get the link for a specific stream.
You can also use the --best flag to automatically select the best stream quality.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var client *http.Client
		var writer *table.Table

		client = cmd.Context().Value("client").(*http.Client)
		writer = cmd.Context().Value("writer").(*table.Table)

		linkUnlockResponse, err := client.LinkUnlock(args[0])
		if err != nil {
			log.Fatal(err)
		}

		if len(linkUnlockResponse.Data.Streams) == 0 {
			fmt.Println("No stream available for this link")
			return
		}

		idToUnlock, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}

		if best, err := cmd.Flags().GetBool("best"); err == nil && best {
			idToUnlock = linkUnlockResponse.Data.Streams[0].Id
		}

		writer.AppendHeader(table.Row{"ID", "Quality", "Format", "Size", "Link"})
		for _, stream := range linkUnlockResponse.Data.Streams {
			link := ""

			if idToUnlock != "" && stream.Id == idToUnlock {
				linkStreamingResponse, err := client.LinkStreaming(linkUnlockResponse.Data.Id, idToUnlock)
				if err != nil {
					log.Fatal(err)
				}

				if linkStreamingResponse.Data.Delayed > 0 {
					linkDelayedStatus := http.DelayedStatusProcessing
					for linkDelayedStatus == http.DelayedStatusProcessing || linkDelayedStatus == 0 {
						time.Sleep(time.Second)
						linkDelayedResponse, err := client.LinkDelayed(strconv.Itoa(linkStreamingResponse.Data.Delayed))
						if err != nil {
							log.Fatal(err)
						}

						linkDelayedStatus = linkDelayedResponse.Data.Status
						link = linkDelayedResponse.Data.Link
					}

					if linkDelayedStatus == http.DelayedStatusError {
						log.Fatal("Error while waiting for the link to be available")
					}

				} else {
					link = linkStreamingResponse.Data.Link
				}
			}

			writer.AppendRow(table.Row{stream.Id, stream.Quality, stream.Ext, stream.Filesize, link})
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
	streamCmd.Flags().String("id", "", "Stream ID")
	streamCmd.Flags().Bool("best", false, "Automatically select the best stream quality")
}
