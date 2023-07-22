# Golang Basic Unit Test

This project is a very basic example unit test in golang. It's for a learning purpose

## Test Case

This project has two case for testing.

1. test odd or even of number included basic test, testTable, subtest, skipTest
2. test simple http request with mocking

## Running Test

1. test `go test ./... -v`
2. coverage test and convert file to html

   `go test ./... -v --coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html`

