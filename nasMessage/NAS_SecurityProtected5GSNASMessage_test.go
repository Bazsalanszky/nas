// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
)

type nasMessageSecurityProtected5GSNASMessageData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet                uint8
	inMessageAuthenticationCode     nasType.MessageAuthenticationCode
	inSequenceNumber                nasType.SequenceNumber
	inPlain5GSNASMessage            nasType.Plain5GSNASMessage
}

var nasMessageSecurityProtected5GSNASMessageTable = []nasMessageSecurityProtected5GSNASMessageData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inMessageAuthenticationCode: nasType.MessageAuthenticationCode{
			Octet: [4]uint8{0x01, 0x01, 0x01, 0x01},
		},
		inSequenceNumber: nasType.SequenceNumber{
			Octet: 0x01,
		},
		inPlain5GSNASMessage: nasType.Plain5GSNASMessage{},
	},
}

func TestNasTypeNewSecurityProtected5GSNASMessage(t *testing.T) {
	a := nasMessage.NewSecurityProtected5GSNASMessage(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewSecurityProtected5GSNASMessageMessage(t *testing.T) {
	for i, table := range nasMessageSecurityProtected5GSNASMessageTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityProtected5GSNASMessage(0)
		b := nasMessage.NewSecurityProtected5GSNASMessage(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)

		a.MessageAuthenticationCode = table.inMessageAuthenticationCode
		a.SequenceNumber = table.inSequenceNumber
		a.Plain5GSNASMessage = table.inPlain5GSNASMessage

		buff := new(bytes.Buffer)
		a.EncodeSecurityProtected5GSNASMessage(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeSecurityProtected5GSNASMessage(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}

// The octets after the security header are the protected plain NAS message, not
// optional IEs, so decoding must accept them.
func TestDecodeSecurityProtected5GSNASMessageWithPayload(t *testing.T) {
	data := []byte{
		nasMessage.Epd5GSMobilityManagementMessage, 0x01, // EPD, security header type
		0x01, 0x02, 0x03, 0x04, // MAC
		0x05,             // sequence number
		0x7e, 0x00, 0x43, // plain Registration Complete
	}
	msg := nasMessage.NewSecurityProtected5GSNASMessage(0)
	if err := msg.DecodeSecurityProtected5GSNASMessage(&data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.SequenceNumber.Octet != 0x05 {
		t.Errorf("sequence number = %d, want 5", msg.SequenceNumber.Octet)
	}
	if msg.MessageAuthenticationCode.Octet != [4]uint8{0x01, 0x02, 0x03, 0x04} {
		t.Errorf("MAC = %x, want 01020304", msg.MessageAuthenticationCode.Octet)
	}
}
