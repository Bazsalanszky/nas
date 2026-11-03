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

type ConfigurationUpdateCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ConfigurationUpdateCommandMessageIdentity
	*nasType.ConfigurationUpdateIndication
	*nasType.GUTI5G
	*nasType.TAIList
	*nasType.AllowedNSSAI
	*nasType.ServiceAreaList
	*nasType.FullNameForNetwork
	*nasType.ShortNameForNetwork
	*nasType.LocalTimeZone
	*nasType.UniversalTimeAndLocalTimeZone
	*nasType.NetworkDaylightSavingTime
	*nasType.LADNInformation
	*nasType.MICOIndication
	*nasType.NetworkSlicingIndication
	*nasType.ConfiguredNSSAI
	*nasType.RejectedNSSAI
	*nasType.OperatordefinedAccessCategoryDefinitions
	*nasType.SMSIndication
	*nasType.CAGInformationList
	*nasType.UERadioCapabilityID
	*nasType.UERadioCapabilityIDDeletionIndicationIE
	*nasType.TruncatedFiveGSTMSIConfiguration
	*nasType.ExtendedRejectedNSSAI
	*nasType.ServiceLevelAAContainer
	*nasType.NSSRGInformation
	*nasType.RegistrationWaitRange
	DisasterReturnWaitRange *nasType.RegistrationWaitRange
	*nasType.ListOfPLMNsForDisasterCondition
	*nasType.ExtendedCAGInformationList
	*nasType.NSAGInformation
	*nasType.ExtendedLADNInformation
	*nasType.RegistrationResult5GS
	*nasType.AdditionalConfigurationIndication
	*nasType.UpdatedPEIPSAssistanceInformation
	*nasType.PriorityIndicator
	*nasType.RANTimingSynchronization
	*nasType.AlternativeNSSAI
	*nasType.SNSSAILocationValidityInformation
	*nasType.SNSSAITimeValidityInformation
	*nasType.DiscontinuousCoverageMaxTimeOffset
	*nasType.PartiallyAllowedNSSAI
	*nasType.PartiallyRejectedNSSAI
	*nasType.FeatureAuthorizationIndication
}

func NewConfigurationUpdateCommand(iei uint8) (configurationUpdateCommand *ConfigurationUpdateCommand) {
	configurationUpdateCommand = &ConfigurationUpdateCommand{}
	return configurationUpdateCommand
}

