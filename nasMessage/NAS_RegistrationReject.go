// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type RegistrationReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationRejectMessageIdentity
	nasType.Cause5GMM
	*nasType.T3346Value
	*nasType.T3502Value
	*nasType.EAPMessage
	*nasType.RejectedNSSAI
	*nasType.CAGInformationList
	*nasType.ExtendedRejectedNSSAI
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ExtendedCAGInformationList
	*nasType.LowerBoundTimerValue
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
	*nasType.N3IWFIdentifier
	*nasType.TNANInformation
}

func NewRegistrationReject(iei uint8) (registrationReject *RegistrationReject) {
	registrationReject = &RegistrationReject{}
	return registrationReject
}

const (
	RegistrationRejectT3346ValueType                    uint8 = 0x5F
	RegistrationRejectT3502ValueType                    uint8 = 0x16
	RegistrationRejectEAPMessageType                    uint8 = 0x78
	RegistrationRejectRejectedNSSAIType                 uint8 = 0x69
	RegistrationRejectCAGInformationListType            uint8 = 0x75
	RegistrationRejectExtendedRejectedNSSAIType         uint8 = 0x68
	RegistrationRejectDisasterReturnWaitRangeType       uint8 = 0x2C
	RegistrationRejectExtendedCAGInformationListType    uint8 = 0x71
	RegistrationRejectLowerBoundTimerValueType          uint8 = 0x3A
	RegistrationRejectForbiddenTAIRoamingType           uint8 = 0x1D
	RegistrationRejectForbiddenTAIRegionalProvisionType uint8 = 0x1E
	RegistrationRejectN3IWFIdentifierType               uint8 = 0x3E
	RegistrationRejectTNANInformationType               uint8 = 0x4D
)

func (a *RegistrationReject) EncodeRegistrationReject(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return
	}
	if a.T3346Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
			return
		}
	}
	if a.T3502Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3502Value.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3502Value.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.T3502Value.Octet); err != nil {
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
	if a.RejectedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RejectedNSSAI.Buffer); err != nil {
			return
		}
	}
	if a.CAGInformationList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.CAGInformationList.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.CAGInformationList.Buffer); err != nil {
			return
		}
	}
	if a.ExtendedRejectedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.ExtendedRejectedNSSAI.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:uint8(a.ExtendedRejectedNSSAI.GetLen())]); err != nil {
			return
		}
	}
	if a.DisasterReturnWaitRange != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.DisasterReturnWaitRange.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.DisasterReturnWaitRange.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:uint8(a.DisasterReturnWaitRange.GetLen())]); err != nil {
			return
		}
	}
	if a.ExtendedCAGInformationList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedCAGInformationList.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedCAGInformationList.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Buffer); err != nil {
			return
		}
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
	if a.ForbiddenTAIRoaming != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Buffer); err != nil {
			return
		}
	}
	if a.ForbiddenTAIRegionalProvision != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Buffer); err != nil {
			return
		}
	}
	if a.N3IWFIdentifier != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.N3IWFIdentifier.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.N3IWFIdentifier.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.N3IWFIdentifier.Buffer); err != nil {
			return
		}
	}
	if a.TNANInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.TNANInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TNANInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.TNANInformation.Buffer); err != nil {
			return
		}
	}
}

func (a *RegistrationReject) DecodeRegistrationReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
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
		case RegistrationRejectT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len); err != nil {
				return err
			}
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
				return err
			}
		case RegistrationRejectT3502ValueType:
			a.T3502Value = nasType.NewT3502Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Len); err != nil {
				return err
			}
			a.T3502Value.SetLen(a.T3502Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Octet); err != nil {
				return err
			}
		case RegistrationRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case RegistrationRejectRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len); err != nil {
				return err
			}
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()]); err != nil {
				return err
			}
		case RegistrationRejectCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len); err != nil {
				return err
			}
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()]); err != nil {
				return err
			}
		case RegistrationRejectExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.ExtendedRejectedNSSAI.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:l]); err != nil {
				return err
			}
		case RegistrationRejectDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.DisasterReturnWaitRange.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:l]); err != nil {
				return err
			}
		case RegistrationRejectExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len); err != nil {
				return err
			}
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()]); err != nil {
				return err
			}
		case RegistrationRejectLowerBoundTimerValueType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len); err != nil {
				return err
			}
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet); err != nil {
				return err
			}
		case RegistrationRejectForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()]); err != nil {
				return err
			}
		case RegistrationRejectForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()]); err != nil {
				return err
			}
		case RegistrationRejectN3IWFIdentifierType:
			a.N3IWFIdentifier = nasType.NewN3IWFIdentifier(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.N3IWFIdentifier.Len); err != nil {
				return err
			}
			a.N3IWFIdentifier.SetLen(a.N3IWFIdentifier.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.N3IWFIdentifier.Buffer[:a.N3IWFIdentifier.GetLen()]); err != nil {
				return err
			}
		case RegistrationRejectTNANInformationType:
			a.TNANInformation = nasType.NewTNANInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.TNANInformation.Len); err != nil {
				return err
			}
			a.TNANInformation.SetLen(a.TNANInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.TNANInformation.Buffer[:a.TNANInformation.GetLen()]); err != nil {
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
