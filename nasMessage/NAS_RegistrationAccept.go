// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/omec-project/nas/v2/nasType"
)

type RegistrationAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationAcceptMessageIdentity
	nasType.RegistrationResult5GS
	*nasType.GUTI5G
	*nasType.EquivalentPlmns
	*nasType.TAIList
	*nasType.AllowedNSSAI
	*nasType.RejectedNSSAI
	*nasType.ConfiguredNSSAI
	*nasType.NetworkFeatureSupport5GS
	*nasType.PDUSessionStatus
	*nasType.PDUSessionReactivationResult
	*nasType.PDUSessionReactivationResultErrorCause
	*nasType.LADNInformation
	*nasType.MICOIndication
	*nasType.NetworkSlicingIndication
	*nasType.ServiceAreaList
	*nasType.T3512Value
	*nasType.Non3GppDeregistrationTimerValue
	*nasType.T3502Value
	*nasType.EmergencyNumberList
	*nasType.ExtendedEmergencyNumberList
	*nasType.SORTransparentContainer
	*nasType.EAPMessage
	*nasType.NSSAIInclusionMode
	*nasType.OperatordefinedAccessCategoryDefinitions
	*nasType.NegotiatedDRXParameters
	*nasType.Non3GPPNWProvidedPolicies
	*nasType.EPSBearerContextStatus
	*nasType.ExtendedDRXParameters
	*nasType.UERadioCapabilityID
	*nasType.UERadioCapabilityIDDeletionIndicationIE
	*nasType.CipheringKeyData
	*nasType.CAGInformationList
	*nasType.TruncatedFiveGSTMSIConfiguration
	*nasType.ExtendedRejectedNSSAI
	*nasType.ServiceLevelAAContainer
	*nasType.FiveGSAdditionalRequestResult
	*nasType.NSSRGInformation
	*nasType.RegistrationWaitRange
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ListOfPLMNsForDisasterCondition
	*nasType.ExtendedCAGInformationList
	*nasType.NSAGInformation
	*nasType.PendingNSSAI
	NegotiatedNBN1ModeDRXParameters      *nasType.NBN1ModeDRXParameters
	NegotiatedWUSAssistanceInformation   *nasType.WUSAssistanceInformation
	NegotiatedPEIPSAssistanceInformation *nasType.PEIPSAssistanceInformation
	ForbiddenTAIRoaming                  *nasType.TAIList
	ForbiddenTAIRegionalProvision        *nasType.TAIList
	*nasType.EquivalentSNPNs
	*nasType.NID
	*nasType.RANTimingSynchronization
	*nasType.AlternativeNSSAI
	*nasType.DiscontinuousCoverageMaxTimeOffset
	*nasType.SNSSAITimeValidityInformation
	UnavailabilityPeriodDuration *nasType.LowerBoundTimerValue
	*nasType.FeatureAuthorizationIndication
}

func NewRegistrationAccept(iei uint8) (registrationAccept *RegistrationAccept) {
	registrationAccept = &RegistrationAccept{}
	return registrationAccept
}

