# Learning Go programming

## Useful links

* [Presentation](https://docs.google.com/presentation/d/1HmMR0fcAK4GcU9IRZZsIIE0gUBAO3mqQebCst0Z-sgA/edit?usp=sharing)
* [Language specification](https://golang.org/ref/spec)
* [Package documentation](https://golang.org/pkg/)
* [Package source code](https://github.com/golang/go/tree/master/src)

## Install environment

### Install Go

https://golang.org/doc/install

Add go/bin directory to you PATH in $HOME/.profile. Test your installation:

```
$ go version
$ go env
```

Check in `go env` the GOPATH if you don't like it you can
[change it](https://github.com/golang/go/wiki/SettingGOPATH).

The main workspace is going to be a different one, the GOPATH will store the downloaded libraries.

Go tool is the swiss army knife. Check out what it can do:
```
$ go help
```

### Install an IDE

#### Visual Studio Code

[Download](https://code.visualstudio.com/)

(A short tour)[https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a]

Check that gopath is correct: ctrl+` (open terminal) and check go env


Add go extension
* `ctrl(mod)+shift+x` (open extensions)
* filter: go lang
* Install rich go language support
* Whenever it asks for installing plugin install.
* add folder to workspace course/
* debug hello/main.go

ctrl-space completion for fmt.Println

#### GoLand

* no need to subscribe, that's only for email updates.
* [Download](https://www.jetbrains.com/go/), run goland/bin/goland.sh
* When you get username/password click on evaulate
* Open project, course/
* Check that you can debug hello.go
* Checkand that you get code completion.


### Let's create the working directory for the codelabs


Choose a working directory where you want to do the codelabs and remember it.
It should be outside of the $GOPATH.

```
$ mkdir gocourse
$ cd gocourse
$ go mod init mycompany.com/gocourse
$ go list -m ...
```

If you like git, create a git repository now:

```
$ git init
```

Add Hello World program:

```
# In gocourse directory.
$ mkdir hello
$ cd hello
$ touch hello.go
```

Edit hello.go:

```
package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}
```

```
$ go run hello.go
$ go build hello.go
$ ./hello
$ go help build
```

You can also run by specifying the whole package name:

```
$ go run go run mycompany.com/gocourse/hello
```

Read documentation from console.

```
$ go doc fmt
$ go doc fmt.Println
```

### get the training repository

The training repository contains the codelabs.

```
$ go cat go.mod
$ go run  github.com/acsaba22/go/hellocourse
$ go cat go.mod

# Notice that a new package was added

$ go list -m ...
```

Check where the dependency was added in your $GOPATH.

```
$ go list -m -f '{{.Dir}}' ...acsaba22/go
# Go to that directory:
$ cd `go list -m -f '{{.Dir}}' ...acsaba22/go`
$ ls
$ cd hellocourse
$ cat main.go
$ go run main.go
```

### Use a library

Go back to your workspace and read the documentation on how to use a function from the training repository:

```
$ cd ~/gocourse
$ go doc github.com/acsaba22/go/hellolib
$ go doc github.com/acsaba22/go/hellolib.Greeting
```

Change hello/hello.go to call the Greeting function and print the result.
Then run again:

```
$ go run hello/hello.go
```

Check the sourcecode of hellolib/greetings.go, it contains the actualGreeting function.

Try using hellolib.actualGreeting.

```
$ go run hello/hello.go
# command-line-arguments
hello/hello.go:12:14: cannot refer to unexported name hellolib.actualGreeting
```

