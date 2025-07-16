package utils

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

type SecretKey struct {
	Key             *ecdsa.PrivateKey
	Privatekey      string // privatekey in hex without "0x"
	Publickey       string // publickey in hex without "0x"
	ACCAountAddress string
	SessionAbi      string
}

// 生成密钥对
func CustomGenerateKey() SecretKey {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 使用`FromECDSA`方法将其转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("public key: ", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)

	s := SecretKey{
		Key:             privateKey,
		Privatekey:      hexutil.Encode(privateKeyBytes)[2:],
		Publickey:       hexutil.Encode(publicKeyBytes)[4:],
		ACCAountAddress: address,
	}
	return s
}
