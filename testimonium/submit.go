package testimonium

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pantos-io/go-ethrelay/typedefs"
)

func (c Client) SubmitHeader(chainId string, header *types.Header) error {
	fmt.Printf("Submitting block: \nNo: %s\nHash: %s\n", header.Number, header.Hash())

	rlpHeader, err := encodeHeaderToRLP(header)
	if err != nil {
		return fmt.Errorf("failed to encode header to RLP: %s", err)
	}

	return c.SubmitRLPHeader(chainId, rlpHeader)
}

func (c Client) SubmitHeaderLive(dstChainId string, srcChainId string, lockTime time.Duration) error {
	/*
		there is much more to care about here:
		- 	if the genesis block of the testimonium contract is not on the current main chain,
			it is unpossible to find a block from where we can start submitting new blocks.
			this should be taken into account when creating a contract by selecting a genesis block
			far enough in the past that is ensured to not get into this scenario. the best case would be
			to start at block 0 of the source blockchain as this can never be a forked orphan branch.
			if this is the case, the search would go to block no. 0 to realize the genesis block is not on
			a given branch. this could be prevented by including a contract-creation search and check if the
			given blockNumber is not smaller than the genesis block, else it is useless and we can stop immediately.
		- 	if new blocks occur on the main chain, it is not safe if thy get into the longest chain or not.
			a participant in the relay-network may not take the risk to submit transactions with a high transaction-
			cost if they will never get any fees for this block. here, the participant with the fastest
			internet-connection (receiving new blocks from the main-chain and sending the blocks to the
			destination chain only requires latency and no hard problems) can relay all new blocks and
			wins the race so all fees go to the single participant. maybe this is not a big problem as relays
			can be deployed in the cloud and run in a very cheap and cost-intensive way.
		-	at the very first time the search and submit-phase takes a very long time because first one
			has to go back all the blocks and than submit all single blocks. here, the only way this could
			be enhanced is batch processing, which the testimonium contract already supports and some kind of
			binary search to effectively search for the latest submitted block which is part of the longest chain
			and part in the testimonium contract
		-	there may be not enough stake to deposit all blocks so one have to wait until the blocks are unlocked and
			the stake is freed again.
		-	another one has already submitted the block between finding this block with the search and submitting the block
		-	implement a live mode by subscribing to new blocks, here the same errors as above can occur
	*/

	// as the relay makes only sense if it's on the latest state of the main chain and it should never get older
	// than a few blocks behind the main chain it should be more performant and easier to do a simple backward-scan
	// starting with the newest block of the main chain, on the other hand, a binary search is always very fast and
	// if in the backwards search more than log2(n) blocks are stored, a binary search is always faster, so maybe
	// implement a binary search as default here

	dstChain, srcChain := c.DstChain(dstChainId), c.SrcChain(srcChainId)

	genesis, err := dstChain.testimonium.GetGenesisBlockHash(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Getting sure ETH Relay genesis block 0x%s from destination chain '%s' really exists on source chain '%s'\n", common.Bytes2Hex(genesis[:]), srcChainId, dstChainId)

	// returns an error if genesis was not found
	_, err = srcChain.client.HeaderByHash(context.Background(), genesis)
	if err != nil {
		log.Fatal(err)
	}

	// at the beginning this is nil - which returns the most recent block
	var blockNumber *big.Int = nil
	one := big.NewInt(1)

	var header *types.Header = nil

	// find the most recent block that was already submitted
	for {
		// get newest, longest header from source chain
		header, err = c.HeaderByNumber(srcChainId, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nSearching for block No. %s from source chain '%s' on destination chain '%s'", header.Number, srcChainId, dstChainId)

		isHeaderStored, err := dstChain.testimonium.IsHeaderStored(nil, header.Hash())
		if err != nil {
			log.Fatal(err)
		}

		if isHeaderStored {
			break
		}

		blockNumber = header.Number

		blockNumber.Sub(blockNumber, one)
	}

	fmt.Printf("\n\nlatest block No. submitted to destination chain: %s\n\n", header.Number)

	requiredStake, err := dstChain.testimonium.GetRequiredStakePerBlock(nil)
	if err != nil {
		log.Fatal(err)
	}

	stake, err := c.GetStake(dstChainId)
	if err != nil {
		log.Fatal(err)
	}

	// check if there is enough stake
	if stake.Cmp(requiredStake) < 0 {
		log.Fatal("not enough stake deposited")
	}

	// has to be bigger than one as we checked above
	// TODO: if stake is currently locked, the routine tries to use it, but an error occurs
	maxBlocksWithStake := big.NewInt(0)
	maxBlocksWithStake.Div(stake, requiredStake)

	// calculate max. block submissions with stake
	var queue []time.Time

	// blockNumber was updated, so the destination chain is a few blocks behind source chain - updating now
	if blockNumber != nil {
		// submit all blocks to the most recent one
		for {
			if len(queue) >= int(maxBlocksWithStake.Uint64()) {
				timeUntilNextBlockIsUnlocked := queue[0].Add(lockTime)
				waitingTime := timeUntilNextBlockIsUnlocked.Sub(time.Now())

				if waitingTime > 0 {
					fmt.Printf("All stake is locked, waiting for %fs to continue\n", waitingTime.Seconds())
					time.Sleep(waitingTime)
				}

				queue = queue[1:]
			}

			// increase by one as we only want blocks that are new
			blockNumber.Add(blockNumber, one)

			header, err := c.HeaderByNumber(srcChainId, blockNumber)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Stake queue-length:", len(queue))

			// TODO: a check for enough free/unlocked stake is required here, though a time based workaround is already implemented
			err = c.SubmitHeader(dstChainId, header)
			if err != nil {
				log.Fatal(err)
			}

			// add now + 1m for latency and whatever
			queue = append(queue, time.Now().Add(time.Second))

			// get newest, longest header from source chain
			header, err = c.HeaderByNumber(srcChainId, nil)
			if err != nil {
				log.Fatal(err)
			}

			// we caught up all the blocks... continue
			if header.Number.Cmp(blockNumber) == 0 {
				break
			}
		}
	}

	fmt.Printf("\nstarting live mode...\n\n")

	headers := make(chan *types.Header)

	sub, err := srcChain.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			if len(queue) >= int(maxBlocksWithStake.Uint64()) {
				timeUntilNextBlockIsUnlocked := queue[0].Add(lockTime)
				waitingTime := timeUntilNextBlockIsUnlocked.Sub(time.Now())

				if waitingTime > 0 {
					fmt.Printf("All stake is locked, waiting for %fs to continue\n", waitingTime.Seconds())
					time.Sleep(waitingTime)
				}

				queue = queue[1:]
			}

			fmt.Println("Stake queue-length: ", len(queue))

			err = c.SubmitHeader(dstChainId, header)
			if err != nil {
				log.Fatal(err)
			}

			queue = append(queue, time.Now().Add(time.Second))
		}
	}
}

