This is a go presentation I built for talking about the basics of go.

To run:

1) Install [go](https://golang.org/doc/install)
2) After making sure that your GOPATH is set to search GOPATH/bin, run 

    go get golang.org/x/net
    go get golang.org/x/tools
    go install golang.org/x/tools/cmd/present

3) Then run `present` in this directory, and you're set to go.
