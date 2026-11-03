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

type nasMessagePDUSessionAuthenticationResultData struct {
	inExtendedProtocolDiscriminator                 uint8
	inPDUSessionID                                  uint8
	inPTI                                           uint8
	inPDUSESSIONAUTHENTICATIONRESULTMessageIdentity uint8
	inEAPMessage                                    nasType.EAPMessage
	inExtendedProtocolConfigurationOptions          nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionAuthenticationResultTable = []nasMessagePDUSessionAuthenticationResultData{
	{
		inExtendedProtocolDiscriminator: nas.MsgTypePDUSessionAuthenticationResult,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONAUTHENTICATIONRESULTMessageIdentity: 0x01,
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.PDUSessionAuthenticationResultEAPMessageType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionAuthenticationResultExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionAuthenticationResult(t *testing.T) {
	a := nasMessage.NewPDUSessionAuthenticationResult(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewPDUSessionAuthenticationResultMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionAuthenticationResultTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionAuthenticationResult(0)
		b := nasMessage.NewPDUSessionAuthenticationResult(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetPDUSessionID(table.inPDUSessionID)
		a.SetPTI(table.inPTI)
		a.SetMessageType(table.inPDUSESSIONAUTHENTICATIONRESULTMessageIdentity)

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.PDUSessionAuthenticationResultEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionAuthenticationResultExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionAuthenticationResult(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodePDUSessionAuthenticationResult(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
