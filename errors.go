package sss

import "errors"

var (
	// ErrInvalidThreshold is returned when the threshold is invalid.
	ErrInvalidThreshold = errors.New("invalid threshold")

	// ErrInvalidNumShares is returned when the number of shares is invalid.
	ErrInvalidNumShares = errors.New("invalid number of shares")

	// ErrInvalidSecret is returned when the secret is invalid.
	ErrInvalidSecret = errors.New("invalid secret")
)
