package sign

import (
	"crypto/ecdsa"
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
