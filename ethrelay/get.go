package ethrelay

import "github.com/ethereum/go-ethereum/common"

func (c Client) GetLongestChainEndpoint(chainId string) (common.Hash, error) {
	return c.DstChain(chainId).ethrelay.GetLongestChainEndpoint(nil)
}