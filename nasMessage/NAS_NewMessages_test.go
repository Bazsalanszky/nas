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

func TestNasTypeNewNetworkSliceSpecificAuthenticationCommand(t *testing.T) {
	a := nasMessage.NewNetworkSliceSpecificAuthenticationCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNetworkSliceSpecificAuthenticationCommandEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: NetworkSliceSpecificAuthenticationCommand---")

	a := nasMessage.NewNetworkSliceSpecificAuthenticationCommand(0)
	b := nasMessage.NewNetworkSliceSpecificAuthenticationCommand(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x50)

	// Set S-NSSAI (1 byte SST)
	a.SNSSAI.Len = 1
	a.SNSSAI.Octet[0] = 0x01

	// Set EAP message
	a.EAPMessage = nasType.EAPMessage{}
	a.EAPMessage.SetLen(4)
	copy(a.Buffer, []byte{0x02, 0x01, 0x00, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeNetworkSliceSpecificAuthenticationCommand(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeNetworkSliceSpecificAuthenticationCommand(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("NetworkSliceSpecificAuthenticationCommand encode/decode mismatch")
	}
}

func TestNasTypeNewNetworkSliceSpecificAuthenticationComplete(t *testing.T) {
	a := nasMessage.NewNetworkSliceSpecificAuthenticationComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNetworkSliceSpecificAuthenticationCompleteEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: NetworkSliceSpecificAuthenticationComplete---")

	a := nasMessage.NewNetworkSliceSpecificAuthenticationComplete(0)
	b := nasMessage.NewNetworkSliceSpecificAuthenticationComplete(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x51)

	// Set S-NSSAI (1 byte SST only)
	a.SNSSAI.Len = 1
	a.SNSSAI.Octet[0] = 0x01

	// Set EAP message
	a.EAPMessage = nasType.EAPMessage{}
	a.EAPMessage.SetLen(4)
	copy(a.Buffer, []byte{0x02, 0x01, 0x00, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeNetworkSliceSpecificAuthenticationComplete(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeNetworkSliceSpecificAuthenticationComplete(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("NetworkSliceSpecificAuthenticationComplete encode/decode mismatch")
	}
}

func TestNasTypeNewNetworkSliceSpecificAuthenticationResult(t *testing.T) {
	a := nasMessage.NewNetworkSliceSpecificAuthenticationResult(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNetworkSliceSpecificAuthenticationResultEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: NetworkSliceSpecificAuthenticationResult---")

	a := nasMessage.NewNetworkSliceSpecificAuthenticationResult(0)
	b := nasMessage.NewNetworkSliceSpecificAuthenticationResult(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x52)

	// Set S-NSSAI (1 byte SST only)
	a.SNSSAI.Len = 1
	a.SNSSAI.Octet[0] = 0x02

	// Set EAP message
	a.EAPMessage = nasType.EAPMessage{}
	a.EAPMessage.SetLen(4)
	copy(a.Buffer, []byte{0x03, 0x01, 0x00, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeNetworkSliceSpecificAuthenticationResult(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeNetworkSliceSpecificAuthenticationResult(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("NetworkSliceSpecificAuthenticationResult encode/decode mismatch")
	}
}

func TestNasTypeNewRelayKeyRequest(t *testing.T) {
	a := nasMessage.NewRelayKeyRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewRelayKeyAccept(t *testing.T) {
	a := nasMessage.NewRelayKeyAccept(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewRelayKeyReject(t *testing.T) {
	a := nasMessage.NewRelayKeyReject(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewRelayAuthenticationRequest(t *testing.T) {
	a := nasMessage.NewRelayAuthenticationRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestNasTypeNewRelayAuthenticationResponse(t *testing.T) {
	a := nasMessage.NewRelayAuthenticationResponse(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestRelayKeyRequestEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RelayKeyRequest---")

	a := nasMessage.NewRelayKeyRequest(0)
	b := nasMessage.NewRelayKeyRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x69)
	a.SetProSeRelayTransactionIdentityValue(0x01)
	a.SetLen(4)
	copy(a.Buffer, []byte{0x01, 0x02, 0x03, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeRelayKeyRequest(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRelayKeyRequest(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RelayKeyRequest encode/decode mismatch")
	}
}

func TestRelayKeyAcceptEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RelayKeyAccept---")

	a := nasMessage.NewRelayKeyAccept(0)
	b := nasMessage.NewRelayKeyAccept(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x6A)
	a.SetProSeRelayTransactionIdentityValue(0x01)
	a.RelayKeyResponseParameters.SetLen(4)
	copy(a.RelayKeyResponseParameters.Buffer, []byte{0x0A, 0x0B, 0x0C, 0x0D})

	buff := new(bytes.Buffer)
	a.EncodeRelayKeyAccept(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRelayKeyAccept(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RelayKeyAccept encode/decode mismatch")
	}
}

func TestRelayKeyRejectEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RelayKeyReject---")

	a := nasMessage.NewRelayKeyReject(0)
	b := nasMessage.NewRelayKeyReject(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x6B)
	a.SetProSeRelayTransactionIdentityValue(0x02)

	buff := new(bytes.Buffer)
	a.EncodeRelayKeyReject(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRelayKeyReject(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RelayKeyReject encode/decode mismatch")
	}
}

func TestRelayAuthenticationRequestEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RelayAuthenticationRequest---")

	a := nasMessage.NewRelayAuthenticationRequest(0)
	b := nasMessage.NewRelayAuthenticationRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x6C)
	a.SetProSeRelayTransactionIdentityValue(0x01)
	a.SetLen(4)
	copy(a.Buffer, []byte{0x02, 0x01, 0x00, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeRelayAuthenticationRequest(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRelayAuthenticationRequest(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RelayAuthenticationRequest encode/decode mismatch")
	}
}

func TestRelayAuthenticationResponseEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RelayAuthenticationResponse---")

	a := nasMessage.NewRelayAuthenticationResponse(0)
	b := nasMessage.NewRelayAuthenticationResponse(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x6D)
	a.SetProSeRelayTransactionIdentityValue(0x01)
	a.SetLen(4)
	copy(a.Buffer, []byte{0x03, 0x01, 0x00, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeRelayAuthenticationResponse(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeRelayAuthenticationResponse(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RelayAuthenticationResponse encode/decode mismatch")
	}
}

func TestNasTypeNewControlPlaneServiceRequest(t *testing.T) {
	a := nasMessage.NewControlPlaneServiceRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
}

func TestControlPlaneServiceRequestEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: ControlPlaneServiceRequest---")

	a := nasMessage.NewControlPlaneServiceRequest(0)
	b := nasMessage.NewControlPlaneServiceRequest(0)
	if a == nil {
		t.Fatal("Expected value not to be nil")
	}
	if b == nil {
		t.Fatal("Expected value not to be nil")
	}

	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetSecurityHeaderType(0x00)
	a.SetSpareHalfOctet(0x00)
	a.SetMessageType(0x4F)
	a.SetControlPlaneServiceType(0x01)
	a.SetNasKeySetIdentifiler(0x07)

	a.UERequestType = nasType.NewUERequestType(nasMessage.ControlPlaneServiceRequestUERequestTypeType)
	a.UERequestType.SetLen(1)
	copy(a.UERequestType.Buffer, []byte{0x01})

	a.PagingRestriction = nasType.NewPagingRestriction(nasMessage.ControlPlaneServiceRequestPagingRestrictionType)
	a.PagingRestriction.SetLen(2)
	copy(a.PagingRestriction.Buffer, []byte{0x01, 0x02})

	buff := new(bytes.Buffer)
	a.EncodeControlPlaneServiceRequest(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	if _, err := buff.Read(data); err != nil {
		t.Fatal(err)
	}
	if err := b.DecodeControlPlaneServiceRequest(&data); err != nil {
		t.Fatal(err)
	}
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ControlPlaneServiceRequest encode/decode mismatch")
	}
}
