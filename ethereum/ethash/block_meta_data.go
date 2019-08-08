// Code has been extracted from SmartPool (https://github.com/smartpool)

package ethash

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pf92/go-testimonium/mtree"
	"github.com/pf92/go-testimonium/typedefs"
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
	indices := Instance.GetVerificationIndices(
		s.blockNumber,
		s.hashNoNonce,
		s.nonce,
	)
	fmt.Printf("indices: %v\n", indices)
	s.dt = mtree.NewDagTree()
	s.dt.RegisterIndex(indices...)
	MakeDAG(s.blockNumber, DefaultDir)
	fullSize := DAGSize(s.blockNumber)
	fullSizeIn128Resolution := fullSize / 128
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	s.dt.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	path := PathToDAG(uint64(s.blockNumber/30000), DefaultDir)
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

func NewBlockMetaData(blockNumber uint64, nonce uint64, rlpHeaderHashWithoutNonce [32]byte) *BlockMetaData {
	return &BlockMetaData{
		blockNumber,
		nonce,
		rlpHeaderHashWithoutNonce,
		nil,
	}
}
