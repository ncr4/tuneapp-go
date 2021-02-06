# Tuneup Technology App Go Client Library

The Go client library for the Tuneup Technology App.

[![Build Status](https://github.com/tuneuptechnology/tuneuptechnology-go/workflows/build/badge.svg)](https://github.com/tuneuptechnology/tuneuptechnology-go/actions)
[![Licence](https://img.shields.io/github/license/tuneuptechnology/tuneuptechnology-go)](LICENSE)

This library allows you to interact with the customers, tickets, inventory, and locations objects without needing to do the hard work of binding your calls and data to endpoints. Simply call an action such as `CreateCustomer` and pass some data and let the library do the rest.

## Install

```bash
go get -u github.com/tuneuptechnology/tuneuptechnology-go
```

## Example

```go
package main

import (
	"github.com/tuneuptechnology/tuneuptechnology-go"
	"os"
)

func main() {
	// Setup your email and API key
	apiEmail := os.Getenv("API_EMAIL")
	apiKey := os.Getenv("API_KEY")

	// Create a customer passing in all required params
	customer := tuneuptechnology.CreateCustomer(
		&tuneuptechnology.Customer{
			Auth:       apiEmail,
			APIKey:     apiKey,
			Firstname:  "Jake",
			Lastname:   "Peralta",
			Email:      "jake@example.com",
			Phone:      "8015551234",
			UserID:     1,
			Notes:      "Believes he is a good detective.",
			LocationID: 1,
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
