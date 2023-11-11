package auth

import (
	"fmt"
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"runtime"
	"time"
)

// AuthCmd represents the auth command
var AuthCmd = &cobra.Command{
	Use:   "auth [token]",
	Short: "Authenticate to Alldebrid",
	Long: `The auth command will authenticate you to Alldebrid. It supports two ways of authenticating.
If you already have a token (generated from the Allebdrid console), you can pass it as an argument and it will be saved to your configuration file.
If you don't have a token, the command will open your browser to the Alldebrid console and ask you to enter a code. Once you have entered the code, the command will save the token to your configuration file.`,
	Args: cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var client *http.Client
		client = cmd.Context().Value("client").(*http.Client)

		if len(args) == 1 {
			viper.Set("auth.token", args[0])
			err = viper.WriteConfig()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Token saved to %s\n", viper.ConfigFileUsed())
			return
		}

		pinResponse, err := client.Pin()
		if err != nil {
			log.Fatal(err)
		}

		url := pinResponse.Data.UserUrl

		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Please visit %s and enter the following code: %s\n", url, pinResponse.Data.Pin)
		pinCheckResponse, err := client.PinCheck(pinResponse.Data.Check, pinResponse.Data.Pin)
		if err != nil {
			log.Fatal(err)
		}

		for !pinCheckResponse.Data.Activated {
			pinCheckResponse, err = client.PinCheck(pinResponse.Data.Check, pinResponse.Data.Pin)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(1 * time.Second)
		}

		if pinCheckResponse.Data.Apikey != "" {
			viper.Set("token", pinCheckResponse.Data.Apikey)
			err = viper.WriteConfig()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Token saved to %s\n", viper.ConfigFileUsed())
		} else {
			log.Fatal("No token received")
		}
	},
}
