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
