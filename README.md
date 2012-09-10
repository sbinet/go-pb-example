go-pb-example
=============

A simple example modeled after [http://goprotobuf.googlecode.com/hg/README](goprotobuf README) and updated for ``go-1``.

Installation
------------

```
$ go get github.com/sbinet/go-pb-example
$ cd github.com/sbinet/go-pb-example
$ make && go-pb-example
protoc --go_out=. example/test.proto
go get ./example
go get ./go-pb-example
2012/09/10 10:54:09 ref-test: label:"hello" type:17 OptionalGroup{RequiredField:"good bye" } 
2012/09/10 10:54:09 new-test: label:"hello" type:17 OptionalGroup{RequiredField:"good bye" } 

```
