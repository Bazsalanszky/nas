// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2"
	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
)

type nasMessageConfigurationUpdateCommandData struct {
	inExtendedProtocolDiscriminator             uint8
	inSecurityHeaderType                        uint8
	inSpareHalfOctet                            uint8
	inConfigurationUpdateCommandMessageIdentity uint8
	inConfigurationUpdateIndication             nasType.ConfigurationUpdateIndication
	inGUTI5G                                    nasType.GUTI5G
	inTAIList                                   nasType.TAIList
	inAllowedNSSAI                              nasType.AllowedNSSAI
	inServiceAreaList                           nasType.ServiceAreaList
	inFullNameForNetwork                        nasType.FullNameForNetwork
	inShortNameForNetwork                       nasType.ShortNameForNetwork
	inLocalTimeZone                             nasType.LocalTimeZone
	inUniversalTimeAndLocalTimeZone             nasType.UniversalTimeAndLocalTimeZone
	inNetworkDaylightSavingTime                 nasType.NetworkDaylightSavingTime
	inLADNInformation                           nasType.LADNInformation
	inMICOIndication                            nasType.MICOIndication
	inNetworkSlicingIndication                  nasType.NetworkSlicingIndication
	inConfiguredNSSAI                           nasType.ConfiguredNSSAI
	inRejectedNSSAI                             nasType.RejectedNSSAI
	inOperatordefinedAccessCategoryDefinitions  nasType.OperatordefinedAccessCategoryDefinitions
	inSMSIndication                             nasType.SMSIndication
}

