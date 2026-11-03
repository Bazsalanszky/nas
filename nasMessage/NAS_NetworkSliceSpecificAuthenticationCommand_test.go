// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/omec-project/nas/v2/nasMessage"
)

// TestNasTypeNetworkSliceSpecificAuthenticationCommandOverlengthSNSSAI guards
// against a declared S-NSSAI length (9) exceeding the fixed 8-octet SNSSAI
// buffer; decoding must return ErrMalformedIE instead of panicking, and since SNSSAI is a mandatory
// (non-pointer) field, the length must be reset so re-encoding cannot panic either.
func TestNasTypeNetworkSliceSpecificAuthenticationCommandOverlengthSNSSAI(t *testing.T) {
	data := []byte{0x00, 0x00, 0x00, 0x09} // EPD, spare/security header type, message type, SNSSAI len=9
	b := nasMessage.NewNetworkSliceSpecificAuthenticationCommand(0)
	if err := b.DecodeNetworkSliceSpecificAuthenticationCommand(&data); !errors.Is(err, nasMessage.ErrMalformedIE) {
		t.Fatalf("error = %v, want ErrMalformedIE", err)
	}

	if b.SNSSAI.GetLen() != 0 {
		t.Errorf("expected overlength SNSSAI to be reset to len 0, got %d", b.SNSSAI.GetLen())
	}

	buff := new(bytes.Buffer)
	b.EncodeNetworkSliceSpecificAuthenticationCommand(buff)
}
