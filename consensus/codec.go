package consensus

import (
	amino "github.com/kava-labs/go-amino"
	"github.com/kava-labs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterConsensusMessages(cdc)
	RegisterWALMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
