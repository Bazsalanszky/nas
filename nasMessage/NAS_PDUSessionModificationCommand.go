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

type PDUSessionModificationCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity
	*nasType.Cause5GSM
	*nasType.SessionAMBR
	*nasType.RQTimerValue
	*nasType.AlwaysonPDUSessionIndication
	*nasType.AuthorizedQosRules
	*nasType.MappedEPSBearerContexts
	*nasType.AuthorizedQosFlowDescriptions
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.ATSSSContainer
	*nasType.IPHeaderCompressionConfiguration
	*nasType.ServingPLMNRateControl
	*nasType.EthernetHeaderCompressionConfiguration
	*nasType.ReceivedMBSContainer
	*nasType.ServiceLevelAAContainer
	AlternativeSNSSAI *nasType.SNSSAI
	*nasType.N3QAI
}

func NewPDUSessionModificationCommand(iei uint8) (pDUSessionModificationCommand *PDUSessionModificationCommand) {
	pDUSessionModificationCommand = &PDUSessionModificationCommand{}
	return pDUSessionModificationCommand
}

const (
	PDUSessionModificationCommandCause5GSMType                            uint8 = 0x59
	PDUSessionModificationCommandSessionAMBRType                          uint8 = 0x2A
	PDUSessionModificationCommandRQTimerValueType                         uint8 = 0x56
	PDUSessionModificationCommandAlwaysonPDUSessionIndicationType         uint8 = 0x08
	PDUSessionModificationCommandAuthorizedQosRulesType                   uint8 = 0x7A
	PDUSessionModificationCommandMappedEPSBearerContextsType              uint8 = 0x7F
	PDUSessionModificationCommandAuthorizedQosFlowDescriptionsType        uint8 = 0x79
	PDUSessionModificationCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionModificationCommandATSSSContainerType                       uint8 = 0x77
	PDUSessionModificationCommandIPHeaderCompressionConfigurationType     uint8 = 0x66
	PDUSessionModificationCommandServingPLMNRateControlType               uint8 = 0x1E
	PDUSessionModificationCommandEthernetHeaderCompressionConfigType      uint8 = 0x1F
	PDUSessionModificationCommandReceivedMBSContainerType                 uint8 = 0x71
	PDUSessionModificationCommandServiceLevelAAContainerType              uint8 = 0x72
	PDUSessionModificationCommandAlternativeSNSSAIType                    uint8 = 0x5A
	PDUSessionModificationCommandN3QAIType                                uint8 = 0x70
)

func (a *PDUSessionModificationCommand) EncodePDUSessionModificationCommand(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity.Octet); err != nil {
		return
	}
	if a.Cause5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
			return
		}
	}
	if a.SessionAMBR != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SessionAMBR.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SessionAMBR.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SessionAMBR.Octet[:a.SessionAMBR.GetLen()]); err != nil {
			return
		}
	}
	if a.RQTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RQTimerValue.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RQTimerValue.Octet); err != nil {
			return
		}
	}
	if a.AlwaysonPDUSessionIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.AlwaysonPDUSessionIndication.Octet); err != nil {
			return
		}
	}
	if a.AuthorizedQosRules != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AuthorizedQosRules.Buffer); err != nil {
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
	if a.AuthorizedQosFlowDescriptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AuthorizedQosFlowDescriptions.Buffer); err != nil {
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
	if a.ATSSSContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ATSSSContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ATSSSContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ATSSSContainer.Buffer); err != nil {
			return
		}
	}
	if a.IPHeaderCompressionConfiguration != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.IPHeaderCompressionConfiguration.Buffer); err != nil {
			return
		}
	}
	if a.ServingPLMNRateControl != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ServingPLMNRateControl.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ServingPLMNRateControl.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ServingPLMNRateControl.Buffer); err != nil {
			return
		}
	}
	if a.EthernetHeaderCompressionConfiguration != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EthernetHeaderCompressionConfiguration.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EthernetHeaderCompressionConfiguration.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Octet); err != nil {
			return
		}
	}
	if a.ReceivedMBSContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ReceivedMBSContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ReceivedMBSContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ReceivedMBSContainer.Buffer); err != nil {
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
	if a.N3QAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.N3QAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.N3QAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.N3QAI.Buffer); err != nil {
			return
		}
	}
}

