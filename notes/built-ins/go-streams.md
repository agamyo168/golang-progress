## io.Writer

In Node.js, you have Writable Streams (like res in an Express handler or process.stdout). In Go, io.Writer is the universal interface for anything you can send data to.

If a type implements the Write method, it is an io.Writer.
|Feature |Node.js Writable Stream |Go io.Writer|
|-----------|---------------------------|------------|
|Core Method| .write(chunk)| .Write(p []byte)|
|Abstraction| Class-based inheritance| Implicit Interface|
|Data Type |Buffers or Strings |Byte Slices ([]byte) only|

The "magic" of io.Writer is that it allows different parts of the standard library to talk to each other without knowing what they are.

For example, the function fmt.Fprintf(writer, "message") doesn't care if it's writing to a file, a network socket, or an HTTP response. As long as it has a Write method, it works.

```go
// Writing to Standard Out (Console)
fmt.Fprintf(os.Stdout, "Hello to the console!")

// Writing to an HTTP Response
fmt.Fprintf(w, "Hello to the browser!")

// Writing to a file
fmt.Fprintf(myFile, "Hello to the disk!")
```

### Common Writers

| Writer              | What it does                   | Express/NodeEquivalent    |
| ------------------- | ------------------------------ | ------------------------- |
| os.Stdout           | Writes to the terminal         | process.stdout            |
| http.ResponseWriter | Sends data to the client       | res (the response object) |
| os.File             | Writes to a file on disk       | fs.createWriteStream()    |
| bytes.Buffer        | Writes to a variable in memory | Buffer.alloc()            |
