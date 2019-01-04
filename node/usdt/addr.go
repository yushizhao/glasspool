package usdt

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// Init in chain.go
var Master *hdkeychain.ExtendedKey

// Reset to 1 during Init. Record it?
var ChildID uint32

// keyToAddr maps the passed private to corresponding p2pkh address.
func keyToAddr(key *btcec.PrivateKey, net *chaincfg.Params) (btcutil.Address, error) {
	serializedKey := key.PubKey().SerializeCompressed()
	pubKeyAddr, err := btcutil.NewAddressPubKey(serializedKey, net)
	if err != nil {
		return nil, err
	}
	return pubKeyAddr.AddressPubKeyHash(), nil
}

func GenerateAddress() (string, error) {
	child, err := Master.Child(ChildID)
	ChildID++
	if err != nil {
		return "error", err
	}
	key, err := child.ECPrivKey()
	if err != nil {
		return "error", err
	}
	addr, err := keyToAddr(key, NET)
	if err != nil {
		return "error", err
	}
	return addr.EncodeAddress(), nil
}
