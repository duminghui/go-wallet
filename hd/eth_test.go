package hd

import (
	"fmt"
	"testing"

	"github.com/duminghui/go-wallet/bytesutil"
	"github.com/duminghui/go-wallet/ec"
)

func TestEthPrivKeyImportable(t *testing.T) {
	privKeyStr := "1f2b77e3a4b50120692912c94b204540ad44404386b10c615786a7efaa065d20"
	byteSlice, err := bytesutil.FromHexStrFixZeroPrefix(privKeyStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	privKey, _ := ec.PrivKeyFromBytes(ec.S256(), byteSlice)
	eth := NewEth(privKey)
	testPrivKeyStr := eth.EthPrivKeyImportable()
	fmt.Println(testPrivKeyStr == privKeyStr)
}

func TestEthEncodeAddress(t *testing.T) {
	privKeyStr := "1f2b77e3a4b50120692912c94b204540ad44404386b10c615786a7efaa065d20"
	targetAddr := "0xabcd68033a72978c1084e2d44d1fa06ddc4a2d57"
	byteSlice, err := bytesutil.FromHexStrFixZeroPrefix(privKeyStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	privKey, _ := ec.PrivKeyFromBytes(ec.S256(), byteSlice)
	eth := NewEth(privKey)

	testAddr := eth.EthEncodeAddress()
	fmt.Println(testAddr)
	fmt.Println(targetAddr == testAddr)

}
