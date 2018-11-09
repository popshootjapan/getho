package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/popshootjapan/getho-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Network struct {
	Address string `json:"address"`
}

type Contract struct {
	ContractName string                   `json:"contractName"`
	ABI          []map[string]interface{} `json:"abi"`
	Networks     map[string]Network       `json:"networks"`
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload your smart contract.",
	Long:  `Upload your smart contract. You'll need to specify one of the compiled smart contract json file from build/contracts directory.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contractFilePath := args[0]

		subdomain, err := cmd.Flags().GetString("subdomain")
		if err != nil || subdomain == "" {
			fmt.Println("Need subdomain. You can set with -s or --subdomain.")
			return
		}

		network, err := cmd.Flags().GetString("network")
		if err != nil {
			network = "1010"
		}

		raw, err := ioutil.ReadFile(contractFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		var contract Contract
		json.Unmarshal(raw, &contract)

		abiBytes, err := json.Marshal(contract.ABI)
		if err != nil {
			fmt.Println(err)
			return
		}

		token := fmt.Sprintf("%v", viper.Get("token"))
		status, err := api.PostContracts(token, subdomain, contract.ContractName, contract.Networks[network].Address, string(abiBytes))
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.Contains(status, "201") {
			fmt.Println("Uploaded")
		} else {
			fmt.Println(status)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.SetArgs([]string{"contract.json"})
	uploadCmd.Flags().StringP("network-id", "n", "1010", "choice your networks in trrufle.js.")
	uploadCmd.Flags().StringP("subdomain", "s", "", "choice your getho node subdomain.")
}
