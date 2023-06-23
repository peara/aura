package keeper

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"strconv"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/aura-nw/aura/x/smartaccount/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InstantiateSmartAccount(
	ctx sdk.Context, 
	keeper Keeper, 
	wasmKeepper *wasmkeeper.PermissionedKeeper, 
	msg *types.MsgActivateAccount,
) (sdk.AccAddress, []byte, cryptotypes.PubKey, error) {

	admin, err := sdk.AccAddressFromBech32(msg.AccountAddress)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(types.ErrAddressFromBech32, err)
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(types.ErrAddressFromBech32, err)
	}

	pubKey, err := types.PubKeyDecode(msg.PubKey)
	if err != nil {
		return nil, nil, nil, err
	}

	salt := types.InstantiateSalt{
		Owner:   msg.Owner,
		CodeID:  msg.CodeID,
		InitMsg: msg.InitMsg,
		PubKey:  pubKey.Bytes(),
	}

	saltBytes, err := json.Marshal(salt)
	if err != nil {
		return nil, nil, nil, err
	}

	salt_hashed := sha512.Sum512(saltBytes)

	// instantiate smartcontract by code id
	address, data, iErr := wasmKeepper.Instantiate2(
		ctx,
		msg.CodeID,
		owner,       // owner
		admin,       // admin
		msg.InitMsg, // message
		fmt.Sprintf("%s/%d", types.ModuleName, keeper.GetAndIncrementNextAccountID(ctx)), // label
		sdk.NewCoins(), // empty funds
		salt_hashed[:], // salt
		true,
	)
	if iErr != nil {
		return nil, nil, nil, fmt.Errorf(types.ErrBadInstantiateMsg, iErr.Error())
	}

	contractAddrStr := address.String()

	ctx.Logger().Info(
		"smart account created",
		types.AttributeKeyCreator, msg.AccountAddress,
		types.AttributeKeyCodeID, msg.CodeID,
		types.AttributeKeyContractAddr, contractAddrStr,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAccountRegistered,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.AccountAddress),
			sdk.NewAttribute(types.AttributeKeyCodeID, strconv.FormatUint(msg.CodeID, 10)),
			sdk.NewAttribute(types.AttributeKeyContractAddr, contractAddrStr),
		),
	)

	return address, data, pubKey, nil
}
