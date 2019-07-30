package ethereum

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pf92/testimonium-cli/testimonium"
	"github.com/probatio-client/ethereum/ethash"
	"github.com/probatio-client/mtree"
	"github.com/probatio-client/typedefs"
	"io"
	"log"
	"math/big"
	"os"
	"time"
)

type BlockMetaData struct {
	blockNumber		uint64
	nonce			uint64
	hashNoNonce		common.Hash
	dt              *mtree.DagTree
}

// SealHash returns the hash of a block prior to it being sealed.
//func SealHash(header *types.Header) (hash common.Hash) {
//	hasher := sha3.NewLegacyKeccak256()
//
//	rlp.Encode(hasher, []interface{}{
//		header.ParentHash,
//		header.UncleHash,
//		header.Coinbase,
//		header.Root,
//		header.TxHash,
//		header.ReceiptHash,
//		header.Bloom,
//		header.Difficulty,
//		header.Number,
//		header.GasLimit,
//		header.GasUsed,
//		header.Time,
//		header.Extra,
//	})
//	hasher.Sum(hash[:0])
//	return hash
//}

//func (s *BlockMetaData) BlockHeader() *types.Header {
//	return s.blockHeader
//}
//
//func (s *BlockMetaData) RlpHeaderWithoutNonce() ([]byte, error) {
//	buffer := new(bytes.Buffer)
//	err := rlp.Encode(buffer, []interface{}{
//		s.BlockHeader().ParentHash,
//		s.BlockHeader().UncleHash,
//		s.BlockHeader().Coinbase,
//		s.BlockHeader().Root,
//		s.BlockHeader().TxHash,
//		s.BlockHeader().ReceiptHash,
//		s.BlockHeader().Bloom,
//		s.BlockHeader().Difficulty,
//		s.BlockHeader().Number,
//		s.BlockHeader().GasLimit,
//		s.BlockHeader().GasUsed,
//		s.BlockHeader().Time,
//		s.BlockHeader().Extra,
//	})
//	fmt.Printf("RLP: 0x%s\n", hex.EncodeToString(buffer.Bytes()))
//	return buffer.Bytes(), err
//}

//func (s *BlockMetaData) Timestamp() *big.Int {
//	return new(big.Int).SetUint64(s.blockHeader.Time)
//}

//func (s *BlockMetaData) Hash() (result typedefs.SPHash) {
//	h := SealHash(s.blockHeader)
//	copy(result[:typedefs.HashLength], h[typedefs.HashLength:])
//	return
//}

func ProcessDuringRead(
	datasetPath string, mt *mtree.DagTree) {
	var f *os.File
	var err error
	for {
		f, err = os.Open(datasetPath)
		if err == nil {
			break
		} else {
			fmt.Printf("Reading DAG file %s failed with %s. Retry in 10s...\n", datasetPath, err.Error())
			time.Sleep(10 * time.Second)
		}
	}
	r := bufio.NewReader(f)
	buf := [128]byte{}
	// ignore first 8 bytes magic number at the beginning
	// of dataset. See more at https://gopkg.in/ethereum/wiki/wiki/Ethash-DAG-Disk-Storage-Format
	_, err = io.ReadFull(r, buf[:8])
	if err != nil {
		log.Fatal(err)
	}
	var i uint32 = 0
	for {
		n, err := io.ReadFull(r, buf[:128])
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if n != 128 {
			log.Fatal("Malformed dataset")
		}
		mt.Insert(typedefs.Word(buf), i)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		i++
	}
}

func (s *BlockMetaData) buildDagTree() {
	indices := ethash.Instance.GetVerificationIndices(
		s.blockNumber,
		s.hashNoNonce,
		s.nonce,
	)
	fmt.Printf("indices: %v\n", indices)
	s.dt = mtree.NewDagTree()
	s.dt.RegisterIndex(indices...)
	ethash.MakeDAG(s.blockNumber, ethash.DefaultDir)
	fullSize := ethash.DAGSize(s.blockNumber)
	fullSizeIn128Resolution := fullSize / 128
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	s.dt.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	path := ethash.PathToDAG(uint64(s.blockNumber/30000), ethash.DefaultDir)
	ProcessDuringRead(path, s.dt)
	s.dt.Finalize()
}

func (s *BlockMetaData) DAGElementArray() []*big.Int {
	if s.dt == nil {
		s.buildDagTree()
	}
	result := []*big.Int{}
	for _, w := range s.dt.AllDAGElements() {
		result = append(result, w.ToUint256Array()...)
	}
	return result
}

func (s *BlockMetaData) DAGProofArray() []*big.Int {
	if s.dt == nil {
		s.buildDagTree()
	}
	result := []*big.Int{}
	for _, be := range s.dt.AllBranchesArray() {
		result = append(result, be.Big())
	}
	return result
}

//func NewBlockMetaData(h *types.Header) *BlockMetaData {
//	return &BlockMetaData{
//		h,
//		nil,
//	}
//}
func NewBlockMetaData(h *testimonium.BlockHeader) *BlockMetaData {
	return &BlockMetaData{
		h.BlockNumber.Uint64(),
		h.Nonce.Uint64(),
		h.RlpHeaderHashWithoutNonce,
		nil,
	}
}
