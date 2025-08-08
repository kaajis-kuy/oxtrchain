package sample

import (
	"cosmossdk.io/crypto/keys/ed25519"
	sdk "cosmossdk.io/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}
