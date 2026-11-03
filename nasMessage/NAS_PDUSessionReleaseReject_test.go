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

type nasMessagePDUSessionReleaseRejectData struct {
	inExtendedProtocolDiscriminator          uint8
	inPDUSessionID                           uint8
	inPTI                                    uint8
	inPDUSESSIONRELEASEREJECTMessageIdentity uint8
	inCause5GSM                              nasType.Cause5GSM
	inExtendedProtocolConfigurationOptions   nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionReleaseRejectTable = []nasMessagePDUSessionReleaseRejectData{
	{
		inExtendedProtocolDiscriminator:          nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID:                           0x01,
		inPTI:                                    0x01,
		inPDUSESSIONRELEASEREJECTMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Octet: 0x01,
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionReleaseRejectExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionReleaseReject(t *testing.T) {
	a := nasMessage.NewPDUSessionReleaseReject(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewPDUSessionReleaseRejectMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionReleaseRejectTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionReleaseReject(0)
		b := nasMessage.NewPDUSessionReleaseReject(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetPDUSessionID(table.inPDUSessionID)
		a.SetPTI(table.inPTI)
		a.SetMessageType(table.inPDUSESSIONRELEASEREJECTMessageIdentity)

		a.Cause5GSM = table.inCause5GSM

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionReleaseRejectExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionReleaseReject(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodePDUSessionReleaseReject(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
