package state_test

import (
	"os"
	"testing"

	"github.com/kava-labs/tendermint/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}
