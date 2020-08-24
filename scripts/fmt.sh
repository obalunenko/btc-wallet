

gofumpt -s -l $(find . -type f -name '*.go' | grep -v 'vendor' |grep -v '.git' )
