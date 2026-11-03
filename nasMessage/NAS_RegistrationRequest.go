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

type RegistrationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationRequestMessageIdentity
	nasType.NgksiAndRegistrationType5GS
	nasType.MobileIdentity5GS
	*nasType.NoncurrentNativeNASKeySetIdentifier
	*nasType.Capability5GMM
	*nasType.UESecurityCapability
	*nasType.RequestedNSSAI
	*nasType.LastVisitedRegisteredTAI
	*nasType.S1UENetworkCapability
	*nasType.UplinkDataStatus
	*nasType.PDUSessionStatus
	*nasType.MICOIndication
	*nasType.UEStatus
	*nasType.AdditionalGUTI
	*nasType.AllowedPDUSessionStatus
	*nasType.UesUsageSetting
	*nasType.RequestedDRXParameters
	*nasType.EPSNASMessageContainer
	*nasType.LADNIndication
	*nasType.PayloadContainer
	*nasType.NetworkSlicingIndication
	*nasType.UpdateType5GS
	*nasType.MobileStationClassmark2
	*nasType.SupportedCodecs
	*nasType.NASMessageContainer
	*nasType.EPSBearerContextStatus
	*nasType.ExtendedDRXParameters
	*nasType.UERadioCapabilityID
	*nasType.RequestedMappedNSSAI
	*nasType.AdditionalInformationRequested
	*nasType.WUSAssistanceInformation
	*nasType.N5GCIndication
	*nasType.NBN1ModeDRXParameters
	*nasType.UERequestType
	*nasType.PagingRestriction
	*nasType.ServiceLevelAAContainer
	*nasType.NID
	*nasType.PLMNIdentityWithDisasterCondition
	*nasType.PEIPSAssistanceInformation
	*nasType.TimeDuration
	*nasType.Non3GPPPathSwitchingInformation
	*nasType.AUN3Indication
	*nasType.T3512Value
}

func NewRegistrationRequest(iei uint8) (registrationRequest *RegistrationRequest) {
	registrationRequest = &RegistrationRequest{}
	return registrationRequest
}

const (
	RegistrationRequestNoncurrentNativeNASKeySetIdentifierType uint8 = 0x0C
	RegistrationRequestCapability5GMMType                      uint8 = 0x10
	RegistrationRequestUESecurityCapabilityType                uint8 = 0x2E
	RegistrationRequestRequestedNSSAIType                      uint8 = 0x2F
	RegistrationRequestLastVisitedRegisteredTAIType            uint8 = 0x52
	RegistrationRequestS1UENetworkCapabilityType               uint8 = 0x17
	RegistrationRequestUplinkDataStatusType                    uint8 = 0x40
	RegistrationRequestPDUSessionStatusType                    uint8 = 0x50
	RegistrationRequestMICOIndicationType                      uint8 = 0x0B
	RegistrationRequestUEStatusType                            uint8 = 0x2B
	RegistrationRequestAdditionalGUTIType                      uint8 = 0x77
	RegistrationRequestAllowedPDUSessionStatusType             uint8 = 0x25
	RegistrationRequestUesUsageSettingType                     uint8 = 0x18
	RegistrationRequestRequestedDRXParametersType              uint8 = 0x51
	RegistrationRequestEPSNASMessageContainerType              uint8 = 0x70
	RegistrationRequestLADNIndicationType                      uint8 = 0x74
	RegistrationRequestPayloadContainerType                    uint8 = 0x7B
	RegistrationRequestNetworkSlicingIndicationType            uint8 = 0x09
	RegistrationRequestUpdateType5GSType                       uint8 = 0x53
	RegistrationRequestNASMessageContainerType                 uint8 = 0x71
	RegistrationRequestEPSBearerContextStatusType              uint8 = 0x60
	RegistrationRequestExtendedDRXParametersType               uint8 = 0x6E
	RegistrationRequestUERadioCapabilityIDType                 uint8 = 0x67
	RegistrationRequestRequestedMappedNSSAIType                uint8 = 0x35
	RegistrationRequestAdditionalInformationRequestedType      uint8 = 0x48
	RegistrationRequestWUSAssistanceInformationType            uint8 = 0x1A
	RegistrationRequestN5GCIndicationType                      uint8 = 0x0A
	RegistrationRequestNBN1ModeDRXParametersType               uint8 = 0x30
	RegistrationRequestUERequestTypeType                       uint8 = 0x29
	RegistrationRequestPagingRestrictionType                   uint8 = 0x28
	RegistrationRequestServiceLevelAAContainerType             uint8 = 0x72
	RegistrationRequestNIDType                                 uint8 = 0x32
	RegistrationRequestPLMNIdentityWithDisasterConditionType   uint8 = 0x16
	RegistrationRequestPEIPSAssistanceInformationType          uint8 = 0x2A
	RegistrationRequestTimeDurationType                        uint8 = 0x3C
	RegistrationRequestNon3GPPPathSwitchingInformationType     uint8 = 0x3F
	RegistrationRequestAUN3IndicationType                      uint8 = 0x56
	RegistrationRequestMobileStationClassmark2Type             uint8 = 0x41
	RegistrationRequestSupportedCodecsType                     uint8 = 0x42
	RegistrationRequestRequestedT3512ValueType                 uint8 = 0x3B
)

