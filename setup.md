# Go workspace vs go modules
We canmanage code on computer via workspace or go modules

## workspace


They way computer  files should be structured

- bin - compiled stuff
- pkg - archived folders so that used packages do not need to be re-compiled
- src - namespacing package managing
  - github.com
    - <github.com username>
      - folder with code for project / repo
      - folder with code for project / repo


In order to get package we use `go get <packagename>`


## env variables
-GOROOT - path to binary installation of go - need to be set up to current workspace
-GOPATH - path to go workspace


we can get variables by terminal `go env`


## creating module

module can be created via 
`go mod init example.com/test`

this creates new namespace


## downloading git code
 go get -d github.com/GoesToEleven/GolangTraining/...

... will download all things recursively


# GO commands

```bash
go mod int example.com/test # init a module
go fmt ./... # in a working directory will format the code. Dots will format recursivelly
go run main.go # runs code without building
go install *.go # builds executable and puts it in a bin file for all */go files
go get # way of downloading repo with go

```

# on modules
when we are creating modulevia go module namespace.com/namespace. We are creating a file go.mod with our namespace

when we add deps by go get <> they are added to go.sum with a cryptographic hash of the versions.
This file should be added to version controll as it will make sure correct and secure packages are downlaoded