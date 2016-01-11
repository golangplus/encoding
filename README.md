# encoding
Plus to the standard `encoding` packages.

Godocs: [jsonp](http://godoc.org/github.com/golangplus/encoding/json)

## Featured
```go
// MarshalToFile encodes a data into a JSON file.
func MarshalToFile(v interface{}, filename string) error {}

// UnmarshalFile decodes a JSON from a file.
func UnmarshalFile(filename string, v interface{}) error {}
```
