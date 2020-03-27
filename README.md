# Tuneup Technology App Go Client Library

The Go client library for the Tuneup Technology App.

[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

This library allows you to interact with the customers, tickets, inventory, and locations objects without needing to do the hard work of binding your calls and data to endpoints. Simply call an action such as `CreateCustomer` and pass some data and let the library do the rest.

## Install

```bash
go get -u github.com/ncr4/tuneuptechnology-go
```

## Example

```go
package main

import (
	"os"
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	api_email := os.Getenv("API_EMAIL")
	api_key := os.Getenv("API_KEY")

	// Create a customer passing in all required params
	customer := tuneuptechnology.CreateCustomer(
		&tuneuptechnology.Customer{
			Auth: api_email,
			APIKey: api_key,
			Firstname: "Jake",
			Lastname: "Peralta",
			Email: "jake@example.com",
			Phone: "8015551234",
			UserId: 1,
			Notes: "Believes he is a good detective.",
			LocationId: 1,
		},
	)

	tuneuptechnology.PrettyPrint(customer)
}
```

Other examples can be found in the `/examples` directory. Alter according to your needs.

## Usage

```bash
API_EMAIL=email@example.com API_KEY=123... go run create_customer.go
```

## Documentation

Up-to-date API documentation can be [found here](https://app.tuneuptechnology.com/docs/api).

## Releasing

As a separate PR from the feature/bug PR:

1. Update the Version constant in `version.go`
1. Update `CHANGELOG`
1. Create a GitHub tag with proper Go version semantics (eg: v1.0.0)
