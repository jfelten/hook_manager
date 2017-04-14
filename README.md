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