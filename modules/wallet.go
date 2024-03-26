package modules

import (
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type WalletType struct {
	Address    string
	PrivateKey string
	Mnemonic   string
}

func Wallets(length int) (WalletType, error) {
	mnemonic, err := hdwallet.NewMnemonic(length) //length 128 for 12 words and 256 for 24 words
	returnType := WalletType{}
	if err != nil {
		returnType.Address = ""
		returnType.PrivateKey = ""
		returnType.Mnemonic = ""
		return returnType, err
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return returnType, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return returnType, err
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return returnType, err
	}

	returnType.Address = account.Address.Hex()
	returnType.PrivateKey = "0x" + privateKey
	returnType.Mnemonic = mnemonic

	return returnType, nil
}
