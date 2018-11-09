// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/popshootjapan/getho-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// contractsCmd represents the contracts command
var contractsCmd = &cobra.Command{
	Use:   "contracts",
	Short: "Get smart contracts in a node.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := fmt.Sprintf("%v", viper.Get("token"))

		subdomain, err := cmd.Flags().GetString("subdomain")
		if err != nil || subdomain == "" {
			fmt.Println("Need subdomain. You can set with -s or --subdomain.")
			return
		}

		contracts, err := api.GetContracts(token, subdomain)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, contract := range contracts {
			fmt.Println(contract.Name, contract.Address)
		}

	},
}

func init() {
	rootCmd.AddCommand(contractsCmd)
	contractsCmd.Flags().StringP("subdomain", "s", "", "choice your getho node subdomain.")
}
