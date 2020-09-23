package core_types

import (
	amino "github.com/kava-labs/go-amino"
	"github.com/kava-labs/tendermint/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
