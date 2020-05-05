package evidence

import (
	amino "github.com/tendermint/go-amino"
	cryptoamino "github.com/kava-labs/tendermint/crypto/encoding/amino"
	"github.com/kava-labs/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMessages(cdc)
	cryptoamino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}
