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

type nasMessageAuthenticationRequestData struct {
	inExtendedProtocolDiscriminator        uint8
	inSecurityHeader                       uint8
	inSpareHalfOctet1                      uint8
	inAuthenticationRequestMessageIdentity uint8
	inTsc                                  uint8
	inNASKeySetIdentifier                  uint8
	inSpareHalfOctet2                      uint8
	inABBA                                 nasType.ABBA
	inAuthenticationParameterRAND          nasType.AuthenticationParameterRAND
	inAuthenticationParameterAUTN          nasType.AuthenticationParameterAUTN
	inEAPMessage                           nasType.EAPMessage
}

var nasMessageAuthenticationRequestTable = []nasMessageAuthenticationRequestData{
	{
		inExtendedProtocolDiscriminator:        0x01,
		inSecurityHeader:                       0x08,
		inSpareHalfOctet1:                      0x01,
		inAuthenticationRequestMessageIdentity: 0x01,
		inTsc:                                  0x01,
		inNASKeySetIdentifier:                  0x07,
		inSpareHalfOctet2:                      0x07,
		inABBA:                                 nasType.ABBA{Iei: 0, Len: 2, Buffer: []byte{0x00, 0x00}},
		inAuthenticationParameterRAND:          nasType.AuthenticationParameterRAND{Iei: nasMessage.AuthenticationRequestAuthenticationParameterRANDType, Octet: [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inAuthenticationParameterAUTN:          nasType.AuthenticationParameterAUTN{Iei: nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, Len: 16, Octet: [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inEAPMessage:                           nasType.EAPMessage{Iei: nasMessage.AuthenticationRequestEAPMessageType, Len: 2, Buffer: []byte{0x00, 0x00}},
	},
}

func TestNasTypeNewAuthenticationRequest(t *testing.T) {
	a := nasMessage.NewAuthenticationRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewAuthenticationRequestMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationRequestMessage---")
	for i, table := range nasMessageAuthenticationRequestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewAuthenticationRequest(0)
		b := nasMessage.NewAuthenticationRequest(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.SetMessageType(table.inAuthenticationRequestMessageIdentity)

		a.ABBA = table.inABBA

		a.AuthenticationParameterRAND = nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
		a.AuthenticationParameterRAND = &table.inAuthenticationParameterRAND

		a.AuthenticationParameterAUTN = nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
		a.AuthenticationParameterAUTN = &table.inAuthenticationParameterAUTN

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRequestEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationRequest(buff)
		logger.NasMsgLog.Debugln("encode:", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeAuthenticationRequest(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("decode:", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
