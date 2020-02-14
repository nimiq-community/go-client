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
	"strconv"
	"strings"
)

// Available LogLevels
const (
	LogLevelTrace   LogLevel = "trace"
	LogLevelVerbose LogLevel = "verbose"
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarn    LogLevel = "warn"
	LogLevelError   LogLevel = "error"
	LogLevelAssert  LogLevel = "assert"
)

// Account types on the blockchain
const (
	AccountTypeBasic   = 0
	AccountTypeVesting = 1
	AccountTypeHTLC    = 2
)

// LogLevel is the level of logging that is enabled on a node
type LogLevel string

// NIM is the token transacted within Nimiq as a store and transfer of value: it acts as digital cash
type NIM string

// FormatNIM is a function to format Luna to NIM
func FormatNIM(l Luna) NIM {
	nimString := strconv.Itoa(int(l))
	for i := len(nimString); i < 5; i++ {
		nimString = "0" + nimString
	}

	switch {
	case len(nimString) == 5:
		nimString = "0." + nimString
	case len(nimString) > 5:
		nimString = nimString[0:len(nimString)-5] + "." + nimString[len(nimString)-5:]
	}

	if nimString[len(nimString)-5:] == "00000" {
		return NIM(nimString[:len(nimString)-6])
	}
	return NIM(nimString)
}

// ToLuna converts NIM to Luna
func (n *NIM) ToLuna() (Luna, error) {
	return FormatLuna(*n)
}

// Luna is the smallest unit of NIM and 100’000 (1e5) Luna equals 1 NIM
type Luna int64

// FormatLuna is a function to format NIM to Luna
func FormatLuna(n NIM) (Luna, error) {
	nimString := string(n)
	dotIndex := strings.Index(nimString, ".")
	switch {
	case dotIndex == -1:
		nimString = nimString + "00000"
	default:
		nimString = nimString[:dotIndex] + nimString[dotIndex+1:]
	}

	luna, err := strconv.Atoi(nimString)
	if err != nil {
		return 0, err
	}

	return Luna(luna), nil
}

// ToNIM converts Luna to NIM
func (l *Luna) ToNIM() NIM {
	return FormatNIM(*l)
}

// Account holds the details on an account
type Account struct {
	ID      string `json:"id"`      // hex-encoded address bytes
	Address string `json:"address"` // user friendly address (NQ-address).
	Balance Luna   `json:"balance"` // Balance of the account in Luna
	Type    int    `json:"type"`    // see AccountType const block

	// Additional fields for AccountTypeVesting
	Owner              string `json:"owner,omitempty"`              // hex-encoded address of contract owner
	OwnerAddress       string `json:"ownerAddress,omitempty"`       // user friendly address of contract owner
	VestingStart       int    `json:"vestingStart,omitempty"`       // the block that the vesting contracted commenced
	VestingStepBlocks  int    `json:"vestingStepBlocks,omitempty"`  // no. of blocks after which some part of the vested funds is released
	VestingStepAmount  int    `json:"vestingStepAmount,omitempty"`  // amount in Luna released every VestingStepBlocks blocks
	VestingTotalAmount int    `json:"vestingTotalAmount,omitempty"` // total amount in Luna that was provided at the contract creation

	// Additional fields for AccountTypeHTLC
	Sender           string `json:"sender,omitempty"`           // hex-encoded address of HTLC sender
	SenderAddress    string `json:"senderAddress,omitempty"`    // user friendly address of HTLC sender
	Recipient        string `json:"recipient,omitempty"`        // hex-encoded address of HTLC recipient
	RecipientAddress string `json:"recipientAddress,omitempty"` // user friendly address of HTLC recipient
	HashRoot         string `json:"hashRoot,omitempty"`         // hex-encoded 32 byte hash root
	HashCount        int    `json:"hashCount,omitempty"`        // no. of hashes this HTLC is split into
	Timeout          int    `json:"timeout,omitempty"`          // block at which the HTLC times out
	TotalAmount      int    `json:"totalAmount,omitempty"`      // total amount in Luna provided at contract creation
}

