package sss

import (
	"crypto/rand"
	"math/big"
)

func generatePolynomial(secretByte byte, k int) ([]byte, error) {
	poly := make([]byte, k)
	poly[0] = secretByte

	for i := 1; i < k; i++ {
		coef, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return nil, err
		}
		poly[i] = byte(coef.Int64())
	}

	return poly, nil
}

func evaluatePolynomial(poly []byte, x byte) byte {
	result := byte(0)

	for i := len(poly) - 1; i >= 0; i-- {
		result = gfMul(result, x)
		result = gfAdd(result, poly[i])
	}

	return result
}

func interpolatePolynomial(shares [][]byte, index int) byte {
	secretByte := byte(0)

	for i := 0; i < len(shares); i++ {
		xi := shares[i][0]
		yi := shares[i][index+1]
		li := byte(1)

		for m := 0; m < len(shares); m++ {
			if i != m {
				xm := shares[m][0]
				li = gfMul(li, gfDiv(xm, gfAdd(xm, xi)))
			}
		}

		secretByte = gfAdd(secretByte, gfMul(yi, li))
	}

	return secretByte
}
