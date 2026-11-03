// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type ServiceReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceRejectMessageIdentity
	nasType.Cause5GMM
	*nasType.PDUSessionStatus
	*nasType.T3346Value
	*nasType.EAPMessage
	*nasType.T3448Value
	*nasType.CAGInformationList
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ExtendedCAGInformationList
	*nasType.LowerBoundTimerValue
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
}

func NewServiceReject(iei uint8) (serviceReject *ServiceReject) {
	serviceReject = &ServiceReject{}
	return serviceReject
}

const (
	ServiceRejectPDUSessionStatusType              uint8 = 0x50
	ServiceRejectT3346ValueType                    uint8 = 0x5F
	ServiceRejectEAPMessageType                    uint8 = 0x78
	ServiceRejectT3448ValueType                    uint8 = 0x6B
	ServiceRejectCAGInformationListType            uint8 = 0x75
	ServiceRejectDisasterReturnWaitRangeType       uint8 = 0x2C
	ServiceRejectExtendedCAGInformationListType    uint8 = 0x71
	ServiceRejectLowerBoundTimerValueType          uint8 = 0x3A
	ServiceRejectForbiddenTAIRoamingType           uint8 = 0x1D
	ServiceRejectForbiddenTAIRegionalProvisionType uint8 = 0x1E
)

func (a *ServiceReject) EncodeServiceReject(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ServiceRejectMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return
	}
	if a.PDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionStatus.Buffer); err != nil {
			return
		}
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
	if a.T3448Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3448Value.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3448Value.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.T3448Value.Octet); err != nil {
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
}

func (a *ServiceReject) DecodeServiceReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ServiceRejectMessageIdentity.Octet); err != nil {
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
		case ServiceRejectPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return err
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()]); err != nil {
				return err
			}
		case ServiceRejectT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len); err != nil {
				return err
			}
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
				return err
			}
		case ServiceRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case ServiceRejectT3448ValueType:
			a.T3448Value = nasType.NewT3448Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3448Value.Len); err != nil {
				return err
			}
			a.T3448Value.SetLen(a.T3448Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3448Value.Octet); err != nil {
				return err
			}
		case ServiceRejectCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len); err != nil {
				return err
			}
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()]); err != nil {
				return err
			}
		case ServiceRejectDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.DisasterReturnWaitRange.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:l]); err != nil {
				return err
			}
		case ServiceRejectExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len); err != nil {
				return err
			}
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()]); err != nil {
				return err
			}
		case ServiceRejectLowerBoundTimerValueType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len); err != nil {
				return err
			}
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet); err != nil {
				return err
			}
		case ServiceRejectForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()]); err != nil {
				return err
			}
		case ServiceRejectForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()]); err != nil {
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
