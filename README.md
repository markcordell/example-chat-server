# example-chat-server
This repo is a small demo of building the APIs and datastructure necessary for a chat service written in go.

NOTE: This is not a fully working chat server, just an example of my go code.


## Build instructions
- Place the code into the directory $GOPATH/src/github.com/markcordell/example-chat-server before running the make commands.
- Dependencies are vendored with Dep, and are in the vendor directory.
- Code was built with go version 1.11.4 and tested with go 1.11.5
- Code was built on a linux machine and tested on macOS
- Run Make build to build the code
- Run Make run to run the server
- Run make alive to test the server
