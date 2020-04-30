package commands

import (
	amino "github.com/tendermint/go-amino"

	cryptoamino "github.com/kava-labs/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoamino.RegisterAmino(cdc)
}
