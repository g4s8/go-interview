It's simple HTTP server which handles all requests, parses `src` query parameter
and write it to local file.

It's working with positive test cases. But sometimes it can crash or work incorrectly.
Also, the code could be improved in terms of readability and performance.

The candidate will be asked to find all mistakes and improve the code.

## Usage

To run and test this app use:
 1. `go build` to build server
 2. start with `./go-interview`
 3. test by sending HTTP request to `http://localhost/any?src=val` (use any parameter to log instead of `val`)
 4. check `/tmp/out.txt` for log messages.

## Solving

This project contains many mistakes. Try to find all bugs, improve code quality, and make some optimizations. Submit the results as a pull-request.
