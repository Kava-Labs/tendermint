package store

import (
	amino "github.com/tendermint/go-amino"

	"github.com/kava-labs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}
