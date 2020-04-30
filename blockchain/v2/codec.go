package v2

import (
	amino "github.com/tendermint/go-amino"

	"github.com/kava-labs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
