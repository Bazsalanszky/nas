// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage_test

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2"
	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
)

type nasMessageSecurityModeCompleteData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inSecurityModeCompleteMessageIdentity uint8
	inIMEISV                              nasType.IMEISV
	inNASMessageContainer                 nasType.NASMessageContainer
}

var nasMessageSecurityModeCompleteTable = []nasMessageSecurityModeCompleteData{
	{
		inExtendedProtocolDiscriminator:       nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                      0x01,
		inSpareHalfOctet:                      0x01,
		inSecurityModeCompleteMessageIdentity: nas.MsgTypeSecurityModeComplete,
		inIMEISV: nasType.IMEISV{
			Iei:   nasMessage.SecurityModeCompleteIMEISVType,
			Len:   9,
			Octet: [9]uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		inNASMessageContainer: nasType.NASMessageContainer{
			Iei:    nasMessage.SecurityModeCompleteNASMessageContainerType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewSecurityModeComplete(t *testing.T) {
	a := nasMessage.NewSecurityModeComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewSecurityModeCompleteMessage(t *testing.T) {
	for i, table := range nasMessageSecurityModeCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityModeComplete(0)
		b := nasMessage.NewSecurityModeComplete(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inSecurityModeCompleteMessageIdentity)

		a.IMEISV = nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
		a.IMEISV = &table.inIMEISV

		a.NASMessageContainer = nasType.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)
		a.NASMessageContainer = &table.inNASMessageContainer

		buff := new(bytes.Buffer)
		a.EncodeSecurityModeComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeSecurityModeComplete(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}

// TestNasTypeSecurityModeCompleteOverlengthIMEISV guards against a crafted
// IMEISV IE whose declared length (10) exceeds the fixed 9-octet IMEISV
// buffer; decoding must not panic with a slice bounds out of range error, and
// the malformed IE must be discarded so re-encoding the result cannot panic either.
func TestNasTypeSecurityModeCompleteOverlengthIMEISV(t *testing.T) {
	data := []byte{
		0x7e, 0x00, 0x5e, // EPD, spare/security header type, message type
		0x77, 0x00, 0x0a, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, // IMEISV, len=10
		0x71, 0x01, 0x00, // NAS message container, len=1
	}
	b := nasMessage.NewSecurityModeComplete(0)
	if err := b.DecodeSecurityModeComplete(&data); !errors.Is(err, nasMessage.ErrMalformedIE) {
		t.Fatalf("error = %v, want ErrMalformedIE", err)
	}

	if b.IMEISV != nil {
		t.Errorf("expected malformed IMEISV to be discarded, got %+v", b.IMEISV)
	}

	buff := new(bytes.Buffer)
	b.EncodeSecurityModeComplete(buff)
}
