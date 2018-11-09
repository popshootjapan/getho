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

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Get nodes.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := fmt.Sprintf("%v", viper.Get("token"))

		nodes, err := api.GetNodes(token)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, node := range nodes {
			fmt.Println(node.Subdomain)
		}
	},
}

func init() {
	rootCmd.AddCommand(nodesCmd)
}
