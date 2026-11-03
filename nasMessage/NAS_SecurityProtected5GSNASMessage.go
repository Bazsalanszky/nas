// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

type SecurityProtected5GSNASMessage struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.MessageAuthenticationCode
	nasType.SequenceNumber
	nasType.Plain5GSNASMessage
}

func NewSecurityProtected5GSNASMessage(iei uint8) (securityProtected5GSNASMessage *SecurityProtected5GSNASMessage) {
	securityProtected5GSNASMessage = &SecurityProtected5GSNASMessage{}
	return securityProtected5GSNASMessage
}

func (a *SecurityProtected5GSNASMessage) EncodeSecurityProtected5GSNASMessage(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.MessageAuthenticationCode.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SequenceNumber.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Plain5GSNASMessage); err != nil {
		return
	}
}

func (a *SecurityProtected5GSNASMessage) DecodeSecurityProtected5GSNASMessage(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.MessageAuthenticationCode.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SequenceNumber.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Plain5GSNASMessage); err != nil {
		return err
	}
	// The octets following the header are the (possibly ciphered) plain 5GS NAS message, not
	// optional information elements, so they must not be parsed here. Plain5GSNASMessage is an
	// empty struct, so the payload is not retained by this type.
	return nil
}
