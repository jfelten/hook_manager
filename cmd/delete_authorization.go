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
	"context"
	"strings"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"github.com/google/go-github/github"

	"github.com/spf13/cobra"
)

var (
	auth_id int
	
)

// delete_authorizationCmd represents the delete_authorization command
var delete_authorizationCmd = &cobra.Command{
	Use:   "delete_authorization",
	Short: "creates an github auth token that has permission to admin webhooks for automated management",
	Long: `This will generate an an authization for any github account, but it is intended for bot accounts. For example:

hook_manager delete_authorization --owner=<MY_BOT_USER>`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		delete_authorization(account, auth_id)
	},
}

func init() {
	RootCmd.AddCommand(delete_authorizationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delete_authorizationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delete_authorizationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	delete_authorizationCmd.Flags().IntVar(&auth_id, "auth_id", 0, "github authorization id")
	delete_authorizationCmd.Flags().StringVar(&account, "account", "", "account user name ( you will be prompted for password )")
	
}


func delete_authorization(account string, id int) {
	

	fmt.Printf("%s's GitHub Password: ",account)
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(account),
		Password: strings.TrimSpace(password),
	}
	
	ctx := context.Background()
	var client = github.NewClient(tp.Client())
	
    resp, err := client.Authorizations.Delete(ctx, id)
    if err != nil {
    	fmt.Printf("\n\x1b[31;1m%s\x1b[0m\n",err.Error())
    } else if (resp == nil) {
    	fmt.Println("\x1b[31;1mNo response from github\x1b[0m")
    	
    } else {
    	fmt.Printf("\n\x1b[32;1m%v deleted\x1b[0m\n",id)
    }

}