const (
	RegistrationAcceptGUTI5GType                                   uint8 = 0x77
	RegistrationAcceptEquivalentPlmnsType                          uint8 = 0x4A
	RegistrationAcceptTAIListType                                  uint8 = 0x54
	RegistrationAcceptAllowedNSSAIType                             uint8 = 0x15
	RegistrationAcceptRejectedNSSAIType                            uint8 = 0x11
	RegistrationAcceptConfiguredNSSAIType                          uint8 = 0x31
	RegistrationAcceptNetworkFeatureSupport5GSType                 uint8 = 0x21
	RegistrationAcceptPDUSessionStatusType                         uint8 = 0x50
	RegistrationAcceptPDUSessionReactivationResultType             uint8 = 0x26
	RegistrationAcceptPDUSessionReactivationResultErrorCauseType   uint8 = 0x72
	RegistrationAcceptLADNInformationType                          uint8 = 0x79
	RegistrationAcceptMICOIndicationType                           uint8 = 0x0B
	RegistrationAcceptNetworkSlicingIndicationType                 uint8 = 0x09
	RegistrationAcceptServiceAreaListType                          uint8 = 0x27
	RegistrationAcceptT3512ValueType                               uint8 = 0x5E
	RegistrationAcceptNon3GppDeregistrationTimerValueType          uint8 = 0x5D
	RegistrationAcceptT3502ValueType                               uint8 = 0x16
	RegistrationAcceptEmergencyNumberListType                      uint8 = 0x34
	RegistrationAcceptExtendedEmergencyNumberListType              uint8 = 0x7A
	RegistrationAcceptSORTransparentContainerType                  uint8 = 0x73
	RegistrationAcceptEAPMessageType                               uint8 = 0x78
	RegistrationAcceptNSSAIInclusionModeType                       uint8 = 0x0A
	RegistrationAcceptOperatordefinedAccessCategoryDefinitionsType uint8 = 0x76
	RegistrationAcceptNegotiatedDRXParametersType                  uint8 = 0x51
	RegistrationAcceptEPSBearerContextStatusType                   uint8 = 0x60
	RegistrationAcceptExtendedDRXParametersType                    uint8 = 0x6E
	RegistrationAcceptUERadioCapabilityIDType                      uint8 = 0x67
	RegistrationAcceptUERadioCapabilityIDDeletionIndicationType    uint8 = 0x0E
	RegistrationAcceptCipheringKeyDataType                         uint8 = 0x74
	RegistrationAcceptCAGInformationListType                       uint8 = 0x75
	RegistrationAcceptTruncatedFiveGSTMSIConfigurationType         uint8 = 0x1B
	RegistrationAcceptExtendedRejectedNSSAIType                    uint8 = 0x68
	RegistrationAcceptServiceLevelAAContainerType                  uint8 = 0x7B
	RegistrationAcceptFiveGSAdditionalRequestResultType            uint8 = 0x35
	RegistrationAcceptNSSRGInformationType                         uint8 = 0x70
	RegistrationAcceptRegistrationWaitRangeType                    uint8 = 0x14
	RegistrationAcceptDisasterReturnWaitRangeType                  uint8 = 0x2C
	RegistrationAcceptListOfPLMNsForDisasterConditionType          uint8 = 0x13
	RegistrationAcceptExtendedCAGInformationListType               uint8 = 0x71
	RegistrationAcceptNSAGInformationType                          uint8 = 0x7C
	RegistrationAcceptNon3GPPNWProvidedPoliciesType                uint8 = 0x0D
	RegistrationAcceptPendingNSSAIType                             uint8 = 0x39
	RegistrationAcceptNegotiatedNBN1ModeDRXParametersType          uint8 = 0x29
	RegistrationAcceptNegotiatedWUSAssistanceInformationType       uint8 = 0x1C
	RegistrationAcceptNegotiatedPEIPSAssistanceInformationType     uint8 = 0x33
	RegistrationAcceptForbiddenTAIRoamingType                      uint8 = 0x1D
	RegistrationAcceptForbiddenTAIRegionalProvisionType            uint8 = 0x1E
	RegistrationAcceptEquivalentSNPNsType                          uint8 = 0x3D
	RegistrationAcceptNIDType                                      uint8 = 0x32
	RegistrationAcceptRANTimingSynchronizationType                 uint8 = 0x4B
	RegistrationAcceptAlternativeNSSAIType                         uint8 = 0x4C
	RegistrationAcceptDiscontinuousCoverageMaxTimeOffsetType       uint8 = 0x4F
	RegistrationAcceptSNSSAITimeValidityInformationType            uint8 = 0x5B
	RegistrationAcceptUnavailabilityPeriodDurationType             uint8 = 0x3C
	RegistrationAcceptFeatureAuthorizationIndicationType           uint8 = 0x5C
)