var nasMessageConfigurationUpdateCommandTable = []nasMessageConfigurationUpdateCommandData{
	{
		inExtendedProtocolDiscriminator:             nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:                        0x01,
		inSpareHalfOctet:                            0x01,
		inConfigurationUpdateCommandMessageIdentity: nas.MsgTypeConfigurationUpdateCommand,
		inConfigurationUpdateIndication: nasType.ConfigurationUpdateIndication{
			Octet: 0xD0,
		},
		inGUTI5G: nasType.GUTI5G{
			Iei:   nasMessage.ConfigurationUpdateCommandGUTI5GType,
			Len:   11,
			Octet: [11]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B},
		},
		inTAIList: nasType.TAIList{
			Iei:    nasMessage.ConfigurationUpdateCommandTAIListType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inAllowedNSSAI: nasType.AllowedNSSAI{
			Iei:    nasMessage.ConfigurationUpdateCommandAllowedNSSAIType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inServiceAreaList: nasType.ServiceAreaList{
			Iei:    nasMessage.ConfigurationUpdateCommandServiceAreaListType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inFullNameForNetwork: nasType.FullNameForNetwork{
			Iei:    nasMessage.ConfigurationUpdateCommandFullNameForNetworkType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inShortNameForNetwork: nasType.ShortNameForNetwork{
			Iei:    nasMessage.ConfigurationUpdateCommandShortNameForNetworkType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inLocalTimeZone: nasType.LocalTimeZone{
			Iei:   nasMessage.ConfigurationUpdateCommandLocalTimeZoneType,
			Octet: 0x01,
		},
		inUniversalTimeAndLocalTimeZone: nasType.UniversalTimeAndLocalTimeZone{
			Iei:   nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType,
			Octet: [7]uint8{0x01},
		},
		inNetworkDaylightSavingTime: nasType.NetworkDaylightSavingTime{
			Iei:   nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType,
			Len:   2,
			Octet: 0x01,
		},
		inLADNInformation: nasType.LADNInformation{
			Iei:    nasMessage.ConfigurationUpdateCommandLADNInformationType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inMICOIndication: nasType.MICOIndication{
			Octet: 0xB0,
		},
		inNetworkSlicingIndication: nasType.NetworkSlicingIndication{
			Octet: 0x90,
		},
		inConfiguredNSSAI: nasType.ConfiguredNSSAI{
			Iei:    nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inRejectedNSSAI: nasType.RejectedNSSAI{
			Iei:    nasMessage.ConfigurationUpdateCommandRejectedNSSAIType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inOperatordefinedAccessCategoryDefinitions: nasType.OperatordefinedAccessCategoryDefinitions{
			Iei:    nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inSMSIndication: nasType.SMSIndication{
			Octet: 0xF0,
		},
	},
}

func TestNasTypeNewConfigurationUpdateCommand(t *testing.T) {
	a := nasMessage.NewConfigurationUpdateCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewConfigurationUpdateCommandMessage(t *testing.T) {
	for i, table := range nasMessageConfigurationUpdateCommandTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewConfigurationUpdateCommand(0)
		b := nasMessage.NewConfigurationUpdateCommand(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inConfigurationUpdateCommandMessageIdentity)

		a.ConfigurationUpdateIndication = nasType.NewConfigurationUpdateIndication(nasMessage.ConfigurationUpdateCommandConfigurationUpdateIndicationType)
		a.ConfigurationUpdateIndication = &table.inConfigurationUpdateIndication

		a.GUTI5G = nasType.NewGUTI5G(nasMessage.ConfigurationUpdateCommandGUTI5GType)
		a.GUTI5G = &table.inGUTI5G

		a.TAIList = nasType.NewTAIList(nasMessage.ConfigurationUpdateCommandTAIListType)
		a.TAIList = &table.inTAIList

		a.AllowedNSSAI = nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
		a.AllowedNSSAI = &table.inAllowedNSSAI

		a.ServiceAreaList = nasType.NewServiceAreaList(nasMessage.ConfigurationUpdateCommandServiceAreaListType)
		a.ServiceAreaList = &table.inServiceAreaList

		a.FullNameForNetwork = nasType.NewFullNameForNetwork(nasMessage.ConfigurationUpdateCommandFullNameForNetworkType)
		a.FullNameForNetwork = &table.inFullNameForNetwork

		a.ShortNameForNetwork = nasType.NewShortNameForNetwork(nasMessage.ConfigurationUpdateCommandShortNameForNetworkType)
		a.ShortNameForNetwork = &table.inShortNameForNetwork

		a.LocalTimeZone = nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
		a.LocalTimeZone = &table.inLocalTimeZone

		a.UniversalTimeAndLocalTimeZone = nasType.NewUniversalTimeAndLocalTimeZone(nasMessage.ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType)
		a.UniversalTimeAndLocalTimeZone = &table.inUniversalTimeAndLocalTimeZone

		a.NetworkDaylightSavingTime = nasType.NewNetworkDaylightSavingTime(nasMessage.ConfigurationUpdateCommandNetworkDaylightSavingTimeType)
		a.NetworkDaylightSavingTime = &table.inNetworkDaylightSavingTime

		a.LADNInformation = nasType.NewLADNInformation(nasMessage.ConfigurationUpdateCommandLADNInformationType)
		a.LADNInformation = &table.inLADNInformation

		a.MICOIndication = nasType.NewMICOIndication(nasMessage.ConfigurationUpdateCommandMICOIndicationType)
		a.MICOIndication = &table.inMICOIndication

		a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(nasMessage.ConfigurationUpdateCommandNetworkSlicingIndicationType)
		a.NetworkSlicingIndication = &table.inNetworkSlicingIndication

		a.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
		a.ConfiguredNSSAI = &table.inConfiguredNSSAI

		a.RejectedNSSAI = nasType.NewRejectedNSSAI(nasMessage.ConfigurationUpdateCommandRejectedNSSAIType)
		a.RejectedNSSAI = &table.inRejectedNSSAI

		a.OperatordefinedAccessCategoryDefinitions = nasType.NewOperatordefinedAccessCategoryDefinitions(nasMessage.ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType)
		a.OperatordefinedAccessCategoryDefinitions = &table.inOperatordefinedAccessCategoryDefinitions

		a.SMSIndication = nasType.NewSMSIndication(nasMessage.ConfigurationUpdateCommandSMSIndicationType)
		a.SMSIndication = &table.inSMSIndication

		buff := new(bytes.Buffer)
		a.EncodeConfigurationUpdateCommand(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeConfigurationUpdateCommand(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Dncode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}

func TestConfigurationUpdateCommandNewIEsEncodeDecode(t *testing.T) {
	a := nasMessage.NewConfigurationUpdateCommand(0)
	b := nasMessage.NewConfigurationUpdateCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(nas.MsgTypeConfigurationUpdateCommand)

	a.UERadioCapabilityID = nasType.NewUERadioCapabilityID(nasMessage.ConfigurationUpdateCommandUERadioCapabilityIDType)
	a.UERadioCapabilityID.SetLen(3)
	copy(a.UERadioCapabilityID.Buffer, []byte{0x01, 0x02, 0x03})

	a.TruncatedFiveGSTMSIConfiguration = nasType.NewTruncatedFiveGSTMSIConfiguration(nasMessage.ConfigurationUpdateCommandTruncatedFiveGSTMSIConfigurationType)
	a.TruncatedFiveGSTMSIConfiguration.SetLen(2)
	copy(a.TruncatedFiveGSTMSIConfiguration.Buffer, []byte{0x01, 0x02})

	a.ExtendedLADNInformation = nasType.NewExtendedLADNInformation(nasMessage.ConfigurationUpdateCommandExtendedLADNInformationType)
	a.ExtendedLADNInformation.SetLen(5)
	copy(a.ExtendedLADNInformation.Buffer, []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE})

	a.UERadioCapabilityIDDeletionIndicationIE = nasType.NewUERadioCapabilityIDDeletionIndicationIE(nasMessage.ConfigurationUpdateCommandUERadioCapabilityIDDeletionIndicationType)
	a.SetDeletionIndicationValue(0x01)

	a.DisasterReturnWaitRange = nasType.NewRegistrationWaitRange(nasMessage.ConfigurationUpdateCommandDisasterReturnWaitRangeType)
	a.DisasterReturnWaitRange.SetLen(2)
	copy(a.DisasterReturnWaitRange.Buffer, []byte{0x05, 0x0A})

	a.ExtendedCAGInformationList = nasType.NewExtendedCAGInformationList(nasMessage.ConfigurationUpdateCommandExtendedCAGInformationListType)
	a.ExtendedCAGInformationList.SetLen(3)
	copy(a.ExtendedCAGInformationList.Buffer, []byte{0x01, 0x02, 0x03})

	buff := new(bytes.Buffer)
	a.EncodeConfigurationUpdateCommand(buff)
	logger.NasMsgLog.Debugln("Encode: ", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeConfigurationUpdateCommand(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("Decode: ", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ConfigurationUpdateCommand new IEs encode/decode mismatch")
	}
}

func TestConfigurationUpdateCommandRel1718IEsEncodeDecode(t *testing.T) {
	a := nasMessage.NewConfigurationUpdateCommand(0)
	b := nasMessage.NewConfigurationUpdateCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(nas.MsgTypeConfigurationUpdateCommand)

	a.RegistrationResult5GS = nasType.NewRegistrationResult5GS(nasMessage.ConfigurationUpdateCommandRegistrationResult5GSType)
	a.RegistrationResult5GS.SetLen(1)
	a.RegistrationResult5GS.Octet = 0x01

	a.AdditionalConfigurationIndication = nasType.NewAdditionalConfigurationIndication(nasMessage.ConfigurationUpdateCommandAdditionalConfigurationIndicationType)
	a.SetSCMR(0x01)

	a.UpdatedPEIPSAssistanceInformation = nasType.NewUpdatedPEIPSAssistanceInformation(nasMessage.ConfigurationUpdateCommandUpdatedPEIPSAssistanceInformationType)
	a.UpdatedPEIPSAssistanceInformation.SetLen(2)
	copy(a.UpdatedPEIPSAssistanceInformation.Buffer, []byte{0xAA, 0xBB})

	a.PriorityIndicator = nasType.NewPriorityIndicator(nasMessage.ConfigurationUpdateCommandPriorityIndicatorType)
	a.SetMPSI(0x01)

	a.RANTimingSynchronization = nasType.NewRANTimingSynchronization(nasMessage.ConfigurationUpdateCommandRANTimingSynchronizationType)
	a.RANTimingSynchronization.SetLen(2)
	copy(a.RANTimingSynchronization.Buffer, []byte{0x01, 0x02})

	buff := new(bytes.Buffer)
	a.EncodeConfigurationUpdateCommand(buff)
	logger.NasMsgLog.Debugln("Encode: ", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeConfigurationUpdateCommand(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("Decode: ", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ConfigurationUpdateCommand Rel-17/18 IEs encode/decode mismatch")
	}
}