// AddressObject holds the representation of a Nimiq address in two formats.
type AddressObject struct {
	ID      string `json:"id"`      // hex-encoded 20 byte address
	Address string `json:"address"` // user friendly address (NQ-address)
}

// Block holds the details on a block
type Block struct {
	Number       int         `json:"number"`       // height of the block
	Hash         string      `json:"hash"`         // block hash
	POW          string      `json:"pow"`          // proof-of-work hash
	ParentHash   string      `json:"parentHash"`   // hash of the predecessor block
	Nonce        int         `json:"nonce"`        // nonce of the block used to fulfill the proof-of-work
	BodyHash     string      `json:"bodyHash"`     // hash of the block body Merkle root
	AccountHash  string      `json:"accountHash"`  // hash of the accounts tree root
	Miner        string      `json:"miner"`        // hex-encoded address of the miner
	MinerAddress string      `json:"minerAddress"` // user friendly address of the miner
	Difficulty   json.Number `json:"difficulty"`   // block difficulty
	ExtraData    string      `json:"extraData"`    // hex-encoded value of the extra data field
	Size         int         `json:"size"`         // block size in bytes
	Timestamp    int         `json:"timestamp"`    // UNIX timestamp of the block

	// Transactions contains either an array of transaction hashes or transaction objects,
	// depending on request parameters.
	// Depending on the result, the TransactionHashes or TransactionObjects fields are set.
	Transactions       json.RawMessage `json:"transactions"`
	TransactionHashes  []string        // hashes of the transactions in the block
	TransactionObjects []Transaction   // detailed transaction objects
}

// BlockTemplate contains details on a block template
type BlockTemplate struct {
	Header    BlockTemplateHeader `json:"header"`
	Interlink string              `json:"interlink"` // hex-encoded interlink
	Body      BlockTemplateBody   `json:"body"`
	Target    int                 `json:"target"` // compact form of hash target to submit a block
}

// BlockTemplateHeader describes the block header of a template
type BlockTemplateHeader struct {
	Version       int    `json:"version"`       // block version
	PrevHash      string `json:"prevHash"`      // hash of the predecessor block
	InterlinkHash string `json:"interlinkHash"` // hash of the interlink
	AccountHash   string `json:"accountHash"`   // hash of the accounts tree root
	NBits         int    `json:"nBits"`         // compact form of hash target for this block
	Height        int    `json:"height"`        // block number
}

// BlockTemplateBody describes the block body of a template
type BlockTemplateBody struct {
	Hash           string   `json:"hash"`           // hash of the block body
	MinerAddr      string   `json:"minerAddr"`      // hex-encoded miner address
	ExtraData      string   `json:"extraData"`      // hex-encoded value of the extra data field
	Transactions   []string `json:"transactions"`   // hex-encoded transactions for this block
	PrunedAccounts []string `json:"prunedAccounts"` // hex-encoded pruned accounts for this block
	// MerkleHashes contains hex-encoded hashes that verify the path of the miner
	// address in the merkle tree. This can be used to change the miner address easily.
	MerkleHashes []string `json:"merkleHashes"`
}

// Mempool holds the details on a mempool.
type Mempool struct {
	// Total number of pending transactions in mempool.
	Total int `json:"total,omitempty"`
	// Buckets is the subset of fee per byte buckets from
	// [10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1, 0]
	// that currently have more than one transaction.
	Buckets []int `json:"buckets,omitempty"`
	// any of the numbers present in buckets: Integer - Number of transaction in the bucket.
	// A transaction is assigned to the highest bucket of a value lower than its fee per byte value.
	Bucket0    int `json:"0,omitempty"`
	Bucket1    int `json:"1,omitempty"`
	Bucket2    int `json:"2,omitempty"`
	Bucket5    int `json:"5,omitempty"`
	Bucket10   int `json:"10,omitempty"`
	Bucket20   int `json:"20,omitempty"`
	Bucket50   int `json:"50,omitempty"`
	Bucket100  int `json:"100,omitempty"`
	Bucket200  int `json:"200,omitempty"`
	Bucket500  int `json:"500,omitempty"`
	Bucket1000 int `json:"1000,omitempty"`
	Bucket2000 int `json:"2000,omitempty"`
	Bucket5000 int `json:"5000,omitempty"`
}

