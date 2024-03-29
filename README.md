get-going
=========

Exercises in [Go](http://tour.golang.org/).

```sh
$ go get github.com/mockturtl/get-going
```

###### Debian

Install the `golang` package.  Source code highlighting is available in `vim-syntax-go`.

Export your workspace (`$GOPATH`) and command paths (`$GOPATH/bin`) in `~/.profile`.

View your environment (`$GOROOT`, etc.) with `go env`.

Use [go-binfmt](https://github.com/str1ngs/go-binfmt) to compile and run on the fly.

```sh
$ ./hello.go  # Hello world!
```

[newton.go](http://tour.golang.org/#45)
-----------

use tangent lines to approximate `sqrt(x)`

[wordcount.go](http://tour.golang.org/#46)
--------------

count words in a string

[draw-pixels.go](http://tour.golang.org/#47)
----------------

generate pixel data and write to `foo.css`

[fibonacci.go](http://tour.golang.org/#48)
--------------

print the Fibonacci sequence

[cubic-newton.go](http://tour.golang.org/#49)
-----------------

use tangent lines to approximate `x^(1/3)`

[image-impl.go](http://tour.golang.org/#60)
---------------

implement `image.Image`

```sh
$ go get code.google.com/p/go-tour/pic
```

[substitution-cipher.go](http://tour.golang.org/#61)
------------------------

read bytes to decode secret messages

[bintrees.go](http://tour.golang.org/#70)
-------------

walk and compare binary trees

[prime-sieve.go](http://golang.org/doc/play/sieve.go)
----------------

print the first `N` prime numbers

```sh
GOMAXPROCS=$(nproc) time ./prime-sieve.go
```

[time-bomb.go](http://tour.golang.org/#68)
--------------

`time` and `select`
