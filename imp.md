# Anonymous functions in Go

This is an example I have encountered a lot when coming to defining goroutines and channels. A lot of those have something similar to this:

```golang

go func() {
    // something here
}()

```

Those are called **anonymous functions** in Golang. [Here](https://jeremybytes.blogspot.com/2021/02/go-golang-anonymous-functions-inlining.html) is an article explaining that concept.

The parentheses at the end are there for passing values for the parameters defined in the function. An actual function using that would look something like this:

```golang

 go func(message string) {
      fmt.Println(message)
    }("hello")

```