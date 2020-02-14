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
	"flag"
	"log"
	"os"
	"testing"
)

var (
	nodeAddr     string
	auth         bool
	authUsername string
	authPassword string
	client       *Client
)

// Initial test setup. These tests require a Nimiq node to test, so flags require to be provided.
// If a node address is not provided, the tests that require a node will be skipped.
// Test cases are skipped on individual basis, so other tests that do not have this dependency
// can still take place.
func TestMain(m *testing.M) {
	flag.StringVar(&nodeAddr, "node-addr", "", "The address of the Nimiq RPC node that can be used to test this library")
	flag.BoolVar(&auth, "auth", false, "Use authentication to the Nimiq RPC node to test RPC calls")
	flag.StringVar(&authUsername, "username", "", "The username for authentication")
	flag.StringVar(&authPassword, "password", "", "The password for authentication")
	flag.Parse()

	client = NewClient(nodeAddr)
	if auth {
		if authPassword == "" || authUsername == "" {
			log.Println("Auth enabled but username or password is not provided")
			os.Exit(1)
		}
		client = NewClientWithAuth(nodeAddr, authUsername, authPassword)
	}

	os.Exit(m.Run())
}

func TestBatch(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.CallBatch: No node address provided")
		t.Skip("No node address provided")
	}

	resp, err := client.CallBatch(
		NewRequest("accounts"),
		NewRequest("hashrate"),
		NewRequest("blockNumber"),
	)

	if err != nil {
		log.Printf("FAILED: *client.CallBatch: %v", err)
		t.FailNow()
	}

	if resp.HasError() {
		log.Printf("FAILED: *client.CallBatch: one of the batch responses has an error")
		t.FailNow()
	}

	log.Println("SUCCES: *client.CallBatch")
}

