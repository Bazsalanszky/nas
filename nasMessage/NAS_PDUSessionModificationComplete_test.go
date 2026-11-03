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

type nasMessagePDUSessionModificationCompleteData struct {
	inExtendedProtocolDiscriminator                 uint8
	inPDUSessionID                                  uint8
	inPTI                                           uint8
	inPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity uint8
	inExtendedProtocolConfigurationOptions          nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionModificationCompleteTable = []nasMessagePDUSessionModificationCompleteData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity: 0x01,
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionModificationCompleteExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionModificationComplete(t *testing.T) {
	a := nasMessage.NewPDUSessionModificationComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewPDUSessionModificationCompleteMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionModificationCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionModificationComplete(0)
		b := nasMessage.NewPDUSessionModificationComplete(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetPDUSessionID(table.inPDUSessionID)
		a.SetPTI(table.inPTI)
		a.SetMessageType(table.inPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity)

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionModificationCompleteExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionModificationComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodePDUSessionModificationComplete(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
