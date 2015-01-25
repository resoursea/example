# A showcase for the Resoursea framework
This is a showcase for a bookstore that demonstrate how to create a REST service using the Resoursea framework.

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

Then compile it:

~~~
go build github.com/resoursea/example
~~~

And run it:

~~~
$GOPATH/src/github.com/resoursea/example/example
~~~
