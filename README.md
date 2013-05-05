Golab
=====

A playground for experimenting with the Go Programming language.

The project is broken down into several workspaces: apps, examples, and servers.  

Apps - contains simple command line apps which explore multiple packages and techniques. However, the purpose is to use go to implement simple real works apps.

Examples - contains example programs which explore packages and programming techinques of Go.

Servers - contains simple example server programs both http, pure socket, and rpc based servers.

To build set GOPATH to contain the root of each workspace. For example, if golab is installed in $HOME/golab, then set GOPATH as follows:

$ export GOPATH=$HOME/golab/exmples:$HOME/golab/apps:$HOME/golab/servers
