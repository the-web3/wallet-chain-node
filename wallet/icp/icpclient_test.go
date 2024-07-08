package icp

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/ic"
	"github.com/aviate-labs/agent-go/principal"
	"log"
	"net/url"
	"testing"
)

var ic0, _ = url.Parse("https://ic0.app/")

func Test_icp_status(t *testing.T) {
	c := agent.NewClient(agent.ClientConfig{Host: ic0})
	status, _ := c.Status()
	fmt.Println(status.Version)
	fmt.Println(status)
}

func Test_icp_QueryBlocks_heignt(t *testing.T) {
	// 我们还不知道最后一个区块是什么，因此首先要查询区块高度。
	ledgerClient, _ := NewLedgerClient(agent.DefaultConfig)
	fmt.Println("ledgerClient ", ledgerClient)
	blockHeight, err := ledgerClient.QueryBlocks(GetBlocksArgs{})
	fmt.Printf("blockHeight: %+v\n", blockHeight)
	if err != nil {
		log.Fatal(err)
	}
	//我们可以查询分类账的第一个区块。
	oldestBlock := blockHeight.FirstBlockIndex
	fmt.Printf("oldestBlock: %+v\n", oldestBlock)
	//我们可以查询分类账的最后一个区块。
	lastBlock := blockHeight.ChainLength
	fmt.Printf("lastBlock: %+v\n", lastBlock)
}

func Test_icp_Get_AccountIdentifier(t *testing.T) {
	// 我们还不知道最后一个区块是什么，因此首先要查询区块高度。
	ledgerClient, _ := NewLedgerClient(agent.DefaultConfig)
	fmt.Println("ledgerClient ", ledgerClient)

	// 十六进制字符串表示的地址
	hexAddress := "57040de54254560c163c6aede27e3466acd5609f1d8f232c1ac66519b4cf78c3"
	// 解码十六进制字符串为字节数组
	addressBytes, err := hex.DecodeString(hexAddress)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}
	// 将字节数组转换为 SubAccount 类型
	subAccount := addressBytes

	account := Account{
		Owner:      ic.LEDGER_PRINCIPAL,
		Subaccount: &subAccount,
	}
	accountResp, _ := ledgerClient.AccountIdentifier(account)
	fmt.Printf("accountResp (hex): %x\n", *accountResp)
}

func Test_icp_Get_Account_Balance(t *testing.T) {
	// 我们还不知道最后一个区块是什么，因此首先要查询区块高度。
	ledgerClient, _ := NewLedgerClient(agent.DefaultConfig)
	fmt.Println("ledgerClient ", ledgerClient)

	// 十六进制字符串表示的地址
	hexAddress := "27bbe9b4f0b00e4b6fe3fb39328358cf82031e82014e0cd0ae60983cc92008f5"
	// 解码十六进制字符串为字节数组
	addressBytes, err := hex.DecodeString(hexAddress)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}
	fmt.Println("addressBytes ", addressBytes)
	fmt.Printf("token (string): %x\n", string(addressBytes))

	accountBalance := AccountBalanceArgs{
		Account: addressBytes,
	}
	token, _ := ledgerClient.AccountBalance(accountBalance)
	fmt.Printf("token %.2f ICP minted.\n", float64(token.E8s)/1e8)
}

func Test_icp_TransferFee(t *testing.T) {
	// 我们还不知道最后一个区块是什么，因此首先要查询区块高度。
	ledgerClient, _ := NewLedgerClient(agent.DefaultConfig)
	fmt.Println("ledgerClient ", ledgerClient)
	transferFee, _ := ledgerClient.TransferFee(TransferFeeArg{})
	fmt.Printf("transferFee %.8f ICP minted.\n", float64(transferFee.TransferFee.E8s)/1e8)
}

func Test_icp_Transfer(t *testing.T) {
	// 我们还不知道最后一个区块是什么，因此首先要查询区块高度。
	ledgerClient, _ := NewLedgerClient(agent.DefaultConfig)
	fmt.Println("ledgerClient ", ledgerClient)

	transferArgs := TransferArgs{}

	transferResult, err := ledgerClient.Transfer(transferArgs)
	if err != nil {
		fmt.Println("Error Test_icp_Transfer :", err)
		return
	}
	fmt.Printf("transferResult: %+v\n", transferResult)
}

func Test_icp_QueryBlocks(t *testing.T) {
	// 我们还不知道最后一个区块是什么，因此首先要查询区块高度。
	ledgerClient, _ := NewLedgerClient(agent.DefaultConfig)
	blockHeight, err := ledgerClient.QueryBlocks(GetBlocksArgs{})
	if err != nil {
		log.Fatal(err)
	}
	//我们可以查询分类账的最后一个区块。
	lastBlock := blockHeight.ChainLength

	// Query the last 10 blocks.
	response, err := ledgerClient.QueryBlocks(GetBlocksArgs{
		Start:  lastBlock - 10,
		Length: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, block := range response.Blocks {
		marshal, err := json.Marshal(block)
		if err != nil {
			return
		}
		println("block ", marshal)

	}
	println("=======================================================")
	println("=======================================================")
	println("=======================================================")
	println("=======================================================")
	println("=======================================================")
	println("=======================================================")
	println("=======================================================")
	println("=======================================================")

	for i, block := range response.Blocks {
		operation := block.Transaction.Operation
		if transfer := operation.Transfer; transfer != nil {
			var from principal.AccountIdentifier
			copy(from[:], transfer.From)

			var to principal.AccountIdentifier
			copy(to[:], transfer.To)

			fmt.Printf("Block %d: %s -> %s: %.2f ICP.\n", int(lastBlock)+i, from, to, float64(transfer.Amount.E8s)/1e8)
		} else if burn := operation.Burn; burn != nil {
			var from principal.AccountIdentifier
			copy(from[:], burn.From)

			fmt.Printf("Block %d: %s: %.2f ICP burned.\n", int(lastBlock)+i, from, float64(burn.Amount.E8s)/1e8)
		} else if mint := operation.Mint; mint != nil {
			var to principal.AccountIdentifier
			copy(to[:], mint.To)

			fmt.Printf("Block %d: %s: %.2f ICP minted.\n", int(lastBlock)+i, to, float64(mint.Amount.E8s)/1e8)
		}
	}
}
