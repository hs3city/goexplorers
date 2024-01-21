# goexplorers

This repository contains various separate modules for [Go](https://go.dev/) self-study coding sessions held at [Hackerspace Tricity](https://github.com/hs3city/).

In order to have multiple modules in a single repo we use [Go workspaces](https://go.dev/doc/tutorial/workspaces).

`go.work` file is already created and checked in to the repository, allowing all session participants to have the same developer experience.

In order to initialize a new Go module in a workspace use the following commands:

```
export NEW_MODULE_DIRECTORY=mymodule
mkdir $NEW_MODULE_DIRECTORY
go work use $NEW_MODULE_DIRECTORY
```

The last command will modify [go.work](./go.work) file. When adding a new module please also remember to add the directory to [linter configuration](./.github/workflows/golangci-lint.yml#39)<br>
The `args` (`jobs.golangci.steps[golangci-lint].args`) key in the linter config defines a set of module directories to be checked by a suite of linters.

Having a linter in place allows us to learn best practices and Go idioms while we're learning to walk 🙂


## How to run modules

### [Quiz](./quiz/)

```
go run ./quiz
```


### [Url Shortener](./urlshort/)

Start the server
```
go run ./urlshort/...
```

Then send a request for a shortened URL to get redirected to it
```
curl -v localhost:8080/hs3
```


### [Choose Your Own Adventure](./cyoa/)

Start the server
```
cd cyoa/main
go run ./...
```

Then point your browser to [localhost:8080/intro](http://localhost:8080/intro) to start the game.


## How to run tests for all modules

```
go test ./cyoa/... ./quiz/... ./urlshort/...
```

From the structure of the command above one can easily guess how to run tests for a single module or how to add more modules.<br>
What if the number of modules is so high that it does not make sense to use that command anymore?<br>
We can use `find` with `maxdepth` flag to select any non-hidden top-level directory, pass it to `sed` to append the dots and finally use `xargs` to invoke the sacred `go` command 💪💪💪

```
find . -maxdepth 1 -type d -not -path '.' -not -path '*/.*' | sed 's|$|/...|' | xargs go test
```

Note that the command above depends on the existence of `go.work` file.

![](https://traust.duckdns.org/api/public/dl/Q9Haw2rT?inline=true)
