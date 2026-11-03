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

type PDUSessionEstablishmentAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity
	nasType.SelectedSSCModeAndSelectedPDUSessionType
	nasType.AuthorizedQosRules
	nasType.SessionAMBR
	*nasType.Cause5GSM
	*nasType.PDUAddress
	*nasType.RQTimerValue
	*nasType.SNSSAI
	*nasType.AlwaysonPDUSessionIndication
	*nasType.MappedEPSBearerContexts
	*nasType.EAPMessage
	*nasType.AuthorizedQosFlowDescriptions
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.DNN
	*nasType.FiveGSMNetworkFeatureSupport
	*nasType.ServingPLMNRateControl
	*nasType.ATSSSContainer
	*nasType.ControlPlaneOnlyIndication
	*nasType.IPHeaderCompressionConfiguration
	*nasType.EthernetHeaderCompressionConfiguration
	*nasType.ServiceLevelAAContainer
	*nasType.ReceivedMBSContainer
	*nasType.N3QAI
}

func NewPDUSessionEstablishmentAccept(iei uint8) (pDUSessionEstablishmentAccept *PDUSessionEstablishmentAccept) {
	pDUSessionEstablishmentAccept = &PDUSessionEstablishmentAccept{}
	return pDUSessionEstablishmentAccept
}

const (
	PDUSessionEstablishmentAcceptCause5GSMType                            uint8 = 0x59
	PDUSessionEstablishmentAcceptPDUAddressType                           uint8 = 0x29
	PDUSessionEstablishmentAcceptRQTimerValueType                         uint8 = 0x56
	PDUSessionEstablishmentAcceptSNSSAIType                               uint8 = 0x22
	PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType         uint8 = 0x08
	PDUSessionEstablishmentAcceptMappedEPSBearerContextsType              uint8 = 0x75
	PDUSessionEstablishmentAcceptEAPMessageType                           uint8 = 0x78
	PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType        uint8 = 0x79
	PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionEstablishmentAcceptDNNType                                  uint8 = 0x25
	PDUSessionEstablishmentAcceptFiveGSMNetworkFeatureSupportType         uint8 = 0x17
	PDUSessionEstablishmentAcceptServingPLMNRateControlType               uint8 = 0x18
	PDUSessionEstablishmentAcceptATSSSContainerType                       uint8 = 0x77
	PDUSessionEstablishmentAcceptControlPlaneOnlyIndicationType           uint8 = 0x0C
	PDUSessionEstablishmentAcceptIPHeaderCompressionConfigurationType     uint8 = 0x66
	PDUSessionEstablishmentAcceptEthernetHeaderCompressionConfigType      uint8 = 0x1F
	PDUSessionEstablishmentAcceptServiceLevelAAContainerType              uint8 = 0x72
	PDUSessionEstablishmentAcceptReceivedMBSContainerType                 uint8 = 0x71
	PDUSessionEstablishmentAcceptN3QAIType                                uint8 = 0x70
)

