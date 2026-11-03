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

type PDUSessionModificationReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONMODIFICATIONREJECTMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.Fivegsmcongestionreattemptindicator
	*nasType.ReAttemptIndicator
}

func NewPDUSessionModificationReject(iei uint8) (pDUSessionModificationReject *PDUSessionModificationReject) {
	pDUSessionModificationReject = &PDUSessionModificationReject{}
	return pDUSessionModificationReject
}

const (
	PDUSessionModificationRejectBackoffTimerValueType                    uint8 = 0x37
	PDUSessionModificationRejectExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionModificationRejectFivegsmcongestionreattemptindicatorType  uint8 = 0x61
	PDUSessionModificationRejectReAttemptIndicatorType                   uint8 = 0x1D
)

func (a *PDUSessionModificationReject) EncodePDUSessionModificationReject(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONREJECTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
		return
	}
	if a.BackoffTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
			return
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
			return
		}
	}
	if a.Fivegsmcongestionreattemptindicator != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Fivegsmcongestionreattemptindicator.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Fivegsmcongestionreattemptindicator.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Octet); err != nil {
			return
		}
	}
	if a.ReAttemptIndicator != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ReAttemptIndicator.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ReAttemptIndicator.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ReAttemptIndicator.Octet); err != nil {
			return
		}
	}
}

func (a *PDUSessionModificationReject) DecodePDUSessionModificationReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONREJECTMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
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
		case PDUSessionModificationRejectBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len); err != nil {
				return err
			}
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
				return err
			}
		case PDUSessionModificationRejectExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return err
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationRejectFivegsmcongestionreattemptindicatorType:
			a.Fivegsmcongestionreattemptindicator = nasType.NewFivegsmcongestionreattemptindicator(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Len); err != nil {
				return err
			}
			a.Fivegsmcongestionreattemptindicator.SetLen(a.Fivegsmcongestionreattemptindicator.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Octet); err != nil {
				return err
			}
		case PDUSessionModificationRejectReAttemptIndicatorType:
			a.ReAttemptIndicator = nasType.NewReAttemptIndicator(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ReAttemptIndicator.Len); err != nil {
				return err
			}
			a.ReAttemptIndicator.SetLen(a.ReAttemptIndicator.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.ReAttemptIndicator.Octet); err != nil {
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
