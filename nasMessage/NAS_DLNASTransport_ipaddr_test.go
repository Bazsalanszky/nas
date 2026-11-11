// SPDX-FileCopyrightText: 2026 Balázs Toldi
// SPDX-License-Identifier: Apache-2.0

package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/omec-project/nas/v2"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
)

const (
	testIID   = "\x11\x22\x33\x44\x55\x66\x77\x88"
	testIPv4  = "\x0a\x2d\x00\x07" // 10.45.0.7
	wantIPv4  = "10.45.0.7"
	msgTypeEA = nas.MsgTypePDUSessionEstablishmentAccept
)

type acceptOpts struct {
	msgType     uint8
	pduType     uint8
	addrInfo    string // PDU address information; no PDU address IE when empty
	qosRulesLen uint16
	cause       bool
}

// encodeAccept builds a 5GSM message with the library encoder.
func encodeAccept(t *testing.T, o acceptOpts) []byte {
	t.Helper()
	m := nasMessage.NewPDUSessionEstablishmentAccept(0)
	m.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	m.SetPDUSessionID(5)
	m.SetPTI(1)
	m.SetMessageType(o.msgType)
	m.SetSSCMode(1)
	m.SetPDUSessionType(o.pduType)
	m.AuthorizedQosRules.SetLen(o.qosRulesLen)
	for i := range m.AuthorizedQosRules.Buffer {
		// Values that look like the cause (0x59) and PDU address (0x29) IEIs,
		// to catch a scanner that does not honour the QoS rules length.
		m.AuthorizedQosRules.Buffer[i] = []byte{0x29, 0x59}[i%2]
	}
	m.SessionAMBR.SetLen(6)
	if o.cause {
		m.Cause5GSM = nasType.NewCause5GSM(nasMessage.PDUSessionEstablishmentAcceptCause5GSMType)
		m.SetCauseValue(nasMessage.Cause5GSMPDUSessionTypeIPv4OnlyAllowed)
	}
	if o.addrInfo != "" {
		m.PDUAddress = nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
		m.PDUAddress.SetLen(uint8(1 + len(o.addrInfo)))
		m.SetPDUSessionTypeValue(o.pduType)
		var info [12]uint8
		copy(info[:], o.addrInfo)
		m.SetPDUAddressInformation(info)
	}
	m.DNN = nasType.NewDNN(nasMessage.PDUSessionEstablishmentAcceptDNNType)
	m.SetDNN([]uint8("internet"))

	buf := new(bytes.Buffer)
	m.EncodePDUSessionEstablishmentAccept(buf)
	return buf.Bytes()
}

// decodeDLNASTransport wraps n1sm in a DL NAS Transport, encodes it and
// decodes it again.
func decodeDLNASTransport(t *testing.T, payloadType uint8, n1sm []byte) *nasMessage.DLNASTransport {
	t.Helper()
	a := nasMessage.NewDLNASTransport(0)
	a.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SetMessageType(nas.MsgTypeDLNASTransport)
	a.SetPayloadContainerType(payloadType)
	a.PayloadContainer.SetLen(uint16(len(n1sm)))
	a.SetPayloadContainerContents(n1sm)
	a.PduSessionID2Value = nasType.NewPduSessionID2Value(nasMessage.DLNASTransportPduSessionID2ValueType)
	a.SetPduSessionID2Value(5)

	buf := new(bytes.Buffer)
	a.EncodeDLNASTransport(buf)
	data := buf.Bytes()
	b := nasMessage.NewDLNASTransport(0)
	if err := b.DecodeDLNASTransport(&data); err != nil {
		t.Fatal(err)
	}
	if b.GetPduSessionID2Value() != 5 {
		t.Fatalf("PDU session ID = %d, the IEs after the payload container were not decoded", b.GetPduSessionID2Value())
	}
	return b
}

func TestDLNASTransportIpaddr(t *testing.T) {
	tests := []struct {
		name string
		opts acceptOpts
		want string
	}{
		{
			name: "IPv4",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4, addrInfo: testIPv4, qosRulesLen: 9},
			want: wantIPv4,
		},
		{
			name: "IPv4 with 5GSM cause",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4, addrInfo: testIPv4, qosRulesLen: 9, cause: true},
			want: wantIPv4,
		},
		{
			name: "IPv4, QoS rules longer than 9 octets",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4, addrInfo: testIPv4, qosRulesLen: 17},
			want: wantIPv4,
		},
		{
			name: "IPv4, QoS rules shorter than 9 octets",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4, addrInfo: testIPv4, qosRulesLen: 6},
			want: wantIPv4,
		},
		{
			name: "IPv6 has only an interface identifier",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv6, addrInfo: testIID, qosRulesLen: 9},
			want: "",
		},
		{
			name: "IPv4v6 gives the IPv4 part",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4IPv6, addrInfo: testIID + testIPv4, qosRulesLen: 9},
			want: wantIPv4,
		},
		{
			name: "IPv4v6 PDU address too short for the IPv4 part",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4IPv6, addrInfo: testIID, qosRulesLen: 9},
			want: "",
		},
		{
			name: "IPv4 PDU address too short",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4, addrInfo: "\x0a\x2d", qosRulesLen: 9},
			want: "",
		},
		{
			name: "no PDU address",
			opts: acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4, qosRulesLen: 9},
			want: "",
		},
		{
			name: "other 5GSM message",
			opts: acceptOpts{msgType: nas.MsgTypePDUSessionModificationCommand, pduType: nasMessage.PDUSessionTypeIPv4, addrInfo: testIPv4, qosRulesLen: 9},
			want: "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := decodeDLNASTransport(t, nasMessage.PayloadContainerTypeN1SMInfo, encodeAccept(t, tc.opts))
			if b.Ipaddr != tc.want {
				t.Errorf("Ipaddr = %q, want %q", b.Ipaddr, tc.want)
			}
		})
	}
}

func TestDLNASTransportIpaddrOnlyForN1SM(t *testing.T) {
	n1sm := encodeAccept(t, acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4, addrInfo: testIPv4, qosRulesLen: 9})
	b := decodeDLNASTransport(t, nasMessage.PayloadContainerTypeSMS, n1sm)
	if b.Ipaddr != "" {
		t.Errorf("Ipaddr = %q for an SMS payload container", b.Ipaddr)
	}
}

// A truncated or garbled payload container must neither fail nor panic the
// DL NAS Transport decode; Ipaddr is then empty or the address that was
// decoded before the damage.
func TestDLNASTransportIpaddrMalformed(t *testing.T) {
	full := encodeAccept(t, acceptOpts{msgType: msgTypeEA, pduType: nasMessage.PDUSessionTypeIPv4IPv6, addrInfo: testIID + testIPv4, qosRulesLen: 9})
	for n := range len(full) {
		b := decodeDLNASTransport(t, nasMessage.PayloadContainerTypeN1SMInfo, full[:n])
		if b.Ipaddr != "" && b.Ipaddr != wantIPv4 {
			t.Errorf("truncated to %d octets: Ipaddr = %q", n, b.Ipaddr)
		}
	}
	garbled := append([]byte(nil), full...)
	garbled[5], garbled[6] = 0xff, 0xff // QoS rules length beyond the message
	if b := decodeDLNASTransport(t, nasMessage.PayloadContainerTypeN1SMInfo, garbled); b.Ipaddr != "" {
		t.Errorf("bad QoS rules length: Ipaddr = %q", b.Ipaddr)
	}
}
