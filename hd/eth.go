package hd

import (
	"encoding/hex"
	"fmt"

	"github.com/duminghui/go-wallet/ec"
	"golang.org/x/crypto/sha3"
)

const (
	HashLength = 32
)

type (
	Hash [HashLength]byte
)

func Keccak256Hash(data ...[]byte) (h Hash) {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(h[:0])
	return h
}

func EthPrivKeyImportable(privKey *ec.PrivateKey) string {
	return hex.EncodeToString(privKey.Serialize())
}

// https://www.codenong.com/cs106540552/
func EthEncodeAddress(pubKey *ec.PublicKey) string {
	keccak256Hash := Keccak256Hash(pubKey.SerializeUncompressed()[1:])
	//addressHash := keccak256Hash[12:]
	//checksum := Keccak256Hash(addressHash)
	//hashCheckSum := hex.EncodeToString(checksum[:])
	h := hex.EncodeToString(keccak256Hash[12:])

	return fmt.Sprintf("0x%s", h)
}
