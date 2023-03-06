package CreateKeystore

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeystore() {
	ks := keystore.NewKeyStore(".", keystore.StandardScryptN, keystore.StandardScryptP)
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
}
