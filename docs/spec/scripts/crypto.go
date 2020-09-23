package main

import (
	"fmt"
	"os"

	amino "github.com/kava-labs/go-amino"
	cryptoAmino "github.com/kava-labs/tendermint/crypto/encoding/amino"
)

func main() {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
