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

type nasMessagePDUSessionAuthenticationCompleteData struct {
	inExtendedProtocolDiscriminator                   uint8
	inPDUSessionID                                    uint8
	inPTI                                             uint8
	inPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity uint8
	inEAPMessage                                      nasType.EAPMessage
	inExtendedProtocolConfigurationOptions            nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionAuthenticationCompleteTable = []nasMessagePDUSessionAuthenticationCompleteData{
	{
		inExtendedProtocolDiscriminator: nas.MsgTypePDUSessionAuthenticationComplete,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity: 0x01,
		inEAPMessage: nasType.EAPMessage{
			Iei:    0,
			Len:    4,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01},
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionAuthenticationCompleteExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionAuthenticationComplete(t *testing.T) {
	a := nasMessage.NewPDUSessionAuthenticationComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewPDUSessionAuthenticationCompleteMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionAuthenticationCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionAuthenticationComplete(0)
		b := nasMessage.NewPDUSessionAuthenticationComplete(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetPDUSessionID(table.inPDUSessionID)
		a.SetPTI(table.inPTI)
		a.SetMessageType(table.inPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity)

		a.EAPMessage = table.inEAPMessage

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionAuthenticationCompleteExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionAuthenticationComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodePDUSessionAuthenticationComplete(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
