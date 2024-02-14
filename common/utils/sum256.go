package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
)

type Fin256Hash struct {
	data [32]byte
}

func SumSha256String(data string, compactJSON bool) (*Fin256Hash, error) {
	dataBytes := []byte(data)
	if compactJSON {
		compactStr := bytes.Buffer{}
		err := json.Compact(&compactStr, []byte(data))
		if err != nil {
			return nil, err
		}
		dataBytes = compactStr.Bytes()
	}
	return SumSha256Bytes(dataBytes), nil
}

func SumSha256Bytes(data []byte) *Fin256Hash {
	h := sha256.New()
	h.Write(data)
	dst := h.Sum([]byte{})
	ss := [32]byte{}
	copy(ss[:], dst)
	return &Fin256Hash{data: ss}
}

func From32Bytes(dst []byte) (*Fin256Hash, error) {
	if len(dst) != 32 {
		return nil, fmt.Errorf("failed to parse to [32]bytes,wrong length hexstring")
	}
	ss := [32]byte{}
	copy(ss[:], dst)
	return &Fin256Hash{data: ss}, nil
}

func FromHexString(data string) (*Fin256Hash, error) {
	dst, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	if len(dst) != 32 {
		return nil, fmt.Errorf("failed to parse to [32]bytes,wrong length hexstring")
	}
	ss := [32]byte{}
	copy(ss[:], dst)
	return &Fin256Hash{data: ss}, nil
}

// from uint256
func FromBigInt(src *big.Int) (*Fin256Hash, error) {
	dst := src.Bytes()
	if len(dst) != 32 {
		return nil, fmt.Errorf("failed to parse to [32]bytes,wrong length hexstring")
	}
	ss := [32]byte{}
	copy(ss[:], dst)
	return &Fin256Hash{data: ss}, nil
}

func (u *Fin256Hash) ToHexString() string {
	return fmt.Sprintf("0x%v", hex.EncodeToString(u.data[:]))
}

func (u *Fin256Hash) ToString() string {
	return u.ToHexString()
}

func (u *Fin256Hash) ToBigInt() *big.Int {
	return big.NewInt(0).SetBytes(u.data[:])
}

func (u *Fin256Hash) To32Bytes() [32]byte {
	return u.data
}
