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
	"os"
	"fmt"
	"context"
	"strings"
	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
)

 var (
	hook_id int
)

// delete_webhookCmd represents the create_webhook command
var delete_webhookCmd = &cobra.Command{
	Use:   "delete_webhook",
	Short: "deletes a webhook on github",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create_webhook a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		 delete_webhook(credentials, repo, hook_id)
	},
}

func init() {
	RootCmd.AddCommand(delete_webhookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delete_webhookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delete_webhookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	delete_webhookCmd.Flags().StringVar(&repo, "repo", "", "The name of the github repository in the format <REPO_OWNER/REPO_NAME")
	delete_webhookCmd.Flags().StringVar(&credentials, "credentials", "", "credentials in format: <GITHUB_USER>:<PASSWORD_OR_TOKEN>")
	delete_webhookCmd.Flags().IntVar(&hook_id, "hook_id", 0, "the id of the hook to delete")

}


func delete_webhook(credentials string, repo string, id int) {
	

	username := strings.Split(credentials,":")[0]
	password := strings.Split(credentials,":")[1]
	repo_owner := strings.Split(repo,"/")[0]
	repo_name := strings.Split(repo,"/")[1]
	
	ctx := context.Background()
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	var client = github.NewClient(tp.Client())
	
    resp, err := client.Repositories.DeleteHook(ctx, repo_owner, repo_name, id)
    if err != nil {
    	fmt.Printf("\n\x1b[31;1m%s\x1b[0m\n",err.Error())
    	os.Exit(1)
    } else if (resp == nil) {
    	fmt.Println("\x1b[31;1mNo response from github\x1b[0m")
    	os.Exit(1)
    } else {
    	fmt.Printf("\x1b[32;1m%v deleted.\x1b[0m\n",id)
    }

}