func (a *PDUSessionModificationCommand) DecodePDUSessionModificationCommand(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity.Octet); err != nil {
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
		case PDUSessionModificationCommandCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
				return err
			}
		case PDUSessionModificationCommandSessionAMBRType:
			a.SessionAMBR = nasType.NewSessionAMBR(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SessionAMBR.Len); err != nil {
				return err
			}
			a.SessionAMBR.SetLen(a.SessionAMBR.GetLen())
			if a.SessionAMBR.GetLen() > uint8(len(a.SessionAMBR.Octet)) {
				a.SessionAMBR = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: SessionAMBR length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.SessionAMBR.Octet[:a.SessionAMBR.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandRQTimerValueType:
			a.RQTimerValue = nasType.NewRQTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RQTimerValue.Octet); err != nil {
				return err
			}
		case PDUSessionModificationCommandAlwaysonPDUSessionIndicationType:
			a.AlwaysonPDUSessionIndication = nasType.NewAlwaysonPDUSessionIndication(ieiN)
			a.AlwaysonPDUSessionIndication.Octet = ieiN
		case PDUSessionModificationCommandAuthorizedQosRulesType:
			a.AuthorizedQosRules = nasType.NewAuthorizedQosRules(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosRules.Len); err != nil {
				return err
			}
			a.AuthorizedQosRules.SetLen(a.AuthorizedQosRules.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AuthorizedQosRules.Buffer[:a.AuthorizedQosRules.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandMappedEPSBearerContextsType:
			a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Len); err != nil {
				return err
			}
			a.MappedEPSBearerContexts.SetLen(a.MappedEPSBearerContexts.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer[:a.MappedEPSBearerContexts.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandAuthorizedQosFlowDescriptionsType:
			a.AuthorizedQosFlowDescriptions = nasType.NewAuthorizedQosFlowDescriptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosFlowDescriptions.Len); err != nil {
				return err
			}
			a.AuthorizedQosFlowDescriptions.SetLen(a.AuthorizedQosFlowDescriptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.Buffer[:a.AuthorizedQosFlowDescriptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return err
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandATSSSContainerType:
			a.ATSSSContainer = nasType.NewATSSSContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ATSSSContainer.Len); err != nil {
				return err
			}
			a.ATSSSContainer.SetLen(a.ATSSSContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ATSSSContainer.Buffer[:a.ATSSSContainer.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandIPHeaderCompressionConfigurationType:
			a.IPHeaderCompressionConfiguration = nasType.NewIPHeaderCompressionConfiguration(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.IPHeaderCompressionConfiguration.Len); err != nil {
				return err
			}
			a.IPHeaderCompressionConfiguration.SetLen(a.IPHeaderCompressionConfiguration.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.Buffer[:a.IPHeaderCompressionConfiguration.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandServingPLMNRateControlType:
			a.ServingPLMNRateControl = nasType.NewServingPLMNRateControl(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServingPLMNRateControl.Len); err != nil {
				return err
			}
			a.ServingPLMNRateControl.SetLen(a.ServingPLMNRateControl.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServingPLMNRateControl.Buffer[:a.ServingPLMNRateControl.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandEthernetHeaderCompressionConfigType:
			a.EthernetHeaderCompressionConfiguration = nasType.NewEthernetHeaderCompressionConfiguration(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Len); err != nil {
				return err
			}
			a.EthernetHeaderCompressionConfiguration.SetLen(a.EthernetHeaderCompressionConfiguration.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Octet); err != nil {
				return err
			}
		case PDUSessionModificationCommandReceivedMBSContainerType:
			a.ReceivedMBSContainer = nasType.NewReceivedMBSContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ReceivedMBSContainer.Len); err != nil {
				return err
			}
			a.ReceivedMBSContainer.SetLen(a.ReceivedMBSContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ReceivedMBSContainer.Buffer[:a.ReceivedMBSContainer.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len); err != nil {
				return err
			}
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()]); err != nil {
				return err
			}
		case PDUSessionModificationCommandAlternativeSNSSAIType:
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
		case PDUSessionModificationCommandN3QAIType:
			a.N3QAI = nasType.NewN3QAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.N3QAI.Len); err != nil {
				return err
			}
			a.N3QAI.SetLen(a.N3QAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.N3QAI.Buffer[:a.N3QAI.GetLen()]); err != nil {
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
