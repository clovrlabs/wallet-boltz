//go:build !testnet && !regtest
// +build !testnet,!regtest

package boltz

import "github.com/btcsuite/btcd/chaincfg"

const (
	apiURL = "https://boltz.exchange/api"
)

var (
	chain = &chaincfg.MainNetParams
)
