# Nimiq Go RPC Client
[![Report card](https://goreportcard.com/badge/github.com/nimiq-community/go-client)](https://goreportcard.com/report/github.com/nimiq-community/go-client)
[![GoDoc](https://godoc.org/github.com/nimiq-community/go-client?status.svg)](https://godoc.org/github.com/nimiq-community/go-client)


> Go implementation of the Nimiq RPC client specs.

## About
A Nimiq RPC client library in Go. This client library implements the [Nimiq RPC specification](https://github.com/nimiq/core-js/wiki/JSON-RPC-API#remotejs-client). This client uses the jsonrpc library to handle JSON-RPC 2.0 requests and responses. For more information about this library see the [jsonrpc documentation](https://godoc.org/github.com/ybbus/jsonrpc)

### What is Nimiq?
A blockchain technology inspired by Bitcoin but designed to run in your browser. It is money by nature but capable to do much more.
See [the Nimiq website](https://nimiq.com) for more information.

### Usage
This library is fully Go module compatible. See the full documentation on [GoDoc](https://godoc.org/github.com/nimiq-community/go-Client)
An example of using this library can be found below:

```
package main

import (
  fmt

  nimiqrpc "github.com/nimiq-community/go-client"
)

func main() {
  // Initialise a new client
  nimiqClient := nimiqrpc.NewClient("address.to.nimiqnode.com")

  // Do an RPC call. For example retrieve the balance of a Nimiq account:
  balance, err := nimiqClient.GetBalance("NQ52 V4BF 52J3 0PM6 BG4M 9QY1 RUYS UAL6 CJD2")
  if err != nil {
      panic(err)
  }
  fmt.Printf("Balance: %v\n", balance)
}
```

## Testing
This library provides several tests to guarantee API consistency. Several tests require an RPC server to test the RPC requests. In order to run theses tests, arguments need to be provided to the go test command. The following arguments are supported:

* <b>node-addr</b> <i>[address]</i> Sets the address of the Nimiq RPC server
* <b>auth</b> Enables the use of authentication to the RPC server
* <b>username</b> <i>[username]</i> Sets the username of the client
* <b>password</b> <i>[password]</i> Sets the password of the client

For example:
```
go test --cover --node-addr "http://seed.nimiq.example:8443" --auth --username "username"  --password "this is an example password: the higher the entropy the better the password"
```

## Contributions

This implementation was originally contributed by [redmaner](https://github.com/redmaner/).

Please send your contributions as pull requests.
Refer to the [issue tracker](issues) for ideas.

## License

[Apache 2.0](LICENSE)
