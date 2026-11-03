// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"errors"
	"fmt"
)

// ErrMalformedIE is returned (wrapped) when an information element cannot be decoded
// because its length is invalid or exceeds the remaining message.
var ErrMalformedIE = errors.New("malformed information element")

// skipUnknownIE consumes the remainder of an optional information element whose IEI
// is not known in the message, so that the following IEs are still decoded correctly.
// TS 24.501 clause 9.11.1: the receiver determines whether the IE is of type 1 or 2
// (one octet, nothing left to skip) or of type 4 (the next octet is the length indicator).
// Assuming (TS 24.007 clause 11.2.4 is not in the spec set used for review): IEIs with
// bit 8 set are one octet long, any other unknown IEI is type 4. Unknown type 6 (TLV-E) IEs
// cannot be told apart from type 4 and are therefore not supported.
func skipUnknownIE(buffer *bytes.Buffer, iei uint8) error {
	if iei >= 0x80 {
		return nil
	}
	length, err := buffer.ReadByte()
	if err != nil {
		return err
	}
	if buffer.Len() < int(length) {
		return fmt.Errorf("%w: unknown IEI 0x%02x has length %d but only %d octets remain",
			ErrMalformedIE, iei, length, buffer.Len())
	}
	buffer.Next(int(length))
	return nil
}
