// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/omec-project/nas/v2/nasMessage"
)

func authRespHeader() []byte {
	return []byte{nasMessage.Epd5GSMobilityManagementMessage, 0x00, 0x57}
}

func authRespParam() []byte {
	// IEI 0x2D, length 16, 16 octets of value
	return append([]byte{nasMessage.AuthenticationResponseAuthenticationResponseParameterType, 0x10},
		bytes.Repeat([]byte{0xab}, 16)...)
}

func TestDecodeSkipsUnknownIEs(t *testing.T) {
	tests := map[string][]byte{
		"unknown type 4 IE":      append(append(authRespHeader(), 0x01, 0x02, 0xaa, 0xbb), authRespParam()...),
		"unknown zero length IE": append(append(authRespHeader(), 0x01, 0x00), authRespParam()...),
		"unknown type 1 IE":      append(append(authRespHeader(), 0xA0), authRespParam()...),
	}
	for name, data := range tests {
		t.Run(name, func(t *testing.T) {
			msg := nasMessage.NewAuthenticationResponse(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
			if err := msg.DecodeAuthenticationResponse(&data); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if msg.AuthenticationResponseParameter == nil {
				t.Fatal("IE following the unknown IE was not decoded")
			}
			if got := msg.AuthenticationResponseParameter.GetLen(); got != 16 {
				t.Errorf("Len = %d, want 16", got)
			}
		})
	}
}

func TestDecodeReturnsErrorOnMalformedMessage(t *testing.T) {
	tests := map[string][]byte{
		"truncated header":         authRespHeader()[:2],
		"truncated known IE":       append(authRespHeader(), 0x2D, 0x10, 0x01, 0x02),
		"missing length octet":     append(authRespHeader(), 0x2D),
		"unknown IE overruns data": append(authRespHeader(), 0x01, 0x05, 0xaa),
		"known IE length too big":  append(authRespHeader(), 0x2D, 0x11, 0x01),
	}
	for name, data := range tests {
		t.Run(name, func(t *testing.T) {
			msg := nasMessage.NewAuthenticationResponse(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
			if err := msg.DecodeAuthenticationResponse(&data); err == nil {
				t.Fatal("expected an error, got nil")
			}
		})
	}
}

func TestDecodeUnknownIEOverrunIsMalformedIE(t *testing.T) {
	data := append(authRespHeader(), 0x01, 0x05, 0xaa)
	msg := nasMessage.NewAuthenticationResponse(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
	if err := msg.DecodeAuthenticationResponse(&data); !errors.Is(err, nasMessage.ErrMalformedIE) {
		t.Fatalf("error = %v, want ErrMalformedIE", err)
	}
}
