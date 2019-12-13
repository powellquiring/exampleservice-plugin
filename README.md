# Template for generated CLI plugins

This template provides the structure to generate code into, in order to have a fully-functioning CLI plugin.

## Getting Started
- Clone this repository
- Run the generator with the language "CLI" and the output path set to the filepath of this repository
- Create a Go module to manage dependencies:
`go mod init <package-name>`

Package name is configurable but should match the package name given for CLI plugin in the API Definition used to generate the repo.

- The IBM Cloud CLI SDK uses a version of `go-i18n` that is incompatible with Go Modules. Install a the following version:
`go get github.com/nicksnyder/go-i18n/i18n@v1.10.1`

- Install the dependencies and build the code:
`go build main.go`

- Install this locally built plugin to be used with the `ibmcloud` CLI tool:
`ibmcloud plugin install main`

If successful, your plugin is now ready for local use.

Note: If you are generating more than one service for a single plugin, the file `plugin/commands/root.go` should be hand-maintained as generating it for each service will overwrite the contents of the last.
