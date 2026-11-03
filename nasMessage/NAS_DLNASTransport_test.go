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

type nasMessageDLNASTransportData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeaderType            uint8
	inSpareHalfOctet1               uint8
	inDLNASTRANSPORTMessageIdentity uint8
	inPayloadContainerType          uint8
	inSpareHalfOctet2               uint8
	inPayloadContainer              nasType.PayloadContainer
	inPduSessionID2Value            nasType.PduSessionID2Value
	inAdditionalInformation         nasType.AdditionalInformation
	inCause5GMM                     nasType.Cause5GMM
	inBackoffTimerValue             nasType.BackoffTimerValue
}

var nasMessageDLNASTransportTable = []nasMessageDLNASTransportData{
	{
		inExtendedProtocolDiscriminator: nas.MsgTypeDLNASTransport,
		inSecurityHeaderType:            0x01,
		inSpareHalfOctet1:               0x01,
		inDLNASTRANSPORTMessageIdentity: 0x01,
		inPayloadContainerType:          0x01,
		inSpareHalfOctet2:               0x01,
		inPayloadContainer: nasType.PayloadContainer{
			Iei:    0,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPduSessionID2Value: nasType.PduSessionID2Value{
			Iei:   nasMessage.DLNASTransportPduSessionID2ValueType,
			Octet: 0x01,
		},
		inAdditionalInformation: nasType.AdditionalInformation{
			Iei:    nasMessage.DLNASTransportAdditionalInformationType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inCause5GMM: nasType.Cause5GMM{
			Iei:   nasMessage.DLNASTransportCause5GMMType,
			Octet: 0xF0,
		},
		inBackoffTimerValue: nasType.BackoffTimerValue{
			Iei:   nasMessage.DLNASTransportBackoffTimerValueType,
			Len:   2,
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewDLNASTransport(t *testing.T) {
	a := nasMessage.NewDLNASTransport(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewDLNASTransportMessage(t *testing.T) {
	for i, table := range nasMessageDLNASTransportTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewDLNASTransport(0)
		b := nasMessage.NewDLNASTransport(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.SetMessageType(table.inDLNASTRANSPORTMessageIdentity)
		a.SetPayloadContainerType(table.inPayloadContainerType)
		a.PayloadContainer = table.inPayloadContainer

		a.PduSessionID2Value = nasType.NewPduSessionID2Value(nasMessage.DLNASTransportPduSessionID2ValueType)
		a.PduSessionID2Value = &table.inPduSessionID2Value

		a.AdditionalInformation = nasType.NewAdditionalInformation(nasMessage.DLNASTransportAdditionalInformationType)
		a.AdditionalInformation = &table.inAdditionalInformation

		a.Cause5GMM = nasType.NewCause5GMM(nasMessage.DLNASTransportCause5GMMType)
		a.Cause5GMM = &table.inCause5GMM

		a.BackoffTimerValue = nasType.NewBackoffTimerValue(nasMessage.DLNASTransportBackoffTimerValueType)
		a.BackoffTimerValue = &table.inBackoffTimerValue

		buff := new(bytes.Buffer)
		a.EncodeDLNASTransport(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeDLNASTransport(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
