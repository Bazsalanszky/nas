// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"net"

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
	if a.SpareHalfOctetAndPayloadContainerType.Octet == uint8(1) {
		// best-effort extraction of the UE IP address; malformed/short ESM content just stops this scan
		func() {
			esmMsg := bytes.NewBuffer(a.PayloadContainer.Buffer)
			var pd uint8
			if err := binary.Read(esmMsg, binary.BigEndian, &pd); err != nil {
				return
			}
			if err := binary.Read(esmMsg, binary.BigEndian, &pd); err != nil {
				return
			}
			if err := binary.Read(esmMsg, binary.BigEndian, &pd); err != nil {
				return
			}
			if err := binary.Read(esmMsg, binary.BigEndian, &pd); err != nil {
				return
			}
			if err := binary.Read(esmMsg, binary.BigEndian, &pd); err != nil {
				return
			}
			var msgLen uint16
			if err := binary.Read(esmMsg, binary.BigEndian, &msgLen); err != nil {
				return
			}
			var qos [9]uint8
			if err := binary.Read(esmMsg, binary.BigEndian, &qos); err != nil {
				return
			}
			var ambr_len uint8
			if err := binary.Read(esmMsg, binary.BigEndian, &ambr_len); err != nil {
				return
			}
			var ambr [6]uint8
			if err := binary.Read(esmMsg, binary.BigEndian, &ambr); err != nil {
				return
			}
		forLoop:
			for esmMsg.Len() > 0 {
				var ieiN uint8
				if err := binary.Read(esmMsg, binary.BigEndian, &ieiN); err != nil {
					return
				}
				switch ieiN {
				case 89:
					var cause uint8
					if err := binary.Read(esmMsg, binary.BigEndian, &cause); err != nil {
						return
					}
				case 41:
					var iplen uint8
					if err := binary.Read(esmMsg, binary.BigEndian, &iplen); err != nil {
						return
					}
					var iptype uint8
					if err := binary.Read(esmMsg, binary.BigEndian, &iptype); err != nil {
						return
					}
					var ipaddr [4]uint8
					if err := binary.Read(esmMsg, binary.BigEndian, &ipaddr); err != nil {
						return
					}
					ip := net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
					a.Ipaddr = ip.String()
					break forLoop // we just need ip address nothing more
				default:
				}
			}
		}()
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
