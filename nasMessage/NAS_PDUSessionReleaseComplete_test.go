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

type nasMessagePDUSessionReleaseCompleteData struct {
	inExtendedProtocolDiscriminator            uint8
	inPDUSessionID                             uint8
	inPTI                                      uint8
	inPDUSESSIONRELEASECOMPLETEMessageIdentity uint8
	inCause5GSM                                nasType.Cause5GSM
	inExtendedProtocolConfigurationOptions     nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionReleaseCompleteTable = []nasMessagePDUSessionReleaseCompleteData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONRELEASECOMPLETEMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   nasMessage.PDUSessionReleaseCompleteCause5GSMType,
			Octet: 0x01,
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionReleaseComplete(t *testing.T) {
	a := nasMessage.NewPDUSessionReleaseComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewPDUSessionReleaseCompleteMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionReleaseCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionReleaseComplete(0)
		b := nasMessage.NewPDUSessionReleaseComplete(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetPDUSessionID(table.inPDUSessionID)
		a.SetPTI(table.inPTI)
		a.SetMessageType(table.inPDUSESSIONRELEASECOMPLETEMessageIdentity)

		a.Cause5GSM = nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
		a.Cause5GSM = &table.inCause5GSM

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionReleaseComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodePDUSessionReleaseComplete(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)
		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
