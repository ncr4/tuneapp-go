# Tuneup Technology App Go Client Library

The Go client library for the Tuneup Technology App.

The Go client library allows you to interact with the customers, tickets, inventory, and locations objects without needing to do the hard work of binding your calls and data to endpoints. Simply call an action such as `CreateCustomer` and pass some data and let the library do the rest.

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
	tuneuptechnology.CreateCustomer(
		// Pass in your email and API key
		&tuneuptechnology.Client{
			Auth: "",
			APIKey: "",
		},

		// Create a customer passing in all required params
		&tuneuptechnology.Customer{
			Firstname: "Go",
			Lastname: "Test",
			Email: "go-test@test.com",
			Phone: "8018981234",
			UserId: 1,
			Notes: "Testing some notes here",
			LocationId: 1,
		},
	)
}
```

Other examples can be found in the `/examples` directory. Alter according to your needs.

## Usage

```bash
go run create_customer.go
```

## Documentation

Up-to-date documentation can be [found here](https://app.tuneuptechnology.com/docs/api).

## Releasing

1. Update the Version constant in `version.go`
1. Update CHANGELOG
1. Create a git tag with proper Go version semantics (eg: v1.0.0)

## TODO

- User-Agent
- Response (JSON pretty printed)
- Error handling
- Consolidate and re-use code
