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