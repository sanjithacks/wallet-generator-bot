# Wallet generator bot

Bot that generate HD multi-coin wallet.

## Commands

- start : Wake up the bot from sleeping
- generate : Generate multi-coin wallet
- source : Shows bot source code repo

## wallet.go

```go

package main

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	mnemonic ,err := hdwallet.NewMnemonic(256)

	if err != nil {
		log.Fatal(err)
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}


	println("New mn: ", mnemonic)

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Account: ", account.Address.Hex())
	fmt.Println("Private Key: ", privateKey)

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}


	privateKey2, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}



	fmt.Println("Account: ", account.Address.Hex())
	fmt.Println("Private Key: ", privateKey2)
}

```

# How to build?

- Build Image

```sh
docker build -t imagetag .
```

- Run container

```sh
docker run -d -e BOT_TOKEN="2989822:xyzdd" IMAGE_ID
```

- Stop container

```sh
docker kill IMAGE_ID -f
```

- Delete Image

```sh
docker rmi IMAGE_ID -f
```
