# Hook Manager

## This is a simple utility designed to automate the management of github hooks. 

This is a series of command line tools designed to help script automated github webhook setup and configuration.  The original intent isto automate giuhub hooks kubernetes prow CI.
The idea is to create hooks that last the life time of a CI system deployment and are then discarded when the CI process completes.

How it works:

set up process --> create new access token (prompt for password) --> Uee token to create webhooks for lsit of repos --> on application shut down remove hooks and revoke access token

## Installation

Make sure $GOPATH is set then:  

```bash
go get github.com/jfelten/hook_manager
go install github.com/jfelten/hook_manager
```

## Usage - see the example scripts: example_create.sh and example_delete.sh

The script creates a new github account (bot) auth token and then uses it to genreate hooks for a list of repositories  The only human interaction is to enter the bot's password.  The generated auth and hook ids are stored in files on the local system while the process is running.

```bash
#!/bin/sh

BOT_USER_ID=jfelten
REPOS_TO_HOOK=( "jfelten/knowhow" "jfelten/knowhow-shell" "jfelten/knowhow-server" "jfelten/knowhow-agent" )
HOOK_URL="https://blah.blah-blah.com/some_hook_receiver"

#first create an auth token
$GOPATH/bin/hook_manager create_authorization --account=$BOT_USER_ID --note="my bot hook cred"

#store the token values as shell vairables
eval $(cat github_cred)

#now create webhook on the repos
HMAC_KEY=`$GOPATH/bin/hook_manager create_hmac` #used by the webhook security
for repo in "${REPOS_TO_HOOK[@]}"
do
   : 
   HOOK_ID=`$GOPATH/bin/hook_manager create_webhook --credentials=${BOT_USER_ID}:${GITHUB_AUTH_TOKEN} --url=${HOOK_URL} --repo=${repo}`
   echo -e "${HOOK_ID}:${repo}" >> created_hooks
done
```

## clean up when done:

```bash
#!/bin/sh

BOT_USER_ID=jfelten
#store the token values as shell vairables - created by create script
eval $(cat github_cred)

##int the create script we created a text file with each line containing <HOOK_ID>:<REPO>
for hook in (created_hooks)
do
    echo "${hook}"
    values=(${hook//:/ }) #split the line into an array
    #now delete each hook
    $GOPATH/bin/hook_manager delete_webhook --credentials=${BOT_USER_ID}:${GITHUB_AUTH_TOKEN} --id=${values[0]} --repo=${values[1]}"
done

#now make sure to remove the created token
$GOPATH/bin/hook_manager delete_authorization --credentials=${BOT_USER_ID}:${GITHUB_AUTH_TOKEN} --id=${GITHUB_AUTH_ID}

```