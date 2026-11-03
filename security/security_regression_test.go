// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package security

import (
	"bytes"
	"sync"
	"testing"
)

func TestNASMacCalculateNIA0ReturnsZeroMAC(t *testing.T) {
	var key [16]byte
	mac, err := NASMacCalculate(AlgIntegrity128NIA0, key, 0, Bearer3GPP, DirectionUplink, []byte{0x01, 0x02})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !bytes.Equal(mac, []byte{0, 0, 0, 0}) {
		t.Errorf("NIA0 MAC = %x, want 00000000", mac)
	}
}

func TestNIA1RejectsEmptyAndShortMessages(t *testing.T) {
	var key [16]byte
	if _, err := NIA1(key, 0, 0, 0, []byte{}, 0); err == nil {
		t.Error("expected error for zero length")
	}
	if _, err := NIA1(key, 0, 0, 0, []byte{0x01}, 16); err == nil {
		t.Error("expected error when message is shorter than length")
	}
}

// Concurrent SNOW 3G users must get the same result as sequential ones.
// Run with -race to also detect shared state.
func TestSnow3gConcurrentUse(t *testing.T) {
	keyA := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	keyB := [16]byte{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	msg := bytes.Repeat([]byte{0xa5}, 64)

	wantEnc, err := NEA1(keyA, 7, 1, 0, msg, uint32(len(msg))*8)
	if err != nil {
		t.Fatal(err)
	}
	wantMac, err := NIA1(keyB, 9, 1, 1, msg, uint64(len(msg))*8)
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	for range 16 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for range 50 {
				got, err := NEA1(keyA, 7, 1, 0, msg, uint32(len(msg))*8)
				if err != nil || !bytes.Equal(got, wantEnc) {
					t.Errorf("NEA1 mismatch under concurrency: err=%v", err)
					return
				}
			}
		}()
		go func() {
			defer wg.Done()
			for range 50 {
				got, err := NIA1(keyB, 9, 1, 1, msg, uint64(len(msg))*8)
				if err != nil || !bytes.Equal(got, wantMac) {
					t.Errorf("NIA1 mismatch under concurrency: err=%v", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}
