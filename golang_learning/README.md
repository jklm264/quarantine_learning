# Learning Go (GoLang)

- Developed in part by Ken Thompson (THEE ken Thompson) and was started on November 10, 2009 and version 1 came out on 3/28/2012.
- Includes memory safety, garbage collection, structure typing, and [Communiccating-Sequential-Processing](https://en.wikipedia.org/wiki/Go_(programming_language)#:~:text=and-,CSP,--style[)-style concurrency. 
  - C-influenced
  - Concurrency/Multiprocessing: [light-weight processees](https://en.wikipedia.org/wiki/Light-weight_process) called goroutines, channels for IPC, and `select`
  - Uses [interfaces](https://en.wikipedia.org/wiki/Interface_(computing)#In_object-oriented_languages) instead of virtual inheritance...idk
- May be compiled via gccgo or self-hosting compiler
  - Also "[transpiler](https://en.wikipedia.org/wiki/Transpiler)" with GopherJS to product Javascript front-end web dev code
  - will product statically linked native binaries (no dependancies)

- Language takes advantage of multicore, networked machines
  - static typing and run-time efficiency
  - high-performance networking and multiprocessing

## Creating your first code

```go
package main

import (
  "fmt"
  
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
```

In order to run this code you need to create a mod file via `$ go mod init <fielname.go>`. This creates a *go.mod* file.

Then to run the file: `$ go run hello.go`. This will also result in in a *go.sum* file.

## Creating your first module

Running `$ go mode init example.com/greetings` which will create a *go.mod* file which identifies your code as a module that might be used from other code.

Once ready, can build via `$ go build` while in the module's directory.

Last, create the monolothic exe via `$ ./hello.exe`

## Running your code

To simply run a file `$ go run <file.go> <args>`

To create an exe use build (`$ go build`) first then run as executable.

## Error Handling

Use `errors` and `log` modules.

## Misc

- A map is akin to a dict in Python

## Writing Tests

Use `testing` module. Also, each testing function needs to start with *Test* and has to utilize the `testing.T` object, usually passed into functions as `func Testxxxx(*testing.T)`.

Example:
```go
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```
This module may also be used for benchmarking.