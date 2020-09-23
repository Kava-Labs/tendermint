package main

import (
	amino "github.com/kava-labs/go-amino"
	ctypes "github.com/kava-labs/tendermint/rpc/core/types"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}
