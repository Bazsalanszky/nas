// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/omec-project/nas/v2/nasType"
)

// NetworkSliceSpecificAuthenticationComplete 8.2.32
type NetworkSliceSpecificAuthenticationComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity
	nasType.SNSSAI
	nasType.EAPMessage
}

func NewNetworkSliceSpecificAuthenticationComplete(iei uint8) (networkSliceSpecificAuthenticationComplete *NetworkSliceSpecificAuthenticationComplete) {
	networkSliceSpecificAuthenticationComplete = &NetworkSliceSpecificAuthenticationComplete{}
	return networkSliceSpecificAuthenticationComplete
}

func (a *NetworkSliceSpecificAuthenticationComplete) EncodeNetworkSliceSpecificAuthenticationComplete(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity.Octet); err != nil {
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

func (a *NetworkSliceSpecificAuthenticationComplete) DecodeNetworkSliceSpecificAuthenticationComplete(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NETWORKSLICESPECIFICAUTHENTICATIONCOMPLETEMessageIdentity.Octet); err != nil {
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
