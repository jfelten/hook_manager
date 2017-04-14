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





