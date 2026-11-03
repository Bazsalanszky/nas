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

type nasMessageAuthenticationRejectData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inAuthenticationRejectMessageIdentity uint8
	inEAPMessage                          nasType.EAPMessage
}

var nasMessageAuthenticationRejectTable = []nasMessageAuthenticationRejectData{
	{
		inExtendedProtocolDiscriminator:       0x01,
		inSecurityHeader:                      0x01,
		inSpareHalfOctet:                      0x01,
		inAuthenticationRejectMessageIdentity: 0x01,
		inEAPMessage:                          nasType.EAPMessage{Iei: nasMessage.AuthenticationRejectEAPMessageType, Len: 2, Buffer: []byte{0x00, 0x00}},
	},
}

func TestNasTypeNewAuthenticationReject(t *testing.T) {
	a := nasMessage.NewAuthenticationReject(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewAuthenticationRejectMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationRejectMessage---")
	for i, table := range nasMessageAuthenticationRejectTable {
		t.Logf("Test Cnt:%d", i)

		a := nasMessage.NewAuthenticationReject(nasMessage.AuthenticationRejectEAPMessageType)
		b := nasMessage.NewAuthenticationReject(nasMessage.AuthenticationRejectEAPMessageType)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inAuthenticationRejectMessageIdentity)

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRejectEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationReject(buff)
		logger.NasMsgLog.Debugln(a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeAuthenticationReject(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("decode:", data)
		logger.NasMsgLog.Infoln(b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
