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
)

type nasMessageIdentityRequestData struct {
	inExtendedProtocolDiscriminator  uint8
	inSecurityHeader                 uint8
	inSpareHalfOctet1                uint8
	inIdentityRequestMessageIdentity uint8
	inIdentityType                   uint8
	inSpareHalfOctet2                uint8
}

var nasMessageIdentityRequestTable = []nasMessageIdentityRequestData{
	{
		inExtendedProtocolDiscriminator:  0x01,
		inSecurityHeader:                 0x08,
		inSpareHalfOctet1:                0x01,
		inIdentityRequestMessageIdentity: nas.MsgTypeIdentityRequest,
		inIdentityType:                   0x01,
		inSpareHalfOctet2:                0x01,
	},
}

func TestNasTypeNewIdentityRequest(t *testing.T) {
	a := nasMessage.NewIdentityRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewIdentityRequestMessage(t *testing.T) {
	for i, table := range nasMessageIdentityRequestTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewIdentityRequest(0)
		b := nasMessage.NewIdentityRequest(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.SetMessageType(table.inIdentityRequestMessageIdentity)
		a.SetTypeOfIdentity(table.inIdentityType)

		buff := new(bytes.Buffer)
		a.EncodeIdentityRequest(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeIdentityRequest(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Dncode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
