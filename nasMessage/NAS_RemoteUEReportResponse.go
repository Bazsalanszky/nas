// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RemoteUEReportResponse 8.3.20
type RemoteUEReportResponse struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.REMOTEUEREPORTRESPONSEMessageIdentity
}

func NewRemoteUEReportResponse(iei uint8) (remoteUEReportResponse *RemoteUEReportResponse) {
	remoteUEReportResponse = &RemoteUEReportResponse{}
	return remoteUEReportResponse
}

func (a *RemoteUEReportResponse) EncodeRemoteUEReportResponse(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.REMOTEUEREPORTRESPONSEMessageIdentity.Octet); err != nil {
		return
	}
}

func (a *RemoteUEReportResponse) DecodeRemoteUEReportResponse(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.REMOTEUEREPORTRESPONSEMessageIdentity.Octet); err != nil {
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
		default:
			if err := skipUnknownIE(buffer, ieiN); err != nil {
				return err
			}
		}
	}
	return nil
}
