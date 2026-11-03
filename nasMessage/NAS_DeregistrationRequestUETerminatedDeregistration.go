// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type DeregistrationRequestUETerminatedDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationRequestMessageIdentity
	nasType.SpareHalfOctetAndDeregistrationType
	*nasType.Cause5GMM
	*nasType.T3346Value
	*nasType.RejectedNSSAI
	*nasType.CAGInformationList
	*nasType.ExtendedRejectedNSSAI
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ExtendedCAGInformationList
	*nasType.LowerBoundTimerValue
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
}

func NewDeregistrationRequestUETerminatedDeregistration(iei uint8) (deregistrationRequestUETerminatedDeregistration *DeregistrationRequestUETerminatedDeregistration) {
	deregistrationRequestUETerminatedDeregistration = &DeregistrationRequestUETerminatedDeregistration{}
	return deregistrationRequestUETerminatedDeregistration
}

const (
	DeregistrationRequestUETerminatedDeregistrationCause5GMMType                     uint8 = 0x58
	DeregistrationRequestUETerminatedDeregistrationT3346ValueType                    uint8 = 0x5F
	DeregistrationRequestUETerminatedDeregistrationRejectedNSSAIType                 uint8 = 0x6D
	DeregistrationRequestUETerminatedDeregistrationCAGInformationListType            uint8 = 0x75
	DeregistrationRequestUETerminatedDeregistrationExtendedRejectedNSSAIType         uint8 = 0x68
	DeregistrationRequestUETerminatedDeregistrationDisasterReturnWaitRangeType       uint8 = 0x2C
	DeregistrationRequestUETerminatedDeregistrationExtendedCAGInformationListType    uint8 = 0x71
	DeregistrationRequestUETerminatedDeregistrationLowerBoundTimerValueType          uint8 = 0x3A
	DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRoamingType           uint8 = 0x1D
	DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRegionalProvisionType uint8 = 0x1E
)

func (a *DeregistrationRequestUETerminatedDeregistration) EncodeDeregistrationRequestUETerminatedDeregistration(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.DeregistrationRequestMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndDeregistrationType.Octet); err != nil {
		return
	}
	if a.Cause5GMM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GMM.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
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
}

func (a *DeregistrationRequestUETerminatedDeregistration) DecodeDeregistrationRequestUETerminatedDeregistration(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndDeregistrationType.Octet); err != nil {
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
		case DeregistrationRequestUETerminatedDeregistrationCause5GMMType:
			a.Cause5GMM = nasType.NewCause5GMM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len); err != nil {
				return err
			}
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len); err != nil {
				return err
			}
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()]); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len); err != nil {
				return err
			}
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()]); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.ExtendedRejectedNSSAI.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:l]); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.DisasterReturnWaitRange.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:l]); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len); err != nil {
				return err
			}
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()]); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationLowerBoundTimerValueType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len); err != nil {
				return err
			}
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()]); err != nil {
				return err
			}
		case DeregistrationRequestUETerminatedDeregistrationForbiddenTAIRegionalProvisionType:
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