func (a *PDUSessionEstablishmentAccept) EncodePDUSessionEstablishmentAccept(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SelectedSSCModeAndSelectedPDUSessionType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.AuthorizedQosRules.Buffer); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SessionAMBR.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SessionAMBR.Octet); err != nil {
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
	if a.PDUAddress != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUAddress.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUAddress.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUAddress.Octet[:a.PDUAddress.GetLen()]); err != nil {
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
	if a.SNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
			return
		}
	}
	if a.AlwaysonPDUSessionIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.AlwaysonPDUSessionIndication.Octet); err != nil {
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
	if a.DNN != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.DNN.Buffer); err != nil {
			return
		}
	}
	if a.FiveGSMNetworkFeatureSupport != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.FiveGSMNetworkFeatureSupport.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.FiveGSMNetworkFeatureSupport.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.FiveGSMNetworkFeatureSupport.Buffer); err != nil {
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
	if a.ControlPlaneOnlyIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.ControlPlaneOnlyIndication.Octet); err != nil {
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

func (a *PDUSessionEstablishmentAccept) DecodePDUSessionEstablishmentAccept(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SelectedSSCModeAndSelectedPDUSessionType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosRules.Len); err != nil {
		return err
	}
	a.AuthorizedQosRules.SetLen(a.AuthorizedQosRules.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosRules.Buffer); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SessionAMBR.Len); err != nil {
		return err
	}
	a.SessionAMBR.SetLen(a.SessionAMBR.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.SessionAMBR.Octet); err != nil {
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
		case PDUSessionEstablishmentAcceptCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptPDUAddressType:
			a.PDUAddress = nasType.NewPDUAddress(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUAddress.Len); err != nil {
				return err
			}
			a.PDUAddress.SetLen(a.PDUAddress.GetLen())
			if a.PDUAddress.GetLen() > uint8(len(a.PDUAddress.Octet)) {
				a.PDUAddress = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: PDUAddress length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.PDUAddress.Octet[:a.PDUAddress.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptRQTimerValueType:
			a.RQTimerValue = nasType.NewRQTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RQTimerValue.Octet); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptSNSSAIType:
			a.SNSSAI = nasType.NewSNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len); err != nil {
				return err
			}
			a.SNSSAI.SetLen(a.SNSSAI.GetLen())
			if a.SNSSAI.GetLen() > uint8(len(a.SNSSAI.Octet)) {
				a.SNSSAI = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: SNSSAI length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType:
			a.AlwaysonPDUSessionIndication = nasType.NewAlwaysonPDUSessionIndication(ieiN)
			a.AlwaysonPDUSessionIndication.Octet = ieiN
		case PDUSessionEstablishmentAcceptMappedEPSBearerContextsType:
			a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Len); err != nil {
				return err
			}
			a.MappedEPSBearerContexts.SetLen(a.MappedEPSBearerContexts.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer[:a.MappedEPSBearerContexts.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType:
			a.AuthorizedQosFlowDescriptions = nasType.NewAuthorizedQosFlowDescriptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosFlowDescriptions.Len); err != nil {
				return err
			}
			a.AuthorizedQosFlowDescriptions.SetLen(a.AuthorizedQosFlowDescriptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.Buffer[:a.AuthorizedQosFlowDescriptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return err
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptDNNType:
			a.DNN = nasType.NewDNN(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.DNN.Len); err != nil {
				return err
			}
			a.DNN.SetLen(a.DNN.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.DNN.Buffer[:a.DNN.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptFiveGSMNetworkFeatureSupportType:
			a.FiveGSMNetworkFeatureSupport = nasType.NewFiveGSMNetworkFeatureSupport(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.FiveGSMNetworkFeatureSupport.Len); err != nil {
				return err
			}
			a.FiveGSMNetworkFeatureSupport.SetLen(a.FiveGSMNetworkFeatureSupport.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.FiveGSMNetworkFeatureSupport.Buffer[:a.FiveGSMNetworkFeatureSupport.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptServingPLMNRateControlType:
			a.ServingPLMNRateControl = nasType.NewServingPLMNRateControl(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServingPLMNRateControl.Len); err != nil {
				return err
			}
			a.ServingPLMNRateControl.SetLen(a.ServingPLMNRateControl.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServingPLMNRateControl.Buffer[:a.ServingPLMNRateControl.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptATSSSContainerType:
			a.ATSSSContainer = nasType.NewATSSSContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ATSSSContainer.Len); err != nil {
				return err
			}
			a.ATSSSContainer.SetLen(a.ATSSSContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ATSSSContainer.Buffer[:a.ATSSSContainer.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptControlPlaneOnlyIndicationType:
			a.ControlPlaneOnlyIndication = nasType.NewControlPlaneOnlyIndication(ieiN)
			a.ControlPlaneOnlyIndication.Octet = ieiN
		case PDUSessionEstablishmentAcceptIPHeaderCompressionConfigurationType:
			a.IPHeaderCompressionConfiguration = nasType.NewIPHeaderCompressionConfiguration(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.IPHeaderCompressionConfiguration.Len); err != nil {
				return err
			}
			a.IPHeaderCompressionConfiguration.SetLen(a.IPHeaderCompressionConfiguration.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.IPHeaderCompressionConfiguration.Buffer[:a.IPHeaderCompressionConfiguration.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptEthernetHeaderCompressionConfigType:
			a.EthernetHeaderCompressionConfiguration = nasType.NewEthernetHeaderCompressionConfiguration(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Len); err != nil {
				return err
			}
			a.EthernetHeaderCompressionConfiguration.SetLen(a.EthernetHeaderCompressionConfiguration.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.EthernetHeaderCompressionConfiguration.Octet); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len); err != nil {
				return err
			}
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptReceivedMBSContainerType:
			a.ReceivedMBSContainer = nasType.NewReceivedMBSContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ReceivedMBSContainer.Len); err != nil {
				return err
			}
			a.ReceivedMBSContainer.SetLen(a.ReceivedMBSContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ReceivedMBSContainer.Buffer[:a.ReceivedMBSContainer.GetLen()]); err != nil {
				return err
			}
		case PDUSessionEstablishmentAcceptN3QAIType:
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
