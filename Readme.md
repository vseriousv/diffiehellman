# GoLang Diffie-Hellman Library

This library for GoLang provides two main functionalities:

1. Generate a transport key using Alice's private key and Bob's public key.
2. Encrypt and decrypt messages using the transport key.

### Installation

To install the library, use the following command:

```shell
$ go get -u github.com/vseriousv/diffiehellman
```

### Usage

Here's an example of how to use the library:

```go
package main

import (
	"fmt"
	"github.com/vseriousv/blockchainkeys"
	"github.com/vseriousv/diffiehellman/pkg/encryption"
	"github.com/vseriousv/diffiehellman/pkg/transport_key"
	"log"
)

func getKeys() (string, string, error) {
	bc, err := blockchainkeys.NewBlockchain(blockchainkeys.Ethereum)
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	privateKey, publicKey, _, err := bc.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", err
	}

	return privateKey, publicKey, nil
}

func main() {
	// Generation ALisa pair
	privateAlisa, publicAlisa, err := getKeys()
	if err != nil {
		log.Fatal(err)
	}

	// Generation Bob pair
	privateBob, publicBob, err := getKeys()
	if err != nil {
		log.Fatal(err)
	}

	// Get transportKey for Alisa,
	// Alisa has publicKey(Bob) and privateKey(Alisa)
	transportKeyOne, err := transport_key.GetTransportKey(publicBob, privateAlisa)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("transportKeyOne", transportKeyOne)

	// Alisa's encryption message for Bob
	message := []byte("Hello Bob, I'm Alisa")
	encryptionMessage, err := encryption.Encrypt(message, []byte(transportKeyOne))
	if err != nil {
		log.Fatal(err)
	}

	// Get transportKey for Bob,
	// Bob has publicKey(Alisa) and privateKey(Bob)
	transportKeyTwo, err := transport_key.GetTransportKey(publicAlisa, privateBob)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("transportKeyTwo", transportKeyTwo)

	// Bob decrypt message from Alisa
	messageResult, err := encryption.Decrypt(encryptionMessage, []byte(transportKeyTwo))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("message from ALisa: ", string(message))
	log.Println("message to Bob :", messageResult)
}
```

## Contribution

Contributions to this library are welcome! If you'd like to contribute, please feel free to fork the repository, make your changes, and submit a pull request. We appreciate your efforts to help improve this project!