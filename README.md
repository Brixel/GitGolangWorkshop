# GitGolangWorkshop
Code from the workshop about Git and Golang on 2019-11-03

We have a workshop about Golang and also some Git on Sunday 3 november. 
This is the repo that we will use for the Git part and it will contain the code we will build.

## What will you need for this workshop

To follow the workshop, you will need some programs to participate:
- [Git](https://git-scm.com/downloads)
- [Golang](https://golang.org/dl/)
- A text editor to program Go (I will use [Visual Studio Code](https://code.visualstudio.com/download))

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

### Golang

The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write 
programs that get the most out of multicore and networked machines, while its novel type system 
enables flexible and modular program construction. Go compiles quickly to machine code yet has 
the convenience of garbage collection and the power of run-time reflection. It's a fast, statically 
typed, compiled language that feels like a dynamically typed, interpreted language.

#### Installation Windows

1. Download the [binary](https://golang.org/dl/) for your operating system.
2. Follow the instructions on [the install page](https://golang.org/doc/install) of golang.

#### Installation Linux

Use package manager or follow the [installation instructions](https://golang.org/doc/install)

#### Test installation

You can test the Golang installation by writing a basic 'hello world' program:
```
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```
To run the program, navigate to the directory in your terminal and run the `go run .` command.

