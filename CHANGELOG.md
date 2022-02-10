# CHANGELOG

## v3.0.1 (2022-02-10)

* Corrects the module namespace by appending `/v3`

## v3.0.0 (2021-06-21)

**NOTE: Do not use this release as it was not packaged properly, use v3.0.1+**

* Updates entire library to be compliant with the new v2 API (endpoints, HTTP methods, etc)
* Build requests via a Client now providing your email and apiKey
* Added optional BaseURL and Timeout options to client
* Module names are now plural
* The Client now checks if an email and apiKey is provided and raises an error if not
* Added unit tests (closes #1)
* Removed `PrettyPrint` helper function
* Adds missing fields to inventory
* Adds missing `created_at`, `updated_at`, `deleted_at` fields to all records
* Corrected various other data types
* Added `omitempty` to all JSON fields in each request

## v2.1.0 (2021-02-20)

* Changes `Response` to `makeHTTPRequest`
* Adds a timeout to requests

## v2.0.0 (2021-02-05)

* Overhauled properties of go objects (eg: `Id` to `ID`, `Url` to `URL` etc) - breaking change
* Added GitHub Actions
* Bumped Go version from 1.14 to 1.15
* Added missing user-agent to headers (closes #2)
* Linted the entire project and used gofmt for better formatting

## v1.0.0 (2020/04/17)

* Removed lots of duplicate code
* Use maps properly and return data instead of assuming the user wanted to print instead
* Use built in structures
* Better examples and documentation

## v0.2.0 (2020/03/25)

* Added Inventory actions
* Added Location actions
* Added Ticket actions
* Consolidated and re-used code in a proper client/response struct and function
* Added a complete set of examples
* This release is considered the MVP

## v0.1.0 (2020/03/24)

* Working customer actions
* Added documentation

## v0.0.1 (2020/03/23)

* Initial release, can retrieve customer records
