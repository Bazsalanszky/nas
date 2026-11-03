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

type SecurityModeComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.SecurityModeCompleteMessageIdentity
	*nasType.IMEISV
	*nasType.NASMessageContainer
	NonIMEISVPEI *nasType.MobileIdentity5GS
}

func NewSecurityModeComplete(iei uint8) (securityModeComplete *SecurityModeComplete) {
	securityModeComplete = &SecurityModeComplete{}
	return securityModeComplete
}

const (
	SecurityModeCompleteIMEISVType              uint8 = 0x77
	SecurityModeCompleteNASMessageContainerType uint8 = 0x71
	SecurityModeCompleteNonIMEISVPEIType        uint8 = 0x78
)

func (a *SecurityModeComplete) EncodeSecurityModeComplete(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet); err != nil {
		return
	}
	if a.IMEISV != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.IMEISV.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.IMEISV.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()]); err != nil {
			return
		}
	}
	if a.NASMessageContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
			return
		}
	}
	if a.NonIMEISVPEI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NonIMEISVPEI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NonIMEISVPEI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NonIMEISVPEI.Buffer); err != nil {
			return
		}
	}
}

func (a *SecurityModeComplete) DecodeSecurityModeComplete(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet); err != nil {
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
		case SecurityModeCompleteIMEISVType:
			a.IMEISV = nasType.NewIMEISV(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.IMEISV.Len); err != nil {
				return err
			}
			a.IMEISV.SetLen(a.IMEISV.GetLen())
			if a.IMEISV.GetLen() > uint16(len(a.IMEISV.Octet)) {
				a.IMEISV = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: IMEISV length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()]); err != nil {
				return err
			}
		case SecurityModeCompleteNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len); err != nil {
				return err
			}
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Buffer[:a.NASMessageContainer.GetLen()]); err != nil {
				return err
			}
		case SecurityModeCompleteNonIMEISVPEIType:
			a.NonIMEISVPEI = nasType.NewMobileIdentity5GS(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NonIMEISVPEI.Len); err != nil {
				return err
			}
			a.NonIMEISVPEI.SetLen(a.NonIMEISVPEI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NonIMEISVPEI.Buffer[:a.NonIMEISVPEI.GetLen()]); err != nil {
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
