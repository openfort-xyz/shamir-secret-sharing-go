package sss

import (
	"errors"
)

var (
	ErrInvalidThreshold = errors.New("invalid threshold")
	ErrInvalidNumShares = errors.New("invalid number of shares")
	ErrInvalidSecret    = errors.New("invalid secret")
)

func Split(n, threshold int, secret []byte) ([][]byte, error) {
	if threshold < 2 || threshold > n || threshold > 255 {
		return nil, ErrInvalidThreshold
	}

	if n < 2 || n > 255 {
		return nil, ErrInvalidNumShares
	}

	if len(secret) == 0 {
		return nil, ErrInvalidSecret
	}

	shares := make([][]byte, n)
	for i := 0; i < n; i++ {
		shares[i] = make([]byte, len(secret)+1)
		shares[i][0] = byte(i + 1) // x-coordinate
	}

	for i, secretByte := range secret {
		polynomial, err := generatePolynomial(secretByte, threshold)
		if err != nil {
			return nil, err
		}

		for j := 0; j < n; j++ {
			x := byte(j + 1)
			shares[j][i+1] = evaluatePolynomial(polynomial, x)
		}
	}

	return shares, nil
}

func Combine(shares [][]byte) ([]byte, error) {
	if len(shares) < 2 {
		return nil, ErrInvalidNumShares
	}

	length := len(shares[0]) - 1
	secret := make([]byte, length)

	for i := 0; i < length; i++ {
		secret[i] = interpolatePolynomial(shares, i)
	}

	return secret, nil
}
