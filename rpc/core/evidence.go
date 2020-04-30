package core

import (
	"github.com/kava-labs/tendermint/evidence"
	ctypes "github.com/kava-labs/tendermint/rpc/core/types"
	rpctypes "github.com/kava-labs/tendermint/rpc/lib/types"
	"github.com/kava-labs/tendermint/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://docs.tendermint.com/master/rpc/#/Info/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	err := evidencePool.AddEvidence(ev)
	if _, ok := err.(evidence.ErrEvidenceAlreadyStored); err == nil || ok {
		return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
	}
	return nil, err
}
