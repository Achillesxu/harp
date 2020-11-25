// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"

	// Import Blake2b
	_ "golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ssh"
	jose "gopkg.in/square/go-jose.v2"

	"github.com/elastic/harp/pkg/sdk/types"
)

const (
	blockTypeRsaPrivateKey     = "RSA PRIVATE KEY"
	blockTypeRsaPublicKey      = "RSA PUBLIC KEY"
	blockTypeEcdsaPrivateKey   = "EC PRIVATE KEY"
	blockTypeEcdsaPublicKey    = "EC PUBLIC KEY"
	blockTypePrivateKey        = "PRIVATE KEY"
	blockTypePublicKey         = "PUBLIC KEY"
	blockTypeOpenSSHPrivateKey = "OPENSSH PRIVATE KEY"
)

// ToJWK encodes given key using JWK.
func ToJWK(key interface{}) (string, error) {
	// Check key
	if types.IsNil(key) {
		return "", fmt.Errorf("unable to encode nil key")
	}

	// Wrap key
	keyWrapper := jose.JSONWebKey{Key: key, KeyID: ""}

	// Generate thumbprint
	thumb, err := keyWrapper.Thumbprint(crypto.BLAKE2b_256)
	if err != nil {
		return "", err
	}

	// Assign thumbprint
	keyWrapper.KeyID = base64.URLEncoding.EncodeToString(thumb)

	// Marshal private as JSON
	payload, err := keyWrapper.MarshalJSON()
	if err != nil {
		return "", err
	}

	// No error
	return string(payload), nil
}

// ToPEM encodes the given key using PEM.
func ToPEM(key interface{}) (string, error) {
	var pemData []byte
	switch k := key.(type) {
	// Private keys ------------------------------------------------------------
	case *rsa.PrivateKey:
		pemData = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypeRsaPrivateKey,
				Bytes: x509.MarshalPKCS1PrivateKey(k),
			},
		)
	case *ecdsa.PrivateKey:
		privkeyBytes, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return "", err
		}
		pemData = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypeEcdsaPrivateKey,
				Bytes: privkeyBytes,
			},
		)
	case ed25519.PrivateKey:
		privkeyBytes, err := x509.MarshalPKCS8PrivateKey(k)
		if err != nil {
			return "", err
		}
		pemData = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypePrivateKey,
				Bytes: privkeyBytes,
			},
		)

	// Private keys ------------------------------------------------------------
	case *rsa.PublicKey:
		pubkeyBytes, err := x509.MarshalPKIXPublicKey(k)
		if err != nil {
			return "", err
		}
		pemData = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypeRsaPublicKey,
				Bytes: pubkeyBytes,
			},
		)

	case *ecdsa.PublicKey:
		pubkeyBytes, err := x509.MarshalPKIXPublicKey(k)
		if err != nil {
			return "", err
		}
		pemData = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypeEcdsaPublicKey,
				Bytes: pubkeyBytes,
			},
		)

	case ed25519.PublicKey:
		pubkeyBytes, err := x509.MarshalPKIXPublicKey(k)
		if err != nil {
			return "", err
		}
		pemData = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypePublicKey,
				Bytes: pubkeyBytes,
			},
		)
	default:
		return "", fmt.Errorf("given key type is not supported")
	}

	return string(pemData), nil
}

// EncryptPEM returns an encrypted PEM block using the given passphrase.
func EncryptPEM(pemData, passphrase string) (string, error) {
	// Check passphrase
	if len(passphrase) < 32 {
		return "", fmt.Errorf("passphrase must contains more than 32 characters, usage of a diceware passphrase is recommended")
	}

	// Decode PEM
	inputBlock, _ := pem.Decode([]byte(pemData))
	if inputBlock == nil {
		return "", fmt.Errorf("unable to parse input PEM")
	}

	// Generate new block
	block, err := x509.EncryptPEMBlock(
		rand.Reader,
		inputBlock.Type,
		inputBlock.Bytes,
		[]byte(passphrase),
		x509.PEMCipherAES256,
	)
	if err != nil {
		return "", fmt.Errorf("unable to encrypt PEM: %w", err)
	}

	// Encode output PEM
	outPem := pem.EncodeToMemory(block)

	// No error
	return string(outPem), nil
}

