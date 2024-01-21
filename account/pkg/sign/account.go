package sign

import (
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"github.com/Vigo-Tea/go-ethereum-ant/crypto"
)

// func Message(ctx context.Context, s3Store string, in []byte) ([]byte, error) {
// 	privateKey, err := crypto.HexToECDSA(string(pk))
// 	if err != nil {
// 		return nil, err
// 	}

// 	signedTx, err := types.SignTx(preSignData.Tx, types.NewEIP155Signer(preSignData.ChainID), privateKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	signedTxBuf := bytes.Buffer{}
// 	err = signedTx.EncodeRLP(&signedTxBuf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	signedData := eth.SignedData{
// 		SignedTx: signedTxBuf.Bytes(),
// 	}

// 	return json.Marshal(signedData)
// }

const (
	// rand string
	defaultFuzzStr = "sdfda213esadf43grs"
)

func GenAccount() (pri, pub string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyBytesHex := make([]byte, len(privateKeyBytes)*2)
	hex.Encode(privateKeyBytesHex, privateKeyBytes)

	priKey := string(privateKeyBytesHex)
	pubKey, err := GetPubKey(priKey)
	if err != nil {
		return "", "", err
	}

	return priKey, pubKey, nil
}

func GetPubKey(hexPri string) (pub string, err error) {
	privateKey, err := crypto.HexToECDSA(hexPri)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("create account error casting public key to ECDSA")
	}

	pubKey := crypto.PubkeyToAddress(*publicKeyECDSA).Hex() // Hex String
	return pubKey, nil
}

func DefaultFuzzStr(srcStr string) string {
	return FuzzStr(srcStr, defaultFuzzStr)
}
func DefaultDefuzzStr(srcStr string) (string, error) {
	return DefuzzStr(srcStr, defaultFuzzStr)
}
func FuzzStr(srcStr, fuzzStr string) string {
	fuzzStr = hex.EncodeToString([]byte(fuzzStr))
	src := []byte(srcStr + fuzzStr)

	for i := 0; i < len(src)/3; i++ {
		src[i], src[i+len(src)/3], src[i+len(src)/3*2] = src[i+len(src)/3], src[i+len(src)/3*2], src[i]
	}
	dst1Len := base64.RawStdEncoding.EncodedLen(len(src))
	dst1 := make([]byte, dst1Len)
	base64.RawStdEncoding.Encode(dst1, src)

	return string(dst1)
}

func DefuzzStr(srcStr, fuzzStr string) (string, error) {
	fuzzStr = hex.EncodeToString([]byte(fuzzStr))
	src, err := base64.RawStdEncoding.DecodeString(string(srcStr))
	if err != nil {
		return "", err
	}

	for i := 0; i < len(src)/3; i++ {
		src[i+len(src)/3], src[i+len(src)/3*2], src[i] = src[i], src[i+len(src)/3], src[i+len(src)/3*2]
	}
	srcStr = string(src[:len(src)-len(fuzzStr)])
	return srcStr, nil
}
