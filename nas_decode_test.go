// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nas

import "testing"

func TestGmmMessageDecodePropagatesDecodeError(t *testing.T) {
	// Authentication response whose parameter IE is truncated
	data := []byte{0x7e, 0x00, MsgTypeAuthenticationResponse, 0x2d, 0x10, 0x01}
	msg := NewMessage()
	if err := msg.PlainNasDecode(&data); err == nil {
		t.Fatal("expected error for truncated IE, got nil")
	}
}
