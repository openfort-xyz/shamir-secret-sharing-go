# Shamir's Secret Sharing
[![Go](https://github.com/openfort-xyz/shamir-secret-sharing-go/actions/workflows/ci.yml/badge.svg)](https://github.com/openfort/shamir-secret-sharing/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/go.openfort.xyz/shamir-secret-sharing-go)](https://goreportcard.com/report/go.openfort.xyz/shamir-secret-sharing-go)

This project implements Shamir's Secret Sharing with zero dependency in Go. Shamir's Secret Sharing is a cryptographic algorithm that allows a secret to be divided into parts, giving each participant its own unique part. To reconstruct the secret, a minimum number of parts are needed. This implementation uses GF(256) arithmetic for secure and efficient operations.

## Installation

To install this package, you need to have [Go](https://golang.org/doc/install) installed on your machine.

```sh
go get -u go.openfort.xyz/shamir-secret-sharing-go
```

## Usage

### Import the package 
```go
import (
    "go.openfort.xyz/shamir-secret-sharing-go"
)
```

### Example: Splitting and Combining a Secret

```go
package main

import (
    "fmt"
    sss "go.openfort.xyz/shamir-secret-sharing-go"
)

func main() {
    secret := []byte("this is a secret")
    n := 5
    threshold := 3

    shares, err := sss.Split(n, threshold, secret)
    if err != nil {
        fmt.Println("Error splitting the secret:", err)
        return
    }

    fmt.Println("Shares:")
    for i, share := range shares {
        fmt.Printf("Share %d: %v\n", i+1, share)
    }

    // Using threshold shares to reconstruct the secret
    recoveredSecret, err := sss.Combine(shares[:threshold])
    if err != nil {
        fmt.Println("Error reconstructing the secret:", err)
        return
    }
    fmt.Printf("Recovered Secret: %s\n", recoveredSecret)
}
```

## Functions
- `Split(n, threshold int, secret []byte) ([][]byte, error)`: Splits the secret into n shares with a minimum threshold of shares needed to reconstruct the secret.
- `Combine(shares [][]byte) ([]byte, error)`: Combines the given shares to reconstruct the secret.