func (a *RegistrationAccept) EncodeRegistrationAccept(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.RegistrationAcceptMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet); err != nil {
		return
	}
	if a.GUTI5G != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.Octet[:a.GUTI5G.GetLen()]); err != nil {
			return
		}
	}
	if a.EquivalentPlmns != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentPlmns.Octet[:a.EquivalentPlmns.GetLen()]); err != nil {
			return
		}
	}
	if a.TAIList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.TAIList.Buffer); err != nil {
			return
		}
	}
	if a.AllowedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AllowedNSSAI.Buffer); err != nil {
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
	if a.ConfiguredNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Buffer); err != nil {
			return
		}
	}
	if a.NetworkFeatureSupport5GS != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.Octet[:a.NetworkFeatureSupport5GS.GetLen()]); err != nil {
			return
		}
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
	if a.LADNInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.LADNInformation.Buffer); err != nil {
			return
		}
	}
	if a.MICOIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.MICOIndication.Octet); err != nil {
			return
		}
	}
	if a.NetworkSlicingIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.NetworkSlicingIndication.Octet); err != nil {
			return
		}
	}
	if a.ServiceAreaList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ServiceAreaList.Buffer); err != nil {
			return
		}
	}
	if a.T3512Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3512Value.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3512Value.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.T3512Value.Octet); err != nil {
			return
		}
	}
	if a.Non3GppDeregistrationTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GppDeregistrationTimerValue.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GppDeregistrationTimerValue.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Octet); err != nil {
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
	if a.EmergencyNumberList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EmergencyNumberList.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EmergencyNumberList.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EmergencyNumberList.Buffer); err != nil {
			return
		}
	}
	if a.ExtendedEmergencyNumberList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedEmergencyNumberList.Buffer); err != nil {
			return
		}
	}
	if a.SORTransparentContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.SORTransparentContainer.Buffer); err != nil {
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
	if a.NSSAIInclusionMode != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.NSSAIInclusionMode.Octet); err != nil {
			return
		}
	}
	if a.OperatordefinedAccessCategoryDefinitions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Buffer); err != nil {
			return
		}
	}
	if a.NegotiatedDRXParameters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedDRXParameters.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedDRXParameters.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Octet); err != nil {
			return
		}
	}
	if a.Non3GPPNWProvidedPolicies != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.Non3GPPNWProvidedPolicies.Octet); err != nil {
			return
		}
	}
	if a.EPSBearerContextStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.EPSBearerContextStatus.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSBearerContextStatus.Buffer[:uint8(a.EPSBearerContextStatus.GetLen())]); err != nil {
			return
		}
	}
	if a.ExtendedDRXParameters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedDRXParameters.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.ExtendedDRXParameters.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedDRXParameters.Buffer[:uint8(a.ExtendedDRXParameters.GetLen())]); err != nil {
			return
		}
	}
	if a.UERadioCapabilityID != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UERadioCapabilityID.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.UERadioCapabilityID.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:uint8(a.UERadioCapabilityID.GetLen())]); err != nil {
			return
		}
	}
	if a.UERadioCapabilityIDDeletionIndicationIE != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.UERadioCapabilityIDDeletionIndicationIE.Octet); err != nil {
			return
		}
	}
	if a.CipheringKeyData != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.CipheringKeyData.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.CipheringKeyData.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.CipheringKeyData.Buffer); err != nil {
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
	if a.TruncatedFiveGSTMSIConfiguration != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.TruncatedFiveGSTMSIConfiguration.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.Buffer[:uint8(a.TruncatedFiveGSTMSIConfiguration.GetLen())]); err != nil {
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
	if a.NSSRGInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NSSRGInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NSSRGInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NSSRGInformation.Buffer); err != nil {
			return
		}
	}
	if a.RegistrationWaitRange != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RegistrationWaitRange.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.RegistrationWaitRange.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RegistrationWaitRange.Buffer[:uint8(a.RegistrationWaitRange.GetLen())]); err != nil {
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
	if a.ListOfPLMNsForDisasterCondition != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.ListOfPLMNsForDisasterCondition.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.Buffer[:uint8(a.ListOfPLMNsForDisasterCondition.GetLen())]); err != nil {
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
	if a.NSAGInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NSAGInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NSAGInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NSAGInformation.Buffer); err != nil {
			return
		}
	}
	if a.PendingNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PendingNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PendingNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PendingNSSAI.Buffer); err != nil {
			return
		}
	}
	if a.NegotiatedNBN1ModeDRXParameters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedNBN1ModeDRXParameters.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedNBN1ModeDRXParameters.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NegotiatedNBN1ModeDRXParameters.Buffer); err != nil {
			return
		}
	}
	if a.NegotiatedWUSAssistanceInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedWUSAssistanceInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.NegotiatedWUSAssistanceInformation.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedWUSAssistanceInformation.Buffer[:uint8(a.NegotiatedWUSAssistanceInformation.GetLen())]); err != nil {
			return
		}
	}
	if a.NegotiatedPEIPSAssistanceInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedPEIPSAssistanceInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.NegotiatedPEIPSAssistanceInformation.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NegotiatedPEIPSAssistanceInformation.Buffer[:uint8(a.NegotiatedPEIPSAssistanceInformation.GetLen())]); err != nil {
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
	if a.EquivalentSNPNs != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentSNPNs.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EquivalentSNPNs.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EquivalentSNPNs.Buffer); err != nil {
			return
		}
	}
	if a.NID != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NID.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.NID.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NID.Buffer[:uint8(a.NID.GetLen())]); err != nil {
			return
		}
	}
	if a.RANTimingSynchronization != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RANTimingSynchronization.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RANTimingSynchronization.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RANTimingSynchronization.Buffer); err != nil {
			return
		}
	}
	if a.AlternativeNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AlternativeNSSAI.Buffer); err != nil {
			return
		}
	}
	if a.DiscontinuousCoverageMaxTimeOffset != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.DiscontinuousCoverageMaxTimeOffset.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DiscontinuousCoverageMaxTimeOffset.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Octet); err != nil {
			return
		}
	}
	if a.SNSSAITimeValidityInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.SNSSAITimeValidityInformation.Buffer); err != nil {
			return
		}
	}
	if a.UnavailabilityPeriodDuration != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UnavailabilityPeriodDuration.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UnavailabilityPeriodDuration.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UnavailabilityPeriodDuration.Octet); err != nil {
			return
		}
	}
	if a.FeatureAuthorizationIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.FeatureAuthorizationIndication.Buffer); err != nil {
			return
		}
	}
}

