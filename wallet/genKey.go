package wallet

import (
	"os"
	"fmt"
	"errors"
	"io/ioutil"
	"encoding/pem"
	"crypto/rand"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// key pair paths
const privPath = "wallet/priv.pem"
const pubPath = "wallet/pub.pem"

// UserChoice() runs the main cli
func UserChoice() {
	var chAccepted int
	if _, err := os.Stat(privPath); err == nil {
		fmt.Print("There is already a pre-saved wallet. What do you want to do?\n\t1. Use the pre-saved wallet.\n\t2. Generate a new one.\n\t3. Exit.\nPlease choose an option: ")
		fmt.Scanf("%d\n", &chAccepted)
		switch chAccepted  {
		case 1:
			genKeyPair()
		case 2:
			fmt.Print("Generating new wallet... What do you want to do?\n\t1. Generate a random key.\n\t2. Recover wallet from mnemonic.\n\t3. Exit.\nPlease choose an option: ")
			fmt.Scanf("%d\n", &chAccepted)
			switch chAccepted  {
			case 1:
				genKeyPair()
			case 2:
				fmt.Println("Coming soon...")
			case 3:
				os.Exit(0)
			default:
				os.Exit(0)
			}
		case 3:
			os.Exit(0)
		default:
			os.Exit(0)
		}
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Print("There is no pre-saved wallet. What do you want to do?\n\t1. Create a new one.\n\t2. Recover wallet from mnemonic.\n\t3. Exit.\nPlease choose an option: ")
		fmt.Scanf("%d\n", &chAccepted)
		switch chAccepted  {
		case 1:
			genKeyPair()
		case 2:
			fmt.Println("Coming soon...")
		case 3:
			os.Exit(0)
		default:
			os.Exit(0)
		}
	}
}

// genKeyPair(): this function generates a random private-public key pair using the secp256k1 curve
// and then saves them to the wallet directory
func genKeyPair() {
	fmt.Println("Generating new key pair...")
	curve256 := secp256k1.S256()
	privateKey, err := ecdsa.GenerateKey(curve256, rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := (curve256.Marshal(privateKey.PublicKey.X,privateKey.PublicKey.Y))
		
	saveKeyPair(privPath, privateKey.D.Bytes(), pubPath, publicKey)
	getBitcoinAddress(publicKey)
}

// saveKeyPair(): this function saves both keys in PEM format
func saveKeyPair(filenamePriv string, privateKey []byte, filenamePub string, publicKey []byte) {
	filePriv, _ := os.Create(filenamePriv)
	defer filePriv.Close()				
	privBlock := &pem.Block{Type: "PRIVATE KEY", Bytes: privateKey}
	err := pem.Encode(filePriv, privBlock)
	if err != nil {
		panic(err)
	}

	filePub, _ := os.Create(filenamePub)
	defer filePub.Close()
	pubBlock := &pem.Block{Type: "PRIVATE KEY", Bytes: publicKey}
	err = pem.Encode(filePub, pubBlock)
	if err != nil {
		panic(err)
	}
}


// saveKeyPair(): this function loads both saved keys
func loadKeys() ([]byte, []byte) {
	privFile, _ := ioutil.ReadFile(privPath)
	privBytes, _  := pem.Decode(privFile)

	pubFile, _ := ioutil.ReadFile(pubPath)
	pubBytes, _ := pem.Decode(pubFile)

	return privBytes.Bytes, pubBytes.Bytes
}