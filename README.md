# Usefull commands
$ go env -w GO111MODULE=on
$ go mod init terratestmodules && go mod tidy
$ go test -v -run TestTerraformBasicExample

$ go test -v -timeout 30m
