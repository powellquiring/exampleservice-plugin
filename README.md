# cli-example-service-plugin

CLI Plugin for the Example Service. This code is automatically generated from an OpenAPI Definition.

## Testing
### Build the code
Assuming the IBM Cloud CLI is already installed:  
- `go build main.go`  
- `ibmcloud plugin install main`

### Authenticate
Currently, the CLI Plugin relies on external configuration for authentication. Either create a file called `ibm-credentials.env` that includes the authentication data or set environment variables.
Either way, the credentials need to have the format `<service-name>_<service-version>_APIKEY`.

e.g. `EXAMPLE_SERVICE_V1_APIKEY=abcd1234`

### Usage
- `ibmcloud my-plugin example-service <command-name> <flags>`

Example:
`ibmcloud my-plugin example-service list-resources`