const (
	ConfigurationUpdateCommandConfigurationUpdateIndicationType            uint8 = 0x0D
	ConfigurationUpdateCommandGUTI5GType                                   uint8 = 0x77
	ConfigurationUpdateCommandTAIListType                                  uint8 = 0x54
	ConfigurationUpdateCommandAllowedNSSAIType                             uint8 = 0x15
	ConfigurationUpdateCommandServiceAreaListType                          uint8 = 0x27
	ConfigurationUpdateCommandFullNameForNetworkType                       uint8 = 0x43
	ConfigurationUpdateCommandShortNameForNetworkType                      uint8 = 0x45
	ConfigurationUpdateCommandLocalTimeZoneType                            uint8 = 0x46
	ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType            uint8 = 0x47
	ConfigurationUpdateCommandNetworkDaylightSavingTimeType                uint8 = 0x49
	ConfigurationUpdateCommandLADNInformationType                          uint8 = 0x79
	ConfigurationUpdateCommandMICOIndicationType                           uint8 = 0x0B
	ConfigurationUpdateCommandNetworkSlicingIndicationType                 uint8 = 0x09
	ConfigurationUpdateCommandConfiguredNSSAIType                          uint8 = 0x31
	ConfigurationUpdateCommandRejectedNSSAIType                            uint8 = 0x11
	ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType uint8 = 0x76
	ConfigurationUpdateCommandSMSIndicationType                            uint8 = 0x0F
	ConfigurationUpdateCommandCAGInformationListType                       uint8 = 0x75
	ConfigurationUpdateCommandUERadioCapabilityIDType                      uint8 = 0x67
	ConfigurationUpdateCommandUERadioCapabilityIDDeletionIndicationType    uint8 = 0x0A
	ConfigurationUpdateCommandTruncatedFiveGSTMSIConfigurationType         uint8 = 0x1B
	ConfigurationUpdateCommandExtendedRejectedNSSAIType                    uint8 = 0x68
	ConfigurationUpdateCommandServiceLevelAAContainerType                  uint8 = 0x72
	ConfigurationUpdateCommandNSSRGInformationType                         uint8 = 0x70
	ConfigurationUpdateCommandRegistrationWaitRangeType                    uint8 = 0x14
	ConfigurationUpdateCommandDisasterReturnWaitRangeType                  uint8 = 0x2C
	ConfigurationUpdateCommandListOfPLMNsForDisasterConditionType          uint8 = 0x13
	ConfigurationUpdateCommandExtendedCAGInformationListType               uint8 = 0x71
	ConfigurationUpdateCommandNSAGInformationType                          uint8 = 0x73
	ConfigurationUpdateCommandExtendedLADNInformationType                  uint8 = 0x78
	ConfigurationUpdateCommandRegistrationResult5GSType                    uint8 = 0x44
	ConfigurationUpdateCommandAdditionalConfigurationIndicationType        uint8 = 0x0C
	ConfigurationUpdateCommandUpdatedPEIPSAssistanceInformationType        uint8 = 0x1F
	ConfigurationUpdateCommandPriorityIndicatorType                        uint8 = 0x0E
	ConfigurationUpdateCommandRANTimingSynchronizationType                 uint8 = 0x4B
	ConfigurationUpdateCommandAlternativeNSSAIType                         uint8 = 0x4C
	ConfigurationUpdateCommandSNSSAILocationValidityInformationType        uint8 = 0x7D
	ConfigurationUpdateCommandSNSSAITimeValidityInformationType            uint8 = 0x5B
	ConfigurationUpdateCommandDiscontinuousCoverageMaxTimeOffsetType       uint8 = 0x4F
	ConfigurationUpdateCommandPartiallyAllowedNSSAIType                    uint8 = 0x74
	ConfigurationUpdateCommandPartiallyRejectedNSSAIType                   uint8 = 0x7A
	ConfigurationUpdateCommandFeatureAuthorizationIndicationType           uint8 = 0x5C
)

