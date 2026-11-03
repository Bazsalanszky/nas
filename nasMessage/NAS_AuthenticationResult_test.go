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

type nasMessageAuthenticationResultData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeaderType            uint8
	inMessageType                   uint8
	inTsc                           uint8
	inNASKeySetIdentifier           uint8
	inEAPLen                        uint16
	inEAPMessage                    []uint8
	inABBA                          nasType.ABBA
}

var aBBATestData = []nasType.ABBA{
	{Iei: nasMessage.AuthenticationResultABBAType, Len: 2, Buffer: []byte{0x00, 0x00}},
	//{Iei: 0x81, Len: 2, Buffer: []byte{0x00, 0x00}},
}

var nasMessageAuthenticationResultTable = []nasMessageAuthenticationResultData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:            0x01,
		inMessageType:                   nas.MsgTypeAuthenticationResult,
		inTsc:                           0x01,
		inNASKeySetIdentifier:           0x01,
		inEAPLen:                        0x02,
		inEAPMessage:                    []uint8{0x10, 0x11},
		inABBA:                          aBBATestData[0],
	},
	/*{inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
	inSecurityHeaderType:  0x01,
	inMessageType:         nas.MsgTypeAuthenticationResult,
	inTsc:                 0x01,
	inNASKeySetIdentifier: 0x01,
	inEAPLen:              0x02,
	inEAPMessage:          []uint8{0x10, 0x11},
	inABBA:                aBBATestData[1]},*/
}

func TestNasTypeNewAuthenticationResult(t *testing.T) {
	a := nasMessage.NewAuthenticationResult(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewAuthenticationResultMessage(t *testing.T) {
	for i, table := range nasMessageAuthenticationResultTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewAuthenticationResult(0)
		b := nasMessage.NewAuthenticationResult(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SetMessageType(table.inMessageType)
		a.SetTSC(table.inTsc)
		a.SetNasKeySetIdentifiler(table.inNASKeySetIdentifier)
		a.EAPMessage.SetLen(table.inEAPLen)
		a.SetEAPMessage(table.inEAPMessage)

		a.ABBA = nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
		a.ABBA = &table.inABBA

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationResult(buff)
		logger.NasMsgLog.Debugln(buff)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeAuthenticationResult(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
