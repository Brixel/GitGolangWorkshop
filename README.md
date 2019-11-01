# GitGolangWorkshop

<img src="https://github.com/Brixel/GitGolangWorkshop/raw/master/.github/git-go.png" alt="Golang Github image" width="200" />

Code from the workshop about Git and Golang on 2019-11-03

We have a workshop about Golang and also some Git on Sunday 3 november.
This is the repo that we will use for the Git part and it will contain the code we will build.

## What will you need for this workshop

To follow the workshop, you will need some programs to participate:

- [Git](https://git-scm.com/downloads)
- [Golang](https://golang.org/dl/)
- A text editor to program Go (I will use [Visual Studio Code](https://code.visualstudio.com/download))

<img src="https://github.com/Brixel/GitGolangWorkshop/raw/master/.github/git.png" alt="Golang Github image" height="75" />

### Git

Git is a free and open source distributed version control system designed to handle everything from
small to very large projects with speed and efficiency.

Git is easy to learn and has a tiny footprint with lightning fast performance. It outclasses SCM tools
like Subversion, CVS, Perforce, and ClearCase with features like cheap local branching, convenient
staging areas, and multiple workflows.

#### Installation Windows/Mac

1. Download the [binary](https://git-scm.com/downloads) for your operating system
2. Install the binary.

#### Installation Linux

- Debian based: `sudo apt install git`
- Arch based: `sudo pacman -S git`
- Red Hat based: `sudo yum install git`

<img src="https://github.com/Brixel/GitGolangWorkshop/raw/master/.github/golang.png" alt="Golang image" height="75" />

<img src="https://github.com/Brixel/GitGolangWorkshop/raw/master/.github/gopher.png" alt="Golang Gopher image" height="75" />

### Golang

The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write
programs that get the most out of multicore and networked machines, while its novel type system
enables flexible and modular program construction. Go compiles quickly to machine code yet has
the convenience of garbage collection and the power of run-time reflection. It's a fast, statically
typed, compiled language that feels like a dynamically typed, interpreted language.

#### Installation Windows/Mac

1. Download the [binary](https://golang.org/dl/) for your operating system.
2. Follow the instructions on [the install page](https://golang.org/doc/install) of golang.

#### Installation Linux

Use package manager or follow the [installation instructions](https://golang.org/doc/install)

#### Test installation

You can test the Golang installation by writing a basic 'hello world' program:

``` golang
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

To run the program, navigate to the directory in your terminal and run the `go run .` command.

### Text Editor

You are free to use whatever text editor you want for Go. After all the code is just plain text.
I prefer to use [Visual Studio Code](https://code.visualstudio.com/). Other common options are
[The Go Playground](https://play.golang.org/), [GoLand](https://www.jetbrains.com/go/),
[Atom](https://ide.atom.io/), [Vim](https://www.vim.org/),...

For all these editors, plugins can be found ([plugins](https://golang.org/doc/editors.html)).

### Debugging

<img src="https://github.com/Brixel/GitGolangWorkshop/raw/master/.github/delve.png" alt="Golang Github image" height="75" />

The most commonly used debbuger for go is [delve](https://github.com/go-delve/delve). You can
install the debugger with a `go get` command: `go get -u github.com/go-delve/delve/cmd/dlv`.
