// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// RemoteUEReport 8.3.19
type RemoteUEReport struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.REMOTEUEREPORTMessageIdentity
	*nasType.RemoteUEContextList                              // Remote UE context connected (IEI 0x76)
	RemoteUEContextDisconnected  *nasType.RemoteUEContextList // Remote UE context disconnected (IEI 0x70)
}

func NewRemoteUEReport(iei uint8) (remoteUEReport *RemoteUEReport) {
	remoteUEReport = &RemoteUEReport{}
	return remoteUEReport
}

const (
	RemoteUEReportRemoteUEContextConnectedType    uint8 = 0x76
	RemoteUEReportRemoteUEContextDisconnectedType uint8 = 0x70
)

func (a *RemoteUEReport) EncodeRemoteUEReport(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.REMOTEUEREPORTMessageIdentity.Octet); err != nil {
		return
	}
	if a.RemoteUEContextList != nil {
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
	if a.RemoteUEContextDisconnected != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RemoteUEContextDisconnected.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RemoteUEContextDisconnected.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RemoteUEContextDisconnected.Buffer); err != nil {
			return
		}
	}
}

func (a *RemoteUEReport) DecodeRemoteUEReport(byteArray *[]byte) error {
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
	if err := binary.Read(buffer, binary.BigEndian, &a.REMOTEUEREPORTMessageIdentity.Octet); err != nil {
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
		case RemoteUEReportRemoteUEContextConnectedType:
			a.RemoteUEContextList = nasType.NewRemoteUEContextList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
				return err
			}
			a.SetLen(a.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Buffer); err != nil {
				return err
			}
		case RemoteUEReportRemoteUEContextDisconnectedType:
			a.RemoteUEContextDisconnected = nasType.NewRemoteUEContextList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RemoteUEContextDisconnected.Len); err != nil {
				return err
			}
			a.RemoteUEContextDisconnected.SetLen(a.RemoteUEContextDisconnected.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RemoteUEContextDisconnected.Buffer); err != nil {
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
