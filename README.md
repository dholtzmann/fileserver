fileserver
========

[Go](http://golang.org) package for serving static files.

The original http.FileServer() will show directory listings, this one does not.

## Example

```go
	func main() {
		fs := simplefileserver.OnlyFilesFilesystem{http.Dir("/tmp/")}
		http.ListenAndServe(":8080", http.FileServer(fs))
	}
```
