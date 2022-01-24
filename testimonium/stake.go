package testimonium

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (c Client) GetStake(chainId string) (*big.Int, error) {
	stake, err := c.DstChain(chainId).testimonium.GetStake(
		&bind.CallOpts{
			From: c.account,
		})
	if err != nil {
		return nil, err
	}
	return stake, nil
}

func (c Client) DepositStake(chainId string, amountInWei *big.Int) error {
	chain := c.DstChain(chainId)
	auth := prepareTransaction(c.account, c.privateKey, &chain.Chain, amountInWei)

	_, err := chain.testimonium.DepositStake(auth, amountInWei)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) WithdrawStake(chainId string, amountInWei *big.Int) error {
	chain := c.DstChain(chainId)
	auth := prepareTransaction(c.account, c.privateKey, &chain.Chain, big.NewInt(0))

	tx, err := chain.testimonium.WithdrawStake(auth, amountInWei)
	if err != nil {
		return err
	}

	receipt, err := awaitTxReceipt(chain.client, tx.Hash())
	if err != nil {
		return err
	}

	if receipt.Status == 0 {
		// Transaction failed
		reason := getFailureReason(chain.client, c.account, tx, receipt.BlockNumber)
		return fmt.Errorf("tx failed: %s", reason)
	}

	// Transaction is successful
	eventIterator, err := chain.testimonium.TestimoniumFilterer.FilterWithdrawStake(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		return err
	}

	if eventIterator.Next() {
		if eventIterator.Event.WithdrawnStake.Cmp(amountInWei) != 0 {
			return errors.New("withdraw not successful, reason: more than 'amount' stake is locked in contract")
		}

		return nil
	}

	return errors.New("uncaught error")
}