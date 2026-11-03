// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type IdentityResponse struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.IdentityResponseMessageIdentity
	nasType.MobileIdentity
}

func NewIdentityResponse(iei uint8) (identityResponse *IdentityResponse) {
	identityResponse = &IdentityResponse{}
	return identityResponse
}

func (a *IdentityResponse) EncodeIdentityResponse(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.IdentityResponseMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return
	}
}

func (a *IdentityResponse) DecodeIdentityResponse(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.IdentityResponseMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
		return err
	}
	a.SetLen(a.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return err
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return err
		}
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		switch tmpIeiN {
		default:
			if err := skipUnknownIE(buffer, ieiN); err != nil {
				return err
			}
		}
	}
	return nil
}
