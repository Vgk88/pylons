package handlers

import (
	"encoding/json"
	"testing"

	"github.com/Pylons-tech/pylons/x/pylons/keep"
	"github.com/Pylons-tech/pylons/x/pylons/msgs"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TestCoinLock is coin lock test
func TestCoinLock(t *testing.T) {
	tci := keep.SetupTestCoinInput()

	lotsOfCoins := sdk.Coins{
		sdk.NewInt64Coin("chair", 100000),
		sdk.NewInt64Coin(types.Pylon, 100000),
	}
	sender1, _, _, _ := keep.SetupTestAccounts(
		t,
		tci,
		lotsOfCoins,
		nil, nil, nil,
	)

	cbData := CreateCookbookResponse{}

	cookbookMsg := msgs.NewMsgCreateCookbook(
		"cookbook-0001",
		"",
		"this has to meet character limits",
		"SketchyCo",
		"1.0.0",
		"example@example.com",
		1,
		msgs.DefaultCostPerBlock,
		sender1,
	)

	cookbookResult, _ := HandlerMsgCreateCookbook(tci.Ctx, tci.PlnK, cookbookMsg)
	err := json.Unmarshal(cookbookResult.Data, &cbData)
	require.True(t, err == nil)
	require.True(t, len(cbData.CookbookID) > 0)

	cases := map[string]struct {
		// Balance of two accounts for test
		coinsAccount1 sdk.Coins
		coinsAccount2 sdk.Coins
		// Create Trade Coin Lock Test
		testCreateTradeLock bool
		createTradeAmount   sdk.Coins
		lockDiffCreateTrade sdk.Coins
		// Execute Receipe Coin Lock Test
		testExecuteRecipeLock  bool
		executeRecipeCoinInput types.CoinInputList
		lockDiffSchedule       sdk.Coins
		// Fulfill Trade Coin Unlock Test
		testFulfillTrade         bool
		fulfillTradeInputItemIDs []string
		lockDiffFulfillTrade     sdk.Coins
		// Check Execution Coin Unlock Test
		testCheckExecution bool
		lockDiffCheckExec  sdk.Coins
		// Send Items Test
		testSendItems       bool
		shouldFailSendItems bool
		// Send Coins Test
		testSendCoins       bool
		shouldFailSendCoins bool
		// Execute Receipe Coin Unlock Test
		testSecondExecuteRecipe   bool
		shouldFailSecondExecution bool
		// Enable trade Coin Lock Test
		testEnableTradeLock bool
		lockDiffEnableTrade sdk.Coins
		// Disable trade Coin Unlock Test
		testDisableTrade     bool
		lockDiffDisableTrade sdk.Coins
	}{
		"create trade and fulfill trade coin lock test": {
			coinsAccount1:            lotsOfCoins,
			coinsAccount2:            lotsOfCoins,
			testCreateTradeLock:      true,
			createTradeAmount:        types.NewPylon(100),
			lockDiffCreateTrade:      types.NewPylon(100),
			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),
		},
		"create trade and send items coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),
			testSendItems:       true,
			shouldFailSendItems: false,
		},
		"create trade and send coins coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),
			testSendCoins:       true,
			shouldFailSendCoins: false,
		},
		"create trade and execute recipe coin lock test": {
			coinsAccount1:             lotsOfCoins,
			coinsAccount2:             lotsOfCoins,
			testCreateTradeLock:       true,
			createTradeAmount:         types.NewPylon(100),
			lockDiffCreateTrade:       types.NewPylon(100),
			testSecondExecuteRecipe:   true,
			shouldFailSecondExecution: false,
		},

		"execute recipe and check execution coin lock test": {
			coinsAccount1:          lotsOfCoins,
			coinsAccount2:          lotsOfCoins,
			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testCheckExecution: true,
			lockDiffCheckExec:  types.NewPylon(100),
		},
		"execute recipe and send items coin lock test": {
			coinsAccount1:          lotsOfCoins,
			coinsAccount2:          lotsOfCoins,
			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testSendItems:       true,
			shouldFailSendItems: false,
		},
		"execute recipe and send coins coin lock test": {
			coinsAccount1:          lotsOfCoins,
			coinsAccount2:          lotsOfCoins,
			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testSendCoins:       true,
			shouldFailSendCoins: false,
		},
		"execute recipe and execute recipe coin lock test": {
			coinsAccount1:          lotsOfCoins,
			coinsAccount2:          lotsOfCoins,
			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testSecondExecuteRecipe:   true,
			shouldFailSecondExecution: false,
		},

		"create trade & fulfill trade & execute recipe & check execution coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testCheckExecution: true,
			lockDiffCheckExec:  types.NewPylon(100),
		},
		"create trade & fulfill trade & execute recipe & check execution and send items coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testCheckExecution: true,
			lockDiffCheckExec:  types.NewPylon(100),

			testSendItems:       true,
			shouldFailSendItems: false,
		},
		"create trade & fulfill trade & execute recipe & check execution and send coins coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testCheckExecution: true,
			lockDiffCheckExec:  types.NewPylon(100),

			testSendCoins:       true,
			shouldFailSendCoins: false,
		},
		"create trade & fulfill trade & execute recipe & check execution and execute recipe coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testCheckExecution: true,
			lockDiffCheckExec:  types.NewPylon(100),

			testSecondExecuteRecipe:   true,
			shouldFailSecondExecution: false,
		},

		"create trade & execute recipe  coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),
		},
		"create trade & execute recipe and send items coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testSendItems:       true,
			shouldFailSendItems: false,
		},
		"create trade & execute recipe and send coins coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testSendCoins:       true,
			shouldFailSendCoins: false,
		},
		"create trade & execute recipe and execute recipe coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testSecondExecuteRecipe:   true,
			shouldFailSecondExecution: false,
		},
		"create trade & disable trade coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testDisableTrade:     true,
			lockDiffDisableTrade: types.NewPylon(100),
		},
		"create trade & disable trade & enable trade coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testDisableTrade:     true,
			lockDiffDisableTrade: types.NewPylon(100),

			testEnableTradeLock: true,
			lockDiffEnableTrade: types.NewPylon(100),
		},
		"create trade & disable trade & enable trade & fulfill trade coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testDisableTrade:     true,
			lockDiffDisableTrade: types.NewPylon(100),

			testEnableTradeLock: true,
			lockDiffEnableTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),
		},
		"create trade & disable trade & enable trade & fulfill trade & execute recipe & check execution coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testDisableTrade:     true,
			lockDiffDisableTrade: types.NewPylon(100),

			testEnableTradeLock: true,
			lockDiffEnableTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testExecuteRecipeLock:  true,
			executeRecipeCoinInput: types.GenCoinInputList("pylon", 100),
			lockDiffSchedule:       types.NewPylon(100),

			testCheckExecution: true,
			lockDiffCheckExec:  types.NewPylon(100),
		},
		"create trade & disable trade & enable trade & fulfill trade & send items coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testDisableTrade:     true,
			lockDiffDisableTrade: types.NewPylon(100),

			testEnableTradeLock: true,
			lockDiffEnableTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testSendItems:       true,
			shouldFailSendItems: false,
		},
		"create trade & disable trade & enable trade & fulfill trade & send coins coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testDisableTrade:     true,
			lockDiffDisableTrade: types.NewPylon(100),

			testEnableTradeLock: true,
			lockDiffEnableTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testSendCoins:       true,
			shouldFailSendCoins: false,
		},
		"create trade & disable trade & enable trade & fulfill trade & second execute recipe coin lock test": {
			coinsAccount1:       lotsOfCoins,
			coinsAccount2:       lotsOfCoins,
			testCreateTradeLock: true,
			createTradeAmount:   types.NewPylon(100),
			lockDiffCreateTrade: types.NewPylon(100),

			testDisableTrade:     true,
			lockDiffDisableTrade: types.NewPylon(100),

			testEnableTradeLock: true,
			lockDiffEnableTrade: types.NewPylon(100),

			testFulfillTrade:         true,
			fulfillTradeInputItemIDs: []string{},
			lockDiffFulfillTrade:     types.NewPylon(100),

			testSecondExecuteRecipe:   true,
			shouldFailSecondExecution: false,
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {

			// setup new account1
			_, account1, err := GenAccount()
			require.True(t, err == nil, err)
			if tc.coinsAccount1 != nil {
				_, err := tci.Bk.AddCoins(tci.Ctx, account1, tc.coinsAccount1.Sort())
				require.True(t, err == nil, err)
			}

			// setup new account2
			_, account2, err := GenAccount()
			require.True(t, err == nil, err)
			if tc.coinsAccount2 != nil {
				_, err := tci.Bk.AddCoins(tci.Ctx, account2, tc.coinsAccount2.Sort())
				require.True(t, err == nil, err)
			}

			tradeData := CreateTradeResponse{}
			scheduleOutput := ExecuteRecipeScheduleOutput{}

			// test create trade coin lock
			if tc.testCreateTradeLock {
				lcFirst := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				tradeData, err = MockTrade(
					tci,
					types.GenCoinInputList("chair", 100),
					types.TradeItemInputList{},
					tc.createTradeAmount,
					types.ItemList{},
					account1,
				)

				require.True(t, err == nil)

				lcAfterCreateTrade := tci.PlnK.GetLockedCoin(tci.Ctx, account1)
				lcDiff := lcAfterCreateTrade.Amount.Sort().Sub(lcFirst.Amount.Sort())

				require.True(t, lcDiff.IsEqual(tc.lockDiffCreateTrade))
			}

			// test disable recipe coin unlock
			if tc.testCreateTradeLock && tc.testDisableTrade {
				lcFirst := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				disableTrdMsg := msgs.NewMsgDisableTrade(tradeData.TradeID, account1)
				_, err := HandlerMsgDisableTrade(tci.Ctx, tci.PlnK, disableTrdMsg)

				require.True(t, err == nil)

				lcAfterDisalbeTrade := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				lcDiff := lcFirst.Amount.Sort().Sub(lcAfterDisalbeTrade.Amount.Sort())

				require.True(t, lcDiff.IsEqual(tc.lockDiffDisableTrade))

			}

			// test enable recipe coin lock
			if tc.testCreateTradeLock && tc.testDisableTrade && tc.testEnableTradeLock {
				lcFirst := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				enableTrdMsg := msgs.NewMsgEnableTrade(tradeData.TradeID, account1)
				_, err := HandlerMsgEnableTrade(tci.Ctx, tci.PlnK, enableTrdMsg)

				require.True(t, err == nil)

				lcAfterEnableTrade := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				lcDiff := lcAfterEnableTrade.Amount.Sort().Sub(lcFirst.Amount.Sort())

				require.True(t, lcDiff.IsEqual(tc.lockDiffEnableTrade))

			}

			// test execute recipe coin lock
			if tc.testExecuteRecipeLock {
				lcFirst := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				item := keep.GenItem(cbData.CookbookID, account1, "Knife")
				err = tci.PlnK.SetItem(tci.Ctx, *item)
				require.True(t, err == nil)

				recipeData := MockRecipe(
					tci,
					"coin lock test recipe",
					tc.executeRecipeCoinInput,
					types.GenItemInputList("Knife"),
					types.GenItemOnlyEntry("KnifeMRG"),
					types.GenOneOutput(1),
					cbData.CookbookID,
					2,
					sender1,
				)

				execRcpResponse, err := MockExecution(
					tci,
					recipeData.RecipeID,
					account1,
					[]string{item.ID},
				)

				require.True(t, err == nil)
				require.True(t, execRcpResponse.Status == "Success")
				require.True(t, execRcpResponse.Message == "scheduled the recipe")

				err = json.Unmarshal(execRcpResponse.Output, &scheduleOutput)
				require.True(t, err == nil)

				lcAfterExecRcp := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				lcDiff := lcAfterExecRcp.Amount.Sort().Sub(lcFirst.Amount.Sort())

				require.True(t, lcDiff.IsEqual(tc.lockDiffSchedule))
			}

			// test send items after coin lock
			if tc.testSendItems {
				item := keep.GenItem(cbData.CookbookID, account1, "sword")
				err = tci.PlnK.SetItem(tci.Ctx, *item)
				require.True(t, err == nil)

				msg := msgs.NewMsgSendItems([]string{item.ID}, account1, account2)
				_, err = HandlerMsgSendItems(tci.Ctx, tci.PlnK, msg)

				if !tc.shouldFailSendItems {
					require.True(t, err == nil)
				} else {
					require.True(t, err != nil)
				}
			}

			// test send coins after coin lock
			if tc.testSendCoins {
				err = keep.SendCoins(tci.PlnK, tci.Ctx, account1, account2, types.NewPylon(100))

				if !tc.shouldFailSendCoins {
					require.True(t, err == nil)
				} else {
					require.True(t, err != nil)
				}
			}

			// test execute recipe after coin lock
			if tc.testSecondExecuteRecipe {
				pylonInputRecipeData := MockRecipe(
					tci, "existing recipe",
					types.GenCoinInputList("pylon", 100),
					types.ItemInputList{},
					types.EntriesList{},
					types.WeightedOutputsList{},
					cbData.CookbookID,
					0,
					sender1,
				)

				msg := msgs.NewMsgExecuteRecipe(
					pylonInputRecipeData.RecipeID,
					account1,
					[]string{},
				)
				_, err := HandlerMsgExecuteRecipe(tci.Ctx, tci.PlnK, msg)

				if !tc.shouldFailSecondExecution {
					require.True(t, err == nil)
				} else {
					require.True(t, err != nil)
				}
			}

			// test fulfill trade coin unlock
			if tc.testCreateTradeLock && tc.testFulfillTrade {
				lcFirst := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				ffMsg := msgs.NewMsgFulfillTrade(
					tradeData.TradeID,
					account2,
					tc.fulfillTradeInputItemIDs,
				)
				_, err = HandlerMsgFulfillTrade(tci.Ctx, tci.PlnK, ffMsg)

				require.True(t, err == nil)

				lcAfterFulfillTrade := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				lcDiff := lcFirst.Amount.Sort().Sub(lcAfterFulfillTrade.Amount.Sort())

				require.True(t, lcDiff.IsEqual(tc.lockDiffFulfillTrade))

			}

			// test check execution coin unlock
			if tc.testExecuteRecipeLock && tc.testCheckExecution {
				lcFirst := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				checkExec := msgs.NewMsgCheckExecution(scheduleOutput.ExecID, false, account1)
				futureContext := tci.Ctx.WithBlockHeight(tci.Ctx.BlockHeight() + 3)
				result, _ := HandlerMsgCheckExecution(futureContext, tci.PlnK, checkExec)
				checkExecResp := CheckExecutionResponse{}

				err = json.Unmarshal(result.Data, &checkExecResp)
				require.True(t, err == nil)
				require.True(t, checkExecResp.Status == "Success")

				lcAfter := tci.PlnK.GetLockedCoin(tci.Ctx, account1)

				lcDiff := lcFirst.Amount.Sort().Sub(lcAfter.Amount.Sort())

				require.True(t, lcDiff.IsEqual(tc.lockDiffCheckExec))

			}
		})
	}
}
