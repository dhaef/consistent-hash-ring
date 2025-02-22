package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
)

func hashMd5(key string) string {
	hash := md5.Sum([]byte(key))
	return hex.EncodeToString(hash[:])
}

func md5ToBigInt(md5 string) *big.Int {
	bi := big.NewInt(0)
	bi.SetString(md5, 16)
	return bi
}

func Hash(key string, denominator int) (int, error) {
	hash := hashMd5(key)

	bigInt := md5ToBigInt(hash)

	_, node := new(big.Int).DivMod(bigInt, big.NewInt(int64(denominator)), new(big.Int))

	if !node.IsInt64() {
		return 0, fmt.Errorf("value is too large to be int64")
	}

	return int(node.Int64()), nil
}

// func Hash(key string, denominator int) int {
// 	h := fnv.New32a()
// 	h.Write([]byte(key))
// 	v := h.Sum32()
// 	return int(v % uint32(denominator))
// }
