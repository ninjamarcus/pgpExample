package main

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"os"
)

type Keys struct {
	Public  string
	Private string
}

func main() {
	keys := generateKeys()
	fmt.Println(keys.Public)
	fmt.Println(keys.Private)
	writeKeyToDisk(keys.Public, "public.key")
	writeKeyToDisk(keys.Private, "private.key")

	msg := encodeMSG("hello world", keys.Public)
	decodeMSG(msg, keys.Private, keys.Public)

}

func generateKeys() *Keys {
	name := "testName"
	email := "devlist@loadbalancer.org"

	eckey, err := crypto.GenerateKey(name, email, "x25519", 0)

	if err != nil {
		fmt.Println(err)
	}
	privateKey, err := eckey.Armor()
	publicKey, err := eckey.GetArmoredPublicKey()

	return &Keys{
		publicKey,
		privateKey,
	}

}

func writeKeyToDisk(keyValue string, filename string) {
	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(keyValue)

	if err != nil {
		panic(err)
	}
}

func readPublicKey() {

}

func readPrivateKey() {

}

func encodeMSG(message string, publicKey string) string {
	armor, _ := helper.EncryptMessageArmored(publicKey, message)
	fmt.Println(armor)
	return armor
}

func decodeMSG(message string, privateKey string, publicKey string) {
	//TODO:
	publicKeyObj, _ := crypto.NewKeyFromArmored(publicKey)
	publicKeyRing, _ := crypto.NewKeyRing(publicKeyObj)

	privateKeyObj, _ := crypto.NewKeyFromArmored(privateKey)
	privateKeyRing, _ := crypto.NewKeyRing(privateKeyObj)

	msg, _ := crypto.NewPGPMessageFromArmored(message)
	decoded, _ := privateKeyRing.Decrypt(msg, publicKeyRing, crypto.GetUnixTime())
	fmt.Println(decoded.GetString())

}
