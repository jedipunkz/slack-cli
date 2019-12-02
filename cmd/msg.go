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

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var msgCmd = &cobra.Command{
	Use:   "msg",
	Short: "send message to slack",
	Long: `Slack-cli is a application which sending a Message to Slack.
For example:

slack-cli msg bot 'hello world'

This example mean that sending 'hello world' message to #bot channel on slack.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			log.Println("%s", len(args))
			return errors.New("Require least 2 arguments.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		channel := args[0]

		api := slack.New(token)

		_, _, err := api.PostMessage(channel, slack.MsgOptionText(args[1], false))
		if err != nil {
			log.Println("%s\n", err)
			return
		}
		log.Println("Message successfully sent to channel.")
	},
}

func init() {
	rootCmd.AddCommand(msgCmd)
}
