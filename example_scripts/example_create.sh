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