func (c Client) SubmitRLPHeader(chainId string, rlpHeader []byte) error {
	chain := c.DstChain(chainId)
	auth := prepareTransaction(c.account, c.privateKey, &chain.Chain, big.NewInt(0))

	tx, err := chain.testimonium.SubmitBlock(auth, rlpHeader)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := awaitTxReceipt(chain.client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		reason := getFailureReason(chain.client, c.account, tx, receipt.BlockNumber)
		return errors.New(reason)
	}

	eventIterator, err := chain.testimonium.TestimoniumFilterer.FilterNewBlock(&bind.FilterOpts{
		Start:   receipt.BlockNumber.Uint64(),
		End:     nil,
		Context: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	// TODO: is this really the next event on the same chain? what if a transaction is included into one block,
	//  but get's overtaken by a submit-event of another fork? if it is guaranteed that the client is not reconnecting
	//  to other nodes within a usage - this may also be the case on every other transaction call
	//  workaround: check that the transaction from eventIterator's event is the same as the submitted transaction above
	if eventIterator.Next() {
		// TODO: this is only 1 special hash value emitted by the contract for too small stake and not a read error code
		if (eventIterator.Event.BlockHash == common.Hash{}) {
			return errors.New("block was not submitted, reason: too small stake deposited")
		}

		return nil
	}

	return errors.New("uncaught error")
}

func (c Client) SetEpochData(chainId string, epochData typedefs.EpochData) {
	chain := c.DstChain(chainId)
	nodes := []*big.Int{}
	start := big.NewInt(0)
	//fmt.Printf("No meaningful nodes: %d\n", len(epochData.MerkleNodes))
	for k, n := range epochData.MerkleNodes {
		nodes = append(nodes, n)
		if len(nodes) == 40 || k == len(epochData.MerkleNodes)-1 {
			mnlen := big.NewInt(int64(len(nodes)))
			fmt.Printf("Going to do tx\n")

			if k < 440 && epochData.Epoch.Uint64() == 128 {
				start.Add(start, mnlen)
				nodes = []*big.Int{}
				continue
			}

			auth := prepareTransaction(c.account, c.privateKey, &chain.Chain, big.NewInt(0))

			tx, err := chain.ethash.SetEpochData(auth, epochData.Epoch, epochData.FullSizeIn128Resolution,
				epochData.BranchDepth, nodes, start, mnlen)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

			receipt, err := awaitTxReceipt(chain.client, tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			if receipt.Status == 0 {
				// Transaction failed
				reason := getFailureReason(chain.client, c.account, tx, receipt.BlockNumber)
				fmt.Printf("Tx failed: %s\n", reason)
				return
			}

			start.Add(start, mnlen)
			nodes = []*big.Int{}
		}
	}
}

func (c Client) RandomizeHeader(chainId string, header *types.Header) *types.Header {
	temp := header.TxHash

	header.TxHash = header.ReceiptHash
	header.ReceiptHash = header.Root
	header.Root = temp

	return header
}