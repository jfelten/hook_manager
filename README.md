# Hook Manager

## This is a simple utility designed to automate the management and life cycle of github hooks. 

This is a command line tool designed for scripting automated github webhook setup and configuration.  The idea is to create hooks that last the life time of a process making it ideal for container based frameworks like kubernetes.

How it works:

set up process --> create new access token (prompt for password) --> Use token to create webhooks for list of repos --> on application shut down remove hooks and revoke access token

## Local Installation

Make sure $GOPATH is set then:  

```bash
go get github.com/jfelten/hook_manager
go install github.com/jfelten/hook_manager
```
There is also a public docker image available.

## Usage 

For independent CI solutions that need to manage webhooks.

Usage:
```bash
  hook_manager [command]
```  
or using docker

```bash
  docker run -it jfelten/hook_manager /hook_manager [command]
```
Available Commands:

```bash
  create_authorization creates an github auth token that has permission to admin webhooks for automated management
  create_hmac          A brief description of your command
  create_webhook       creates a webhook on github
  delete_authorization creates an github auth token that has permission to admin webhooks for automated management
  delete_webhook       deletes a webhook on github
  help                 Help about any command

Flags:
      --config string   config file (default is $HOME/.hook_manager.yaml)
  -t, --toggle          Help message for toggle

Use "hook_manager [command] --help" for more information about a command.
```

### Example scripts: example_create.sh and example_delete.sh

The script creates a new github account (bot) auth token and then uses it to genreate hooks for a list of repositories  The only human interaction is to enter the bot's password.  The generated auth and hook ids are stored in files on the local system while the process is running.

```bash
#!/bin/sh

BOT_USER_ID=jfelten
REPOS_TO_HOOK=( "jfelten/knowhow" "jfelten/knowhow-shell" "jfelten/knowhow-server" "jfelten/knowhow-agent" )
HOOK_URL="https://blah.blah-blah.com/some_hook_receiver"

#first create an auth token
$GOPATH/bin/hook_manager create_authorization --account=$BOT_USER_ID --note="bot hook cred"

#store the token values as shell vairables
eval $(cat github_cred)

#now create webhook on the repos
HMAC_KEY=`$GOPATH/bin/hook_manager create_hmac` #used by the webhook security
rm -f ./created_hooks
for repo in "${REPOS_TO_HOOK[@]}"
do
   : 
   HOOK_ID=`$GOPATH/bin/hook_manager create_webhook --credentials=${BOT_USER_ID}:${GITHUB_AUTH_TOKEN} --url=${HOOK_URL} --repo=${repo}`
   echo "${HOOK_ID}:${repo}" >> created_hooks
done

```

### clean up when done:

```bash
#!/bin/sh

BOT_USER_ID=jfelten
#store the token values as shell vairables - created by create script
eval $(cat github_cred)

##int the create script we created a text file with each line containing <HOOK_ID>:<REPO>
while read hook; do
    echo "${hook}"
    values=( ${hook//:/ } ) #split the line into an array
    echo "$GOPATH/bin/hook_manager delete_webhook --credentials=${BOT_USER_ID}:${GITHUB_AUTH_TOKEN} --hook_id=${values[0]} --repo="${values[1]}"
"
    #now delete each hook
    $GOPATH/bin/hook_manager delete_webhook --credentials=${BOT_USER_ID}:${GITHUB_AUTH_TOKEN} --hook_id=${values[0]} --repo="${values[1]}"
done <./created_hooks

#now make sure to remove the created token
$GOPATH/bin/hook_manager delete_authorization --account=${BOT_USER_ID} --auth_id=${GITHUB_AUTH_ID}

```

### To generate a github token as a kubernetes secret using the docker image:

```bash
#!/bin/bash

BOT_USER_ID=jfelten

#first create an auth token
CONTAINER_ID=`docker run -d jfelten/hook_manager tail -f /dev/null`
docker exec -it $CONTAINER_ID /hook_manager create_authorization --account=$BOT_USER_ID --note="${BOT_USER_ID} - generated by hook manager"
eval $(docker exec -it $CONTAINER_ID cat ./github_cred)
docker kill $CONTAINER_ID

# Store as kubernetes secret
kubectl delete secret hookmanager-cred
cat <<EOF | kubectl create -f -
apiVersion: v1
kind: Secret
metadata:
  name: hookmanager-cred
type: Opaque
data:
  auth_id: $(echo "$GITHUB_AUTH_ID" | tr -d '\n\r' | base64)
  auth_token: $(echo "$GITHUB_AUTH_TOKEN" | tr -d '\n\r' | base64)
EOF
```
