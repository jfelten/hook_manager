#!/bin/sh

BOT_USER_ID=jfelten
GITHUB_AUTH_ID=`kubectl get secret hookmanager-cred --output=jsonpath={.data.auth_id} | base64 --decode | tr -d '\n\r'`
#delete the token on github
docker run -it jfelten/hook_manager /hook_manager delete_authorization --account=${BOT_USER_ID} --auth_id=${GITHUB_AUTH_ID}

#delete the token and secret used by this cluster
kubectl delete secret hookmanager-cred