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

type nasMessageStatus5GMMData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet                uint8
	inStatus5GMMMessageIdentity     uint8
	inCause5GMM                     nasType.Cause5GMM
}

var nasMessageStatus5GMMTable = []nasMessageStatus5GMMData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inStatus5GMMMessageIdentity:     nas.MsgTypeStatus5GMM,
		inCause5GMM: nasType.Cause5GMM{
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewStatus5GMM(t *testing.T) {
	a := nasMessage.NewStatus5GMM(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewStatus5GMMMessage(t *testing.T) {
	for i, table := range nasMessageStatus5GMMTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewStatus5GMM(0)
		b := nasMessage.NewStatus5GMM(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)

		a.SetMessageType(table.inStatus5GMMMessageIdentity)

		a.Cause5GMM = table.inCause5GMM

		buff := new(bytes.Buffer)
		a.EncodeStatus5GMM(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeStatus5GMM(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
