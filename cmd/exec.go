/*
Copyright © 2019 Tomokazu HIRAI <tomokazu.hirai@gmail.com>

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
	"errors"
	"log"
	"os/exec"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "send a result of execution command line.",
	Long: `Send a result of execution command line tool via shell.
and usage of using your command. For example:

slack-cli exec bot 'ls /tmp'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			log.Println("%s", len(args))
			return errors.New("Required least 2 arguments.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		channel := args[0]
		command := args[1]

		out, err := exec.Command("sh", "-c", command[:]).Output()
		if err != nil {
			log.Println("Fatal error: %s", err)
		}

		token := viper.GetString("token")
		api := slack.New(token)

		_, _, errslack := api.PostMessage(
			channel, slack.MsgOptionText(
				"command: "+command+"\n"+"```"+string(out)+"```", false))
		if errslack != nil {
			log.Println("Fatal error: %s", errslack)
			return
		}

		log.Println("Message successfully sent to channel.")
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
