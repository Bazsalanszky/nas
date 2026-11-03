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

type AuthenticationResult struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationResultMessageIdentity
	nasType.SpareHalfOctetAndNgksi
	nasType.EAPMessage
	*nasType.ABBA
	*nasType.MasterSessionKey
}

func NewAuthenticationResult(iei uint8) (authenticationResult *AuthenticationResult) {
	authenticationResult = &AuthenticationResult{}
	return authenticationResult
}

const (
	AuthenticationResultABBAType             uint8 = 0x38
	AuthenticationResultMasterSessionKeyType uint8 = 0x55
)

func (a *AuthenticationResult) EncodeAuthenticationResult(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.AuthenticationResultMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer); err != nil {
		return
	}
	if a.ABBA != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ABBA.Buffer); err != nil {
			return
		}
	}
	if a.MasterSessionKey != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MasterSessionKey.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MasterSessionKey.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.MasterSessionKey.Buffer); err != nil {
			return
		}
	}
}

func (a *AuthenticationResult) DecodeAuthenticationResult(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationResultMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
		return err
	}
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer); err != nil {
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
		case AuthenticationResultABBAType:
			a.ABBA = nasType.NewABBA(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ABBA.Len); err != nil {
				return err
			}
			a.ABBA.SetLen(a.ABBA.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ABBA.Buffer[:a.ABBA.GetLen()]); err != nil {
				return err
			}
		case AuthenticationResultMasterSessionKeyType:
			a.MasterSessionKey = nasType.NewMasterSessionKey(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MasterSessionKey.Len); err != nil {
				return err
			}
			a.MasterSessionKey.SetLen(a.MasterSessionKey.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.MasterSessionKey.Buffer[:a.MasterSessionKey.GetLen()]); err != nil {
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
