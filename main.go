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
	readPublicKeyFromFile()
	decodeMSG(msg, readPrivateKeyFromFile(), readPublicKeyFromFile())
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

func readPublicKeyFromFile() *crypto.KeyRing {

	f, err := os.Open("public.key")
	if err != nil {
		panic(err)
	}
	publicKeyObj, err := crypto.NewKeyFromArmoredReader(f)
	publicKeyRing, err := crypto.NewKeyRing(publicKeyObj)
	return publicKeyRing
}

func readPrivateKeyFromFile() *crypto.KeyRing {
	f, err := os.Open("private.key")
	if err != nil {
		panic(err)
	}
	privateKeyOjb, _ := crypto.NewKeyFromArmoredReader(f)
	privateKeyRing, _ := crypto.NewKeyRing(privateKeyOjb)
	return privateKeyRing
}

func encodeMSG(message string, publicKey string) string {
	armor, _ := helper.EncryptMessageArmored(publicKey, message)
	fmt.Println(armor)
	return armor
}

func decodeMSG(message string, privateKeyRing *crypto.KeyRing, publicKeyRing *crypto.KeyRing) {

	msg, _ := crypto.NewPGPMessageFromArmored(message)
	decoded, _ := privateKeyRing.Decrypt(msg, publicKeyRing, crypto.GetUnixTime())
	privateKeyRing.ClearPrivateParams()
	fmt.Println(decoded.GetString())

}
