package ethrelay

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type PoWValidationResult int
const (
	PoWValid		= 0
	PoWEpoch		= 1
	PoWDifficulty	= 2
)

func (event EthrelayRemoveBranch) String() string {
	return fmt.Sprintf("branch with root hash %s removed", common.BytesToHash(event.Root[:]))
}

func (event EthrelayPoWValidationResult) String() string {
	switch event.ReturnCode.Int64() {
	case PoWValid:
		return "PoW was successfully validated"
	case PoWEpoch:
		return fmt.Sprintf("epoch data for epoch %d not set", event.ErrorInfo)
	case PoWDifficulty:
		return fmt.Sprintf("calculated difficulty of %d too low", event.ErrorInfo)
	default:
		return fmt.Sprintf("PoWValidationResultEvent: { returnCode: %d, errorInfo: %d }", event.ReturnCode, event.ErrorInfo)
	}
}