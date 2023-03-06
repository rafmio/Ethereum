// https://translated.turbopages.org/proxy_u/en-ru.ru.1b7d3c8e-63d55f78-3c3e3677-74722d776562/https/ethereum.stackexchange.com/questions/63390/golang-equivalent-to-list-accounts-on-node-web3-eth-accounts
import (
    "log"

    "github.com/ethereum/go-ethereum/accounts"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/common"
)

type KeyStore struct {
    Handle *keystore.KeyStore
}

func SetUpKeyStore(kp string) *KeyStore {
    ks := &KeyStore{}
    ks.Handle = keystore.NewKeyStore(kp, keystore.LightScryptN, keystore.LightScryptP)
    return ks
}

func (ks *KeyStore) CreateNewKeys(password string) accounts.Account {
    account, err := ks.Handle.NewAccount(password)
    if err != nil {
        log.Panic(err)
    }
    return account
}

func (ks *KeyStore) GetKeysByAddress(address string) accounts.Account {

    var account accounts.Account
    var err error
    if ks.Handle.HasAddress(common.HexToAddress(address)) {
        if account, err = ks.Handle.Find(accounts.Account{Address: common.HexToAddress(address)}); err != nil {
            log.Panic(err)
        }
    }
    return account
}

func (ks *KeyStore) GetAllKeys() []accounts.Account {

    return ks.Handle.Accounts()
}
