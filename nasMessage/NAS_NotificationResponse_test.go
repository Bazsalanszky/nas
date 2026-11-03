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

type nasMessageNotificationResponseData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inNotificationResponseMessageIdentity uint8
	inPDUSessionStatus                    nasType.PDUSessionStatus
}

var nasMessageNotificationResponseTable = []nasMessageNotificationResponseData{
	{
		inExtendedProtocolDiscriminator:       0x01,
		inSecurityHeader:                      0x08,
		inSpareHalfOctet:                      0x01,
		inNotificationResponseMessageIdentity: 0x01,
		inPDUSessionStatus: nasType.PDUSessionStatus{
			Iei:    nasMessage.NotificationResponsePDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewNotificationResponse(t *testing.T) {
	a := nasMessage.NewNotificationResponse(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewNotificationResponseMessage(t *testing.T) {
	for i, table := range nasMessageNotificationResponseTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewNotificationResponse(0)
		b := nasMessage.NewNotificationResponse(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inNotificationResponseMessageIdentity)

		a.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.NotificationResponsePDUSessionStatusType)
		a.PDUSessionStatus = &table.inPDUSessionStatus

		buff := new(bytes.Buffer)
		a.EncodeNotificationResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeNotificationResponse(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
