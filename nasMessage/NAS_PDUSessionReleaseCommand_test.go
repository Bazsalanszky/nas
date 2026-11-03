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

type nasMessagePDUSessionReleaseCommandData struct {
	inExtendedProtocolDiscriminator           uint8
	inPDUSessionID                            uint8
	inPTI                                     uint8
	inPDUSessionReleaseCommandMessageIdentity uint8
	inCause5GSM                               nasType.Cause5GSM
	inBackoffTimerValue                       nasType.BackoffTimerValue
	inEAPMessage                              nasType.EAPMessage
	inExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionReleaseCommandTable = []nasMessagePDUSessionReleaseCommandData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSessionReleaseCommandMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   0,
			Octet: 0x01,
		},
		inBackoffTimerValue: nasType.BackoffTimerValue{
			Iei:   nasMessage.PDUSessionReleaseCommandBackoffTimerValueType,
			Len:   2,
			Octet: 0x01,
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.PDUSessionReleaseCommandEAPMessageType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionReleaseCommand(t *testing.T) {
	a := nasMessage.NewPDUSessionReleaseCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewPDUSessionReleaseCommandMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionReleaseCommandTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionReleaseCommand(0)
		b := nasMessage.NewPDUSessionReleaseCommand(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetPDUSessionID(table.inPDUSessionID)
		a.SetPTI(table.inPTI)
		a.SetMessageType(table.inPDUSessionReleaseCommandMessageIdentity)

		a.Cause5GSM = table.inCause5GSM

		a.BackoffTimerValue = nasType.NewBackoffTimerValue(nasMessage.PDUSessionReleaseCommandBackoffTimerValueType)
		a.BackoffTimerValue = &table.inBackoffTimerValue

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.PDUSessionReleaseCommandEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionReleaseCommand(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodePDUSessionReleaseCommand(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)
		// logger.NasMsgLog.Debugln(a.Cause5GSM)
		// logger.NasMsgLog.Debugln(b.Cause5GSM)
		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
