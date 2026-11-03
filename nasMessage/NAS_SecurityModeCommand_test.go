// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

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

type nasMessageSecurityModeCommandData struct {
	inExtendedProtocolDiscriminator      uint8
	inSecurityHeader                     uint8
	inSpareHalfOctet                     uint8
	inSecurityModeCommandMessageIdentity uint8
	inSelectedNASSecurityAlgorithms      nasType.SelectedNASSecurityAlgorithms
	inNgksi                              uint8
	inReplayedUESecurityCapabilities     nasType.ReplayedUESecurityCapabilities
	inIMEISVRequest                      nasType.IMEISVRequest
	inSelectedEPSNASSecurityAlgorithms   nasType.SelectedEPSNASSecurityAlgorithms
	inAdditional5GSecurityInformation    nasType.Additional5GSecurityInformation
	inEAPMessage                         nasType.EAPMessage
	inABBA                               nasType.ABBA
	inReplayedS1UESecurityCapabilities   nasType.ReplayedS1UESecurityCapabilities
}

var nasMessageSecurityModeCommandTable = []nasMessageSecurityModeCommandData{
	{
		inExtendedProtocolDiscriminator:      nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                     0x01,
		inSpareHalfOctet:                     0x01,
		inSecurityModeCommandMessageIdentity: nas.MsgTypeSecurityModeCommand,
		inSelectedNASSecurityAlgorithms: nasType.SelectedNASSecurityAlgorithms{
			Octet: 0x01,
		},
		inNgksi: 0x01,
		inReplayedUESecurityCapabilities: nasType.ReplayedUESecurityCapabilities{
			Len:    8,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		inIMEISVRequest: nasType.IMEISVRequest{
			Octet: 0xE0,
		},
		inSelectedEPSNASSecurityAlgorithms: nasType.SelectedEPSNASSecurityAlgorithms{
			Iei:   nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType,
			Octet: 0x01,
		},
		inAdditional5GSecurityInformation: nasType.Additional5GSecurityInformation{
			Iei:   nasMessage.SecurityModeCommandAdditional5GSecurityInformationType,
			Len:   2,
			Octet: 0x01,
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.SecurityModeCommandEAPMessageType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inABBA: nasType.ABBA{
			Iei:    nasMessage.SecurityModeCommandABBAType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inReplayedS1UESecurityCapabilities: nasType.ReplayedS1UESecurityCapabilities{
			Iei:    nasMessage.SecurityModeCommandReplayedS1UESecurityCapabilitiesType,
			Len:    5,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01, 0x01},
		},
	},
}

func TestNasTypeNewSecurityModeCommand(t *testing.T) {
	a := nasMessage.NewSecurityModeCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewSecurityModeCommandMessage(t *testing.T) {
	for i, table := range nasMessageSecurityModeCommandTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityModeCommand(0)
		b := nasMessage.NewSecurityModeCommand(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inSecurityModeCommandMessageIdentity)

		a.SelectedNASSecurityAlgorithms = table.inSelectedNASSecurityAlgorithms
		a.SpareHalfOctetAndNgksi.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetNasKeySetIdentifiler(table.inNgksi)

		a.ReplayedUESecurityCapabilities = table.inReplayedUESecurityCapabilities

		a.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		a.IMEISVRequest = &table.inIMEISVRequest

		a.SelectedEPSNASSecurityAlgorithms = nasType.NewSelectedEPSNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
		a.SelectedEPSNASSecurityAlgorithms = &table.inSelectedEPSNASSecurityAlgorithms

		a.Additional5GSecurityInformation = nasType.NewAdditional5GSecurityInformation(nasMessage.SecurityModeCommandAdditional5GSecurityInformationType)
		a.Additional5GSecurityInformation = &table.inAdditional5GSecurityInformation

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.SecurityModeCommandEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		a.ABBA = nasType.NewABBA(nasMessage.SecurityModeCommandABBAType)
		a.ABBA = &table.inABBA

		a.ReplayedS1UESecurityCapabilities = nasType.NewReplayedS1UESecurityCapabilities(nasMessage.SecurityModeCommandReplayedS1UESecurityCapabilitiesType)
		a.ReplayedS1UESecurityCapabilities = &table.inReplayedS1UESecurityCapabilities

		buff := new(bytes.Buffer)
		a.EncodeSecurityModeCommand(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeSecurityModeCommand(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
