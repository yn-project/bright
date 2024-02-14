package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/golang/protobuf/jsonpb"
	"yun.tea/block/bright/proto/bright/datafin"
)

func main() {
	// baseURL := "https://api.f2pool.com"

	// accessToken := "wyrhuvsac5iaej9s3q1qx3l2lwvuoso1sdxvxzx1rju6tr27bqiujey9sj5ng546"

	// cli := client.NewClient(baseURL, accessToken)
	// resp, err := cli.MiningUserGet(context.Background(), &types.MiningUserGetReq{
	// MiningUserName: "cococonut3",
	// })

	// fmt.Println(utils.PrettyStruct(resp))
	// fmt.Println(err)

	pbM := jsonpb.Marshaler{}
	fmt.Println(pbM.MarshalToString(&datafin.DataItemReq{DataID: "ssss", Data: []byte("Adfasdf")}))
}

type Fin256Hash struct {
	data [32]byte
}

func SumSha256JsonString() *Fin256Hash {

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

func (u *Fin256Hash) ToBigInt() *big.Int {
	return big.NewInt(0).SetBytes(u.data[:])
}

func (u *Fin256Hash) To32Bytes() [32]byte {
	return u.data
}