func (a *RegistrationAccept) DecodeRegistrationAccept(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationAcceptMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Len); err != nil {
		return err
	}
	a.RegistrationResult5GS.SetLen(a.RegistrationResult5GS.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet); err != nil {
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
		case RegistrationAcceptGUTI5GType:
			a.GUTI5G = nasType.NewGUTI5G(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.GUTI5G.Len); err != nil {
				return err
			}
			a.GUTI5G.SetLen(a.GUTI5G.GetLen())
			if a.GUTI5G.GetLen() > uint16(len(a.GUTI5G.Octet)) {
				a.GUTI5G = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: GUTI5G length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.GUTI5G.Octet[:a.GUTI5G.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptEquivalentPlmnsType:
			a.EquivalentPlmns = nasType.NewEquivalentPlmns(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EquivalentPlmns.Len); err != nil {
				return err
			}
			a.EquivalentPlmns.SetLen(a.EquivalentPlmns.GetLen())
			if a.EquivalentPlmns.GetLen() > uint8(len(a.EquivalentPlmns.Octet)) {
				a.EquivalentPlmns = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: EquivalentPlmns length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.EquivalentPlmns.Octet[:a.EquivalentPlmns.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptTAIListType:
			a.TAIList = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.TAIList.Len); err != nil {
				return err
			}
			a.TAIList.SetLen(a.TAIList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.TAIList.Buffer[:a.TAIList.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptAllowedNSSAIType:
			a.AllowedNSSAI = nasType.NewAllowedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedNSSAI.Len); err != nil {
				return err
			}
			a.AllowedNSSAI.SetLen(a.AllowedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer[:a.AllowedNSSAI.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len); err != nil {
				return err
			}
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptConfiguredNSSAIType:
			a.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Len); err != nil {
				return err
			}
			a.ConfiguredNSSAI.SetLen(a.ConfiguredNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer[:a.ConfiguredNSSAI.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNetworkFeatureSupport5GSType:
			a.NetworkFeatureSupport5GS = nasType.NewNetworkFeatureSupport5GS(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NetworkFeatureSupport5GS.Len); err != nil {
				return err
			}
			a.NetworkFeatureSupport5GS.SetLen(a.NetworkFeatureSupport5GS.GetLen())
			if a.NetworkFeatureSupport5GS.GetLen() > uint8(len(a.NetworkFeatureSupport5GS.Octet)) {
				a.NetworkFeatureSupport5GS = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: NetworkFeatureSupport5GS length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.NetworkFeatureSupport5GS.Octet[:a.NetworkFeatureSupport5GS.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return err
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptPDUSessionReactivationResultType:
			a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResult.Len); err != nil {
				return err
			}
			a.PDUSessionReactivationResult.SetLen(a.PDUSessionReactivationResult.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResult.Buffer[:a.PDUSessionReactivationResult.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptPDUSessionReactivationResultErrorCauseType:
			a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionReactivationResultErrorCause.Len); err != nil {
				return err
			}
			a.PDUSessionReactivationResultErrorCause.SetLen(a.PDUSessionReactivationResultErrorCause.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionReactivationResultErrorCause.Buffer[:a.PDUSessionReactivationResultErrorCause.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptLADNInformationType:
			a.LADNInformation = nasType.NewLADNInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LADNInformation.Len); err != nil {
				return err
			}
			a.LADNInformation.SetLen(a.LADNInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.LADNInformation.Buffer[:a.LADNInformation.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case RegistrationAcceptNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case RegistrationAcceptServiceAreaListType:
			a.ServiceAreaList = nasType.NewServiceAreaList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceAreaList.Len); err != nil {
				return err
			}
			a.ServiceAreaList.SetLen(a.ServiceAreaList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceAreaList.Buffer[:a.ServiceAreaList.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptT3512ValueType:
			a.T3512Value = nasType.NewT3512Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3512Value.Len); err != nil {
				return err
			}
			a.T3512Value.SetLen(a.T3512Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3512Value.Octet); err != nil {
				return err
			}
		case RegistrationAcceptNon3GppDeregistrationTimerValueType:
			a.Non3GppDeregistrationTimerValue = nasType.NewNon3GppDeregistrationTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Len); err != nil {
				return err
			}
			a.Non3GppDeregistrationTimerValue.SetLen(a.Non3GppDeregistrationTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Non3GppDeregistrationTimerValue.Octet); err != nil {
				return err
			}
		case RegistrationAcceptT3502ValueType:
			a.T3502Value = nasType.NewT3502Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Len); err != nil {
				return err
			}
			a.T3502Value.SetLen(a.T3502Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3502Value.Octet); err != nil {
				return err
			}
		case RegistrationAcceptEmergencyNumberListType:
			a.EmergencyNumberList = nasType.NewEmergencyNumberList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EmergencyNumberList.Len); err != nil {
				return err
			}
			a.EmergencyNumberList.SetLen(a.EmergencyNumberList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EmergencyNumberList.Buffer[:a.EmergencyNumberList.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptExtendedEmergencyNumberListType:
			a.ExtendedEmergencyNumberList = nasType.NewExtendedEmergencyNumberList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedEmergencyNumberList.Len); err != nil {
				return err
			}
			a.ExtendedEmergencyNumberList.SetLen(a.ExtendedEmergencyNumberList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedEmergencyNumberList.Buffer[:a.ExtendedEmergencyNumberList.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptSORTransparentContainerType:
			a.SORTransparentContainer = nasType.NewSORTransparentContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SORTransparentContainer.Len); err != nil {
				return err
			}
			a.SORTransparentContainer.SetLen(a.SORTransparentContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SORTransparentContainer.Buffer[:a.SORTransparentContainer.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNSSAIInclusionModeType:
			a.NSSAIInclusionMode = nasType.NewNSSAIInclusionMode(ieiN)
			a.NSSAIInclusionMode.Octet = ieiN
		case RegistrationAcceptOperatordefinedAccessCategoryDefinitionsType:
			a.OperatordefinedAccessCategoryDefinitions = nasType.NewOperatordefinedAccessCategoryDefinitions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Len); err != nil {
				return err
			}
			a.OperatordefinedAccessCategoryDefinitions.SetLen(a.OperatordefinedAccessCategoryDefinitions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer[:a.OperatordefinedAccessCategoryDefinitions.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNegotiatedDRXParametersType:
			a.NegotiatedDRXParameters = nasType.NewNegotiatedDRXParameters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Len); err != nil {
				return err
			}
			a.NegotiatedDRXParameters.SetLen(a.NegotiatedDRXParameters.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.NegotiatedDRXParameters.Octet); err != nil {
				return err
			}
		case RegistrationAcceptEPSBearerContextStatusType:
			a.EPSBearerContextStatus = nasType.NewEPSBearerContextStatus(ieiN)
			var lenN0 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN0); err != nil {
				return err
			}
			a.EPSBearerContextStatus.SetLen(uint16(lenN0))
			if err := binary.Read(buffer, binary.BigEndian, a.EPSBearerContextStatus.Buffer[:lenN0]); err != nil {
				return err
			}
		case RegistrationAcceptExtendedDRXParametersType:
			a.ExtendedDRXParameters = nasType.NewExtendedDRXParameters(ieiN)
			var lenN1 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN1); err != nil {
				return err
			}
			a.ExtendedDRXParameters.SetLen(uint16(lenN1))
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedDRXParameters.Buffer[:lenN1]); err != nil {
				return err
			}
		case RegistrationAcceptUERadioCapabilityIDType:
			a.UERadioCapabilityID = nasType.NewUERadioCapabilityID(ieiN)
			var lenN2 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN2); err != nil {
				return err
			}
			a.UERadioCapabilityID.SetLen(uint16(lenN2))
			if err := binary.Read(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:lenN2]); err != nil {
				return err
			}
		case RegistrationAcceptUERadioCapabilityIDDeletionIndicationType:
			a.UERadioCapabilityIDDeletionIndicationIE = nasType.NewUERadioCapabilityIDDeletionIndicationIE(ieiN)
			a.UERadioCapabilityIDDeletionIndicationIE.Octet = ieiN
		case RegistrationAcceptCipheringKeyDataType:
			a.CipheringKeyData = nasType.NewCipheringKeyData(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CipheringKeyData.Len); err != nil {
				return err
			}
			a.CipheringKeyData.SetLen(a.CipheringKeyData.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.CipheringKeyData.Buffer[:a.CipheringKeyData.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len); err != nil {
				return err
			}
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptTruncatedFiveGSTMSIConfigurationType:
			a.TruncatedFiveGSTMSIConfiguration = nasType.NewTruncatedFiveGSTMSIConfiguration(ieiN)
			var lenN3 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN3); err != nil {
				return err
			}
			a.TruncatedFiveGSTMSIConfiguration.SetLen(uint16(lenN3))
			if err := binary.Read(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.Buffer[:lenN3]); err != nil {
				return err
			}
		case RegistrationAcceptExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var lenN4 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN4); err != nil {
				return err
			}
			a.ExtendedRejectedNSSAI.SetLen(uint16(lenN4))
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:lenN4]); err != nil {
				return err
			}
		case RegistrationAcceptServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len); err != nil {
				return err
			}
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptFiveGSAdditionalRequestResultType:
			a.FiveGSAdditionalRequestResult = nasType.NewFiveGSAdditionalRequestResult(ieiN)
			var lenN5 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN5); err != nil {
				return err
			}
			a.FiveGSAdditionalRequestResult.SetLen(uint16(lenN5))
			if err := binary.Read(buffer, binary.BigEndian, a.FiveGSAdditionalRequestResult.Buffer[:lenN5]); err != nil {
				return err
			}
		case RegistrationAcceptNSSRGInformationType:
			a.NSSRGInformation = nasType.NewNSSRGInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NSSRGInformation.Len); err != nil {
				return err
			}
			a.NSSRGInformation.SetLen(a.NSSRGInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NSSRGInformation.Buffer[:a.NSSRGInformation.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptRegistrationWaitRangeType:
			a.RegistrationWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN6 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN6); err != nil {
				return err
			}
			a.RegistrationWaitRange.SetLen(uint16(lenN6))
			if err := binary.Read(buffer, binary.BigEndian, a.RegistrationWaitRange.Buffer[:lenN6]); err != nil {
				return err
			}
		case RegistrationAcceptDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN6a uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN6a); err != nil {
				return err
			}
			a.DisasterReturnWaitRange.SetLen(uint16(lenN6a))
			if err := binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:lenN6a]); err != nil {
				return err
			}
		case RegistrationAcceptListOfPLMNsForDisasterConditionType:
			a.ListOfPLMNsForDisasterCondition = nasType.NewListOfPLMNsForDisasterCondition(ieiN)
			var lenN7 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN7); err != nil {
				return err
			}
			a.ListOfPLMNsForDisasterCondition.SetLen(uint16(lenN7))
			if err := binary.Read(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.Buffer[:lenN7]); err != nil {
				return err
			}
		case RegistrationAcceptExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len); err != nil {
				return err
			}
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNSAGInformationType:
			a.NSAGInformation = nasType.NewNSAGInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NSAGInformation.Len); err != nil {
				return err
			}
			a.NSAGInformation.SetLen(a.NSAGInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NSAGInformation.Buffer[:a.NSAGInformation.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNon3GPPNWProvidedPoliciesType:
			a.Non3GPPNWProvidedPolicies = nasType.NewNon3GPPNWProvidedPolicies(ieiN)
			a.Non3GPPNWProvidedPolicies.Octet = ieiN
		case RegistrationAcceptPendingNSSAIType:
			a.PendingNSSAI = nasType.NewPendingNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PendingNSSAI.Len); err != nil {
				return err
			}
			a.PendingNSSAI.SetLen(a.PendingNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PendingNSSAI.Buffer[:a.PendingNSSAI.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNegotiatedNBN1ModeDRXParametersType:
			a.NegotiatedNBN1ModeDRXParameters = nasType.NewNBN1ModeDRXParameters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NegotiatedNBN1ModeDRXParameters.Len); err != nil {
				return err
			}
			a.NegotiatedNBN1ModeDRXParameters.SetLen(a.NegotiatedNBN1ModeDRXParameters.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NegotiatedNBN1ModeDRXParameters.Buffer[:a.NegotiatedNBN1ModeDRXParameters.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNegotiatedWUSAssistanceInformationType:
			a.NegotiatedWUSAssistanceInformation = nasType.NewWUSAssistanceInformation(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.NegotiatedWUSAssistanceInformation.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.NegotiatedWUSAssistanceInformation.Buffer[:l]); err != nil {
				return err
			}
		case RegistrationAcceptNegotiatedPEIPSAssistanceInformationType:
			a.NegotiatedPEIPSAssistanceInformation = nasType.NewPEIPSAssistanceInformation(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.NegotiatedPEIPSAssistanceInformation.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.NegotiatedPEIPSAssistanceInformation.Buffer[:l]); err != nil {
				return err
			}
		case RegistrationAcceptForbiddenTAIRoamingType:
			a.ForbiddenTAIRoaming = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRoaming.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRoaming.SetLen(a.ForbiddenTAIRoaming.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRoaming.Buffer[:a.ForbiddenTAIRoaming.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptForbiddenTAIRegionalProvisionType:
			a.ForbiddenTAIRegionalProvision = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ForbiddenTAIRegionalProvision.Len); err != nil {
				return err
			}
			a.ForbiddenTAIRegionalProvision.SetLen(a.ForbiddenTAIRegionalProvision.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ForbiddenTAIRegionalProvision.Buffer[:a.ForbiddenTAIRegionalProvision.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptEquivalentSNPNsType:
			a.EquivalentSNPNs = nasType.NewEquivalentSNPNs(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EquivalentSNPNs.Len); err != nil {
				return err
			}
			a.EquivalentSNPNs.SetLen(a.EquivalentSNPNs.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EquivalentSNPNs.Buffer[:a.EquivalentSNPNs.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptNIDType:
			a.NID = nasType.NewNID(ieiN)
			var l uint8
			if err := binary.Read(buffer, binary.BigEndian, &l); err != nil {
				return err
			}
			a.NID.SetLen(uint16(l))
			if err := binary.Read(buffer, binary.BigEndian, a.NID.Buffer[:l]); err != nil {
				return err
			}
		case RegistrationAcceptRANTimingSynchronizationType:
			a.RANTimingSynchronization = nasType.NewRANTimingSynchronization(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RANTimingSynchronization.Len); err != nil {
				return err
			}
			a.RANTimingSynchronization.SetLen(a.RANTimingSynchronization.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RANTimingSynchronization.Buffer[:a.RANTimingSynchronization.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptAlternativeNSSAIType:
			a.AlternativeNSSAI = nasType.NewAlternativeNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AlternativeNSSAI.Len); err != nil {
				return err
			}
			a.AlternativeNSSAI.SetLen(a.AlternativeNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AlternativeNSSAI.Buffer[:a.AlternativeNSSAI.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptDiscontinuousCoverageMaxTimeOffsetType:
			a.DiscontinuousCoverageMaxTimeOffset = nasType.NewDiscontinuousCoverageMaxTimeOffset(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Len); err != nil {
				return err
			}
			a.DiscontinuousCoverageMaxTimeOffset.SetLen(a.DiscontinuousCoverageMaxTimeOffset.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Octet); err != nil {
				return err
			}
		case RegistrationAcceptSNSSAITimeValidityInformationType:
			a.SNSSAITimeValidityInformation = nasType.NewSNSSAITimeValidityInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAITimeValidityInformation.Len); err != nil {
				return err
			}
			a.SNSSAITimeValidityInformation.SetLen(a.SNSSAITimeValidityInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.Buffer[:a.SNSSAITimeValidityInformation.GetLen()]); err != nil {
				return err
			}
		case RegistrationAcceptUnavailabilityPeriodDurationType:
			a.UnavailabilityPeriodDuration = nasType.NewLowerBoundTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UnavailabilityPeriodDuration.Len); err != nil {
				return err
			}
			a.UnavailabilityPeriodDuration.SetLen(a.UnavailabilityPeriodDuration.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.UnavailabilityPeriodDuration.Octet); err != nil {
				return err
			}
		case RegistrationAcceptFeatureAuthorizationIndicationType:
			a.FeatureAuthorizationIndication = nasType.NewFeatureAuthorizationIndication(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.FeatureAuthorizationIndication.Len); err != nil {
				return err
			}
			a.FeatureAuthorizationIndication.SetLen(a.FeatureAuthorizationIndication.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.FeatureAuthorizationIndication.Buffer[:a.FeatureAuthorizationIndication.GetLen()]); err != nil {
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
