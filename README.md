# Tuneup Technology App Go Client Library

The Go client library for the Tuneup Technology App.

[![Build Status](https://github.com/tuneuptechnology/tuneuptechnology-go/workflows/build/badge.svg)](https://github.com/tuneuptechnology/tuneuptechnology-go/actions)
[![Coverage Status](https://coveralls.io/repos/github/tuneuptechnology/tuneuptechnology-go/badge.svg?branch=main)](https://coveralls.io/github/tuneuptechnology/tuneuptechnology-go?branch=main)
[![Version](https://img.shields.io/github/v/tag/tuneuptechnology/tuneuptechnology-go)](https://github.com/tuneuptechnology/tuneuptechnology-go/releases)
[![Licence](https://img.shields.io/github/license/tuneuptechnology/tuneuptechnology-go)](LICENSE)

This library allows you to interact with the customers, tickets, inventory, and locations objects without needing to do the hard work of binding your calls and data to endpoints. Simply call an action such as `CreateCustomer` and pass some data and let the library do the rest.

## Install

```bash
go get -u github.com/tuneuptechnology/tuneuptechnology-go/v3
```

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go/v3"
)

func main() {
	client := tuneuptechnology.New(os.Getenv("API_EMAIL"), os.Getenv("API_KEY"))

	customer := client.CreateCustomer(
		&tuneuptechnology.Customer{
			Firstname:  "Jake",
			Lastname:   "Peralta",
			Email:      "jake@example.com",
			Phone:      "8015551234",
			UserID:     1,
			Notes:      "Believes he is a good detective.",
			LocationID: 1,
		},
	)

	prettyJSON, err := json.MarshalIndent(customer, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
```

Other examples can be found in the `/examples` directory. Alter according to your needs.

## Usage

```bash
API_EMAIL=email@example.com API_KEY=123... go run create_customer.go
```

## Documentation

Up-to-date API documentation can be [found here](https://app.tuneuptechnology.com/docs/api).

## Development

When re-recording cassettes, comment out the `delete` tests until all other tests have recorded cassettes. This ensures the records that are required to exist will still be active. Once the retrieve and update cassettes have been recorded, uncomment the delete tests and run again to record those cassettes.

```bash
# Build the project
make build

# Install the project globally from source
make install

# Clean the executables
make clean

# Test the project
API_EMAIL=email@example.com API_KEY=123... make test

## Get test coverage
API_EMAIL=email@example.com API_KEY=123... make coverage

# Lint the project (requires golangci-lint be installed)
make lint
```

## Releasing

As a separate PR from the feature/bug PR:

1. Update the Version constant in `client.go`
1. Update `CHANGELOG`
1. Create a GitHub tag with proper Go version semantics (eg: v1.0.0)
