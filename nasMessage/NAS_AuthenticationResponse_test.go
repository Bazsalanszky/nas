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

type nasMessageAuthenticationResponseData struct {
	inExtendedProtocolDiscriminator         uint8
	inSecurityHeader                        uint8
	inSpareHalfOctet                        uint8
	inAuthenticationResponseMessageIdentity uint8
	inAuthenticationResponseParameter       nasType.AuthenticationResponseParameter
	inEAPMessage                            nasType.EAPMessage
}

var nasMessageAuthenticationResponseTable = []nasMessageAuthenticationResponseData{
	{
		inExtendedProtocolDiscriminator:         0x01,
		inSecurityHeader:                        0x08,
		inSpareHalfOctet:                        0x01,
		inAuthenticationResponseMessageIdentity: 0x01,
		inAuthenticationResponseParameter:       nasType.AuthenticationResponseParameter{Iei: nasMessage.AuthenticationResponseAuthenticationResponseParameterType, Len: 16, Octet: [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inEAPMessage:                            nasType.EAPMessage{Iei: nasMessage.AuthenticationResponseEAPMessageType, Len: 2, Buffer: []uint8{0x01, 0x01}},
	},
}

func TestNasTypeNewAuthenticationResponse(t *testing.T) {
	a := nasMessage.NewAuthenticationResponse(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewAuthenticationResponseMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationResponseMessage---")
	for i, table := range nasMessageAuthenticationResponseTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewAuthenticationResponse(0)
		b := nasMessage.NewAuthenticationResponse(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inAuthenticationResponseMessageIdentity)

		a.AuthenticationResponseParameter = nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
		a.AuthenticationResponseParameter = &table.inAuthenticationResponseParameter

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationResponseEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeAuthenticationResponse(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