func (a *RegistrationRequest) EncodeRegistrationRequest(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.RegistrationRequestMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.NgksiAndRegistrationType5GS.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.MobileIdentity5GS.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.MobileIdentity5GS.Buffer); err != nil {
		return
	}
	if a.NoncurrentNativeNASKeySetIdentifier != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.NoncurrentNativeNASKeySetIdentifier.Octet); err != nil {
			return
		}
	}
	if a.Capability5GMM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GMM.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GMM.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GMM.Octet[:a.Capability5GMM.GetLen()]); err != nil {
			return
		}
	}
	if a.UESecurityCapability != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UESecurityCapability.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UESecurityCapability.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UESecurityCapability.Buffer); err != nil {
			return
		}
	}
	if a.RequestedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RequestedNSSAI.Buffer); err != nil {
			return
		}
	}
	if a.LastVisitedRegisteredTAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LastVisitedRegisteredTAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.LastVisitedRegisteredTAI.Octet); err != nil {
			return
		}
	}
	if a.S1UENetworkCapability != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.S1UENetworkCapability.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.S1UENetworkCapability.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.S1UENetworkCapability.Buffer); err != nil {
			return
		}
	}
	if a.UplinkDataStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UplinkDataStatus.Buffer); err != nil {
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
	if a.MICOIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.MICOIndication.Octet); err != nil {
			return
		}
	}
	if a.UEStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UEStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UEStatus.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UEStatus.Octet); err != nil {
			return
		}
	}
	if a.AdditionalGUTI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalGUTI.Octet[:a.AdditionalGUTI.GetLen()]); err != nil {
			return
		}
	}
	if a.AllowedPDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Buffer); err != nil {
			return
		}
	}
	if a.UesUsageSetting != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UesUsageSetting.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UesUsageSetting.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UesUsageSetting.Octet); err != nil {
			return
		}
	}
	if a.RequestedDRXParameters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedDRXParameters.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedDRXParameters.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RequestedDRXParameters.Octet); err != nil {
			return
		}
	}
	if a.EPSNASMessageContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EPSNASMessageContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EPSNASMessageContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EPSNASMessageContainer.Buffer); err != nil {
			return
		}
	}
	if a.LADNIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LADNIndication.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNIndication.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.LADNIndication.Buffer); err != nil {
			return
		}
	}
	if a.PayloadContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PayloadContainer.Buffer); err != nil {
			return
		}
	}
	if a.NetworkSlicingIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.NetworkSlicingIndication.Octet); err != nil {
			return
		}
	}
	if a.UpdateType5GS != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UpdateType5GS.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UpdateType5GS.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UpdateType5GS.Octet); err != nil {
			return
		}
	}
	if a.MobileStationClassmark2 != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MobileStationClassmark2.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.MobileStationClassmark2.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MobileStationClassmark2.Buffer[:uint8(a.MobileStationClassmark2.GetLen())]); err != nil {
			return
		}
	}
	if a.SupportedCodecs != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SupportedCodecs.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.SupportedCodecs.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SupportedCodecs.Buffer[:uint8(a.SupportedCodecs.GetLen())]); err != nil {
			return
		}
	}
	if a.NASMessageContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer); err != nil {
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
	if a.RequestedMappedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedMappedNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.RequestedMappedNSSAI.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedMappedNSSAI.Buffer[:uint8(a.RequestedMappedNSSAI.GetLen())]); err != nil {
			return
		}
	}
	if a.AdditionalInformationRequested != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalInformationRequested.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.AdditionalInformationRequested.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalInformationRequested.Buffer[:uint8(a.AdditionalInformationRequested.GetLen())]); err != nil {
			return
		}
	}
	if a.WUSAssistanceInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.WUSAssistanceInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.WUSAssistanceInformation.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.WUSAssistanceInformation.Buffer[:uint8(a.WUSAssistanceInformation.GetLen())]); err != nil {
			return
		}
	}
	if a.N5GCIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.N5GCIndication.Octet); err != nil {
			return
		}
	}
	if a.NBN1ModeDRXParameters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NBN1ModeDRXParameters.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.NBN1ModeDRXParameters.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NBN1ModeDRXParameters.Buffer[:uint8(a.NBN1ModeDRXParameters.GetLen())]); err != nil {
			return
		}
	}
	if a.UERequestType != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UERequestType.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.UERequestType.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UERequestType.Buffer[:uint8(a.UERequestType.GetLen())]); err != nil {
			return
		}
	}
	if a.PagingRestriction != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PagingRestriction.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.PagingRestriction.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PagingRestriction.Buffer[:uint8(a.PagingRestriction.GetLen())]); err != nil {
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
	if a.PLMNIdentityWithDisasterCondition != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PLMNIdentityWithDisasterCondition.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.PLMNIdentityWithDisasterCondition.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PLMNIdentityWithDisasterCondition.Buffer[:uint8(a.PLMNIdentityWithDisasterCondition.GetLen())]); err != nil {
			return
		}
	}
	if a.PEIPSAssistanceInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PEIPSAssistanceInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.PEIPSAssistanceInformation.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PEIPSAssistanceInformation.Buffer[:uint8(a.PEIPSAssistanceInformation.GetLen())]); err != nil {
			return
		}
	}
	if a.TimeDuration != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.TimeDuration.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.TimeDuration.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TimeDuration.Buffer[:uint8(a.TimeDuration.GetLen())]); err != nil {
			return
		}
	}
	if a.Non3GPPPathSwitchingInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GPPPathSwitchingInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.Non3GPPPathSwitchingInformation.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GPPPathSwitchingInformation.Buffer[:uint8(a.Non3GPPPathSwitchingInformation.GetLen())]); err != nil {
			return
		}
	}
	if a.AUN3Indication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AUN3Indication.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AUN3Indication.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AUN3Indication.Octet); err != nil {
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
}

