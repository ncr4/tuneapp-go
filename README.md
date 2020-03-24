# Tuneup Technology App Go Client Library

The Go client library for the Tuneup Technology App.

**Currently in beta**, the Go client library allows you to interact with the customers, tickets, inventory, and locations objects of the application.

## Install

```bash
go get -u github.com/ncr4/tuneuptechnology-go
```

## Example

```go
package main

import (
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Define your credentials
	var auth string = "YOUR_EMAIL_HERE"
	var api_key string = "YOUR_API_KEY_HERE"
	
	// Call an action
	tuneuptechnology.RetrieveCustomers(auth, api_key)
}
```

## Usage

```bash
go run my_file.go
```

## Documentation

Up-to-date documentation can be [found here](https://app.tuneuptechnology.com/docs/api).

## TODO

- User-Agent
- Client setup (API Key/Auth)
- Response (JSON pretty printed)
- Error handling
- Consolidate and re-use code
