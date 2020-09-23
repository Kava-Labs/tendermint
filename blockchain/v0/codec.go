package v0

import (
	amino "github.com/kava-labs/go-amino"
	"github.com/kava-labs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
