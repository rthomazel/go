module github.com/rthomazel/go/httpmisc

go 1.26.1

require (
	github.com/rthomazel/go/logging v0.3.0
	github.com/rthomazel/go/misc v0.1.5
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rthomazel/go/hue v0.2.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/rthomazel/go/hue => ../hue
	github.com/rthomazel/go/logging => ../logging
	github.com/rthomazel/go/misc => ../misc
)
