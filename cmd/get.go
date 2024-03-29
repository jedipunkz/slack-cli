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
	"fmt"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get slack information",
	Long:  "Get Slack Informations.",
}

// userCmd represents the user command
var getUserCmd = &cobra.Command{
	Use:   "user",
	Short: "get user's info",
	Long: `Get information about slack user which you specified.
For example:

slack-cli user U800M5M53

This example means that getting information of slack user who has user id 'U800M5M53'.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Required only 1 argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getUser(args[0])
	},
}

var getUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "get users",
	Long:  "Get all users list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New("no need any argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getUsers()
	},
}

var getChannelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "get channels",
	Long:  "Get all channels list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New("no need any argument.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getChannels()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getUserCmd)
	getCmd.AddCommand(getUsersCmd)
	getCmd.AddCommand(getChannelsCmd)
}

func getUser(userid string) {
	token := viper.GetString("token")

	api := slack.New(token)

	user, err := api.GetUserInfo(userid)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n",
		user.ID, user.Profile.RealName, user.Profile.Email)
}

func getUsers() {
	token := viper.GetString("token")

	api := slack.New(token)

	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("ID: %s, Name: %s\n",
			user.ID, user.Name)
	}
}

func getChannels() {
	token := viper.GetString("token")

	api := slack.New(token)

	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Printf("ID: %s, Name: %s, IsPrivate: %v\n",
			channel.ID, channel.Name, channel.IsPrivate)
	}
}
