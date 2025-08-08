package types

import (
    
	codectypes "cosmossdk.io/codec/types"
	sdk "cosmossdk.io/types"
	"cosmossdk.io/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
