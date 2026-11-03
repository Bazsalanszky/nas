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

type nasMessageDeregistrationRequestUEOriginatingDeregistrationData struct {
	inExtendedProtocolDiscriminator        uint8
	inSecurityHeaderType                   uint8
	inSpareHalfOctet                       uint8
	inDeregistrationRequestMessageIdentity uint8
	inNgksiAndDeregistrationType           nasType.NgksiAndDeregistrationType
	inMobileIdentity5GS                    nasType.MobileIdentity5GS
}

var nasMessageDeregistrationRequestUEOriginatingDeregistrationTable = []nasMessageDeregistrationRequestUEOriginatingDeregistrationData{
	{
		inExtendedProtocolDiscriminator:        nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:                   0x01,
		inSpareHalfOctet:                       0x01,
		inDeregistrationRequestMessageIdentity: 0x01,
		inNgksiAndDeregistrationType: nasType.NgksiAndDeregistrationType{
			Octet: 0xFF,
		},
		inMobileIdentity5GS: nasType.MobileIdentity5GS{
			Iei:    0,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewDeregistrationRequestUEOriginatingDeregistration(t *testing.T) {
	a := nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewDeregistrationRequestUEOriginatingDeregistrationMessage(t *testing.T) {
	for i, table := range nasMessageDeregistrationRequestUEOriginatingDeregistrationTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(0)
		b := nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inDeregistrationRequestMessageIdentity)

		a.NgksiAndDeregistrationType = table.inNgksiAndDeregistrationType

		a.MobileIdentity5GS = table.inMobileIdentity5GS

		buff := new(bytes.Buffer)
		a.EncodeDeregistrationRequestUEOriginatingDeregistration(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeDeregistrationRequestUEOriginatingDeregistration(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
