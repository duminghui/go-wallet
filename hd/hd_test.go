package hd

import (
	"fmt"
	"testing"

	"github.com/duminghui/go-wallet/d/chaincfg"
)

func TestWallet_NormalPubKey(t *testing.T) {
	mnemonic := "expand photo unfold meat drive summer wool auto good crystal deposit quick"
	w, err := NewWalletFromMnemonic(mnemonic, "", &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(1, err)
		return
	}
	pi, err := w.Path("m/44'/0'/0'/0/0")
	if err != nil {
		fmt.Println(2, err)
		return
	}
	pubKey, err := pi.ExtendedPubKey()
	if err != nil {
		fmt.Println(4, err)
		return
	}
	fmt.Println(pubKey)
	privKey, err := pi.ExtendedPrivKey()
	if err != nil {
		fmt.Println(5, err)
		return
	}
	fmt.Println(privKey)

	privKeyImport, err := pi.PrivKeyImportable()
	if err != nil {
		fmt.Println(6, err)
		return
	}
	fmt.Println(privKeyImport)

	addr, err := pi.AddressNormal()
	if err != nil {
		fmt.Println(7, err)
		return
	}
	fmt.Println("addr1", addr)

	addr3, err := pi.AddressNestedSegwit()
	if err != nil {
		fmt.Println(8, err)
		return
	}
	fmt.Println("addr3", addr3)

	addrBC, err := pi.AddressNativeSegwit()
	if err != nil {
		fmt.Println(9, err)
		return
	}
	fmt.Println("addrBC", addrBC)
}

func TestPathWrapper_EthAddress(t *testing.T) {

	mnemonic := "expand photo unfold meat drive summer wool auto good crystal deposit quick"
	w, err := NewWalletFromMnemonic(mnemonic, "", &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(1, err)
		return
	}
	pi, err := w.Path("m/44'/60'/0'/0/0")
	if err != nil {
		fmt.Println(2, err)
		return
	}
	ethAddr := pi.EthAddress()

	targetAddr := "0xd124Aa9c2EcaB77939211f03A0926cF47367D6e1"

	fmt.Println(ethAddr, targetAddr == ethAddr)

	privKey := pi.EthImportablePrivKey()

	targetPrivKey := "18b59bc5d18aac1041974e6ffe21ddbee0d307df0acc126ecf47d06309638c33"
	fmt.Println(privKey, targetPrivKey == privKey)

}
