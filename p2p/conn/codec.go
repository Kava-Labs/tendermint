package conn

import (
	amino "github.com/kava-labs/go-amino"
	cryptoAmino "github.com/kava-labs/tendermint/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
