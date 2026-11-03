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

type PDUSessionModificationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONMODIFICATIONREQUESTMessageIdentity
	*nasType.Capability5GSM
	*nasType.Cause5GSM
	*nasType.MaximumNumberOfSupportedPacketFilters
	*nasType.AlwaysonPDUSessionRequested
	*nasType.IntegrityProtectionMaximumDataRate
	*nasType.RequestedQosRules
	*nasType.RequestedQosFlowDescriptions
	*nasType.MappedEPSBearerContexts
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionModificationRequest(iei uint8) (pDUSessionModificationRequest *PDUSessionModificationRequest) {
	pDUSessionModificationRequest = &PDUSessionModificationRequest{}
	return pDUSessionModificationRequest
}

const (
	PDUSessionModificationRequestCapability5GSMType                        uint8 = 0x28
	PDUSessionModificationRequestCause5GSMType                             uint8 = 0x59
	PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType uint8 = 0x55
	PDUSessionModificationRequestAlwaysonPDUSessionRequestedType           uint8 = 0x0B
	PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType    uint8 = 0x13
	PDUSessionModificationRequestRequestedQosRulesType                     uint8 = 0x7A
	PDUSessionModificationRequestRequestedQosFlowDescriptionsType          uint8 = 0x79
	PDUSessionModificationRequestMappedEPSBearerContextsType               uint8 = 0x7F
	PDUSessionModificationRequestExtendedProtocolConfigurationOptionsType  uint8 = 0x7B
)

func (a *PDUSessionModificationRequest) EncodePDUSessionModificationRequest(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONREQUESTMessageIdentity.Octet); err != nil {
		return
	}
	if a.Capability5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
			return
		}
	}
	if a.Cause5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
			return
		}
	}
	if a.MaximumNumberOfSupportedPacketFilters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.MaximumNumberOfSupportedPacketFilters.Octet); err != nil {
			return
		}
	}
	if a.AlwaysonPDUSessionRequested != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.AlwaysonPDUSessionRequested.Octet); err != nil {
			return
		}
	}
	if a.IntegrityProtectionMaximumDataRate != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.IntegrityProtectionMaximumDataRate.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.IntegrityProtectionMaximumDataRate.Octet); err != nil {
			return
		}
	}
	if a.RequestedQosRules != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosRules.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosRules.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RequestedQosRules.Buffer); err != nil {
			return
		}
	}
	if a.RequestedQosFlowDescriptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosFlowDescriptions.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosFlowDescriptions.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RequestedQosFlowDescriptions.Buffer); err != nil {
			return
		}
	}
	if a.MappedEPSBearerContexts != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Buffer); err != nil {
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
}

func (a *PDUSessionModificationRequest) DecodePDUSessionModificationRequest(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONREQUESTMessageIdentity.Octet); err != nil {
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
		case PDUSessionModificationRequestCapability5GSMType:
			a.Capability5GSM = nasType.NewCapability5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Capability5GSM.Len); err != nil {
				return err
			}
			a.Capability5GSM.SetLen(a.Capability5GSM.GetLen())
			if a.Capability5GSM.GetLen() > uint8(len(a.Capability5GSM.Octet)) {
				a.Capability5GSM = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: Capability5GSM length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationRequestCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
				return err
			}
		case PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType:
			a.MaximumNumberOfSupportedPacketFilters = nasType.NewMaximumNumberOfSupportedPacketFilters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MaximumNumberOfSupportedPacketFilters.Octet); err != nil {
				return err
			}
		case PDUSessionModificationRequestAlwaysonPDUSessionRequestedType:
			a.AlwaysonPDUSessionRequested = nasType.NewAlwaysonPDUSessionRequested(ieiN)
			a.AlwaysonPDUSessionRequested.Octet = ieiN
		case PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType:
			a.IntegrityProtectionMaximumDataRate = nasType.NewIntegrityProtectionMaximumDataRate(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.IntegrityProtectionMaximumDataRate.Octet); err != nil {
				return err
			}
		case PDUSessionModificationRequestRequestedQosRulesType:
			a.RequestedQosRules = nasType.NewRequestedQosRules(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedQosRules.Len); err != nil {
				return err
			}
			a.RequestedQosRules.SetLen(a.RequestedQosRules.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RequestedQosRules.Buffer[:a.RequestedQosRules.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationRequestRequestedQosFlowDescriptionsType:
			a.RequestedQosFlowDescriptions = nasType.NewRequestedQosFlowDescriptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedQosFlowDescriptions.Len); err != nil {
				return err
			}
			a.RequestedQosFlowDescriptions.SetLen(a.RequestedQosFlowDescriptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RequestedQosFlowDescriptions.Buffer[:a.RequestedQosFlowDescriptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationRequestMappedEPSBearerContextsType:
			a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Len); err != nil {
				return err
			}
			a.MappedEPSBearerContexts.SetLen(a.MappedEPSBearerContexts.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer[:a.MappedEPSBearerContexts.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationRequestExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return err
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
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
