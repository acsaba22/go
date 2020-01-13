# Learning Go programming

## Useful links

* [Presentation](https://docs.google.com/presentation/d/1u4F_frmczfpIRFgloP2UzJAid8W5BkQQwAgXf14ih6M/edit?usp=sharing)
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

Check in `go env` the GOROOT and the GOPATH if they are the ones you like.

Create workspace directory. Default is $HOME/go
Set [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
if you use a different one.

### go tool, the Swiss army knife

#### Write hello world the old way

```
$ cd $GOPATH
$ mkdir -p src/course/hello
$ cd src/course/hello
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
$ go list ...
$ go list course/...
$ go list ...syscall...
$ go help
$ go help list
```

Read documentation from console.

```
$ go doc fmt
$ go doc fmt.Println
```

#### Optional: copy the std packages

If you want to browse the source code of the standard packages
on the web you can do it [here](https://github.com/golang/go/tree/master/src).

If you want to have it locally then create a directory and clone it:

```
git clone https://github.com/golang/go.git
```


### get the training repository

```
$ cd $GOPATH
$ go get -u github.com/acsaba22/go/hellocourse
$ ./bin/hellocourse
```

Another way to run it:

```
$ go run github.com/acsaba22/go/hellocourse
```

Look around.

```
$ go list github.com/acsaba22/go/...
$ cd src/github.com/acsaba22/go
$ ls
$ git log -3
$ cd hellocourse
$ cat main.go
$ go run main.go
```

### Use a library


Edit your hello.go to call the function from the hello. See the documentation:

```
$ go doc github.com/acsaba22/go/hellolib
$ go doc github.com/acsaba22/go/hellolib.Greeting
```

After you finished coding:

```
$ go run course/hello/hello.go
```

Check the sourcecode of hellolib/greetings.go, it contains the actualGreeting function.

Try using hellolib.actualGreeting.

```
$ go run course/hello/hello.go
# command-line-arguments
course/hello/hello.go:12:14: cannot refer to unexported name hellolib.actualGreeting
```

List package details. Observe the dependencies.

```
go list -json course/hello
```

Let's look at building why building/compilation.
Make it compile again.

Look into `pkg/<YOURPLATFORM>/github.com/acsaba22/go` if there is a hellolib.a remove it.

```
$ rm pkg/linux_amd64/github.com/acsaba22/go/hellolib.a
```


```
$ cd $GOPATH
$ go build course/hello/hello.go
$ ./hello
$ ls pkg/linux_amd64/github.com/acsaba22/go
# no hellolib.a
```

Everything is statically linked.

Make compilation faster with -i for install:
```
$ ls pkg/linux_amd64/github.com/acsaba22/go
# no hellolib.a or remove it
$ go install github.com/acsaba22/go/hellolib
$ ls pkg/linux_amd64/github.com/acsaba22/go
# hellolib.a is there, build is faster:
$ go build course/hello/hello.go
$ rm pkg/linux_amd64/github.com/acsaba22/go/hellolib.a
$ go build -i course/hello/hello.go
$ ls pkg/linux_amd64/github.com/acsaba22/go
# hellolib.a is there, subsequent builds are faster
```


## Install an IDE

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

look around in $GOPATH/src and $GOPATH/bin for what was added.

#### GoLand

* no need to subscribe, that's only for email updates.
* [Download](https://www.jetbrains.com/go/), run goland/bin/goland.sh
* When you get username/password click on evaulate
* Open project, course/
* Check that you can debug hello.go
* Checkand that you get code completion.

#### Let's move to modules

Go somewhere outside of $GOPATH.

```
$ mkdir coursemod
$ cd coursemod
$ go mod init mycompany.com/firstgo
$ go list -m
```

If you like git, create a git repository now:

```
$ git init
```

Add this new folder to your VSCode/GoLand environment.

* Create the helloworld program again here.
* Try debugging/launching.
* Try importing and using github.com/acsaba22/go/hellolib

```
$ go get -u github.com/acsaba22/go
$ go list -m ...
<YOURPACKAGE>
github.com/acsaba22/go v1.0.0
go run github.com/acsaba22/go/hellocourse
```

Check where the dependency was added, look around.

```
$ go list -m -f '{{.Dir}}' ...acsaba22/go
```

