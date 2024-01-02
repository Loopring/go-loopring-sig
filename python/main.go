package main

import "C"
import (
	"fmt"
	"math/big"
	"strings"

	"github.com/loopring/go-loopring-sig/loopring"
	"github.com/loopring/go-loopring-sig/poseidon"
	"github.com/loopring/go-loopring-sig/utils"
)

// func FudgeyBenchmark() (string, error) {
// 	iterations := 1000
// 	pk := loopring.NewPrivateKeyFromString("0x4485ade3c570854e240c72e9a9162e629f8e30db4d8130856da31787e7400f0")

// 	var signedMessage string
// 	for i := 0; i < iterations; i++ {
// 		hash, err := poseidon.Hash([]*big.Int{
// 			utils.NewIntFromString("11111111111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("91111111111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99111111111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99911111111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99991111111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99999111111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99999911111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99999991111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99999999111111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99999999911111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99999999991111111111111111111111111111111111111111111111111111111111111111"),
// 			utils.NewIntFromString("99999999999111111111111111111111111111111111111111111111111111111111111111"),
// 		})
// 		if err != nil {
// 			return "", err
// 		}

// 		sig := pk.SignPoseidon(hash)

// 		signedMessage = "0x" +
// 			fmt.Sprintf("%064s", sig.R8.X.Text(16)) +
// 			fmt.Sprintf("%064s", sig.R8.Y.Text(16)) +
// 			fmt.Sprintf("%064s", sig.S.Text(16))
// 	}

// 	return signedMessage, nil
// }


//export SignEddsa
func SignEddsa(privateKey *C.char, hash *C.char) *C.char {
	//fmt.Println("enter go SignEddsa!")

	privateKey2 := C.GoString(privateKey)
	hash2 := C.GoString(hash)

	//fmt.Printf("privateKey: %s\n", privateKey2)
	//fmt.Printf("hash: %s\n", hash2)
	
	pk := loopring.NewPrivateKeyFromString(privateKey2)
	sig := pk.SignPoseidon(utils.NewIntFromString(hash2))

	return C.CString("0x" +
		fmt.Sprintf("%064s", sig.R8.X.Text(16)) +
		fmt.Sprintf("%064s", sig.R8.Y.Text(16)) +
		fmt.Sprintf("%064s", sig.S.Text(16)))
}

//export PoseidonHash
func PoseidonHash(input *C.char) *C.char {
	input2 := C.GoString(input)
	data := strings.Split(input2, ",")
	size := len(data)
	bigData := make([]*big.Int, size)
	for i := range bigData {
		bigData[i] = utils.NewIntFromString(data[i])
	}

	hash, err := poseidon.Hash(bigData)
	if err != nil {
		return C.CString("")
	}

	return C.CString(hash.String())
}

//export SignRequest
func SignRequest(privateKey *C.char, method *C.char, url *C.char, data *C.char) *C.char {
	privateKey2 := C.GoString(privateKey)
	method2 := C.GoString(method)
	url2 := C.GoString(url)
	data2 := C.GoString(data)
	result, err := loopring.SignRequest(
		loopring.NewPrivateKeyFromString(privateKey2),
		method2,
		url2,
		"",
		data2,
	)
	if err != nil {
		return C.CString("")
	}
	return C.CString(result)
}

func main() {}
