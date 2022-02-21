package wallet

import (
	"fmt"
	"crypto/sha256"
	
	"golang.org/x/crypto/ripemd160"
	"github.com/btcsuite/btcutil/base58"
)

func getBitcoinAddress(publicKey []byte) {
	p2kh := getP2KH(publicKey)
	address := base58Check(p2kh)
	fmt.Println(address)
}


func getP2KH(publicKey []byte) []byte{
	newSha256 := sha256.New()
	newSha256.Write(publicKey)

	newRipemd160 := ripemd160.New()
	newRipemd160.Write(newSha256.Sum(nil))
	
	p2kh := newRipemd160.Sum(nil)
	return p2kh
}

func base58Check(p2kh []byte) string{
	p2khPrefix := append([]byte{0},p2kh...)
	fistSha256 := sha256.New()
	secondSha256 := sha256.New()

	fistSha256.Write(p2khPrefix)	
	secondSha256.Write(fistSha256.Sum(nil))

	checksum := secondSha256.Sum(nil)
	preEncode := append(p2khPrefix, checksum[0:4]...)

	return base58.Encode(preEncode)
}