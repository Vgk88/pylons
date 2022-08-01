package keeper

import (
	"cosmossdk.io/math"
	"github.com/Pylons-tech/pylons/x/pylons/types/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CoinIssuers returns the CoinIssuers param
func (k Keeper) CoinIssuers(ctx sdk.Context) (res []v1beta1.CoinIssuer) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyCoinIssuers, &res)
	return
}

// PaymentProcessors returns the CoinIssuers param
func (k Keeper) PaymentProcessors(ctx sdk.Context) (res []v1beta1.PaymentProcessor) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyPaymentProcessors, &res)
	return
}

// CoinIssuedDenomsList returns the CoinIssuedList param
func (k Keeper) CoinIssuedDenomsList(ctx sdk.Context) (res []string) {
	coinIssuers := k.CoinIssuers(ctx)
	for _, ci := range coinIssuers {
		res = append(res, ci.CoinDenom)
	}
	return
}

// RecipeFeePercentage returns the RecipeFeePercentage param
func (k Keeper) RecipeFeePercentage(ctx sdk.Context) (res sdk.Dec) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyRecipeFeePercentage, &res)
	return
}

// ItemTransferFeePercentage returns the CoinIssuedList param
func (k Keeper) ItemTransferFeePercentage(ctx sdk.Context) (res sdk.Dec) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyItemTransferFeePercentage, &res)
	return
}

// UpdateItemStringFee returns the UpdateItemStringFee param
func (k Keeper) UpdateItemStringFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyUpdateItemStringFee, &res)
	return
}

// UpdateUsernameFee returns the UpdateUsernameFee param
func (k Keeper) UpdateUsernameFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyUpdateUsernameFee, &res)
	return
}

// MinTransferFee returns the MinTransferFee param
func (k Keeper) MinTransferFee(ctx sdk.Context) (res math.Int) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyMinTransferFee, &res)
	return
}

// MaxTransferFee returns the MaxTransferFee param
func (k Keeper) MaxTransferFee(ctx sdk.Context) (res math.Int) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyMaxTransferFee, &res)
	return
}

// DistrEpochIdentifier returns the DistrEpochIdentifier param
func (k Keeper) DistrEpochIdentifier(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyDistrEpochIdentifier, &res)
	return
}

// EngineVersion returns the EngineVersion param
func (k Keeper) EngineVersion(ctx sdk.Context) (res uint64) {
	k.paramSpace.Get(ctx, v1beta1.ParamStoreKeyEngineVersion, &res)
	return
}

// GetParams returns the total set of pylons parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params v1beta1.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the pylons parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params v1beta1.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
