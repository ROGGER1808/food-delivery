package common

import (
	"github.com/btcsuite/btcutil/base58"
	"strconv"
)

func UnixToBase58(unix int64) string {
	return base58.Encode([]byte(strconv.Itoa(int(unix))))
}

func Base58ToUnixInt(bs58 string) (int, error) {
	str := string(base58.Decode(bs58))
	return strconv.Atoi(str)
}

func ValidBase582Int(bs58 string) bool {
	str := string(base58.Decode(bs58))
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