func (a *RegistrationRequest) DecodeRegistrationRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationRequestMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NgksiAndRegistrationType5GS.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Len); err != nil {
		return err
	}
	a.MobileIdentity5GS.SetLen(a.MobileIdentity5GS.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Buffer); err != nil {
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
		case RegistrationRequestNoncurrentNativeNASKeySetIdentifierType:
			a.NoncurrentNativeNASKeySetIdentifier = nasType.NewNoncurrentNativeNASKeySetIdentifier(ieiN)
			a.NoncurrentNativeNASKeySetIdentifier.Octet = ieiN
		case RegistrationRequestCapability5GMMType:
			a.Capability5GMM = nasType.NewCapability5GMM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Capability5GMM.Len); err != nil {
				return err
			}
			a.Capability5GMM.SetLen(a.Capability5GMM.GetLen())
			if a.Capability5GMM.GetLen() > uint8(len(a.Capability5GMM.Octet)) {
				a.Capability5GMM = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: Capability5GMM length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.Capability5GMM.Octet[:a.Capability5GMM.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestUESecurityCapabilityType:
			a.UESecurityCapability = nasType.NewUESecurityCapability(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UESecurityCapability.Len); err != nil {
				return err
			}
			a.UESecurityCapability.SetLen(a.UESecurityCapability.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.UESecurityCapability.Buffer[:a.UESecurityCapability.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestRequestedNSSAIType:
			a.RequestedNSSAI = nasType.NewRequestedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedNSSAI.Len); err != nil {
				return err
			}
			a.RequestedNSSAI.SetLen(a.RequestedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RequestedNSSAI.Buffer[:a.RequestedNSSAI.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestLastVisitedRegisteredTAIType:
			a.LastVisitedRegisteredTAI = nasType.NewLastVisitedRegisteredTAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LastVisitedRegisteredTAI.Octet); err != nil {
				return err
			}
		case RegistrationRequestS1UENetworkCapabilityType:
			a.S1UENetworkCapability = nasType.NewS1UENetworkCapability(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.S1UENetworkCapability.Len); err != nil {
				return err
			}
			a.S1UENetworkCapability.SetLen(a.S1UENetworkCapability.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.S1UENetworkCapability.Buffer[:a.S1UENetworkCapability.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestUplinkDataStatusType:
			a.UplinkDataStatus = nasType.NewUplinkDataStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UplinkDataStatus.Len); err != nil {
				return err
			}
			a.UplinkDataStatus.SetLen(a.UplinkDataStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer[:a.UplinkDataStatus.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return err
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case RegistrationRequestUEStatusType:
			a.UEStatus = nasType.NewUEStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UEStatus.Len); err != nil {
				return err
			}
			a.UEStatus.SetLen(a.UEStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.UEStatus.Octet); err != nil {
				return err
			}
		case RegistrationRequestAdditionalGUTIType:
			a.AdditionalGUTI = nasType.NewAdditionalGUTI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AdditionalGUTI.Len); err != nil {
				return err
			}
			a.AdditionalGUTI.SetLen(a.AdditionalGUTI.GetLen())
			if a.AdditionalGUTI.GetLen() > uint16(len(a.AdditionalGUTI.Octet)) {
				a.AdditionalGUTI = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: AdditionalGUTI length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.AdditionalGUTI.Octet[:a.AdditionalGUTI.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestAllowedPDUSessionStatusType:
			a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Len); err != nil {
				return err
			}
			a.AllowedPDUSessionStatus.SetLen(a.AllowedPDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer[:a.AllowedPDUSessionStatus.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestUesUsageSettingType:
			a.UesUsageSetting = nasType.NewUesUsageSetting(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UesUsageSetting.Len); err != nil {
				return err
			}
			a.UesUsageSetting.SetLen(a.UesUsageSetting.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.UesUsageSetting.Octet); err != nil {
				return err
			}
		case RegistrationRequestRequestedDRXParametersType:
			a.RequestedDRXParameters = nasType.NewRequestedDRXParameters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedDRXParameters.Len); err != nil {
				return err
			}
			a.RequestedDRXParameters.SetLen(a.RequestedDRXParameters.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedDRXParameters.Octet); err != nil {
				return err
			}
		case RegistrationRequestEPSNASMessageContainerType:
			a.EPSNASMessageContainer = nasType.NewEPSNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EPSNASMessageContainer.Len); err != nil {
				return err
			}
			a.EPSNASMessageContainer.SetLen(a.EPSNASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EPSNASMessageContainer.Buffer[:a.EPSNASMessageContainer.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestLADNIndicationType:
			a.LADNIndication = nasType.NewLADNIndication(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LADNIndication.Len); err != nil {
				return err
			}
			a.LADNIndication.SetLen(a.LADNIndication.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.LADNIndication.Buffer[:a.LADNIndication.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestPayloadContainerType:
			a.PayloadContainer = nasType.NewPayloadContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Len); err != nil {
				return err
			}
			a.PayloadContainer.SetLen(a.PayloadContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PayloadContainer.Buffer[:a.PayloadContainer.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case RegistrationRequestUpdateType5GSType:
			a.UpdateType5GS = nasType.NewUpdateType5GS(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UpdateType5GS.Len); err != nil {
				return err
			}
			a.UpdateType5GS.SetLen(a.UpdateType5GS.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.UpdateType5GS.Octet); err != nil {
				return err
			}
		case RegistrationRequestNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len); err != nil {
				return err
			}
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer[:a.NASMessageContainer.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestEPSBearerContextStatusType:
			a.EPSBearerContextStatus = nasType.NewEPSBearerContextStatus(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.EPSBearerContextStatus.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.EPSBearerContextStatus.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestExtendedDRXParametersType:
			a.ExtendedDRXParameters = nasType.NewExtendedDRXParameters(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.ExtendedDRXParameters.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedDRXParameters.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestUERadioCapabilityIDType:
			a.UERadioCapabilityID = nasType.NewUERadioCapabilityID(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.UERadioCapabilityID.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestRequestedMappedNSSAIType:
			a.RequestedMappedNSSAI = nasType.NewRequestedMappedNSSAI(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.RequestedMappedNSSAI.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.RequestedMappedNSSAI.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestAdditionalInformationRequestedType:
			a.AdditionalInformationRequested = nasType.NewAdditionalInformationRequested(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.AdditionalInformationRequested.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.AdditionalInformationRequested.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestWUSAssistanceInformationType:
			a.WUSAssistanceInformation = nasType.NewWUSAssistanceInformation(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.WUSAssistanceInformation.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.WUSAssistanceInformation.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestN5GCIndicationType:
			a.N5GCIndication = nasType.NewN5GCIndication(ieiN)
			a.N5GCIndication.Octet = ieiN
		case RegistrationRequestNBN1ModeDRXParametersType:
			a.NBN1ModeDRXParameters = nasType.NewNBN1ModeDRXParameters(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.NBN1ModeDRXParameters.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.NBN1ModeDRXParameters.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestUERequestTypeType:
			a.UERequestType = nasType.NewUERequestType(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.UERequestType.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.UERequestType.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestPagingRestrictionType:
			a.PagingRestriction = nasType.NewPagingRestriction(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.PagingRestriction.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.PagingRestriction.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len); err != nil {
				return err
			}
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()]); err != nil {
				return err
			}
		case RegistrationRequestNIDType:
			a.NID = nasType.NewNID(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.NID.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.NID.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestPLMNIdentityWithDisasterConditionType:
			a.PLMNIdentityWithDisasterCondition = nasType.NewPLMNIdentityWithDisasterCondition(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.PLMNIdentityWithDisasterCondition.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.PLMNIdentityWithDisasterCondition.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestPEIPSAssistanceInformationType:
			a.PEIPSAssistanceInformation = nasType.NewPEIPSAssistanceInformation(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.PEIPSAssistanceInformation.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.PEIPSAssistanceInformation.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestTimeDurationType:
			a.TimeDuration = nasType.NewTimeDuration(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.TimeDuration.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.TimeDuration.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestNon3GPPPathSwitchingInformationType:
			a.Non3GPPPathSwitchingInformation = nasType.NewNon3GPPPathSwitchingInformation(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.Non3GPPPathSwitchingInformation.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.Non3GPPPathSwitchingInformation.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestAUN3IndicationType:
			a.AUN3Indication = nasType.NewAUN3Indication(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AUN3Indication.Len); err != nil {
				return err
			}
			a.AUN3Indication.SetLen(a.AUN3Indication.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.AUN3Indication.Octet); err != nil {
				return err
			}
		case RegistrationRequestMobileStationClassmark2Type:
			a.MobileStationClassmark2 = nasType.NewMobileStationClassmark2(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.MobileStationClassmark2.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.MobileStationClassmark2.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestSupportedCodecsType:
			a.SupportedCodecs = nasType.NewSupportedCodecs(ieiN)
			var lenN uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN); err != nil {
				return err
			}
			a.SupportedCodecs.SetLen(uint16(lenN))
			if err := binary.Read(buffer, binary.BigEndian, a.SupportedCodecs.Buffer[:lenN]); err != nil {
				return err
			}
		case RegistrationRequestRequestedT3512ValueType:
			a.T3512Value = nasType.NewT3512Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3512Value.Len); err != nil {
				return err
			}
			a.T3512Value.SetLen(a.T3512Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3512Value.Octet); err != nil {
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
