package sss

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	secret := []byte("secret")
	n := 5
	threshold := 3

	tc := []struct {
		name      string
		n         int
		threshold int
		valid     bool
		error     error
	}{
		{
			name:      "invalid threshold, to low",
			n:         n,
			threshold: 1,
			valid:     false,
			error:     ErrInvalidThreshold,
		},
		{
			name:      "invalid threshold, to high",
			n:         n,
			threshold: 256,
			valid:     false,
			error:     ErrInvalidThreshold,
		},
		{
			name:      "invalid threshold, higher than n",
			n:         n,
			threshold: 6,
			valid:     false,
			error:     ErrInvalidThreshold,
		},
		{
			name:      "invalid number of shares, to high",
			n:         256,
			threshold: threshold,
			valid:     false,
			error:     ErrInvalidNumShares,
		},
		{
			name:      "valid",
			n:         n,
			threshold: threshold,
			valid:     true,
			error:     nil,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			shares, err := Split(tt.n, tt.threshold, secret)
			if !errors.Is(err, tt.error) {
				t.Errorf("Split(%d, %d, %s) = %v; expected %v", tt.n, tt.threshold, secret, err, tt.error)
			}

			if tt.valid && err != nil {
				t.Errorf("Split(%d, %d, %s) = %v; expected no error", tt.n, tt.threshold, secret, err)
			}

			if !tt.valid && err == nil {
				t.Errorf("Split(%d, %d, %s) = %v; expected error", tt.n, tt.threshold, secret, err)
			}

			if tt.valid && len(shares) != n {
				t.Errorf("Split(%d, %d, %s) = %v; expected %d shares", tt.n, tt.threshold, secret, len(shares), n)
			}

			if tt.valid {
				combined, err := Combine(shares[:threshold])
				if err != nil {
					t.Errorf("Combine(%v) = %v; expected no error", shares[:threshold], err)
				}

				if !bytes.Equal(combined, secret) {
					t.Errorf("Combine(%v) = %s; expected %s", shares[:threshold], combined, secret)
				}
			}
		})

	}
}

func TestCombine(t *testing.T) {
	secret := []byte("secret")
	n := 5
	threshold := 3

	shares, err := Split(n, threshold, secret)
	if err != nil {
		t.Fatalf("Split returned error: %v", err)
	}

	invalidShares, err := Split(n, threshold, []byte("invalid"))
	if err != nil {
		t.Fatalf("Split returned error: %v", err)
	}

	tc := []struct {
		name   string
		shares [][]byte
		valid  bool
		error  error
	}{
		{
			name:   "invalid shares",
			shares: invalidShares[:threshold],
			valid:  false,
		},
		{
			name:   "not enough shares",
			shares: shares[:threshold-1],
			valid:  false,
		},
		{
			name:   "too many shares but valid",
			shares: shares[:threshold+1],
			valid:  true,
		},
		{
			name:   "all shares",
			shares: shares,
			valid:  true,
		},
		{
			name:   "empty shares",
			shares: [][]byte{},
			valid:  false,
			error:  ErrInvalidNumShares,
		},
		{
			name:   "one share",
			shares: shares[:1],
			valid:  false,
			error:  ErrInvalidNumShares,
		},
	}

	allCombinations := generateCombinations(shares, threshold)
	for i, comb := range allCombinations {
		tc = append(tc, struct {
			name   string
			shares [][]byte
			valid  bool
			error  error
		}{
			name:   fmt.Sprintf("combination %d", i+1),
			shares: comb,
			valid:  true,
		})
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Combine(tt.shares)
			if !errors.Is(err, tt.error) {
				t.Errorf("Combine(%v) = %v; expected %v", tt.shares, err, tt.error)
			}

			if tt.valid && !bytes.Equal(result, secret) {
				t.Errorf("Combine(%v) = %s; expected %s", tt.shares, result, secret)
			}

			if !tt.valid && bytes.Equal(result, secret) {
				t.Errorf("Combine(%v) = %s; expected error", tt.shares, result)
			}
		})
	}
}
