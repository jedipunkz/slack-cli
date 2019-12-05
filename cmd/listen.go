/*
Copyright © 2019 Tomokazu HIRAI <tomokazu.hirai@gmaill.com>

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
	"os"
	"os/exec"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen incoming messages",
	Long:  `Listen incoming slack messages.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			log.Println("%s", len(args))
			return errors.New("No required argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		api := slack.New(token)
		os.Exit(run(api))
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		msg := <-rtm.IncomingEvents
		log.Printf("MSG: %#v\n", msg.Data)

		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			log.Printf("Start up!")

		case *slack.MessageEvent:
			// if strings.Contains(ev.Text, "!terraform") {
			if strings.HasPrefix(ev.Text, "terraform") {
				out, err := exec.Command("sh", "-c", ev.Text[:]).Output()
				if err != nil {
					log.Println("Fatal error : %s \n", err)
				}

				rtm.SendMessage(rtm.NewOutgoingMessage("```"+string(out)+"```", ev.Channel))
			}
		}
	}
}
