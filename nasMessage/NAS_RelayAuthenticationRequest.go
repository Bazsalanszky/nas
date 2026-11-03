// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RelayAuthenticationRequest 8.2.37
type RelayAuthenticationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RELAYAUTHENTICATIONREQUESTMessageIdentity
	nasType.ProSeRelayTransactionIdentity
	nasType.EAPMessage
}

func NewRelayAuthenticationRequest(iei uint8) (relayAuthenticationRequest *RelayAuthenticationRequest) {
	relayAuthenticationRequest = &RelayAuthenticationRequest{}
	return relayAuthenticationRequest
}

func (a *RelayAuthenticationRequest) EncodeRelayAuthenticationRequest(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.RELAYAUTHENTICATIONREQUESTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return
	}
}

func (a *RelayAuthenticationRequest) DecodeRelayAuthenticationRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RELAYAUTHENTICATIONREQUESTMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
		return err
	}
	a.SetLen(a.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return err
	}
	return nil
}
