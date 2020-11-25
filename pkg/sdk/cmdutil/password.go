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

package cmdutil

import (
	"fmt"
	"syscall"

	"github.com/awnumar/memguard"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/elastic/harp/pkg/sdk/security"
)

// ReadSecret reads password from Stdin and returns a lockedbuffer.
func ReadSecret(prompt string, confirmation bool) (*memguard.LockedBuffer, error) {
	var (
		err             error
		password        []byte
		passwordConfirm []byte
	)
	defer memguard.WipeBytes(password)
	defer memguard.WipeBytes(passwordConfirm)

	// Ask to password
	fmt.Printf("%s: ", prompt)
	//nolint:unconvert // stdin doesn't share same type on each platform
	password, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, fmt.Errorf("unable to read secret")
	}

	fmt.Print("\n")

	// Check if confirmation is requested
	if !confirmation {
		// Return locked buffer
		return memguard.NewBufferFromBytes(password), nil
	}

	fmt.Printf("%s (confirmation): ", prompt)
	//nolint:unconvert // stdin doesn't share same type on each platform
	passwordConfirm, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, fmt.Errorf("unable to read secret confirmation")
	}

	fmt.Print("\n")

	// Compare if equal
	if !security.SecureCompare(password, passwordConfirm) {
		return nil, fmt.Errorf("passphrase doesn't match")
	}

	// Return locked buffer
	return memguard.NewBufferFromBytes(password), nil
}
