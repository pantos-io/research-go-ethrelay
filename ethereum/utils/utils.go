package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func WeiToEther(wei *big.Int) string {
	return new(big.Float).Quo(new(big.Float).SetInt(wei),big.NewFloat(params.Ether)).Text('f', 18)
}