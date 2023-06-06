package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAccount{}, "smartaccount/CreateAccount", nil)
	cdc.RegisterConcrete(&MsgUpdateKey{}, "smartaccount/UpdateKey", nil)
// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*authtypes.AccountI)(nil), &SmartAccount{})
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgUpdateKey{},
)
// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
