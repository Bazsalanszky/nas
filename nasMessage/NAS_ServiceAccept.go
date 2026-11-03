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

type ServiceAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceAcceptMessageIdentity
	*nasType.PDUSessionStatus
	*nasType.PDUSessionReactivationResult
	*nasType.PDUSessionReactivationResultErrorCause
	*nasType.EAPMessage
	*nasType.T3448Value
	*nasType.FiveGSAdditionalRequestResult
	ForbiddenTAIRoaming           *nasType.TAIList
	ForbiddenTAIRegionalProvision *nasType.TAIList
}

func NewServiceAccept(iei uint8) (serviceAccept *ServiceAccept) {
	serviceAccept = &ServiceAccept{}
	return serviceAccept
}

const (
	ServiceAcceptPDUSessionStatusType                       uint8 = 0x50
	ServiceAcceptPDUSessionReactivationResultType           uint8 = 0x26
	ServiceAcceptPDUSessionReactivationResultErrorCauseType uint8 = 0x72
	ServiceAcceptEAPMessageType                             uint8 = 0x78
	ServiceAcceptT3448ValueType                             uint8 = 0x6B
	ServiceAcceptFiveGSAdditionalRequestResultType          uint8 = 0x34
	ServiceAcceptForbiddenTAIRoamingType                    uint8 = 0x1D
	ServiceAcceptForbiddenTAIRegionalProvisionType          uint8 = 0x1E
)

func (a *ServiceAccept) EncodeServiceAccept(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ServiceAcceptMessageIdentity.Octet); err != nil {
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
	if a.PDUSessionReactivationResult != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResult.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Buffer); err != nil {
			return
		}
	}
	if a.PDUSessionReactivationResultErrorCause != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Buffer); err != nil {
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
	if a.FiveGSAdditionalRequestResult != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.FiveGSAdditionalRequestResult.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.Buffer[:uint8(a.FiveGSAdditionalRequestResult.GetLen())]); err != nil {
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

func (a *ServiceAccept) DecodeServiceAccept(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ServiceAcceptMessageIdentity.Octet); err != nil {
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
		case ServiceAcceptPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return err
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()]); err != nil {
				return err
			}
		case ServiceAcceptPDUSessionReactivationResultType:
			a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Len); err != nil {
				return err
			}
			a.PDUSessionReactivationResult.SetLen(a.PDUSessionReactivationResult.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer[:a.PDUSessionReactivationResult.GetLen()]); err != nil {
				return err
			}
		case ServiceAcceptPDUSessionReactivationResultErrorCauseType:
			a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Len); err != nil {
				return err
			}
			a.PDUSessionReactivationResultErrorCause.SetLen(a.PDUSessionReactivationResultErrorCause.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer[:a.PDUSessionReactivationResultErrorCause.GetLen()]); err != nil {
				return err
			}
		case ServiceAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case ServiceAcceptT3448ValueType:
			a.T3448Value = nasType.NewT3448Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3448Value.Len); err != nil {
				return err
			}
			a.T3448Value.SetLen(a.T3448Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3448Value.Octet); err != nil {
				return err
			}
		case ServiceAcceptFiveGSAdditionalRequestResultType:
			a.FiveGSAdditionalRequestResult = nasType.NewFiveGSAdditionalRequestResult(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.FiveGSAdditionalRequestResult.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.Buffer[:l]); err != nil {
				return err
			}
		case ServiceAcceptForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()]); err != nil {
				return err
			}
		case ServiceAcceptForbiddenTAIRegionalProvisionType:
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
