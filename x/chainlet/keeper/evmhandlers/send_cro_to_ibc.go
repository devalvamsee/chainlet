package evmhandler

import (
	"math/big"

	chainletkeeper "github.com/devalvamsee/chainlet/x/chainlet/keeper"
	"github.com/devalvamsee/chainlet/x/chainlet/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.EvmLogHandler = SendCltToIbcHandler{}

const SendCltToIbcEventName = "__ChainletSendCltToIbc"

// SendCltToIbcEvent represent the signature of
// `event __ChainletSendCltToIbc(string recipient, uint256 amount)`
var SendCltToIbcEvent abi.Event

func init() {
	addressType, _ := abi.NewType("address", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	stringType, _ := abi.NewType("string", "", nil)

	SendCltToIbcEvent = abi.NewEvent(
		SendCltToIbcEventName,
		SendCltToIbcEventName,
		false,
		abi.Arguments{abi.Argument{
			Name:    "sender",
			Type:    addressType,
			Indexed: false,
		}, abi.Argument{
			Name:    "recipient",
			Type:    stringType,
			Indexed: false,
		}, abi.Argument{
			Name:    "amount",
			Type:    uint256Type,
			Indexed: false,
		}},
	)
}

// SendCltToIbcHandler handles `__ChainletSendCltToIbc` log
type SendCltToIbcHandler struct {
	bankKeeper   types.BankKeeper
	chainletKeeper chainletkeeper.Keeper
}

func NewSendCltToIbcHandler(bankKeeper types.BankKeeper, chainletKeeper chainletkeeper.Keeper) *SendCltToIbcHandler {
	return &SendCltToIbcHandler{
		bankKeeper:   bankKeeper,
		chainletKeeper: chainletKeeper,
	}
}

func (h SendCltToIbcHandler) EventID() common.Hash {
	return SendCltToIbcEvent.ID
}

func (h SendCltToIbcHandler) Handle(
	ctx sdk.Context,
	contract common.Address,
	topics []common.Hash,
	data []byte,
	_ func(contractAddress common.Address, logSig common.Hash, logData []byte),
) error {
	unpacked, err := SendCltToIbcEvent.Inputs.Unpack(data)
	if err != nil {
		// log and ignore
		h.chainletKeeper.Logger(ctx).Error("log signature matches but failed to decode", "error", err)
		return nil
	}

	contractAddr := sdk.AccAddress(contract.Bytes())
	sender := sdk.AccAddress(unpacked[0].(common.Address).Bytes())
	recipient := unpacked[1].(string)
	amount := sdkmath.NewIntFromBigInt(unpacked[2].(*big.Int))
	evmDenom := h.chainletKeeper.GetEvmParams(ctx).EvmDenom
	coins := sdk.NewCoins(sdk.NewCoin(evmDenom, amount))
	// First, transfer IBC coin to user so that he will be the refunded address if transfer fails
	if err = h.bankKeeper.SendCoins(ctx, contractAddr, sender, coins); err != nil {
		return err
	}
	// Initiate IBC transfer from sender account
	if err = h.chainletKeeper.IbcTransferCoins(ctx, sender.String(), recipient, coins, ""); err != nil {
		return err
	}
	return nil
}
