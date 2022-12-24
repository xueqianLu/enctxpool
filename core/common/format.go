package common

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"
	encservicetype "github.com/xueqianLu/enctxpool/protocol/generate/encservice/v1"
	"math/big"
)

func ParseBlockHeader(response *encservicetype.LatestHeaderResponse) *types.Header {
	var header *types.Header
	if len(response.HeaderJson) > 0 {
		err := json.Unmarshal(response.HeaderJson, &header)
		if err != nil {
			return nil
		}
	}
	return header
}

func ParseBlockData(blockdata []byte) *types.Block {
	var block *types.Block
	if len(blockdata) > 0 {
		err := json.Unmarshal(blockdata, &block)
		if err != nil {
			return nil
		}
	}
	return block
}

func ParseBalance(response *encservicetype.BalanceResponse) *big.Int {
	v := new(big.Int)
	v.SetBytes(response.Balance)
	return v
}

func ParseNonce(response *encservicetype.NonceResponse) uint64 {
	return response.Nonce
}

func ParseTxData(rlptx []byte) *types.Transaction {
	tx := new(types.Transaction)
	tx.UnmarshalBinary(rlptx)
	return tx
}