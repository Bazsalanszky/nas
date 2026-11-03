// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// ServiceLevelAuthenticationComplete 8.3.18
type ServiceLevelAuthenticationComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity
	nasType.ServiceLevelAAContainer
}

func NewServiceLevelAuthenticationComplete(iei uint8) (serviceLevelAuthenticationComplete *ServiceLevelAuthenticationComplete) {
	serviceLevelAuthenticationComplete = &ServiceLevelAuthenticationComplete{}
	return serviceLevelAuthenticationComplete
}

func (a *ServiceLevelAuthenticationComplete) EncodeServiceLevelAuthenticationComplete(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Buffer); err != nil {
		return
	}
}

func (a *ServiceLevelAuthenticationComplete) DecodeServiceLevelAuthenticationComplete(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.SERVICELEVELAUTHENTICATIONCOMPLETEMessageIdentity.Octet); err != nil {
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
