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

type nasMessageServiceRejectData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet                uint8
	inServiceRejectMessageIdentity  uint8
	inCause5GMM                     nasType.Cause5GMM
	inPDUSessionStatus              nasType.PDUSessionStatus
	inT3346Value                    nasType.T3346Value
	inEAPMessage                    nasType.EAPMessage
}

var nasMessageServiceRejectTable = []nasMessageServiceRejectData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inServiceRejectMessageIdentity:  nas.MsgTypeServiceReject,
		inCause5GMM: nasType.Cause5GMM{
			Octet: 0x01,
		},
		inPDUSessionStatus: nasType.PDUSessionStatus{
			Iei:    nasMessage.ServiceRejectPDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inT3346Value: nasType.T3346Value{
			Iei:   nasMessage.ServiceRejectT3346ValueType,
			Len:   2,
			Octet: 0x01,
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.ServiceRejectEAPMessageType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewServiceReject(t *testing.T) {
	a := nasMessage.NewServiceReject(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewServiceRejectMessage(t *testing.T) {
	for i, table := range nasMessageServiceRejectTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewServiceReject(0)
		b := nasMessage.NewServiceReject(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inServiceRejectMessageIdentity)

		a.Cause5GMM = table.inCause5GMM

		a.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.ServiceRejectPDUSessionStatusType)
		a.PDUSessionStatus = &table.inPDUSessionStatus

		a.T3346Value = nasType.NewT3346Value(nasMessage.ServiceRejectT3346ValueType)
		a.T3346Value = &table.inT3346Value

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.ServiceRejectEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeServiceReject(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeServiceReject(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
