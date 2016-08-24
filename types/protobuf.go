package types

import (
	"github.com/tendermint/tmsp/types"
)

// Convert tendermint types to protobuf types
var TM2PB = tm2pb{}

type tm2pb struct{}

func (tm2pb) Header(header *Header) *types.Header {
	return &types.Header{
		ChainId:        header.ChainID,
		Height:         uint64(header.Height),
		Time:           uint64(header.Time.Unix()),
		NumTxs:         uint64(header.NumTxs),
		LastBlockHash:  header.LastBlockHash,
		LastBlockParts: TM2PB.PartSetHeader(header.LastBlockParts),
		LastCommitHash: header.LastCommitHash,
		DataHash:       header.DataHash,
		AppHash:        header.AppHash,
	}
}

func (tm2pb) PartSetHeader(partSetHeader PartSetHeader) *types.PartSetHeader {
	return &types.PartSetHeader{
		Total: uint64(partSetHeader.Total),
		Hash:  partSetHeader.Hash,
	}
}

func (tm2pb) Validator(val *Validator) *types.Validator {
	return &types.Validator{
		PubKey: val.PubKey.Bytes(),
		Power:  uint64(val.VotingPower),
	}
}
