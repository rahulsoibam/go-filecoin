package types

import (
	logging "gx/ipfs/QmcuXC5cxs79ro2cUuHs4HQ2bkDLJUYokwL8aivcX6HW3C/go-log"

	"github.com/filecoin-project/go-filecoin/address"
	wutil "github.com/filecoin-project/go-filecoin/wallet/util"
)

var log = logging.Logger("types")

// Signature is the result of a cryptographic sign operation.
type Signature = Bytes

// IsValidSignature cryptographically verifies that 'sig' is the signed hash of 'data' with
// the public key belonging to `addr`.
func IsValidSignature(data []byte, addr address.Address, sig Signature) bool {
	maybePk, err := wutil.Ecrecover(data, sig)
	if err != nil {
		// Any error returned from Ecrecover means this signature is not valid.
		log.Infof("error in signature validation: %s", err)
		return false
	}
	maybeAddrHash := address.Hash(maybePk)

	return address.NewMainnet(maybeAddrHash) == addr
}