// ToSSH encodes the given key as SSH key.
func ToSSH(key interface{}) (string, error) {
	var result []byte

	switch k := key.(type) {
	// Private keys ------------------------------------------------------------
	case *rsa.PrivateKey:
		result = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypeRsaPrivateKey,
				Bytes: x509.MarshalPKCS1PrivateKey(k),
			},
		)
	case *ecdsa.PrivateKey:
		privkeyBytes, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return "", err
		}
		result = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypeEcdsaPrivateKey,
				Bytes: privkeyBytes,
			},
		)
	case ed25519.PrivateKey:
		privkeyBytes := marshalED25519PrivateKey(k)
		result = pem.EncodeToMemory(
			&pem.Block{
				Type:  blockTypeOpenSSHPrivateKey,
				Bytes: privkeyBytes,
			},
		)

	// Private keys ------------------------------------------------------------
	case *rsa.PublicKey, *ecdsa.PublicKey, ed25519.PublicKey:
		pubKey, err := ssh.NewPublicKey(k)
		if err != nil {
			return "", fmt.Errorf("unable to convert key as ssh public key: %w", err)
		}
		result = ssh.MarshalAuthorizedKey(pubKey)

	default:
		return "", fmt.Errorf("given key type is not supported")
	}

	// No error
	return string(result), nil
}

// -----------------------------------------------------------------------------

// Writes ed25519 private keys into the new OpenSSH private key format.
func marshalED25519PrivateKey(key ed25519.PrivateKey) []byte {
	// Add our key header (followed by a null byte)
	magic := append([]byte("openssh-key-v1"), 0)

	var w struct {
		CipherName   string
		KdfName      string
		KdfOpts      string
		NumKeys      uint32
		PubKey       []byte
		PrivKeyBlock []byte
	}

	// Fill out the private key fields
	pk1 := struct {
		Check1  uint32
		Check2  uint32
		Keytype string
		Pub     []byte
		Priv    []byte
		Comment string
		Pad     []byte `ssh:"rest"`
	}{}

	// Set our check ints
	ci, err := randUInt32()
	if err != nil {
		panic(err)
	}

	pk1.Check1 = ci
	pk1.Check2 = ci

	// Set our key type
	pk1.Keytype = ssh.KeyAlgoED25519

	// Add the pubkey to the optionally-encrypted block
	pk, ok := key.Public().(ed25519.PublicKey)
	if !ok {
		return nil
	}
	pubKey := []byte(pk)
	pk1.Pub = pubKey

	// Add our private key
	pk1.Priv = []byte(key)

	// Might be useful to put something in here at some point
	pk1.Comment = ""

	// Add some padding to match the encryption block size within PrivKeyBlock (without Pad field)
	// 8 doesn't match the documentation, but that's what ssh-keygen uses for unencrypted keys. *shrug*
	bs := 8
	blockLen := len(ssh.Marshal(pk1))
	padLen := (bs - (blockLen % bs)) % bs
	pk1.Pad = make([]byte, padLen)

	// Padding is a sequence of bytes like: 1, 2, 3...
	for i := 0; i < padLen; i++ {
		pk1.Pad[i] = byte(i + 1)
	}

	// Generate the pubkey prefix "\0\0\0\nssh-ed25519\0\0\0 "
	prefix := []byte{0x0, 0x0, 0x0, 0x0b}
	prefix = append(prefix, []byte(ssh.KeyAlgoED25519)...)
	prefix = append(prefix, []byte{0x0, 0x0, 0x0, 0x20}...)

	// Only going to support unencrypted keys for now
	w.CipherName = "none"
	w.KdfName = "none"
	w.KdfOpts = ""
	w.NumKeys = 1
	// nolint
	w.PubKey = append(prefix, pubKey...)
	w.PrivKeyBlock = ssh.Marshal(pk1)

	magic = append(magic, ssh.Marshal(w)...)

	return magic
}

func randUInt32() (uint32, error) {
	var buf [4]byte
	_, err := rand.Read(buf[:])
	inputInt := big.NewInt(0).SetBytes(buf[:])
	return uint32(inputInt.Uint64()), err
}
