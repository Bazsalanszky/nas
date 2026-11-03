// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// ErrMalformedIE is returned (wrapped) when an information element cannot be decoded
// because its length is invalid or exceeds the remaining message.
var ErrMalformedIE = errors.New("malformed information element")

// skipUnknownIE consumes the remainder of an optional information element whose IEI
// is not known in the message, so that the following IEs are still decoded correctly.
// TS 24.501 clause 9.11.1: the receiver determines whether the IE is of type 1 or 2
// (one octet, nothing left to skip), of type 4 (the next octet is the length indicator)
// or of type 6 (the next two octets are the length indicator).
// Assuming (TS 24.007 clause 11.2.4 is not in the spec set used for review): IEIs with
// bit 8 set are one octet long, IEIs in 0x70-0x7F are type 6 (TLV-E) and any other
// unknown IEI is type 4. This matches the IE tables of TS 24.501 clause 8, where every
// TLV-E IE has an IEI in 0x70-0x7D and no TLV IE has an IEI in 0x70-0x7F.
func skipUnknownIE(buffer *bytes.Buffer, iei uint8) error {
	if iei >= 0x80 {
		return nil
	}
	var length int
	if iei >= 0x70 {
		var lenOctets [2]byte
		if _, err := io.ReadFull(buffer, lenOctets[:]); err != nil {
			return err
		}
		length = int(binary.BigEndian.Uint16(lenOctets[:]))
	} else {
		lenOctet, err := buffer.ReadByte()
		if err != nil {
			return err
		}
		length = int(lenOctet)
	}
	if buffer.Len() < length {
		return fmt.Errorf("%w: unknown IEI 0x%02x has length %d but only %d octets remain",
			ErrMalformedIE, iei, length, buffer.Len())
	}
	buffer.Next(length)
	return nil
}
