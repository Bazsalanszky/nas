// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/omec-project/nas/v2/nasType"
)

type AuthenticationResponse struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationResponseMessageIdentity
	*nasType.AuthenticationResponseParameter
	*nasType.EAPMessage
}

func NewAuthenticationResponse(iei uint8) (authenticationResponse *AuthenticationResponse) {
	authenticationResponse = &AuthenticationResponse{}
	return authenticationResponse
}

const (
	AuthenticationResponseAuthenticationResponseParameterType uint8 = 0x2D
	AuthenticationResponseEAPMessageType                      uint8 = 0x78
)

func (a *AuthenticationResponse) EncodeAuthenticationResponse(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.AuthenticationResponseMessageIdentity.Octet); err != nil {
		return
	}
	if a.AuthenticationResponseParameter != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationResponseParameter.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationResponseParameter.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationResponseParameter.Octet[:a.AuthenticationResponseParameter.GetLen()]); err != nil {
			return
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
			return
		}
	}
}

func (a *AuthenticationResponse) DecodeAuthenticationResponse(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationResponseMessageIdentity.Octet); err != nil {
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
		case AuthenticationResponseAuthenticationResponseParameterType:
			a.AuthenticationResponseParameter = nasType.NewAuthenticationResponseParameter(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationResponseParameter.Len); err != nil {
				return err
			}
			a.AuthenticationResponseParameter.SetLen(a.AuthenticationResponseParameter.GetLen())
			if a.AuthenticationResponseParameter.GetLen() > uint8(len(a.AuthenticationResponseParameter.Octet)) {
				a.AuthenticationResponseParameter = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: AuthenticationResponseParameter length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.AuthenticationResponseParameter.Octet[:a.AuthenticationResponseParameter.GetLen()]); err != nil {
				return err
			}
		case AuthenticationResponseEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		default:
			if err := skipUnknownIE(buffer, ieiN); err != nil {
				return err
			}
		}
	}
	return nil
}
