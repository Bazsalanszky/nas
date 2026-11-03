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

type nasMessageIdentityResponseData struct {
	inExtendedProtocolDiscriminator   uint8
	inSecurityHeader                  uint8
	inSpareHalfOctet                  uint8
	inIdentityResponseMessageIdentity uint8
	inMobileIdentity                  nasType.MobileIdentity
}

var nasMessageIdentityResponseTable = []nasMessageIdentityResponseData{
	{
		inExtendedProtocolDiscriminator:   0x01,
		inSecurityHeader:                  0x08,
		inSpareHalfOctet:                  0x01,
		inIdentityResponseMessageIdentity: nas.MsgTypeIdentityResponse,
		inMobileIdentity: nasType.MobileIdentity{
			Iei:    0,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewIdentityResponse(t *testing.T) {
	a := nasMessage.NewIdentityResponse(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewIdentityResponseMessage(t *testing.T) {
	for i, table := range nasMessageIdentityResponseTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewIdentityResponse(0)
		b := nasMessage.NewIdentityResponse(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inIdentityResponseMessageIdentity)

		a.MobileIdentity = table.inMobileIdentity

		buff := new(bytes.Buffer)
		a.EncodeIdentityResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeIdentityResponse(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
