// Code has been extracted from SmartPool (https://github.com/smartpool)

package typedefs

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const (
	HashLength          = 16
	WordLength          = 128
	BranchElementLength = 32
)

type (
	Word          [WordLength]byte
	SPHash        [HashLength]byte
	BranchElement [BranchElementLength]byte
)

type EpochData struct {
	Epoch                   *big.Int
	FullSizeIn128Resolution *big.Int
	BranchDepth             *big.Int
	MerkleNodes             []*big.Int
}

func BytesToBig(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)

	return n
}

func rev(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func (w Word) ToUint256Array() []*big.Int {
	result := []*big.Int{}
	for i := 0; i < WordLength/32; i++ {
		z := big.NewInt(0)
		// reverse the bytes because contract expects
		// big Int is construct in little endian
		z.SetBytes(rev(w[i*32 : (i+1)*32]))
		result = append(result, z)
	}
	return result
}

func base62DigitsToString(digits []byte) string {
	for i := 0; i < len(digits); i++ {
		if 0 <= digits[i] && digits[i] <= 9 {
			digits[i] += 48
		} else {
			if 10 <= digits[i] && digits[i] <= 9+26 {
				digits[i] += 97 - 10
			} else {
				if 9+26+1 <= digits[i] && digits[i] <= 9+26+26 {
					digits[i] += 65 - 36
				}
			}
		}
	}
	return string(digits)
}

// return 11 chars base 62 representation of a big int
// base chars are 0-9 a-z A-Z
func BigToBase62(num *big.Int) string {
	digits := []byte{}
	n := big.NewInt(0)
	n.Add(n, num)
	zero := big.NewInt(0)
	base := big.NewInt(62)
	for {
		mod := big.NewInt(0)
		n, mod = n.DivMod(n, base, mod)
		mBytes := mod.Bytes()
		if len(mBytes) == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, mod.Bytes()[0])
		}
		if n.Cmp(zero) == 0 {
			break
		}
	}
	l := len(digits)
	for i := 0; i < 11-l; i++ {
		digits = append(digits, 0)
	}
	return base62DigitsToString(digits)
}

func (h SPHash) Str() string   { return string(h[:]) }
func (h SPHash) Bytes() []byte { return h[:] }
func (h SPHash) Big() *big.Int { return BytesToBig(h[:]) }
func (h SPHash) Hex() string   { return hexutil.Encode(h[:]) }

func (h BranchElement) Str() string   { return string(h[:]) }
func (h BranchElement) Bytes() []byte { return h[:] }
func (h BranchElement) Big() *big.Int { return BytesToBig(h[:]) }
func (h BranchElement) Hex() string   { return hexutil.Encode(h[:]) }

func BranchElementFromHash(a, b SPHash) BranchElement {
	result := BranchElement{}
	copy(result[:], append(a[:], b[:]...)[:BranchElementLength])
	return result
}
