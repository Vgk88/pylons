package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

// Execution is a recipe execution used for tracking the execution - specifically a
// scheduled execution
type Execution struct {
	ID          string
	RecipeID    string // the recipe guid
	CoinInputs  sdk.Coins
	CoinOutputs sdk.Coins
	ItemInputs  []Item
	ItemOutputs []Item
	BlockHeight int64
	Sender      sdk.AccAddress
	Completed   bool
}

// KeyGen generates key for the execution
func (exec Execution) KeyGen() string {
	id := uuid.New()
	return exec.Sender.String() + id.String()
}
