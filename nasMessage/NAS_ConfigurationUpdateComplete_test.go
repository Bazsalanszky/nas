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

type nasMessageConfigurationUpdateCompleteData struct {
	inExtendedProtocolDiscriminator              uint8
	inSecurityHeaderType                         uint8
	inSpareHalfOctet                             uint8
	inConfigurationUpdateCompleteMessageIdentity uint8
}

var nasMessageConfigurationUpdateCompleteTable = []nasMessageConfigurationUpdateCompleteData{
	{
		inExtendedProtocolDiscriminator:              nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:                         0x01,
		inSpareHalfOctet:                             0x01,
		inConfigurationUpdateCompleteMessageIdentity: nas.MsgTypeConfigurationUpdateComplete,
	},
}

func TestNasTypeNewConfigurationUpdateComplete(t *testing.T) {
	a := nasMessage.NewConfigurationUpdateComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewConfigurationUpdateCompleteMessage(t *testing.T) {
	for i, table := range nasMessageConfigurationUpdateCompleteTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewConfigurationUpdateComplete(0)
		b := nasMessage.NewConfigurationUpdateComplete(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inConfigurationUpdateCompleteMessageIdentity)

		buff := new(bytes.Buffer)
		a.EncodeConfigurationUpdateComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeConfigurationUpdateComplete(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
