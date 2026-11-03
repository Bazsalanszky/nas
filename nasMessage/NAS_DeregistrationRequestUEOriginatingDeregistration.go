// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type DeregistrationRequestUEOriginatingDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationRequestMessageIdentity
	nasType.NgksiAndDeregistrationType
	nasType.MobileIdentity5GS
	*nasType.LowerBoundTimerValue
	*nasType.NASMessageContainer
}

func NewDeregistrationRequestUEOriginatingDeregistration(iei uint8) (deregistrationRequestUEOriginatingDeregistration *DeregistrationRequestUEOriginatingDeregistration) {
	deregistrationRequestUEOriginatingDeregistration = &DeregistrationRequestUEOriginatingDeregistration{}
	return deregistrationRequestUEOriginatingDeregistration
}

const (
	DeregistrationRequestUEOriginatingDeregistrationUnavailabilityPeriodType uint8 = 0x3C
	DeregistrationRequestUEOriginatingDeregistrationNASMessageContainerType  uint8 = 0x71
)

func (a *DeregistrationRequestUEOriginatingDeregistration) EncodeDeregistrationRequestUEOriginatingDeregistration(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.DeregistrationRequestMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.NgksiAndDeregistrationType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.MobileIdentity5GS.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.MobileIdentity5GS.Buffer); err != nil {
		return
	}
	if a.LowerBoundTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LowerBoundTimerValue.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LowerBoundTimerValue.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet); err != nil {
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
		if err := binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer); err != nil {
			return
		}
	}
}

func (a *DeregistrationRequestUEOriginatingDeregistration) DecodeDeregistrationRequestUEOriginatingDeregistration(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.DeregistrationRequestMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NgksiAndDeregistrationType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Len); err != nil {
		return err
	}
	a.MobileIdentity5GS.SetLen(a.MobileIdentity5GS.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Buffer); err != nil {
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
		case DeregistrationRequestUEOriginatingDeregistrationUnavailabilityPeriodType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len); err != nil {
				return err
			}
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet); err != nil {
				return err
			}
		case DeregistrationRequestUEOriginatingDeregistrationNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len); err != nil {
				return err
			}
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer); err != nil {
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
