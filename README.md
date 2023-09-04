# NIK Extractor
An REST API program for extract and validate any Indonesian ID Card number pattern from a text (use a dummy data).

# Requirements
1. Golang 1.20.6
2. MySQL 8.0.30

# How to run
1. Clone this repository
2. Restore the database from `dump-nik_extractor-202309040857.sql` file
3. Change the database configuration in `main.go` file
4. Run `go install`
5. Run `go run main.go`

# How to use
Please see the documentation in [here](https://app.swaggerhub.com/apis-docs/zarszz/nik-extractor/1.0.0)
