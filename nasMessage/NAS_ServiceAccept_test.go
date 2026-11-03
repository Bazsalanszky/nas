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

type nasMessageServiceAcceptData struct {
	inExtendedProtocolDiscriminator          uint8
	inSecurityHeader                         uint8
	inSpareHalfOctet                         uint8
	inServiceAcceptMessageIdentity           uint8
	inPDUSessionStatus                       nasType.PDUSessionStatus
	inPDUSessionReactivationResult           nasType.PDUSessionReactivationResult
	inPDUSessionReactivationResultErrorCause nasType.PDUSessionReactivationResultErrorCause
	inEAPMessage                             nasType.EAPMessage
}

var nasMessageServiceAcceptTable = []nasMessageServiceAcceptData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inServiceAcceptMessageIdentity:  nas.MsgTypeServiceAccept,
		inPDUSessionStatus: nasType.PDUSessionStatus{
			Iei:    nasMessage.ServiceAcceptPDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPDUSessionReactivationResult: nasType.PDUSessionReactivationResult{
			Iei:    nasMessage.ServiceAcceptPDUSessionReactivationResultType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPDUSessionReactivationResultErrorCause: nasType.PDUSessionReactivationResultErrorCause{
			Iei:    nasMessage.ServiceAcceptPDUSessionReactivationResultErrorCauseType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.ServiceAcceptEAPMessageType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewServiceAccept(t *testing.T) {
	a := nasMessage.NewServiceAccept(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewServiceAcceptMessage(t *testing.T) {
	for i, table := range nasMessageServiceAcceptTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewServiceAccept(0)
		b := nasMessage.NewServiceAccept(0)
		if a == nil {
			t.Fatal("Expected value not to be nil")
		}
		if b == nil {
			t.Fatal("Expected value not to be nil")
		}

		a.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SetSecurityHeaderType(table.inSecurityHeader)
		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SetMessageType(table.inServiceAcceptMessageIdentity)

		a.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.ServiceAcceptPDUSessionStatusType)
		a.PDUSessionStatus = &table.inPDUSessionStatus

		a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(nasMessage.ServiceAcceptPDUSessionReactivationResultType)
		a.PDUSessionReactivationResult = &table.inPDUSessionReactivationResult

		a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(nasMessage.ServiceAcceptPDUSessionReactivationResultErrorCauseType)
		a.PDUSessionReactivationResultErrorCause = &table.inPDUSessionReactivationResultErrorCause

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.ServiceAcceptEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeServiceAccept(buff)
		logger.NasMsgLog.Debugln("encode:", a)

		data := make([]byte, buff.Len())
		if _, err := buff.Read(data); err != nil {
			t.Fatal(err)
		}
		if err := b.DecodeServiceAccept(&data); err != nil {
			t.Fatal(err)
		}
		logger.NasMsgLog.Debugln("decode:", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
