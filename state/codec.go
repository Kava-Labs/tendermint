package state

import (
	amino "github.com/kava-labs/go-amino"
	cryptoAmino "github.com/kava-labs/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
