// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"gitlab.lampiris.be/j.garciagonzalez/gitlab-cli/gitlab"

	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a build variable",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if viper.Get("project") == "" {
			fmt.Println("No project name specified")
			os.Exit(1)
		}

		if len(args) == 0 {
			fmt.Println("no variables specified")
			os.Exit(1)
		}

		project := gitlab.GetProject(viper.GetString("project"))

		err := project.AddVariable(args[0])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	variablesCmd.AddCommand(addCmd)
}