func (a *ConfigurationUpdateCommand) EncodeConfigurationUpdateCommand(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ConfigurationUpdateCommandMessageIdentity.Octet); err != nil {
		return
	}
	if a.ConfigurationUpdateIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.ConfigurationUpdateIndication.Octet); err != nil {
			return
		}
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
	if a.FullNameForNetwork != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.FullNameForNetwork.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.FullNameForNetwork.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.FullNameForNetwork.Buffer); err != nil {
			return
		}
	}
	if a.ShortNameForNetwork != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ShortNameForNetwork.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ShortNameForNetwork.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ShortNameForNetwork.Buffer); err != nil {
			return
		}
	}
	if a.LocalTimeZone != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LocalTimeZone.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.LocalTimeZone.Octet); err != nil {
			return
		}
	}
	if a.UniversalTimeAndLocalTimeZone != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UniversalTimeAndLocalTimeZone.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UniversalTimeAndLocalTimeZone.Octet); err != nil {
			return
		}
	}
	if a.NetworkDaylightSavingTime != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkDaylightSavingTime.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkDaylightSavingTime.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Octet); err != nil {
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
	if a.SMSIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.SMSIndication.Octet); err != nil {
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
	if a.ExtendedLADNInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedLADNInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedLADNInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedLADNInformation.Buffer); err != nil {
			return
		}
	}
	if a.RegistrationResult5GS != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RegistrationResult5GS.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet); err != nil {
			return
		}
	}
	if a.AdditionalConfigurationIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.AdditionalConfigurationIndication.Octet); err != nil {
			return
		}
	}
	if a.UpdatedPEIPSAssistanceInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UpdatedPEIPSAssistanceInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UpdatedPEIPSAssistanceInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UpdatedPEIPSAssistanceInformation.Buffer); err != nil {
			return
		}
	}
	if a.PriorityIndicator != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.PriorityIndicator.Octet); err != nil {
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
	if a.SNSSAILocationValidityInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAILocationValidityInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAILocationValidityInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.SNSSAILocationValidityInformation.Buffer); err != nil {
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
	if a.PartiallyAllowedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PartiallyAllowedNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PartiallyAllowedNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PartiallyAllowedNSSAI.Buffer); err != nil {
			return
		}
	}
	if a.PartiallyRejectedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PartiallyRejectedNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PartiallyRejectedNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PartiallyRejectedNSSAI.Buffer); err != nil {
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

func (a *ConfigurationUpdateCommand) DecodeConfigurationUpdateCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ConfigurationUpdateCommandMessageIdentity.Octet); err != nil {
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
		case ConfigurationUpdateCommandConfigurationUpdateIndicationType:
			a.ConfigurationUpdateIndication = nasType.NewConfigurationUpdateIndication(ieiN)
			a.ConfigurationUpdateIndication.Octet = ieiN
		case ConfigurationUpdateCommandGUTI5GType:
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
		case ConfigurationUpdateCommandTAIListType:
			a.TAIList = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.TAIList.Len); err != nil {
				return err
			}
			a.TAIList.SetLen(a.TAIList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.TAIList.Buffer[:a.TAIList.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandAllowedNSSAIType:
			a.AllowedNSSAI = nasType.NewAllowedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedNSSAI.Len); err != nil {
				return err
			}
			a.AllowedNSSAI.SetLen(a.AllowedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer[:a.AllowedNSSAI.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandServiceAreaListType:
			a.ServiceAreaList = nasType.NewServiceAreaList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceAreaList.Len); err != nil {
				return err
			}
			a.ServiceAreaList.SetLen(a.ServiceAreaList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceAreaList.Buffer[:a.ServiceAreaList.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandFullNameForNetworkType:
			a.FullNameForNetwork = nasType.NewFullNameForNetwork(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.FullNameForNetwork.Len); err != nil {
				return err
			}
			a.FullNameForNetwork.SetLen(a.FullNameForNetwork.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.FullNameForNetwork.Buffer[:a.FullNameForNetwork.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandShortNameForNetworkType:
			a.ShortNameForNetwork = nasType.NewShortNameForNetwork(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ShortNameForNetwork.Len); err != nil {
				return err
			}
			a.ShortNameForNetwork.SetLen(a.ShortNameForNetwork.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ShortNameForNetwork.Buffer[:a.ShortNameForNetwork.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandLocalTimeZoneType:
			a.LocalTimeZone = nasType.NewLocalTimeZone(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LocalTimeZone.Octet); err != nil {
				return err
			}
		case ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType:
			a.UniversalTimeAndLocalTimeZone = nasType.NewUniversalTimeAndLocalTimeZone(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UniversalTimeAndLocalTimeZone.Octet); err != nil {
				return err
			}
		case ConfigurationUpdateCommandNetworkDaylightSavingTimeType:
			a.NetworkDaylightSavingTime = nasType.NewNetworkDaylightSavingTime(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Len); err != nil {
				return err
			}
			a.NetworkDaylightSavingTime.SetLen(a.NetworkDaylightSavingTime.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Octet); err != nil {
				return err
			}
		case ConfigurationUpdateCommandLADNInformationType:
			a.LADNInformation = nasType.NewLADNInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LADNInformation.Len); err != nil {
				return err
			}
			a.LADNInformation.SetLen(a.LADNInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.LADNInformation.Buffer[:a.LADNInformation.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case ConfigurationUpdateCommandNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case ConfigurationUpdateCommandConfiguredNSSAIType:
			a.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Len); err != nil {
				return err
			}
			a.ConfiguredNSSAI.SetLen(a.ConfiguredNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer[:a.ConfiguredNSSAI.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len); err != nil {
				return err
			}
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer[:a.RejectedNSSAI.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType:
			a.OperatordefinedAccessCategoryDefinitions = nasType.NewOperatordefinedAccessCategoryDefinitions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Len); err != nil {
				return err
			}
			a.OperatordefinedAccessCategoryDefinitions.SetLen(a.OperatordefinedAccessCategoryDefinitions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer[:a.OperatordefinedAccessCategoryDefinitions.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandSMSIndicationType:
			a.SMSIndication = nasType.NewSMSIndication(ieiN)
			a.SMSIndication.Octet = ieiN
		case ConfigurationUpdateCommandCAGInformationListType:
			a.CAGInformationList = nasType.NewCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CAGInformationList.Len); err != nil {
				return err
			}
			a.CAGInformationList.SetLen(a.CAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.CAGInformationList.Buffer[:a.CAGInformationList.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandUERadioCapabilityIDType:
			a.UERadioCapabilityID = nasType.NewUERadioCapabilityID(ieiN)
			var lenN0 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN0); err != nil {
				return err
			}
			a.UERadioCapabilityID.SetLen(uint16(lenN0))
			if err := binary.Read(buffer, binary.BigEndian, a.UERadioCapabilityID.Buffer[:lenN0]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandUERadioCapabilityIDDeletionIndicationType:
			a.UERadioCapabilityIDDeletionIndicationIE = nasType.NewUERadioCapabilityIDDeletionIndicationIE(ieiN)
			a.UERadioCapabilityIDDeletionIndicationIE.Octet = ieiN
		case ConfigurationUpdateCommandTruncatedFiveGSTMSIConfigurationType:
			a.TruncatedFiveGSTMSIConfiguration = nasType.NewTruncatedFiveGSTMSIConfiguration(ieiN)
			var lenN1 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN1); err != nil {
				return err
			}
			a.TruncatedFiveGSTMSIConfiguration.SetLen(uint16(lenN1))
			if err := binary.Read(buffer, binary.BigEndian, a.TruncatedFiveGSTMSIConfiguration.Buffer[:lenN1]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandExtendedRejectedNSSAIType:
			a.ExtendedRejectedNSSAI = nasType.NewExtendedRejectedNSSAI(ieiN)
			var lenN2 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN2); err != nil {
				return err
			}
			a.ExtendedRejectedNSSAI.SetLen(uint16(lenN2))
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedRejectedNSSAI.Buffer[:lenN2]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandServiceLevelAAContainerType:
			a.ServiceLevelAAContainer = nasType.NewServiceLevelAAContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceLevelAAContainer.Len); err != nil {
				return err
			}
			a.ServiceLevelAAContainer.SetLen(a.ServiceLevelAAContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceLevelAAContainer.Buffer[:a.ServiceLevelAAContainer.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandNSSRGInformationType:
			a.NSSRGInformation = nasType.NewNSSRGInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NSSRGInformation.Len); err != nil {
				return err
			}
			a.NSSRGInformation.SetLen(a.NSSRGInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NSSRGInformation.Buffer[:a.NSSRGInformation.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandRegistrationWaitRangeType:
			a.RegistrationWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN3 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN3); err != nil {
				return err
			}
			a.RegistrationWaitRange.SetLen(uint16(lenN3))
			if err := binary.Read(buffer, binary.BigEndian, a.RegistrationWaitRange.Buffer[:lenN3]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandDisasterReturnWaitRangeType:
			a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(ieiN)
			var lenN3a uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN3a); err != nil {
				return err
			}
			a.DisasterReturnWaitRange.SetLen(uint16(lenN3a))
			if err := binary.Read(buffer, binary.BigEndian, a.DisasterReturnWaitRange.Buffer[:lenN3a]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandListOfPLMNsForDisasterConditionType:
			a.ListOfPLMNsForDisasterCondition = nasType.NewListOfPLMNsForDisasterCondition(ieiN)
			var lenN4 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN4); err != nil {
				return err
			}
			a.ListOfPLMNsForDisasterCondition.SetLen(uint16(lenN4))
			if err := binary.Read(buffer, binary.BigEndian, a.ListOfPLMNsForDisasterCondition.Buffer[:lenN4]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandExtendedCAGInformationListType:
			a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedCAGInformationList.Len); err != nil {
				return err
			}
			a.ExtendedCAGInformationList.SetLen(a.ExtendedCAGInformationList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedCAGInformationList.Buffer[:a.ExtendedCAGInformationList.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandNSAGInformationType:
			a.NSAGInformation = nasType.NewNSAGInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NSAGInformation.Len); err != nil {
				return err
			}
			a.NSAGInformation.SetLen(a.NSAGInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NSAGInformation.Buffer[:a.NSAGInformation.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandExtendedLADNInformationType:
			a.ExtendedLADNInformation = nasType.NewExtendedLADNInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedLADNInformation.Len); err != nil {
				return err
			}
			a.ExtendedLADNInformation.SetLen(a.ExtendedLADNInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedLADNInformation.Buffer[:a.ExtendedLADNInformation.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandRegistrationResult5GSType:
			a.RegistrationResult5GS = nasType.NewRegistrationResult5GS(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Len); err != nil {
				return err
			}
			a.RegistrationResult5GS.SetLen(a.RegistrationResult5GS.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.RegistrationResult5GS.Octet); err != nil {
				return err
			}
		case ConfigurationUpdateCommandAdditionalConfigurationIndicationType:
			a.AdditionalConfigurationIndication = nasType.NewAdditionalConfigurationIndication(ieiN)
			a.AdditionalConfigurationIndication.Octet = ieiN
		case ConfigurationUpdateCommandUpdatedPEIPSAssistanceInformationType:
			a.UpdatedPEIPSAssistanceInformation = nasType.NewUpdatedPEIPSAssistanceInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UpdatedPEIPSAssistanceInformation.Len); err != nil {
				return err
			}
			a.UpdatedPEIPSAssistanceInformation.SetLen(a.UpdatedPEIPSAssistanceInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.UpdatedPEIPSAssistanceInformation.Buffer[:a.UpdatedPEIPSAssistanceInformation.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandPriorityIndicatorType:
			a.PriorityIndicator = nasType.NewPriorityIndicator(ieiN)
			a.PriorityIndicator.Octet = ieiN
		case ConfigurationUpdateCommandRANTimingSynchronizationType:
			a.RANTimingSynchronization = nasType.NewRANTimingSynchronization(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RANTimingSynchronization.Len); err != nil {
				return err
			}
			a.RANTimingSynchronization.SetLen(a.RANTimingSynchronization.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RANTimingSynchronization.Buffer[:a.RANTimingSynchronization.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandAlternativeNSSAIType:
			a.AlternativeNSSAI = nasType.NewAlternativeNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AlternativeNSSAI.Len); err != nil {
				return err
			}
			a.AlternativeNSSAI.SetLen(a.AlternativeNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AlternativeNSSAI.Buffer[:a.AlternativeNSSAI.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandSNSSAILocationValidityInformationType:
			a.SNSSAILocationValidityInformation = nasType.NewSNSSAILocationValidityInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAILocationValidityInformation.Len); err != nil {
				return err
			}
			a.SNSSAILocationValidityInformation.SetLen(a.SNSSAILocationValidityInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SNSSAILocationValidityInformation.Buffer[:a.SNSSAILocationValidityInformation.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandSNSSAITimeValidityInformationType:
			a.SNSSAITimeValidityInformation = nasType.NewSNSSAITimeValidityInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAITimeValidityInformation.Len); err != nil {
				return err
			}
			a.SNSSAITimeValidityInformation.SetLen(a.SNSSAITimeValidityInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SNSSAITimeValidityInformation.Buffer[:a.SNSSAITimeValidityInformation.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandDiscontinuousCoverageMaxTimeOffsetType:
			a.DiscontinuousCoverageMaxTimeOffset = nasType.NewDiscontinuousCoverageMaxTimeOffset(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Len); err != nil {
				return err
			}
			a.DiscontinuousCoverageMaxTimeOffset.SetLen(a.DiscontinuousCoverageMaxTimeOffset.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.DiscontinuousCoverageMaxTimeOffset.Octet); err != nil {
				return err
			}
		case ConfigurationUpdateCommandPartiallyAllowedNSSAIType:
			a.PartiallyAllowedNSSAI = nasType.NewPartiallyAllowedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PartiallyAllowedNSSAI.Len); err != nil {
				return err
			}
			a.PartiallyAllowedNSSAI.SetLen(a.PartiallyAllowedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PartiallyAllowedNSSAI.Buffer[:a.PartiallyAllowedNSSAI.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandPartiallyRejectedNSSAIType:
			a.PartiallyRejectedNSSAI = nasType.NewPartiallyRejectedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PartiallyRejectedNSSAI.Len); err != nil {
				return err
			}
			a.PartiallyRejectedNSSAI.SetLen(a.PartiallyRejectedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PartiallyRejectedNSSAI.Buffer[:a.PartiallyRejectedNSSAI.GetLen()]); err != nil {
				return err
			}
		case ConfigurationUpdateCommandFeatureAuthorizationIndicationType:
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
