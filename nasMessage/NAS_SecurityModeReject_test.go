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

type nasMessageSecurityModeRejectData struct {
	inExtendedProtocolDiscriminator     uint8
	inSecurityHeader                    uint8
	inSpareHalfOctet                    uint8
	inSecurityModeRejectMessageIdentity uint8
	inCause5GMM                         nasType.Cause5GMM
}

var nasMessageSecurityModeRejectTable = []nasMessageSecurityModeRejectData{
	{
		inExtendedProtocolDiscriminator:     nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                    0x01,
		inSpareHalfOctet:                    0x01,
		inSecurityModeRejectMessageIdentity: nas.MsgTypeSecurityModeReject,
		inCause5GMM: nasType.Cause5GMM{
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewSecurityModeReject(t *testing.T) {
	a := nasMessage.NewSecurityModeReject(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewSecurityModeRejectMessage(t *testing.T) {
	for i, table := range nasMessageSecurityModeRejectTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityModeReject(0)
		b := nasMessage.NewSecurityModeReject(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inSecurityModeRejectMessageIdentity)

		a.Cause5GMM = table.inCause5GMM

		buff := new(bytes.Buffer)
		a.EncodeSecurityModeReject(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeSecurityModeReject(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
