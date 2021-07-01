package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateId() int {
	generatedString := ""
	for len(generatedString) < 5 {
		nBig, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			panic(err)
		}
		generatedString += nBig.String()
	}
	generatedId, err := strconv.Atoi(generatedString)
	if err != nil {
		return 0
	}
	return generatedId
}
