// Copyright 2020 Nimiq community.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*

Package nimiqrpc provides a Nimiq RPC client library in Go.

This client library implements the Nimiq RPC specification
Which can be found here: https://github.com/nimiq/core-js/wiki/JSON-RPC-API#remotejs-client
This package provides a Client and all the necessary types to interact with the Nimiq RPC server.

This client uses the jsonrpc library to handle JSON-RPC 2.0 requests and responses. For more information
about this library see the documentation https://godoc.org/github.com/ybbus/jsonrpc

How to use this package:

  // Initialise a new client
  nimiqClient := nimiqrpc.NewClient("address.to.nimiqnode.com")

  // Do an RPC call. For example retrieve the balance of a Nimiq account:
  balance, err := nimiqClient.GetBalance("NQ52 V4BF 52J3 0PM6 BG4M 9QY1 RUYS UAL6 CJD2")
  if err != nil {
      panic(err)
  }
  fmt.Printf("Balance: %v\n", balance)

*/
package nimiqrpc