// Test client.Accounts()
func TestClientAccounts(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.Accounts: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.Accounts()
	if err != nil {
		log.Printf("FAILED: *client.Accounts: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.Accounts")
}

// Test client.BlockNumber()
func TestClientBlockNumber(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.BlockNumber: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.Accounts()
	if err != nil {
		log.Printf("FAILED: *client.BlockNumber: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.BlockNumber")
}

// Test client.Consensus()
func TestClientConsensus(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.Consensus: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.Consensus()
	if err != nil {
		log.Printf("FAILED: *client.Consensus: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.Consensus")
}

// Test client.GetAccount()
func TestClientGetAccount(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetAccount: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetAccount("NQ52 V4BF 52J3 0PM6 BG4M 9QY1 RUYS UAL6 CJD2")
	if err != nil {
		log.Printf("FAILED: *client.GetAccount: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetAccount")
}

// Test client.GetBalance()
func TestClientGetBalance(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetBalance: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetBalance("NQ52 V4BF 52J3 0PM6 BG4M 9QY1 RUYS UAL6 CJD2")
	if err != nil {
		log.Printf("FAILED: *client.GetBalance: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetBalance")
}

// Test client.GetBlockByHash()
func TestClientGetBlockByHash(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetBlockByHash: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetBlockByHash("2f04512d8c80193f465aedf48160c6b5346f8c061a947c3f5734ad821e482065", false)
	if err != nil {
		log.Printf("FAILED: *client.GetBlockByHash: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetBlockByHash")
}

// Test client.GetBlockByNumber()
func TestClientGetBlockByNumber(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetBlockByNumber: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetBlockByNumber(684057, false)
	if err != nil {
		log.Printf("FAILED: *client.GetBlockByNumber: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetBlockByNumber")
}

// Test client.GetBlockTemplate()
func TestClientGetBlockTemplate(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetBlockTemplate: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetBlockTemplate()
	if err != nil {
		log.Printf("FAILED: *client.GetBlockTemplate: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetBlockTemplate")
}

// Test client.GetBlockTransactionCountByHash()
func TestClientGetBlockTransactionCountByHash(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetBlockTransactionCountByHash: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetBlockTransactionCountByHash("2f04512d8c80193f465aedf48160c6b5346f8c061a947c3f5734ad821e482065")
	if err != nil {
		log.Printf("FAILED: *client.GetBlockTransactionCountByHash: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetBlockTransactionCountByHash")
}

// Test client.GetBlockTransactionCountByNumber()
func TestClientGetBlockTransactionCountByNumber(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetBlockTransactionCountByNumber: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetBlockTransactionCountByNumber(684057)
	if err != nil {
		log.Printf("FAILED: *client.GetBlockTransactionCountByNumber: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetBlockTransactionCountByNumber")
}

// Test client.GetTransactionByBlockHashAndIndex
func TestClientGetTransactionByBlockHashAndIndex(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetTransactionByBlockHashAndIndex: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetTransactionByBlockHashAndIndex("2f04512d8c80193f465aedf48160c6b5346f8c061a947c3f5734ad821e482065", 0)
	if err != nil {
		log.Printf("FAILED: *client.GetTransactionByBlockHashAndIndex: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetTransactionByBlockHashAndIndex")
}

// Test client.GetTransactionByBlockNumberAndIndex
func TestClientGetTransactionByBlockNumberAndIndex(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetTransactionByBlockNumberAndIndex: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetTransactionByBlockNumberAndIndex(684057, 0)
	if err != nil {
		log.Printf("FAILED: *client.GetTransactionByBlockNumberAndIndex: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetTransactionByBlockNumberAndIndex")
}

// Test client.GetTransactionByHash
func TestClientGetTransactionByHash(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetTransactionByHash: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetTransactionByHash("ec8f1de17ce277af66d5a4e7c4d490313879115cd68f38284e79096b758aeb70")
	if err != nil {
		log.Printf("FAILED: *client.GetTransactionByHash: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetTransactionByHash")
}

// Test client.GetTransactionReceipt
func TestClientGetTransactionReceipt(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetTransactionReceipt: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetTransactionReceipt("ec8f1de17ce277af66d5a4e7c4d490313879115cd68f38284e79096b758aeb70")
	if err != nil {
		log.Printf("FAILED: *client.GetTransactionReceipt: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetTransactionReceipt")
}

// Test client.GetWork
func TestClientGetWork(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.GetWork: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.GetWork()
	if err != nil {
		log.Printf("FAILED: *client.GetWork: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.GetWork")
}

// Test client.Hashrate
func TestClientHashrate(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.Hashrate: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.Hashrate()
	if err != nil {
		log.Printf("FAILED: *client.Hashrate: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.Hashrate")
}

// Test client.Mempool
func TestClientMempool(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.Mempool: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.Mempool()
	if err != nil {
		log.Printf("FAILED: *client.Mempool: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.Mempool")
}

// Test client.MempoolContent
func TestClientMempoolContent(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.MempoolContent: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.MempoolContent(true)
	if err != nil {
		log.Printf("FAILED: *client.MempoolContent: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.MempoolContent")
}

// Test client.Mining
func TestClientMining(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.Mining: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.Mining()
	if err != nil {
		log.Printf("FAILED: *client.Mining: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.Mining")
}

// Test client.MinerAddress
func TestClientMinerAddress(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.MinerAddress: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.MinerAddress()
	if err != nil {
		log.Printf("FAILED: *client.MinerAddress: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.MinerAddress")
}

// Test client.MinerThreads
func TestClientMinerThreads(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.MinerThreads: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.MinerThreads()
	if err != nil {
		log.Printf("FAILED: *client.MinerThreadsg: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.MinerThreads")
}

// Test client.PeerCount
func TestClientPeerCount(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.PeerCount: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.PeerCount()
	if err != nil {
		log.Printf("FAILED: *client.PeerCount: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.PeerCount")
}

// Test client.PeerList
func TestClientPeerList(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.PeerList: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.PeerList()
	if err != nil {
		log.Printf("FAILED: *client.PeerList: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.PeerList")
}

// Test client.Pool
func TestClientPool(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.Pool: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.Pool()
	if err != nil {
		log.Printf("FAILED: *client.Pool: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.Pool")
}

// Test client.PoolConnectionState
func TestClientPoolConnectionState(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.PoolConnectionState: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.PoolConnectionState()
	if err != nil {
		log.Printf("FAILED: *client.PoolConnectionState: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.PoolConnectionState")
}

// Test client.PoolConfirmedBalance
func TestClientPoolConfirmedBalance(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.PoolConfirmedBalance: No node address provided")
		t.Skip("No node address provided")
	}

	_, err := client.PoolConfirmedBalance()
	if err != nil {
		log.Printf("FAILED: *client.PoolConfirmedBalance: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.PoolConfirmedBalance")
}

// Test client.Syncing
func TestClientSyncing(t *testing.T) {
	if nodeAddr == "" {
		log.Println("SKIPPED: *client.Syncing: No node address provided")
		t.Skip("No node address provided")
	}

	_, _, err := client.Syncing()
	if err != nil {
		log.Printf("FAILED: *client.Syncing: %v", err)
		t.FailNow()
	}
	log.Println("SUCCES: *client.Syncing")
}
