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

type DeregistrationAcceptUEOriginatingDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationAcceptMessageIdentity
}

func NewDeregistrationAcceptUEOriginatingDeregistration(iei uint8) (deregistrationAcceptUEOriginatingDeregistration *DeregistrationAcceptUEOriginatingDeregistration) {
	deregistrationAcceptUEOriginatingDeregistration = &DeregistrationAcceptUEOriginatingDeregistration{}
	return deregistrationAcceptUEOriginatingDeregistration
}

func (a *DeregistrationAcceptUEOriginatingDeregistration) EncodeDeregistrationAcceptUEOriginatingDeregistration(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.DeregistrationAcceptMessageIdentity.Octet); err != nil {
		return
	}
}

func (a *DeregistrationAcceptUEOriginatingDeregistration) DecodeDeregistrationAcceptUEOriginatingDeregistration(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.DeregistrationAcceptMessageIdentity.Octet); err != nil {
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
