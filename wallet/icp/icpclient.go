package icp

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/ic"
	"github.com/aviate-labs/agent-go/principal"
)

// client for the "ledger" canister.
type LedgerClient struct {
	agentClient *agent.Agent
	canisterId  principal.Principal
}

// NewAgent creates a new agent for the "ledger" canister.
func NewLedgerClient(config agent.Config) (*LedgerClient, error) {
	agentClient, err := agent.New(config)
	if err != nil {
		return nil, err
	}
	return &LedgerClient{
		agentClient: agentClient,
		canisterId:  ic.LEDGER_PRINCIPAL,
	}, nil
}

// QueryBlocks calls the "query_blocks" method on the "ledger" canister.
func (ledgerClient LedgerClient) QueryBlocks(getBlockReq GetBlocksArgs) (*QueryBlocksResponse, error) {
	var resp QueryBlocksResponse
	if err := ledgerClient.agentClient.Query(
		ledgerClient.canisterId,
		"query_blocks",
		[]any{getBlockReq},
		[]any{&resp},
	); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountIdentifier calls the "account_identifier" method on the "ledger" canister.
func (ledgerClient LedgerClient) AccountIdentifier(accountReq Account) (*AccountIdentifier, error) {
	var accountResp AccountIdentifier
	if err := ledgerClient.agentClient.Query(
		ledgerClient.canisterId,
		"account_identifier",
		[]any{accountReq},
		[]any{&accountResp},
	); err != nil {
		return nil, err
	}
	return &accountResp, nil
}

// AccountBalance calls the "account_balance" method on the "ledger" canister.
func (ledgerClient LedgerClient) AccountBalance(accountBalance AccountBalanceArgs) (*Tokens, error) {
	var token Tokens
	if err := ledgerClient.agentClient.Query(
		ledgerClient.canisterId,
		"account_balance",
		[]any{accountBalance},
		[]any{&token},
	); err != nil {
		return nil, err
	}
	return &token, nil
}

// TransferFee calls the "transfer_fee" method on the "ledger" canister.
func (ledgerClient LedgerClient) TransferFee(transferFeeReq TransferFeeArg) (*TransferFee, error) {
	var transferFeeResp TransferFee
	if err := ledgerClient.agentClient.Query(
		ledgerClient.canisterId,
		"transfer_fee",
		[]any{transferFeeReq},
		[]any{&transferFeeResp},
	); err != nil {
		return nil, err
	}
	return &transferFeeResp, nil
}

// Transfer calls the "transfer" method on the "ledger" canister.
func (ledgerClient LedgerClient) Transfer(transferArg TransferArgs) (*TransferResult, error) {
	var transferResult TransferResult
	if err := ledgerClient.agentClient.Call(
		ledgerClient.canisterId,
		"transfer",
		[]any{transferArg},
		[]any{&transferResult},
	); err != nil {
		return nil, err
	}
	return &transferResult, nil
}

// SendDfx calls the "send_dfx" method on the "ledger" canister.
func (ledgerClient LedgerClient) SendDfx(arg0 SendArgs) (*BlockIndex, error) {
	var r0 BlockIndex
	if err := ledgerClient.agentClient.Call(
		ledgerClient.canisterId,
		"send_dfx",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}
