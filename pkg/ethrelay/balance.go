package ethrelay

import (
	"context"
	"math/big"
)

func (c Client) TotalBalance() (*big.Int, error) {
	var totalBalance = new(big.Int)
	for k := range c.chains {
		balance, err := c.Balance(k)
		if err != nil {
			return nil, err
		}
		totalBalance.Add(totalBalance, balance)
	}
	return totalBalance, nil
}

func (c Client) Balance(chainId string) (*big.Int, error) {
	var totalBalance = new(big.Int)

	balance, err := c.Chain(chainId).client.BalanceAt(context.Background(), c.account, nil)
	if err != nil {
		return nil, err
	}

	totalBalance.Add(totalBalance, balance)

	return totalBalance, nil
}