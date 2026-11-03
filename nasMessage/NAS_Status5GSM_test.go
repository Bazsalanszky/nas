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

type nasMessageStatus5GSMData struct {
	inExtendedProtocolDiscriminator uint8
	inPDUSessionID                  nasType.PDUSessionID
	inPTI                           nasType.PTI
	inStatus5GSMMessageIdentity     uint8
	inCause5GSM                     nasType.Cause5GSM
}

var nasMessageStatus5GSMTable = []nasMessageStatus5GSMData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID: nasType.PDUSessionID{
			Octet: 0x01,
		},
		inPTI: nasType.PTI{
			Octet: 0x01,
		},
		inStatus5GSMMessageIdentity: nas.MsgTypeStatus5GSM,
		inCause5GSM: nasType.Cause5GSM{
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewStatus5GSM(t *testing.T) {
	a := nasMessage.NewStatus5GSM(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewStatus5GSMMessage(t *testing.T) {
	for i, table := range nasMessageStatus5GSMTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewStatus5GSM(0)
		b := nasMessage.NewStatus5GSM(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID = table.inPDUSessionID
		a.PTI = table.inPTI

		a.SetMessageType(table.inStatus5GSMMessageIdentity)

		a.Cause5GSM = table.inCause5GSM

		buff := new(bytes.Buffer)
		a.EncodeStatus5GSM(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeStatus5GSM(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
