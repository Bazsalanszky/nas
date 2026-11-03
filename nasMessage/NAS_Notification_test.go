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

type nasMessageNotificationData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet1               uint8
	inNotificationMessageIdentity   uint8
	inAccessType                    uint8
	inSpareHalfOctet2               uint8
}

var nasMessageNotificationTable = []nasMessageNotificationData{
	{
		inExtendedProtocolDiscriminator: 0x01,
		inSecurityHeader:                0x08,
		inSpareHalfOctet1:               0x01,
		inNotificationMessageIdentity:   nas.MsgTypeNotification,
		inAccessType:                    0x01,
		inSpareHalfOctet2:               0x01,
	},
}

func TestNasTypeNewNotification(t *testing.T) {
	a := nasMessage.NewNotification(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewNotificationMessage(t *testing.T) {
	for i, table := range nasMessageNotificationTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewNotification(0)
		b := nasMessage.NewNotification(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.SetMessageType(table.inNotificationMessageIdentity)
		a.SetAccessType(table.inAccessType)

		buff := new(bytes.Buffer)
		a.EncodeNotification(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeNotification(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
