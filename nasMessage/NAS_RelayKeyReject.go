// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RelayKeyReject 8.2.36
type RelayKeyReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RELAYKEYREJECTMessageIdentity
	nasType.ProSeRelayTransactionIdentity
	*nasType.EAPMessage
}

func NewRelayKeyReject(iei uint8) (relayKeyReject *RelayKeyReject) {
	relayKeyReject = &RelayKeyReject{}
	return relayKeyReject
}

const (
	RelayKeyRejectEAPMessageType uint8 = 0x78
)

func (a *RelayKeyReject) EncodeRelayKeyReject(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.RELAYKEYREJECTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet); err != nil {
		return
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
			return
		}
	}
}

func (a *RelayKeyReject) DecodeRelayKeyReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RELAYKEYREJECTMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ProSeRelayTransactionIdentity.Octet); err != nil {
		return err
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return err
		}
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		switch tmpIeiN {
		case RelayKeyRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
				return err
			}
			a.SetLen(a.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Buffer[:a.GetLen()]); err != nil {
				return err
			}
		default:
			if err := skipUnknownIE(buffer, ieiN); err != nil {
				return err
			}
		}
	}
	return nil
}
