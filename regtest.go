//go:build regtest
// +build regtest

package boltz

import "github.com/btcsuite/btcd/chaincfg"

const (
	apiURL = "http://localhost:9003"
)

var (
	chain = &chaincfg.RegressionNetParams
)
