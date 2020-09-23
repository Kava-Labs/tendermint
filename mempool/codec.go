package mempool

import (
	amino "github.com/kava-labs/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMempoolMessages(cdc)
}
