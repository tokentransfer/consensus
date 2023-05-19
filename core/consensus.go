package core

import (
	"github.com/tokentransfer/interfaces/block"
	"github.com/tokentransfer/interfaces/core"
)

type ConsensusService interface {
	VerifyTransaction(rootAccount core.Address, tx block.Transaction) (bool, error)
	ProcessTransaction(rootAccount core.Address, tx block.Transaction) (block.TransactionResult, []block.State, error)
}
