package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionsHash(t *testing.T) {
	executeTxHashHex := "0xb382f51137675c47e5af5e22025204b38dfb90ecba32be3efe6e9971f66ce877"
	salt, _ := new(big.Int).SetString("78875051745042265906378644240522153635671489603671628324097685565218866261548", 10)
	signerAddress := common.HexToAddress("0xdf6aea8ec6fe69d8ddfc1a47649dd875dc600fda")
	data := common.FromHex("0x64a3bc15000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000038d7ea4c6800000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000051d93f709ea363f4539b2b92ac4c0dae22d60232000000000000000000000000df6aea8ec6fe69d8ddfc1a47649dd875dc600fda00000000000000000000000000000000000000000000000000000000000000000000000000000000000000003a1306efbb246f079f8f2528b3efc788b90f7a130000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000038d7ea4c6800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005bf380bd7c865235883bfe2e4c950cdc6875d1484fae66aa149a77f27c797d3aab7e7ab6000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000024f47261b0000000000000000000000000a7c8cd567a2321ed6c31395de25bcbcc4ebfce17000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024f47261b0000000000000000000000000b5e7d9c04bbef1108a158856cbc823ea7bc00c660000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000561c55e1476a08be7c6b6ec9712ae76c0a50f798ada63c3869976543e963532bb4f15f2a8c7abc643a0476c6f66732ce76313a1b75d6864f7ae9c1698295226fcff1a06b79e655db7d7c3b3e7b2cceeb068c3259d0c90400000000000000000000")

	transactionHash := GetTransactionHash(salt, signerAddress, data, common.HexToAddress("0x35dd2932454449b14cee11a94d3674a936d5d7b2"))
	assert.Equal(t, common.ToHex(transactionHash), executeTxHashHex)
}
