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
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"io/ioutil"
	"github.com/google/go-github/github"

	"github.com/spf13/cobra"
)

var (
	account string
	note string
)

// create_authorizationCmd represents the create_authorization command
var create_authorizationCmd = &cobra.Command{
	Use:   "create_authorization",
	Short: "creates an github auth token that has permission to admin webhooks for automated management",
	Long: `This will generate an an authization for any github account, but it is intended for bot accounts. For example:

hook_manager create_authorization --owner=<MY_BOT_USER>`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		create_authorization(account)
	},
}

func init() {
	RootCmd.AddCommand(create_authorizationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// create_authorizationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// create_authorizationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	create_authorizationCmd.Flags().StringVar(&account, "account", "", "github account")
	create_authorizationCmd.Flags().StringVar(&note, "note", "generated by hook_manager", "used by github to scribe the auth")
	
}


func create_authorization(account string) {
	

	fmt.Printf("%s's GitHub Password: ",account)
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(account),
		Password: strings.TrimSpace(password),
	}
	
	ctx := context.Background()
	var client = github.NewClient(tp.Client())
	
	// get all pages of results
	scopes := []github.Scope{ "public_repo", "admin:repo_hook", "notifications"}
	auth_request := &github.AuthorizationRequest {
		Scopes:       scopes,
    	Note:         &note,
	}
    Authorization, resp, err := client.Authorizations.Create(ctx, auth_request)
    if err != nil {
    	fmt.Printf("\n\x1b[31;1m%s\x1b[0m\n",err.Error())
    } else if (resp == nil) {
    	fmt.Println("\x1b[31;1mNo response from github\x1b[0m")
    	
    } else {
    	key_content := fmt.Sprintf("GITHUB_AUTH_ID=%v\nGITHUB_AUTH_TOKEN=%v",*Authorization.ID,*Authorization.Token)
    	bytes := []byte(key_content)
    	err := ioutil.WriteFile("./github_cred",bytes,0644)
    	if err != nil {
    		fmt.Println("\x1b[31;1munable to write results to file\x1b[0m")
    		os.Exit(1)
    	} else {
    		fmt.Printf("\n\x1b[32;1mwrote results to file: github_cred\x1b[0m")
    	}
    	fmt.Printf("\n%v\n",key_content)
    	
    }

}
