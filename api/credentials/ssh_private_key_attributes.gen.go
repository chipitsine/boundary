// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credentials

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type SshPrivateKeyAttributes struct {
	Username                 string `json:"username,omitempty"`
	PrivateKey               string `json:"private_key,omitempty"`
	PrivateKeyHmac           string `json:"private_key_hmac,omitempty"`
	PrivateKeyPassphrase     string `json:"private_key_passphrase,omitempty"`
	PrivateKeyPassphraseHmac string `json:"private_key_passphrase_hmac,omitempty"`
}

func AttributesMapToSshPrivateKeyAttributes(in map[string]interface{}) (*SshPrivateKeyAttributes, error) {
	if in == nil {
		return nil, fmt.Errorf("nil input map")
	}
	var out SshPrivateKeyAttributes
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &out,
		TagName: "json",
	})
	if err != nil {
		return nil, fmt.Errorf("error creating mapstructure decoder: %w", err)
	}
	if err := dec.Decode(in); err != nil {
		return nil, fmt.Errorf("error decoding: %w", err)
	}
	return &out, nil
}

func (pt *Credential) GetSshPrivateKeyAttributes() (*SshPrivateKeyAttributes, error) {
	if pt.Type != "sshprivatekey" {
		return nil, fmt.Errorf("asked to fetch %s-type attributes but credential is of type %s", "sshprivatekey", pt.Type)
	}
	return AttributesMapToSshPrivateKeyAttributes(pt.Attributes)
}
