package Commands

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	vt "github.com/VirusTotal/vt-go"
	"github.com/spf13/cobra"
)

var vtkey string
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a url using the Virus total api",

	Run: func(cmd *cobra.Command, args []string) {

		if vtkey == "" {
			fmt.Println("Must pass --vtkey argument.")
			os.Exit(0)
		}

		var urlID = base64.RawURLEncoding.EncodeToString([]byte(args[0]))

		client := vt.NewClient(vtkey)

		analyseslink := "https://www.virustotal.com/api/v3/urls/" + urlID

		analysisurl, parseurlerr := url.Parse(analyseslink)
		fmt.Println(analyseslink)
		if parseurlerr != nil {
			log.Fatal(parseurlerr)
		}

		getId, idErr := client.Get(analysisurl)

		if idErr != nil {
			log.Fatal(idErr)
		}
		var dat map[string]interface{}

		decodeerr := json.Unmarshal(getId.Data, &dat)
		if decodeerr != nil {
			log.Panic(decodeerr)
		}
		// Iterate through data
		for key, value := range dat {
			fmt.Printf("Key: %s\n", key)

			switch v := value.(type) {
			case string:
				fmt.Printf("Value: %s\n", v)
			case float64:
				fmt.Printf("Value: %f\n", v)
			case map[string]interface{}:
				// Recursively handle nested maps
				for nestedKey, nestedValue := range v {

					switch nv := nestedValue.(type) {

					case map[string]interface{}:

						for nestedKey, nestedValue := range nv {
							switch fn := nestedValue.(type) {

							case map[string]interface{}:
								for fkey, nvalue := range fn {
									fmt.Printf("  ----: %s", fkey)
									fmt.Printf("%v\n", nvalue)
								}
							}

							fmt.Printf("  --------: %s", nestedKey)
							fmt.Printf(": %v\n", nestedValue)

						}
						fmt.Printf("  --: %s", nestedKey)
						fmt.Printf(": %v\n", nestedValue)
					}
				}
			default:
				fmt.Printf("Unknown type: %T\n", v)
			}
		}
	},
}

func init() {
	scanCmd.Flags().StringVarP(&vtkey, "vtkey", "v", "", "Virus Total key for api calls")
	rootCmd.AddCommand(scanCmd)
}
