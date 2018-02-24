# vgo

__version of go pkg, go 1.10+ requires.__

Once run `vgo build`, file `go.mod` is created with contents:

```
module "github.com/toukii/goutils"

require (
	"github.com/everfore/exc" v0.0.1
	"github.com/fatih/color" v1.5.0
)
```

## install & config

 - 1. download tool

```
go get -u golang.org/x/vgo
```

 - 2. config github.com token

 [https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)

__~/.netrc__


```
machine api.github.com login toukii password xxxx
```

 - 3. type `vgo help`

if success, it is like `go help`.
