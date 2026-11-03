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

type nasMessageAuthenticationFailureData struct {
	inExtendedProtocolDiscriminator        uint8
	inSecurityHeader                       uint8
	inSpareHalfOctet                       uint8
	inAuthenticationFailureMessageIdentity uint8
	in5GMMCause                            nasType.Cause5GMM
	inAuthenticationFailureParameter       nasType.AuthenticationFailureParameter
}

var nasMessageAuthenticationFailureTable = []nasMessageAuthenticationFailureData{
	{
		inExtendedProtocolDiscriminator:        0x01,
		inSecurityHeader:                       0x08,
		inSpareHalfOctet:                       0x01,
		inAuthenticationFailureMessageIdentity: 0x01,
		in5GMMCause:                            nasType.Cause5GMM{Iei: 0, Octet: 0xff},
		inAuthenticationFailureParameter:       nasType.AuthenticationFailureParameter{Iei: nasMessage.AuthenticationFailureAuthenticationFailureParameterType, Len: 14, Octet: [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	},
	{
		inExtendedProtocolDiscriminator:        0x01,
		inSecurityHeader:                       0x08,
		inSpareHalfOctet:                       0x01,
		inAuthenticationFailureMessageIdentity: 0x01,
		in5GMMCause:                            nasType.Cause5GMM{Iei: 0, Octet: 0xff},
		inAuthenticationFailureParameter:       nasType.AuthenticationFailureParameter{Iei: 0x30, Len: 14, Octet: [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	},
}

func TestNasTypeNewAuthenticationFailure(t *testing.T) {
	a := nasMessage.NewAuthenticationFailure(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewAuthenticationFailureMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationFailureMessage---")
	for i, table := range nasMessageAuthenticationFailureTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewAuthenticationFailure(0)
		b := nasMessage.NewAuthenticationFailure(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inAuthenticationFailureMessageIdentity)
		a.Cause5GMM = table.in5GMMCause
		a.AuthenticationFailureParameter = nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
		a.AuthenticationFailureParameter = &table.inAuthenticationFailureParameter

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationFailure(buff)
		logger.NasMsgLog.Debugln("encode:", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("data:", data)
		if err := b.DecodeAuthenticationFailure(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("decode:", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
