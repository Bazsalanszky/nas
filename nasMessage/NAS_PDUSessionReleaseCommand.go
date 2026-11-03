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

type PDUSessionReleaseCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONRELEASECOMMANDMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.EAPMessage
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.Fivegsmcongestionreattemptindicator
	*nasType.SpareHalfOctetAndAccessType
	*nasType.ServiceLevelAAContainer
	AlternativeSNSSAI *nasType.SNSSAI
}

func NewPDUSessionReleaseCommand(iei uint8) (pDUSessionReleaseCommand *PDUSessionReleaseCommand) {
	pDUSessionReleaseCommand = &PDUSessionReleaseCommand{}
	return pDUSessionReleaseCommand
}

const (
	PDUSessionReleaseCommandBackoffTimerValueType                    uint8 = 0x37
	PDUSessionReleaseCommandEAPMessageType                           uint8 = 0x78
	PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionReleaseCommandFivegsmcongestionreattemptindicatorType  uint8 = 0x61
	PDUSessionReleaseCommandAccessTypeType                           uint8 = 0x0D
	PDUSessionReleaseCommandServiceLevelAAContainerType              uint8 = 0x72
	PDUSessionReleaseCommandAlternativeSNSSAIType                    uint8 = 0x5A
)

func (a *PDUSessionReleaseCommand) EncodePDUSessionReleaseCommand(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet); err != nil {
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
	if a.SpareHalfOctetAndAccessType != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndAccessType.Octet); err != nil {
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
	if a.AlternativeSNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.Octet[:a.AlternativeSNSSAI.GetLen()]); err != nil {
			return
		}
	}
}

func (a *PDUSessionReleaseCommand) DecodePDUSessionReleaseCommand(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet); err != nil {
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
		case PDUSessionReleaseCommandBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len); err != nil {
				return err
			}
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
				return err
			}
		case PDUSessionReleaseCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return err
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionReleaseCommandFivegsmcongestionreattemptindicatorType:
			a.Fivegsmcongestionreattemptindicator = nasType.NewFivegsmcongestionreattemptindicator(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Len); err != nil {
				return err
			}
			a.Fivegsmcongestionreattemptindicator.SetLen(a.Fivegsmcongestionreattemptindicator.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Fivegsmcongestionreattemptindicator.Octet); err != nil {
				return err
			}
		case PDUSessionReleaseCommandAccessTypeType:
			a.SpareHalfOctetAndAccessType = nasType.NewSpareHalfOctetAndAccessType()
			a.SpareHalfOctetAndAccessType.Octet = ieiN
		case PDUSessionReleaseCommandServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len); err != nil {
				return err
			}
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()]); err != nil {
				return err
			}
		case PDUSessionReleaseCommandAlternativeSNSSAIType:
			a.AlternativeSNSSAI = nasType.NewSNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AlternativeSNSSAI.Len); err != nil {
				return err
			}
			a.AlternativeSNSSAI.SetLen(a.AlternativeSNSSAI.GetLen())
			if a.AlternativeSNSSAI.GetLen() > uint8(len(a.AlternativeSNSSAI.Octet)) {
				a.AlternativeSNSSAI = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: AlternativeSNSSAI length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.AlternativeSNSSAI.Octet[:a.AlternativeSNSSAI.GetLen()]); err != nil {
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
