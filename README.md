# [Resoursea](http://resoursea.com)

A high productivity web framework for quickly writing resource based services fully implementing the REST architectural style.

## What is it?

It is an example of a service written in Go using the Resoursea framework

## Getting Started

First [install Go](https://golang.org/doc/install) and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH).

If you added the workspace's `bin` subdirectory to your `PATH`, the fastest way to run the server for this service is:

~~~
go install github.com/soursea/example && example
~~~

But you also can fallow the  steps above...

Install this service using the `go get` tool:

~~~
go get github.com/resoursea/example
~~~

Go to the installed package folder:

~~~
cd $GOPATH/src/github.com/resoursea/example/
~~~

Then compile it:

~~~
go build
~~~

And run it:

~~~
./example
~~~

## Larn More

[The concept, Samples, Documentation, interfaces and resources to use...](http://resoursea.com)