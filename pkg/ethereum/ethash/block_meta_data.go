// Code has been extracted from SmartPool (https://github.com/smartpool)

package ethash

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pantos-io/go-ethrelay/pkg/mtree"
	"github.com/pantos-io/go-ethrelay/pkg/typedefs"
)

type BlockMetaData struct {
	blockNumber uint64
	nonce       uint64
	hashNoNonce common.Hash
	DagTree     *mtree.DagTree
}

type BufferedDag struct {
	Buf []byte
	Indices []uint32
	IndexCnt uint
}

func ProcessAllDuringRead(
	datasetPath string, metaDataArray []*BlockMetaData) {
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
		fmt.Println("before read")
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
		fmt.Println("before write")
		for _, metaData := range metaDataArray {
			metaData.DagTree.Insert(typedefs.Word(buf), i)
		}
		fmt.Println("after write")
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		i++
	}
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

func ProcessBufferedDag(bufferedDag *BufferedDag, mt *mtree.DagTree) {

	for i:= uint(0); i < bufferedDag.IndexCnt; i++ {
		buf := [128]byte{}
		copy(buf[:], bufferedDag.Buf[i*128:(i+1)*128])
		mt.Insert(typedefs.Word(buf), bufferedDag.Indices[i])
	}
	//var f *os.File
	//var err error
	//for {
	//	f, err = os.Open(datasetPath)
	//	if err == nil {
	//		break
	//	} else {
	//		fmt.Printf("Reading DAG file %s failed with %s. Retry in 10s...\n", datasetPath, err.Error())
	//		time.Sleep(10 * time.Second)
	//	}
	//}
	//r := bufio.NewReader(f)
	//buf := [128]byte{}
	//// ignore first 8 bytes magic number at the beginning
	//// of dataset. See more at https://gopkg.in/ethereum/wiki/wiki/Ethash-DAG-Disk-Storage-Format
	//_, err = io.ReadFull(r, buf[:8])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var i uint32 = 0
	//for {
	//	n, err := io.ReadFull(r, buf[:128])
	//	if n == 0 {
	//		if err == nil {
	//			continue
	//		}
	//		if err == io.EOF {
	//			break
	//		}
	//		log.Fatal(err)
	//	}
	//	if n != 128 {
	//		log.Fatal("Malformed dataset")
	//	}
	//	mt.Insert(typedefs.Word(buf), i)
	//	if err != nil && err != io.EOF {
	//		log.Fatal(err)
	//	}
	//	i++
	//}
}

// make sure to only call with meta data of same epoch
func BuildDagTrees(metaDataArray []*BlockMetaData) {
	fmt.Println("step 1")
	MakeDAG(metaDataArray[0].blockNumber, DefaultDir)
	fmt.Println("step 2")
	fullSize := DAGSize(metaDataArray[0].blockNumber)
	fullSizeIn128Resolution := fullSize / 128
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	fmt.Println("step 3")

	for _, s := range metaDataArray {
		indices := Instance.GetVerificationIndices(
			s.blockNumber,
			s.hashNoNonce,
			s.nonce,
		)
		fmt.Printf("indices: %v\n", indices)
		s.DagTree = mtree.NewDagTree()
		s.DagTree.RegisterIndex(indices...)
		s.DagTree.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	}
	fmt.Println("step 4")

	path := PathToDAG(metaDataArray[0].blockNumber/30000, DefaultDir)
	ProcessAllDuringRead(path, metaDataArray)
	fmt.Println("step 5")
	for _, s := range metaDataArray {
		s.DagTree.Finalize()
	}
	fmt.Println("step 6")
}

func (s *BlockMetaData) buildDagTree() {
	indices := Instance.GetVerificationIndices(
		s.blockNumber,
		s.hashNoNonce,
		s.nonce,
	)
	fmt.Printf("indices: %v\n", indices)
	s.DagTree = mtree.NewDagTree()
	s.DagTree.RegisterIndex(indices...)
	MakeDAG(s.blockNumber, DefaultDir)
	fullSize := DAGSize(s.blockNumber)
	fullSizeIn128Resolution := fullSize / 128
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	s.DagTree.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	path := PathToDAG(uint64(s.blockNumber/30000), DefaultDir)
	ProcessDuringRead(path, s.DagTree)
	s.DagTree.Finalize()
}

func (s *BlockMetaData) BuildDagTree(dag *BufferedDag) {
	indices := Instance.GetVerificationIndices(
		s.blockNumber,
		s.hashNoNonce,
		s.nonce,
	)
	//fmt.Printf("indices: %v\n", indices)
	s.DagTree = mtree.NewDagTree()
	s.DagTree.RegisterIndex(indices...)
	//MakeDAG(s.blockNumber, DefaultDir)
	fullSize := DAGSize(s.blockNumber)
	fullSizeIn128Resolution := fullSize / 128
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	s.DagTree.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	//path := PathToDAG(uint64(s.blockNumber/30000), DefaultDir)
	//ProcessDuringRead(path, s.DagTree)
	ProcessBufferedDag(dag, s.DagTree)
	s.DagTree.Finalize()
}

func (s *BlockMetaData) DAGElementArray() []*big.Int {
	if s.DagTree == nil {
		s.buildDagTree()
	}
	result := []*big.Int{}
	for _, w := range s.DagTree.AllDAGElements() {
		result = append(result, w.ToUint256Array()...)
	}
	return result
}

func (s *BlockMetaData) DAGProofArray() []*big.Int {
	if s.DagTree == nil {
		s.buildDagTree()
	}
	result := []*big.Int{}
	for _, be := range s.DagTree.AllBranchesArray() {
		result = append(result, be.Big())
	}
	return result
}

func NewBlockMetaData(blockNumber uint64, nonce uint64, rlpHeaderHashWithoutNonce common.Hash) *BlockMetaData {
	return &BlockMetaData{
		blockNumber,
		nonce,
		rlpHeaderHashWithoutNonce,
		nil,
	}
}
