module github.com/underopsco/go-action/examples/hello_world

go 1.21.0

toolchain go1.22.7

replace github.com/underopsco/go-action => ../../

require github.com/underopsco/go-action v0.0.0-20220215115724-031faebab42a

require (
	github.com/google/go-github/v66 v66.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
)