// Peer holds the details of a peer
type Peer struct {
	ID              string `json:"id,omitempty"`
	Address         string `json:"address,omitempty"`
	AddressState    int    `json:"addressState,omitempty"`
	ConnectionState int    `json:"connectionState,omitempty"`
	Version         int    `json:"version,omitempty"`
	TimeOffset      int    `json:"timeOffset,omitempty"`
	HeadHash        string `json:"headHash,omitempty"`
	Latency         int    `json:"latency,omitempty"`
	RX              int    `json:"rx,omitempty"`
	TX              int    `json:"tx,omitempty"`
}

// Transaction holds the details on a transaction
type Transaction struct {
	Hash             string `json:"hash"`                       // hash of transaction
	BlockHash        string `json:"blockHash"`                  // hash of containing block
	BlockNumber      int    `json:"blockNumber,omitempty"`      // number of containing block
	Timestamp        int    `json:"timestamp,omitempty"`        // UNIX timestamp of containing block
	Confirmations    int    `json:"confirmations"`              // number of blocks since transaction was mined (0 if not mined)
	TransactionIndex int    `json:"transactionIndex,omitempty"` // index of transaction within block

	From        string `json:"from"`                // hex-encoded address of sending account
	FromAddress string `json:"fromAddress"`         // user friendly address of sending account
	To          string `json:"to"`                  // hex-encoded address of recipient account
	ToAddress   string `json:"toAddress,omitempty"` // user friendly address of recipient account

	Value Luna   `json:"value"`
	Fee   Luna   `json:"fee"`
	Data  string `json:"data"`  // hex-encoded contract parameters or a message
	Flags int    `json:"flags"` // bit-encoded transaction flags
}

// TransactionReceipt holds the details on a transaction receipt.
type TransactionReceipt struct {
	TransactionHash  string `json:"transactionHash"`     // hash of transaction
	TransactionIndex int    `json:"transactionIndex"`    // transaction index within block
	BlockHash        string `json:"blockHash"`           // hash of containing block
	BlockNumber      int    `json:"blockNumber"`         // number of containing block
	Confirmations    int    `json:"confirmations"`       // number of blocks since transaction was mined
	Timestamp        int    `json:"timestamp,omitempty"` // timestamp of containing block
}

// OutgoingTransaction holds the details on a transaction that is not yet sent.
type OutgoingTransaction struct {
	From     string `json:"from"`               // address of sending account
	FromType int    `json:"fromType,omitempty"` // AccountType of sending address (default AccountTypeBasic)
	To       string `json:"to"`                 // address of recipient account
	ToType   int    `json:"toType,omitempty"`   // AccountType of recipient address (default AccountTypeBasic)

	Value Luna   `json:"value"`
	Fee   Luna   `json:"fee"`
	Data  string `json:"data,omitempty"` // hex-encoded contract parameters or a message
}

// SyncStatus holds information about the sync status.
type SyncStatus struct {
	StartingBlock int `json:"startingBlock"` // block at which import started
	CurrentBlock  int `json:"currentBlock"`
	HighestBlock  int `json:"highestBlock"` // estimated
}

// Wallet holds the details on a wallet.
type Wallet struct {
	ID         string `json:"id"`                   // hex-encoded 20 byte address
	Address    string `json:"address"`              // user friendly address (NQ-address)
	PublicKey  string `json:"publicKey"`            // hex-encoded Ed25519 public key
	PrivateKey string `json:"privateKey,omitempty"` // hex-encoded Ed25519 private key (optional)
}

// Work holds the instructions to mine the next block
// For instructions on how to use, see https://github.com/nimiq/core-js/wiki/JSON-RPC-API#getwork
type Work struct {
	Data      string `json:"data"`      // hex-encoded block header
	Suffix    string `json:"suffix"`    // block without header
	Target    int    `json:"target"`    // compact form of hash target
	Algorithm string `json:"algorithm"` // hash algorithm. always "nimiq-argon2" for now
}
