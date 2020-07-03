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
	"encoding/json"
	"fmt"
)

// Accounts returns a list of addresses owned by client.
func (nc *Client) Accounts() (accounts []Account, err error) {
	rpcResp, err := nc.Call("accounts", nil)
	if err != nil {
		return nil, err
	}

	err = rpcResp.GetObject(&accounts)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// BlockNumber returns the height of most recent block.
func (nc *Client) BlockNumber() (blockHeight int, err error) {
	rpcResp, err := nc.Call("blockNumber", nil)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&blockHeight)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// Consensus returns information on the current consensus state.
func (nc *Client) Consensus() (consensus string, err error) {
	rpcResp, err := nc.Call("consensus", nil)
	if err != nil {
		return "", err
	}

	err = rpcResp.GetObject(&consensus)
	if err != nil {
		return "", fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// CreateAccount creates a new account and stores its private key in the client store.
func (nc *Client) CreateAccount() (wallet *Wallet, err error) {
	rpcResp, err := nc.Call("createAccount", nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Wallet
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// CreateRawTransaction creates and signs a transaction without sending it.
// The transaction can then be send via sendRawTransaction without accidentally replaying it.
func (nc *Client) CreateRawTransaction(trn OutgoingTransaction) (transactionHex string, err error) {
	rpcResp, err := nc.Call("createRawTransaction", trn)
	if err != nil {
		return "", err
	}

	err = rpcResp.GetObject(&transactionHex)
	if err != nil {
		return "", err
	}

	return
}

// GetAccount returns details for the account of given address.
func (nc *Client) GetAccount(address string) (account *Account, err error) {
	rpcResp, err := nc.Call("getAccount", address)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Account
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// GetBalance returns the balance of the account of given address.
func (nc *Client) GetBalance(address string) (balance Luna, err error) {
	rpcResp, err := nc.Call("getBalance", address)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&balance)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// GetBlockByHash returns information about a block by block hash.
// If fullTransactions is true it returns a block with the full transaction objects,
// if false only the hashes of the transactions will be returned.
func (nc *Client) GetBlockByHash(blockHash string, fullTransactions bool) (block *Block, err error) {
	rpcResp, err := nc.Call("getBlockByHash", []interface{}{
		blockHash, fullTransactions,
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Block
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, ErrResultUnexpected
	}

	// Transaction result
	switch {
	case fullTransactions:
		err = json.Unmarshal(result.Transactions, &result.TransactionObjects)
	default:
		err = json.Unmarshal(result.Transactions, &result.TransactionHashes)
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// GetBlockByNumber returns information about a block by block number.
// If fullTransactions is true it returns a block with the full transaction objects,
// if false only the hashes of the transactions will be returned.
func (nc *Client) GetBlockByNumber(blockNumber int, fullTransactions bool) (block *Block, err error) {
	rpcResp, err := nc.Call("getBlockByNumber", []interface{}{
		blockNumber, fullTransactions,
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Block
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, ErrResultUnexpected
	}

	// Transaction result
	switch {
	case fullTransactions:
		err = json.Unmarshal(result.Transactions, &result.TransactionObjects)
	default:
		err = json.Unmarshal(result.Transactions, &result.TransactionHashes)
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// GetBlockTemplate returns a template to build the next block for mining.
// This will consider pool instructions when connected to a pool.
// Optional parameters: (1) The address to use as a miner for this block.
// This overrides the address provided during startup or from the pool.
// and (2)  Hex-encoded value for the extra data field. This overrides the address
// provided during startup or from the pool.
func (nc *Client) GetBlockTemplate(params ...interface{}) (template *BlockTemplate, err error) {
	rpcResp, err := nc.Call("getBlockTemplate", params)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result BlockTemplate
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// GetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (nc *Client) GetBlockTransactionCountByHash(blockHash string) (transactionCount int, err error) {
	rpcResp, err := nc.Call("getBlockTransactionCountByHash", blockHash)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&transactionCount)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// GetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block number.
func (nc *Client) GetBlockTransactionCountByNumber(blockNumber int) (transactionCount int, err error) {
	rpcResp, err := nc.Call("getBlockTransactionCountByNumber", blockNumber)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&transactionCount)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// GetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (nc *Client) GetTransactionByBlockHashAndIndex(blockHash string, index int) (transaction *Transaction, err error) {
	rpcResp, err := nc.Call("getTransactionByBlockHashAndIndex", []interface{}{
		blockHash, index,
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionByBlockNumberAndIndex returns information about a transaction by block hash and transaction index position.
func (nc *Client) GetTransactionByBlockNumberAndIndex(blockNumber int, index int) (transaction *Transaction, err error) {
	rpcResp, err := nc.Call("getTransactionByBlockNumberAndIndex", []interface{}{
		blockNumber, index,
	})
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionByHash Returns the information about a transaction requested by transaction hash.
func (nc *Client) GetTransactionByHash(transactionHash string) (transaction *Transaction, err error) {
	rpcResp, err := nc.Call("getTransactionByHash", transactionHash)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionReceipt returns the receipt of a transaction by transaction hash.
func (nc *Client) GetTransactionReceipt(transactionHash string) (transactionReceipt *TransactionReceipt, err error) {
	rpcResp, err := nc.Call("getTransactionReceipt", transactionHash)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result TransactionReceipt
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.TransactionHash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionsByAddress returns the latest transactions successfully performed by or for an address.
// The array will not contain more than maxEntries, but might contain less, even when more transactions happened.
// Any interpretation of the length of this array might result in worng assumptions.
func (nc *Client) GetTransactionsByAddress(address string, maxEntries int) (transactions []Transaction, err error) {
	rpcResp, err := nc.Call("getTransactionsByAddress", []interface{}{
		address, maxEntries,
	})
	if err != nil {
		return nil, err
	}

	err = rpcResp.GetObject(&transactions)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// GetWork returns instructions to mine the next block. This will consider pool instructions when connected to a pool.
// Optional arameters: (1)  The address to use as a miner for this block. This overrides the address provided during startup or from the pool.
// and (2) Hex-encoded value for the extra data field. This overrides the address provided during startup or from the pool
func (nc *Client) GetWork(params ...interface{}) (work *Work, err error) {
	rpcResp, err := nc.Call("getWork", params)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Work
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Data == "" {
		return nil, nil
	}

	return &result, nil
}

// Hashrate returns the number of hashes per second that the node is mining with.
func (nc *Client) Hashrate() (hashrate float64, err error) {
	rpcResp, err := nc.Call("hashrate", nil)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&hashrate)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// Log sets the log level of the node.
func (nc *Client) Log(tag string, level LogLevel) (succes bool, err error) {
	rpcResp, err := nc.Call("log", []interface{}{
		tag, level,
	})
	if err != nil {
		return false, err
	}

	err = rpcResp.GetObject(&succes)
	if err != nil {
		return false, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// Mempool Returns information on the current mempool situation.
// This will provide an overview of the number of transactions sorted into buckets
// based on their fee per byte (in smallest unit).
func (nc *Client) Mempool() (mempool *Mempool, err error) {
	rpcResp, err := nc.Call("mempool", nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Mempool
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Total == 0 && len(result.Buckets) == 0 {
		return nil, nil
	}

	return &result, nil
}

// MempoolContent Returns all transactions currently in the mempool.
// If fullTransactions is true it returns a block with the full transaction objects,
// if false only the hashes of the transactions will be returned.
func (nc *Client) MempoolContent(fullTransactions bool) (transactions interface{}, err error) {
	rpcResp, err := nc.Call("mempoolContent", fullTransactions)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	switch {
	case fullTransactions:
		var result []Transaction
		err = rpcResp.GetObject(&result)
		if err != nil {
			return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
		}

		return result, nil
	default:
		var result []string
		err = rpcResp.GetObject(&result)
		if err != nil {
			return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
		}

		return result, nil
	}
}

// MinFeePerByte returns the minimum fee per byte
// Optionally newFee can be given as a parameter, which sets the minimum fee per byte
// to the value of newFee
func (nc *Client) MinFeePerByte(newFee ...int64) (fee int64, err error) {
	var params []interface{}
	if len(newFee) > 0 {
		params = append(params, newFee[0])
	}

	rpcResp, err := nc.Call("minFeePerByte", params)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&fee)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// Mining returns if client is actively mining new blocks.
// Optionally state can be given as a parameter which enables or disables mining.
func (nc *Client) Mining(state ...bool) (status bool, err error) {
	var params []interface{}
	if len(state) > 0 {
		params = append(params, state[0])
	}

	rpcResp, err := nc.Call("mining", params)
	if err != nil {
		return false, err
	}

	err = rpcResp.GetObject(&status)
	if err != nil {
		return false, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// MinerAddress returns the user friendly miner address.
func (nc *Client) MinerAddress() (address string, err error) {
	rpcResp, err := nc.Call("minerAddress", nil)
	if err != nil {
		return "", err
	}

	err = rpcResp.GetObject(&address)
	if err != nil {
		return "", fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// MinerThreads returns the number of CPU threads the miner is using
// Optionally number can be given as a parameter which sets the threads of the miner
// to the given number.
func (nc *Client) MinerThreads(number ...int) (threads int, err error) {
	var params []interface{}
	if len(number) > 0 {
		params = append(params, number[0])
	}

	rpcResp, err := nc.Call("minerThreads", params)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&threads)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// PeerCount returns number of peers currently connected to the client.
func (nc *Client) PeerCount() (peers int, err error) {
	rpcResp, err := nc.Call("peerCount", nil)
	if err != nil {
		return 0, err
	}

	err = rpcResp.GetObject(&peers)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// PeerList returns a list of peers currently connected to the client
func (nc *Client) PeerList() (peers []Peer, err error) {
	rpcResp, err := nc.Call("peerList", nil)
	if err != nil {
		return []Peer{}, err
	}

	err = rpcResp.GetObject(&peers)
	if err != nil {
		return []Peer{}, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// PeerState returns the state for the given peer address. If update is set, the state
// of the peer will be set to the given update parameter
func (nc *Client) PeerState(peerAddress string, update ...string) (peer *Peer, err error) {
	params := []interface{}{peerAddress}
	if len(update) > 0 {
		if update[0] == "ban" || update[0] == "unban" || update[0] == "connect" || update[0] == "disconnect" {
			params = append(params, update[0])
		}
	}
	rpcResp, err := nc.Call("peerState", params)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Peer
	err = rpcResp.GetObject(&result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// Pool returns the current pool address.
// Optionally newAddress can be given as a parameter, which sets the address to the
// value of newAddress.
func (nc *Client) Pool(newAddress ...string) (address string, err error) {
	var params []interface{}
	if len(newAddress) > 0 {
		params = append(params, newAddress[0])
	}
	rpcResp, err := nc.Call("pool", params)
	if err != nil {
		return "", err
	}

	// Unmarshal result
	err = rpcResp.GetObject(&address)
	if err != nil {
		return "", fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// PoolConnectionState returns the pool connection state.
// Possible connection states: 0 - connected, 1 - connecting, 2 - closed
func (nc *Client) PoolConnectionState() (state int, err error) {
	rpcResp, err := nc.Call("poolConnectionState", nil)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	err = rpcResp.GetObject(&state)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// PoolConfirmedBalance returns the miner balance confirmed by the pool.
func (nc *Client) PoolConfirmedBalance() (balance Luna, err error) {
	rpcResp, err := nc.Call("poolConfirmedBalance", nil)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	err = rpcResp.GetObject(&balance)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// SendRawTransaction sends a signed message call transaction or a contract creation, if the data field contains code.
func (nc *Client) SendRawTransaction(signedTransaction string) (transactionHash string, err error) {
	rpcResp, err := nc.Call("sendRawTransaction", signedTransaction)
	if err != nil {
		return "", err
	}

	err = rpcResp.GetObject(&transactionHash)
	if err != nil {
		return "", fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return
}

// SendTransaction creates new message call transaction or a contract creation, if the data field contains code.
func (nc *Client) SendTransaction(trn OutgoingTransaction) (transactionHash string, err error) {
	rpcResp, err := nc.Call("sendTransaction", trn)
	if err != nil {
		return "", err
	}

	err = rpcResp.GetObject(&transactionHash)
	if err != nil {
		return "", fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}
	return
}

// SubmitBlock submits a block to the node. When the block is valid, the node will forward it to other nodes in the network.
// The argument is a hex-encoded full block (including header, interlink and body).
// When submitting work from getWork, remember to include the suffix.
func (nc *Client) SubmitBlock(fullBlock string) (err error) {
	_, err = nc.Call("submitBlock", fullBlock)
	return
}

// Syncing returns whether the node is syncing and when it is syncing, data about the sync status.
func (nc *Client) Syncing() (syncing bool, syncStatus *SyncStatus, err error) {
	rpcResp, err := nc.Call("syncing", nil)
	if err != nil {
		return false, nil, err
	}

	// Unmarshal result
	var result SyncStatus
	err = rpcResp.GetObject(&result)
	if err != nil {
		var boolResult bool
		err = rpcResp.GetObject(&boolResult)
		if err != nil {
			return false, nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
		}

		return false, nil, nil
	}

	return true, &result, nil
}
