# Hook Manager

## This is a simple utility designed to automate the management of github hooks. 

The purpose is to make it possible to deploy webhook consumers in a kubernetes cluster with automatic consfiguration and set-up.
The idea is to create hooks that last the life time of a CI system deployment and are then discarded when the CI process completes.

## Installation

Make sure $GOPATH is set then:  

```bash
go install https://github.com/jfelten/hook_manager
```

## Usage - see the example scripts: example_create.sh and example_delete.sh

A new github auth token is generated for a bot account and hooks are created using the generate auth token under the bots account.  The only human interaction is to enter the bot's password.  The generate auth and hook ids are stored in files on the local system while the process is running.

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

now clean up when done:

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