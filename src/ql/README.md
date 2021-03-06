QL - Edward Poot
===========

## Generated files
Please note that the files in the following directories are generated by GOCC (Go Parser Generator):

* parser
* lexer
* errors (except the file `typecheckerrors.go`)
* token
* util (except the files `stmtlistcomparison.go` and `strconv.go`)

## Generate lexer/parser
To generate the lexer/parser from the grammar, run the following from the root src directory:

```
gocc -a grammar/ql.bnf
```

All files in the aforementioned folders will be replaced if they already exist.

## Run tests
To run all tests, execute:

```
make test
```
This will run all tests.

## Build
To build the program, execute:

```
make build
```

## Clean
To clean up all binary executables, execute:

```
make clean
```

## Running
Due to the GUI library used (inclusion of dylib file path issue), the `go run` command can not be used to build and run afterwards. Just first do `go build main.go` and then execute `./main`.

```
go build main.go
./main
```

As a shortcut, you also simply run this to achieve the same:

```
make run
```

## Viewing documentation
All documentation can be viewed in your browser by visiting the url [`http://localhost:8080`](http://localhost:8080) after you've executed:

```
make doc
```

## Update/install dependencies
Install the glide packanagement system:

```
brew install glide
```

Install dependencies:

```
glide install
```

Update dependencies:

```
glide update
```

