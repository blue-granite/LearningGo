`go build ./...` at a module root will build module
`go build -o dir/dir/filename.ext ./...` at a module root will build module and save it as the specified file.
`go install ./...` at a module root will build and install the module to GOBIN (if set), else $GOPATH/bin (if set) else $HOME/go/bin.