// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/omec-project/nas/v2/nasType"
)

// NetworkSliceSpecificAuthenticationCommand 8.2.31
type NetworkSliceSpecificAuthenticationCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity
	nasType.SNSSAI
	nasType.EAPMessage
}

func NewNetworkSliceSpecificAuthenticationCommand(iei uint8) (networkSliceSpecificAuthenticationCommand *NetworkSliceSpecificAuthenticationCommand) {
	networkSliceSpecificAuthenticationCommand = &NetworkSliceSpecificAuthenticationCommand{}
	return networkSliceSpecificAuthenticationCommand
}

func (a *NetworkSliceSpecificAuthenticationCommand) EncodeNetworkSliceSpecificAuthenticationCommand(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return
	}
}

func (a *NetworkSliceSpecificAuthenticationCommand) DecodeNetworkSliceSpecificAuthenticationCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len); err != nil {
		return err
	}
	if a.SNSSAI.GetLen() > uint8(len(a.SNSSAI.Octet)) {
		a.SNSSAI.SetLen(0) // SNSSAI is mandatory (not a pointer) - reset it so a later Encode cannot re-panic on it
		return fmt.Errorf("%w: SNSSAI length exceeds maximum", ErrMalformedIE)
	}
	if err := binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
		return err
	}
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return err
	}
	return nil
}
