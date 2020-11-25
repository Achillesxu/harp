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

package secret

import (
	"encoding/asn1"
	"fmt"
)

// Pack a secret value.
func Pack(value interface{}) ([]byte, error) {
	// Encode the payload
	// JSON is not used to prevents Base64 double encoding.
	payload, err := asn1.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("unable to pack secret value: %w", err)
	}

	return payload, nil
}

// Unpack a secret value.
func Unpack(in []byte, out interface{}) error {
	// Decode the value
	if _, err := asn1.Unmarshal(in, out); err != nil {
		return fmt.Errorf("unable to upack secret value: %w", err)
	}

	return nil
}
