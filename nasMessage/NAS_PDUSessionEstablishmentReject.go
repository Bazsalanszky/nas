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

type PDUSessionEstablishmentReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONESTABLISHMENTREJECTMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.AllowedSSCMode
	*nasType.EAPMessage
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.Fivegsmcongestionreattemptindicator
	*nasType.ReAttemptIndicator
	*nasType.ServiceLevelAAContainer
}

func NewPDUSessionEstablishmentReject(iei uint8) (pDUSessionEstablishmentReject *PDUSessionEstablishmentReject) {
	pDUSessionEstablishmentReject = &PDUSessionEstablishmentReject{}
	return pDUSessionEstablishmentReject
}

const (
	PDUSessionEstablishmentRejectBackoffTimerValueType                    uint8 = 0x37
	PDUSessionEstablishmentRejectAllowedSSCModeType                       uint8 = 0x0F
	PDUSessionEstablishmentRejectEAPMessageType                           uint8 = 0x78
	PDUSessionEstablishmentRejectExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionEstablishmentRejectFivegsmcongestionreattemptindicatorType  uint8 = 0x61
	PDUSessionEstablishmentRejectReAttemptIndicatorType                   uint8 = 0x1D
	PDUSessionEstablishmentRejectServiceLevelAAContainerType              uint8 = 0x72
)

func (a *PDUSessionEstablishmentReject) EncodePDUSessionEstablishmentReject(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREJECTMessageIdentity.Octet); err != nil {
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
	if a.AllowedSSCMode != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.AllowedSSCMode.Octet); err != nil {
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
		if err := binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer); err != nil {
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
		if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
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
	if a.ServiceLevelAAContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceLevelAAContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Buffer); err != nil {
			return
		}
	}
}

func (a *PDUSessionEstablishmentReject) DecodePDUSessionEstablishmentReject(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREJECTMessageIdentity.Octet); err != nil {
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
		case PDUSessionEstablishmentRejectBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len); err != nil {
				return err
			}
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
				return err
			}
		case PDUSessionEstablishmentRejectAllowedSSCModeType:
			a.AllowedSSCMode = nasType.NewAllowedSSCMode(ieiN)
			a.AllowedSSCMode.Octet = ieiN
		case PDUSessionEstablishmentRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentRejectExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return err
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentRejectFivegsmcongestionreattemptindicatorType:
			a.Fivegsmcongestionreattemptindicator = nasType.NewFivegsmcongestionreattemptindicator(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Len); err != nil {
				return err
			}
			a.Fivegsmcongestionreattemptindicator.SetLen(a.Fivegsmcongestionreattemptindicator.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Octet); err != nil {
				return err
			}
		case PDUSessionEstablishmentRejectReAttemptIndicatorType:
			a.ReAttemptIndicator = nasType.NewReAttemptIndicator(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ReAttemptIndicator.Len); err != nil {
				return err
			}
			a.ReAttemptIndicator.SetLen(a.ReAttemptIndicator.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.ReAttemptIndicator.Octet); err != nil {
				return err
			}
		case PDUSessionEstablishmentRejectServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len); err != nil {
				return err
			}
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()]); err != nil {
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
