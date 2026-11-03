// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
)

func TestNasTypeNewServiceLevelAuthenticationCommand(t *testing.T) {
	a := nasMessage.NewServiceLevelAuthenticationCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestServiceLevelAuthenticationCommandEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: ServiceLevelAuthenticationCommand---")

	a := nasMessage.NewServiceLevelAuthenticationCommand(0)
	b := nasMessage.NewServiceLevelAuthenticationCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.SetPDUSessionID(0x01)
	a.SetPTI(0x01)
	a.SetMessageType(0xD8)

	// Set ServiceLevelAAContainer
	a.ServiceLevelAAContainer = nasType.ServiceLevelAAContainer{}
	a.SetLen(4)
	copy(a.Buffer, []byte{0x01, 0x02, 0x03, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeServiceLevelAuthenticationCommand(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeServiceLevelAuthenticationCommand(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ServiceLevelAuthenticationCommand encode/decode mismatch")
	}
}

func TestNasTypeNewServiceLevelAuthenticationComplete(t *testing.T) {
	a := nasMessage.NewServiceLevelAuthenticationComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestServiceLevelAuthenticationCompleteEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: ServiceLevelAuthenticationComplete---")

	a := nasMessage.NewServiceLevelAuthenticationComplete(0)
	b := nasMessage.NewServiceLevelAuthenticationComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.SetPDUSessionID(0x01)
	a.SetPTI(0x01)
	a.SetMessageType(0xD9)

	// Set ServiceLevelAAContainer
	a.SetLen(4)
	copy(a.Buffer, []byte{0x05, 0x06, 0x07, 0x08})

	buff := new(bytes.Buffer)
	a.EncodeServiceLevelAuthenticationComplete(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeServiceLevelAuthenticationComplete(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ServiceLevelAuthenticationComplete encode/decode mismatch")
	}
}

func TestNasTypeNewRemoteUEReport(t *testing.T) {
	a := nasMessage.NewRemoteUEReport(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestRemoteUEReportEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RemoteUEReport---")

	a := nasMessage.NewRemoteUEReport(0)
	b := nasMessage.NewRemoteUEReport(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.SetPDUSessionID(0x01)
	a.SetPTI(0x02)
	a.SetMessageType(0xDA)

	buff := new(bytes.Buffer)
	a.EncodeRemoteUEReport(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRemoteUEReport(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RemoteUEReport encode/decode mismatch")
	}
}

func TestRemoteUEReportEncodeDecodeWithOptionalIEs(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RemoteUEReport (with optional IEs)---")

	a := nasMessage.NewRemoteUEReport(0)
	b := nasMessage.NewRemoteUEReport(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.SetPDUSessionID(0x01)
	a.SetPTI(0x02)
	a.SetMessageType(0xDA)

	// Set Remote UE Context Connected (IEI 0x76)
	a.RemoteUEContextList = nasType.NewRemoteUEContextList(nasMessage.RemoteUEReportRemoteUEContextConnectedType)
	a.SetLen(4)
	copy(a.Buffer, []byte{0x01, 0x02, 0x03, 0x04})

	// Set Remote UE Context Disconnected (IEI 0x70)
	a.RemoteUEContextDisconnected = nasType.NewRemoteUEContextList(nasMessage.RemoteUEReportRemoteUEContextDisconnectedType)
	a.RemoteUEContextDisconnected.SetLen(4)
	copy(a.RemoteUEContextDisconnected.Buffer, []byte{0x05, 0x06, 0x07, 0x08})

	buff := new(bytes.Buffer)
	a.EncodeRemoteUEReport(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRemoteUEReport(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RemoteUEReport (with optional IEs) encode/decode mismatch")
	}
}

func TestNasTypeNewRemoteUEReportResponse(t *testing.T) {
	a := nasMessage.NewRemoteUEReportResponse(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestRemoteUEReportResponseEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RemoteUEReportResponse---")

	a := nasMessage.NewRemoteUEReportResponse(0)
	b := nasMessage.NewRemoteUEReportResponse(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.SetPDUSessionID(0x01)
	a.SetPTI(0x02)
	a.SetMessageType(0xDB)

	buff := new(bytes.Buffer)
	a.EncodeRemoteUEReportResponse(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRemoteUEReportResponse(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RemoteUEReportResponse encode/decode mismatch")
	}
}
