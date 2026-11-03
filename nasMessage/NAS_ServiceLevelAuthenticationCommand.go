// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// ServiceLevelAuthenticationCommand 8.3.17
type ServiceLevelAuthenticationCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity
	nasType.ServiceLevelAAContainer
}

func NewServiceLevelAuthenticationCommand(iei uint8) (serviceLevelAuthenticationCommand *ServiceLevelAuthenticationCommand) {
	serviceLevelAuthenticationCommand = &ServiceLevelAuthenticationCommand{}
	return serviceLevelAuthenticationCommand
}

func (a *ServiceLevelAuthenticationCommand) EncodeServiceLevelAuthenticationCommand(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return
	}
}

func (a *ServiceLevelAuthenticationCommand) DecodeServiceLevelAuthenticationCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
		return err
	}
	a.SetLen(a.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return err
	}
	return nil
}
