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
	"time"
	"strings"
	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
)

var (
	hmac string
	repo string
	url string
	credentials string
)

// create_webhookCmd represents the create_webhook command
var create_webhookCmd = &cobra.Command{
	Use:   "create_webhook",
	Short: "creates a webhook on github",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create_webhook a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		 create_webhook(credentials, repo, url, hmac)
	},
}

func init() {
	RootCmd.AddCommand(create_webhookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// create_webhookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// create_webhookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	create_webhookCmd.Flags().StringVar(&hmac, "hmac", "", "hmac key used by the hook")
	create_webhookCmd.Flags().StringVar(&repo, "repo", "", "The name of the git repository in the format <OWNER>/<REPO_NAME>")
	create_webhookCmd.Flags().StringVar(&url, "url", "", "The url that receives hook events")
	create_webhookCmd.Flags().StringVar(&credentials, "credentials", "", "credentials in the format <USER>:<PASSWORD_OR_TOKEN>")

}


func create_webhook(credentials string, repo string, url string, hmac string) {
	username := strings.Split(credentials,":")[0]
	password := strings.Split(credentials,":")[1]
	repo_owner := strings.Split(repo,"/")[0]
	repo_name := strings.Split(repo,"/")[1]
	
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	
	ctx := context.Background()
	var client = github.NewClient(tp.Client())
	
	// get all pages of results
	t := time.Now()
	name := "web"
	active := true
	id := 1
	config := make(map[string]interface{})
	config["url"]=url
	config["secret"]=hmac
	hook := &github.Hook {
		CreatedAt: &t,
    	UpdatedAt: &t,
    	Name:      &name,
    	URL:       &url,
    	Events:    []string{"push", "pull-request"},
    	Active:    &active,
    	ID:        &id,
		Config:    config,
	}
    Hook, resp, err := client.Repositories.CreateHook(ctx, repo_owner, repo_name, hook)
    if err != nil {
    	fmt.Printf("\n\x1b[31;1m%s\x1b[0m\n",err.Error())
    	os.Exit(1)
    } else if (resp == nil) {
    	fmt.Println("\x1b[31;1mNo response from github\x1b[0m")
    	os.Exit(1)
    } else {
    	fmt.Printf("%v\n",*Hook.ID)
    }

}
