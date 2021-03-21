/*
Copyright © 2021 HASEB ANSARI ansari-haseb

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seeree-box",
	Short: "A CLI tool to use TMDB APIs to show TV Shows details",
	Long: `The CLI tool is an iteractive executable which starts a session with the user
			and asks for the title of the TV Show to search for. The user then gets the list of the searched 
			TV Shows in the form of prompt Select list. The user can choose any TV Show from the list and then
			the available list of Season for that TV Show will appear. Choosing any of the Season from the list will 
			take the user to that Season's list of Episodes with the title of the episode. Then finally, once the user selects
			any Episode from the list, the title and the summary of the Episode will be visible to the user in the console.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seeree-box.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

}
