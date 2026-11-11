// SPDX-FileCopyrightText: 2026 Balázs Toldi
// SPDX-License-Identifier: Apache-2.0

package nasConvert

import (
	"bytes"
	"net"
	"testing"

	"github.com/omec-project/nas/v2/nasMessage"
)

const (
	testDNSv6   = "2001:4860:4860::8888"
	testPCSCFv6 = "2001:db8:5::10"
)

// The container identifiers of TS 24.008 Table 10.5.154 (MS to network
// direction, which the DL constants share).
func TestPCOContainerIDs(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  uint16
		want uint16
	}{
		{"P-CSCF IPv6 Address", nasMessage.PCSCFIPv6AddressDL, 0x0001},
		{"DNS Server IPv6 Address", nasMessage.DNSServerIPv6AddressDL, 0x0003},
		{"P-CSCF IPv4 Address", nasMessage.PCSCFIPv4AddressDL, 0x000C},
		{"DNS Server IPv4 Address", nasMessage.DNSServerIPv4AddressDL, 0x000D},
	} {
		if tc.got != tc.want {
			t.Errorf("%s = %#04x, want %#04x", tc.name, tc.got, tc.want)
		}
	}
}

func TestAddIPv6Addresses(t *testing.T) {
	for _, tc := range []struct {
		name string
		add  func(*ProtocolConfigurationOptions, net.IP) error
		ip   string
		id   uint16
	}{
		{"DNS", (*ProtocolConfigurationOptions).AddDNSServerIPv6Address, testDNSv6, 0x0003},
		{"P-CSCF", (*ProtocolConfigurationOptions).AddPCSCFIPv6Address, testPCSCFv6, 0x0001},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pco := NewProtocolConfigurationOptions()
			ip := net.ParseIP(tc.ip)
			if err := tc.add(pco, ip); err != nil {
				t.Fatal(err)
			}
			want := append([]byte{0x80, byte(tc.id >> 8), byte(tc.id), 16}, ip...)
			if got := pco.Marshal(); !bytes.Equal(got, want) {
				t.Errorf("Marshal() = % x, want % x", got, want)
			}
			// The container must not alias the caller's slice.
			ip[15]++
			if c := pco.ProtocolOrContainerList[0].Contents; c[15] == ip[15] {
				t.Error("container contents alias the argument")
			}
		})
	}
}

func TestAddIPv6AddressesRejectIPv4(t *testing.T) {
	for _, ipStr := range []string{
		"8.8.8.8",          // net.ParseIP returns the 16-byte IPv4-mapped form
		"::ffff:8.8.8.8",   // IPv4-mapped, written as IPv6
		"::ffff:808:808",   // the same, in hex
		"",                 // nil
		"4-byte",           // net.IPv4(...).To4()
		"short",            // malformed slice
		"0102030405060708", // 8 bytes
	} {
		var ip net.IP
		switch ipStr {
		case "4-byte":
			ip = net.IPv4(8, 8, 8, 8).To4()
		case "short":
			ip = net.IP{0x20, 0x01}
		case "0102030405060708":
			ip = net.IP{1, 2, 3, 4, 5, 6, 7, 8}
		default:
			ip = net.ParseIP(ipStr)
		}
		pco := NewProtocolConfigurationOptions()
		if err := pco.AddDNSServerIPv6Address(ip); err == nil {
			t.Errorf("AddDNSServerIPv6Address(%q) accepted", ipStr)
		}
		if err := pco.AddPCSCFIPv6Address(ip); err == nil {
			t.Errorf("AddPCSCFIPv6Address(%q) accepted", ipStr)
		}
		if n := len(pco.ProtocolOrContainerList); n != 0 {
			t.Errorf("%q: %d containers added on error", ipStr, n)
		}
	}
}

// The IPv4 helpers are unchanged; they still accept both net.IP forms.
func TestAddIPv4Addresses(t *testing.T) {
	pco := NewProtocolConfigurationOptions()
	if err := pco.AddDNSServerIPv4Address(net.ParseIP("8.8.8.8")); err != nil {
		t.Fatal(err)
	}
	if err := pco.AddPCSCFIPv4Address(net.IPv4(10, 0, 0, 5).To4()); err != nil {
		t.Fatal(err)
	}
	if err := pco.AddPCSCFIPv4Address(net.ParseIP(testPCSCFv6)); err == nil {
		t.Error("AddPCSCFIPv4Address accepted an IPv6 address")
	}
	want := []byte{0x80, 0x00, 0x0d, 4, 8, 8, 8, 8, 0x00, 0x0c, 4, 10, 0, 0, 5}
	if got := pco.Marshal(); !bytes.Equal(got, want) {
		t.Errorf("Marshal() = % x, want % x", got, want)
	}
}

// A dual-stack PCO, as the SMF sends it for an IPv4v6 session, decodes back
// to the same containers.
func TestDualStackPCORoundTrip(t *testing.T) {
	pco := NewProtocolConfigurationOptions()
	for _, err := range []error{
		pco.AddDNSServerIPv4Address(net.ParseIP("8.8.8.8")),
		pco.AddDNSServerIPv6Address(net.ParseIP(testDNSv6)),
		pco.AddPCSCFIPv4Address(net.ParseIP("10.0.0.5")),
		pco.AddPCSCFIPv6Address(net.ParseIP(testPCSCFv6)),
		pco.AddIPv4LinkMTU(1400),
	} {
		if err != nil {
			t.Fatal(err)
		}
	}
	decoded := NewProtocolConfigurationOptions()
	if err := decoded.UnMarshal(pco.Marshal()); err != nil {
		t.Fatal(err)
	}
	if len(decoded.ProtocolOrContainerList) != len(pco.ProtocolOrContainerList) {
		t.Fatalf("%d containers decoded, want %d", len(decoded.ProtocolOrContainerList), len(pco.ProtocolOrContainerList))
	}
	for i, c := range decoded.ProtocolOrContainerList {
		w := pco.ProtocolOrContainerList[i]
		if c.ProtocolOrContainerID != w.ProtocolOrContainerID || c.LengthOfContents != w.LengthOfContents || !bytes.Equal(c.Contents, w.Contents) {
			t.Errorf("container %d = %+v, want %+v", i, c, w)
		}
	}
	if got := net.IP(decoded.ProtocolOrContainerList[3].Contents).String(); got != testPCSCFv6 {
		t.Errorf("P-CSCF IPv6 = %s", got)
	}
}
