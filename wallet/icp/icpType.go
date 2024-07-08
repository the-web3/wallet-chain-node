package icp

import (
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

type BlockIndex = uint64

type GetBlocksArgs = struct {
	Start  BlockIndex `ic:"start"`
	Length uint64     `ic:"length"`
}

type QueryBlocksResponse = struct {
	ChainLength     uint64                `ic:"chain_length"`
	Certificate     *[]byte               `ic:"certificate,omitempty"`
	Blocks          []Block               `ic:"blocks"`
	FirstBlockIndex BlockIndex            `ic:"first_block_index"`
	ArchivedBlocks  []ArchivedBlocksRange `ic:"archived_blocks"`
}

type Block = struct {
	ParentHash  *[]byte     `ic:"parent_hash,omitempty"`
	Transaction Transaction `ic:"transaction"`
	Timestamp   TimeStamp   `ic:"timestamp"`
}

type Transaction = struct {
	Memo          Memo       `ic:"memo"`
	Icrc1Memo     *[]byte    `ic:"icrc1_memo,omitempty"`
	Operation     *Operation `ic:"operation,omitempty"`
	CreatedAtTime TimeStamp  `ic:"created_at_time"`
}

type Memo = uint64

type TimeStamp = struct {
	TimestampNanos uint64 `ic:"timestamp_nanos"`
}

type Operation = struct {
	Mint *struct {
		To     AccountIdentifier `ic:"to"`
		Amount Tokens            `ic:"amount"`
	} `ic:"Mint,variant"`
	Burn *struct {
		From    AccountIdentifier  `ic:"from"`
		Spender *AccountIdentifier `ic:"spender,omitempty"`
		Amount  Tokens             `ic:"amount"`
	} `ic:"Burn,variant"`
	Transfer *struct {
		From    AccountIdentifier `ic:"from"`
		To      AccountIdentifier `ic:"to"`
		Amount  Tokens            `ic:"amount"`
		Fee     Tokens            `ic:"fee"`
		Spender *[]uint8          `ic:"spender,omitempty"`
	} `ic:"Transfer,variant"`
	Approve *struct {
		From              AccountIdentifier `ic:"from"`
		Spender           AccountIdentifier `ic:"spender"`
		AllowanceE8s      idl.Int           `ic:"allowance_e8s"`
		Allowance         Tokens            `ic:"allowance"`
		Fee               Tokens            `ic:"fee"`
		ExpiresAt         *TimeStamp        `ic:"expires_at,omitempty"`
		ExpectedAllowance *Tokens           `ic:"expected_allowance,omitempty"`
	} `ic:"Approve,variant"`
}

type Tokens = struct {
	E8s uint64 `ic:"e8s"`
}

type AccountIdentifier = []byte

type ArchivedBlocksRange = struct {
	Start    BlockIndex     `ic:"start"`
	Length   uint64         `ic:"length"`
	Callback QueryArchiveFn `ic:"callback"`
}
type QueryArchiveFn = struct { /* NOT SUPPORTED */
}

type Account = struct {
	Owner      principal.Principal `ic:"owner"`
	Subaccount *SubAccount         `ic:"subaccount,omitempty"`
}

type SubAccount = []byte

type AccountBalanceArgs = struct {
	Account AccountIdentifier `ic:"account"`
}

type TransferFeeArg = struct {
}

type TransferFee = struct {
	TransferFee Tokens `ic:"transfer_fee"`
}

type SendArgs = struct {
	Memo           Memo                  `ic:"memo"`
	Amount         Tokens                `ic:"amount"`
	Fee            Tokens                `ic:"fee"`
	FromSubaccount *SubAccount           `ic:"from_subaccount,omitempty"`
	To             TextAccountIdentifier `ic:"to"`
	CreatedAtTime  *TimeStamp            `ic:"created_at_time,omitempty"`
}

type TextAccountIdentifier = string

type TransferArgs = struct {
	Memo           Memo              `ic:"memo"`
	Amount         Tokens            `ic:"amount"`
	Fee            Tokens            `ic:"fee"`
	FromSubaccount *SubAccount       `ic:"from_subaccount,omitempty"`
	To             AccountIdentifier `ic:"to"`
	CreatedAtTime  *TimeStamp        `ic:"created_at_time,omitempty"`
}

type TransferResult = struct {
	Ok  *BlockIndex    `ic:"Ok,variant"`
	Err *TransferError `ic:"Err,variant"`
}

type TransferError = struct {
	BadFee *struct {
		ExpectedFee Tokens `ic:"expected_fee"`
	} `ic:"BadFee,variant"`
	InsufficientFunds *struct {
		Balance Tokens `ic:"balance"`
	} `ic:"InsufficientFunds,variant"`
	TxTooOld *struct {
		AllowedWindowNanos uint64 `ic:"allowed_window_nanos"`
	} `ic:"TxTooOld,variant"`
	TxCreatedInFuture *struct{} `ic:"TxCreatedInFuture,variant"`
	TxDuplicate       *struct {
		DuplicateOf BlockIndex `ic:"duplicate_of"`
	} `ic:"TxDuplicate,variant"`
}
