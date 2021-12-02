package bip39

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"math/big"
	"strings"
	"testing"
)

func TestNewEntropy(t *testing.T) {
	bitSizeSlice := []int{
		128, 160, 192, 224, 256,
	}
	for _, bitSize := range bitSizeSlice {
		entropy, err := NewEntropy(bitSize)
		fmt.Println(bitSize, bitSize/32)
		if err != nil {
			fmt.Println(bitSize, err)
		} else {
			fmt.Println(bitSize, entropy)
		}
	}
}

func TestNewMnemonic(t *testing.T) {
	bitSizeSlice := []int{
		128, 160, 192, 224, 256,
	}
	for _, bitSize := range bitSizeSlice {
		bytes, err := NewEntropy(bitSize)
		if err != nil {
			fmt.Println(err)
			return
		}
		mnemonic, err := NewMnemonic(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("mnemonic:", mnemonic, len(strings.Split(mnemonic, " ")))
		seed := NewSeed(mnemonic, "")
		fmt.Println(len(seed))
	}
}

func TestNewSeed(t *testing.T) {
	entropy, _ := NewEntropy(256)
	fmt.Printf("entropy:%x\n", entropy)
	mnemonic, _ := NewMnemonic(entropy)
	fmt.Println("mnemonic:", mnemonic, len(strings.Split(mnemonic, " ")))
	seed := NewSeed(mnemonic, "11111")
	fmt.Printf("seed:%x\n", seed)
}

func TestIsMnemonic(t *testing.T) {
	isMnemonic := IsMnemonicValid("tail reason eagle piece shrimp before annual creek exact insect carry police special rival promote lounge clever click slab toward bridge")
	fmt.Println(isMnemonic)
}

func TestNewSeedWithErrorChecking(t *testing.T) {
	// t.SkipNow()
	entropy, _ := NewEntropy(128)
	fmt.Printf("entropy: %x\n", entropy)
	mnemonic, _ := NewMnemonic(entropy)
	// mnemonic = "much local guess refuse cannon project march dwarf color sleep fringe safe"
	// mnemonic = "army van defense carry jealous true garbage claim echo media make crunch"
	fmt.Println("mnemonic:", mnemonic)
	seed, _ := NewSeedWithErrorChecking(mnemonic, "")
	fmt.Printf("seed: %x , %d\n", seed, len(seed))
	hmac512 := hmac.New(sha512.New, []byte("Bitcoin seed"))
	hmac512.Write(seed)
	masterKey := hmac512.Sum(nil)
	fmt.Printf("masterKey: %x\n", masterKey)
	masterPriKey := masterKey[:32]
	fmt.Printf("masterPriKey: %x\n", masterPriKey)
	fmt.Println(len(masterPriKey))
	masterChainCode := masterKey[32:]
	fmt.Printf("masterChainCode: %x\n", masterChainCode)
}

func TestFunction(t *testing.T) {
	var tmp big.Int
	n, err := fmt.Sscan("0xF", &tmp)
	fmt.Println(tmp, n, err)
	fmt.Println(fmt.Sprintf("%d", tmp.Bytes()))
	fmt.Println(fmt.Sprintf("%x", tmp.String()))
}

//// https://github.com/iancoleman/bip39/issues/58
//// 17rxURoF96VhmkcEGCj5LNQkmN9HVhWb7F
//func TestVector3_2(t *testing.T) {
//	mnemnic := "fruit wave dwarf banana earth journey tattoo true farm silk olive fence"
//	seed, _ := NewSeedWithErrorChecking(mnemnic, "banana")
//	key, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
//
//	fmt.Println(key.String())
//	childKey, _ := key.DerivePath("m/44'/0'/0'/0/0")
//	address, _ := childKey.Address(&chaincfg.MainNetParams)
//	fmt.Println(address.EncodeAddress())
//}
