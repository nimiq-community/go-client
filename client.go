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

package nimiqrpc

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/ybbus/jsonrpc"
)

var (
	// ErrRespBodyEmpty is returned when the underlying HTTP response body is empty
	ErrRespBodyEmpty = errors.New("the HTTP response body was empty")

	// ErrResultUnexpected is returned when the expected result could not be read
	ErrResultUnexpected = errors.New("unexpected result")

	// ErrUnauthorized is returned when the user is not authorized to call the function
	ErrUnauthorized = errors.New("unauthorized")

	// ErrNotAuthenticated is returned when the user is required to be authenticated
	ErrNotAuthenticated = errors.New("not authenticated")
)

// Client contains a Nimiq RPC client
type Client struct {
	rpcClient jsonrpc.RPCClient
}

// NewClient returns a new Nimiq RPC client
func NewClient(address string) *Client {
	return &Client{
		rpcClient: jsonrpc.NewClient(address),
	}
}

// NewClientWithAuth returns a RPC client with the given username and password set for
// authentication
func NewClientWithAuth(address, username, password string) *Client {
	authHeader := make(map[string]string)
	authHeader["Authorization"] = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))
	return &Client{
		rpcClient: jsonrpc.NewClientWithOpts(address, &jsonrpc.RPCClientOpts{
			CustomHeaders: authHeader,
		}),
	}
}

// Call can be used to send a JSON-RPC request by setting the method and the parameters.
//
// This function is used internally to handle all the RPC functions provided by the client.
// Therefore this function should generally not be used. It does however provide the functionality to do
// RPC requests that are (not yet) implemented by this client.
//
// This function returns a *jsonrpc.RPCResponse. Please see the documentation for more information
// on how to unmarshall this RPCResponse. https://godoc.org/github.com/ybbus/jsonrpc#RPCResponse
func (nc *Client) Call(method string, params interface{}) (*jsonrpc.RPCResponse, error) {
	return nc.rpcClient.Call(method, params)
}

// CallBatch invokes a list of RPCRequests in a single batch request. This function is for more
// advanced use cases.
//
// Most convenient is to use the following form:
//
//   CallBatch(
//     NewRequest("hashrate"),
//     NewRequest("accounts"),
//   })
//
// Returns jsonrpc.RPCResponses that is of type []*jsonrpc.RPCResponse
// - note that a list of RPCResponses can be received unordered so it can happen that: responses[i] != responses[i].ID
// - RPCPersponses is enriched with helper functions e.g.: responses.HasError() returns  true if one of the responses holds an RPCError
// Please see the documenation on how to handle jsonrpc.RPCResonses: https://godoc.org/github.com/ybbus/jsonrpc#RPCResponses
func (nc *Client) CallBatch(reqs ...*jsonrpc.RPCRequest) (jsonrpc.RPCResponses, error) {
	return nc.rpcClient.CallBatch(jsonrpc.RPCRequests(reqs))
}

// NewRequest returns a *jsonrpc.RPCRequest that can be used as a parameter to the CallBatch function.
func NewRequest(method string, params ...interface{}) *jsonrpc.RPCRequest {
	return jsonrpc.NewRequest(method, params)
}
