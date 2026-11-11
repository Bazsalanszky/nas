// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasType"
)

type DLNASTransport struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DLNASTRANSPORTMessageIdentity
	nasType.SpareHalfOctetAndPayloadContainerType
	nasType.PayloadContainer
	*nasType.PduSessionID2Value
	*nasType.AdditionalInformation
	*nasType.Cause5GMM
	*nasType.BackoffTimerValue
	*nasType.LowerBoundTimerValue
	// Ipaddr is set by DecodeDLNASTransport, on a best-effort basis, to the
	// UE IPv4 address of a PDU Session Establishment Accept carried in the
	// payload container: the PDU address of an IPv4 session, or the IPv4 part
	// of an IPv4v6 session's PDU address. It is empty for every other message
	// and for IPv6 sessions, whose PDU address holds only an interface
	// identifier.
	Ipaddr string
}

func NewDLNASTransport(iei uint8) (dLNASTransport *DLNASTransport) {
	dLNASTransport = &DLNASTransport{}
	return dLNASTransport
}

const (
	DLNASTransportPduSessionID2ValueType    uint8 = 0x12
	DLNASTransportAdditionalInformationType uint8 = 0x24
	DLNASTransportCause5GMMType             uint8 = 0x58
	DLNASTransportBackoffTimerValueType     uint8 = 0x37
	DLNASTransportLowerBoundTimerValueType  uint8 = 0x3A
)

func (a *DLNASTransport) EncodeDLNASTransport(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.DLNASTRANSPORTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndPayloadContainerType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PayloadContainer.Buffer); err != nil {
		return
	}
	if a.PduSessionID2Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PduSessionID2Value.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet); err != nil {
			return
		}
	}
	if a.AdditionalInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AdditionalInformation.Buffer); err != nil {
			return
		}
	}
	if a.Cause5GMM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GMM.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
			return
		}
	}
	if a.BackoffTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
			return
		}
	}
	if a.LowerBoundTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LowerBoundTimerValue.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LowerBoundTimerValue.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet); err != nil {
			return
		}
	}
}

func (a *DLNASTransport) DecodeDLNASTransport(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.DLNASTRANSPORTMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndPayloadContainerType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Len); err != nil {
		return err
	}
	a.PayloadContainer.SetLen(a.PayloadContainer.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Buffer); err != nil {
		return err
	}
	if a.SpareHalfOctetAndPayloadContainerType.Octet == PayloadContainerTypeN1SMInfo {
		a.Ipaddr = ueIPv4AddressFromN1SM(a.PayloadContainer.Buffer)
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
		case DLNASTransportPduSessionID2ValueType:
			a.PduSessionID2Value = nasType.NewPduSessionID2Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet); err != nil {
				return err
			}
		case DLNASTransportAdditionalInformationType:
			a.AdditionalInformation = nasType.NewAdditionalInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AdditionalInformation.Len); err != nil {
				return err
			}
			a.AdditionalInformation.SetLen(a.AdditionalInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AdditionalInformation.Buffer[:a.AdditionalInformation.GetLen()]); err != nil {
				return err
			}
		case DLNASTransportCause5GMMType:
			a.Cause5GMM = nasType.NewCause5GMM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
				return err
			}
		case DLNASTransportBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len); err != nil {
				return err
			}
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet); err != nil {
				return err
			}
		case DLNASTransportLowerBoundTimerValueType:
			a.LowerBoundTimerValue = nasType.NewLowerBoundTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Len); err != nil {
				return err
			}
			a.LowerBoundTimerValue.SetLen(a.LowerBoundTimerValue.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.LowerBoundTimerValue.Octet); err != nil {
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

// msgTypePDUSessionEstablishmentAccept is the 5GSM message type of a PDU
// Session Establishment Accept (TS 24.501 Table 9.7.1). The nas package
// defines it as MsgTypePDUSessionEstablishmentAccept, but nasMessage cannot
// import nas.
const msgTypePDUSessionEstablishmentAccept uint8 = 0xc2

// ueIPv4AddressFromN1SM returns the UE IPv4 address assigned by a PDU Session
// Establishment Accept, or "" if n1sm is not one, carries no PDU address, or
// the session has no IPv4 address. The PDU address information is laid out
// by PDU session type (TS 24.501 clause 9.11.4.10): IPv4 address (4 octets),
// IPv6 interface identifier (8 octets), or interface identifier followed by
// the IPv4 address (12 octets).
func ueIPv4AddressFromN1SM(n1sm []byte) string {
	// Extended protocol discriminator, PDU session ID, PTI, message type.
	const msgTypeOffset = 3
	if len(n1sm) <= msgTypeOffset || n1sm[msgTypeOffset] != msgTypePDUSessionEstablishmentAccept {
		return ""
	}
	accept := NewPDUSessionEstablishmentAccept(0)
	if err := accept.DecodePDUSessionEstablishmentAccept(&n1sm); err != nil {
		// A PDU address decoded before the error (a later IE is truncated
		// or malformed) is still used.
		logger.NasMsgLog.Debugf("decode PDU Session Establishment Accept in DL NAS Transport: %v", err)
	}
	pduAddress := accept.PDUAddress
	if pduAddress == nil {
		return ""
	}
	info := pduAddress.GetPDUAddressInformation()
	var ip net.IP
	switch pduAddress.GetPDUSessionTypeValue() {
	case PDUSessionTypeIPv4:
		if pduAddress.GetLen() >= 1+net.IPv4len {
			ip = net.IP(info[0:net.IPv4len])
		}
	case PDUSessionTypeIPv4IPv6:
		if pduAddress.GetLen() >= 1+8+net.IPv4len {
			ip = net.IP(info[8 : 8+net.IPv4len])
		}
	}
	if ip == nil {
		return ""
	}
	return ip.String()
}
