// Adopted from https://github.com/soberm/ccsc/blob/4114220216ff59978809086d081c8e7eccc37a07/ccsc_go/pkg/ccsc/merkleproof.go

package ethrelay

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type ByteString []byte

type MerkleProof struct {
	Value ByteString
	Path  ByteString
	Nodes ByteString
}

func (s ByteString) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint("0x", common.Bytes2Hex(s)))
}

func NewMerkleProof(list types.DerivableList, index uint) (*MerkleProof, error) {
	merkleTrie := new(trie.Trie)
	types.DeriveSha(list, merkleTrie)

	path, err := rlp.EncodeToBytes(index)
	if err != nil {
		return nil, err
	}

	nodes, err := LeafProofByKey(merkleTrie, path)
	if err != nil {
		return nil, err
	}

	return &MerkleProof{
		Value: merkleTrie.Get(path),
		Path:  path,
		Nodes: nodes,
	}, nil
}

func LeafProofByKey(mtrie *trie.Trie, key []byte) ([]byte, error) {
	it := mtrie.NodeIterator(nil)
	for it.Next(true) {
		if it.Leaf() && bytes.Equal(it.LeafKey(), key) {
			enc, err := rlp.EncodeToBytes(it.LeafProof());
			if err != nil {
				return nil, err
			}
			return enc, nil
		}
	}
	//lint:ignore ST1005 Merkle is a proper noun
	return nil, fmt.Errorf("Merkle Tree does not contain key %s", common.Bytes2Hex(key))
}