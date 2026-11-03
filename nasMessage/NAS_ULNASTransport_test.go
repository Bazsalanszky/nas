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

type nasMessageULNASTransportData struct {
	inExtendedProtocolDiscriminator         uint8
	inSecurityHeader                        uint8
	inSpareHalfOctet                        uint8
	inULNASTRANSPORTMessageIdentity         uint8
	inSpareHalfOctetAndPayloadContainerType nasType.SpareHalfOctetAndPayloadContainerType
	inPayloadContainer                      nasType.PayloadContainer
	inPduSessionID2Value                    nasType.PduSessionID2Value
	inOldPDUSessionID                       nasType.OldPDUSessionID
	inRequestType                           nasType.RequestType
	inSNSSAI                                nasType.SNSSAI
	inDNN                                   nasType.DNN
	inAdditionalInformation                 nasType.AdditionalInformation
}

var nasMessageULNASTransportTable = []nasMessageULNASTransportData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inULNASTRANSPORTMessageIdentity: nas.MsgTypeULNASTransport,
		inSpareHalfOctetAndPayloadContainerType: nasType.SpareHalfOctetAndPayloadContainerType{
			Octet: 0x01,
		},
		inPayloadContainer: nasType.PayloadContainer{
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPduSessionID2Value: nasType.PduSessionID2Value{
			Iei:   nasMessage.ULNASTransportPduSessionID2ValueType,
			Octet: 0x01,
		},
		inOldPDUSessionID: nasType.OldPDUSessionID{
			Iei:   nasMessage.ULNASTransportOldPDUSessionIDType,
			Octet: 0x01,
		},
		inRequestType: nasType.RequestType{
			Octet: 0x80,
		},
		inDNN: nasType.DNN{
			Iei:    nasMessage.ULNASTransportDNNType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inSNSSAI: nasType.SNSSAI{
			Iei:   nasMessage.ULNASTransportSNSSAIType,
			Len:   8,
			Octet: [8]uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		inAdditionalInformation: nasType.AdditionalInformation{
			Iei:    nasMessage.ULNASTransportAdditionalInformationType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewULNASTransport(t *testing.T) {
	a := nasMessage.NewULNASTransport(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewULNASTransportMessage(t *testing.T) {
	for i, table := range nasMessageULNASTransportTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewULNASTransport(0)
		b := nasMessage.NewULNASTransport(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inULNASTRANSPORTMessageIdentity)

		a.SpareHalfOctetAndPayloadContainerType = table.inSpareHalfOctetAndPayloadContainerType

		a.PayloadContainer = table.inPayloadContainer

		a.PduSessionID2Value = nasType.NewPduSessionID2Value(nasMessage.ULNASTransportPduSessionID2ValueType)
		a.PduSessionID2Value = &table.inPduSessionID2Value

		a.OldPDUSessionID = nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
		a.OldPDUSessionID = &table.inOldPDUSessionID

		a.RequestType = nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
		a.RequestType = &table.inRequestType

		a.SNSSAI = nasType.NewSNSSAI(nasMessage.ULNASTransportSNSSAIType)
		a.SNSSAI = &table.inSNSSAI

		a.DNN = nasType.NewDNN(nasMessage.ULNASTransportDNNType)
		a.DNN = &table.inDNN

		a.AdditionalInformation = nasType.NewAdditionalInformation(nasMessage.ULNASTransportAdditionalInformationType)
		a.AdditionalInformation = &table.inAdditionalInformation

		buff := new(bytes.Buffer)
		a.EncodeULNASTransport(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeULNASTransport(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}

func TestULNASTransportNewIEsEncodeDecode(t *testing.T) {
	a := nasMessage.NewULNASTransport(0)
	b := nasMessage.NewULNASTransport(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(nas.MsgTypeULNASTransport)
	a.SpareHalfOctetAndPayloadContainerType.Octet = 0x01
	a.PayloadContainer.SetLen(2)
	copy(a.PayloadContainer.Buffer, []uint8{0x01, 0x02})

	a.MAPDUSessionInformation = nasType.NewMAPDUSessionInformation(nasMessage.ULNASTransportMAPDUSessionInformationType)
	a.SetMAPSI(0x03)

	a.ReleaseAssistanceIndication = nasType.NewReleaseAssistanceIndication(nasMessage.ULNASTransportReleaseAssistanceIndicationType)
	a.SetPDDEI(0x02)

	a.Non3GPPAccessPathSwitchingIndication = nasType.NewNon3GPPAccessPathSwitchingIndication(nasMessage.ULNASTransportNon3GPPAccessPathSwitchingIndicationType)
	a.Non3GPPAccessPathSwitchingIndication.SetLen(1)
	a.SetNAPS(0x01)

	a.AlternativeSNSSAI = nasType.NewSNSSAI(nasMessage.ULNASTransportAlternativeSNSSAIType)
	a.AlternativeSNSSAI.SetLen(2)
	a.AlternativeSNSSAI.Octet[0] = 0x11
	a.AlternativeSNSSAI.Octet[1] = 0x22

	a.PayloadContainerInformation = nasType.NewPayloadContainerInformation(nasMessage.ULNASTransportPayloadContainerInformationType)
	a.SetPRU(0x05)

	buff := new(bytes.Buffer)
	a.EncodeULNASTransport(buff)
	logger.NasMsgLog.Debugln("Encode: ", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln(data)
	if err := b.DecodeULNASTransport(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("Decode: ", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ULNASTransport new IEs encode/decode mismatch")
	}
}
