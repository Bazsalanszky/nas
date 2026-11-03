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

type ServiceRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceRequestMessageIdentity
	nasType.ServiceTypeAndNgksi
	nasType.TMSI5GS
	*nasType.UplinkDataStatus
	*nasType.PDUSessionStatus
	*nasType.AllowedPDUSessionStatus
	*nasType.NASMessageContainer
}

func NewServiceRequest(iei uint8) (serviceRequest *ServiceRequest) {
	serviceRequest = &ServiceRequest{}
	return serviceRequest
}

const (
	ServiceRequestUplinkDataStatusType        uint8 = 0x40
	ServiceRequestPDUSessionStatusType        uint8 = 0x50
	ServiceRequestAllowedPDUSessionStatusType uint8 = 0x25
	ServiceRequestNASMessageContainerType     uint8 = 0x71
)

func (a *ServiceRequest) EncodeServiceRequest(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ServiceRequestMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ServiceTypeAndNgksi.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.TMSI5GS.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.TMSI5GS.Octet); err != nil {
		return
	}
	if a.UplinkDataStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.UplinkDataStatus.Buffer); err != nil {
			return
		}
	}
	if a.PDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.PDUSessionStatus.Buffer); err != nil {
			return
		}
	}
	if a.AllowedPDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Buffer); err != nil {
			return
		}
	}
	if a.NASMessageContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer); err != nil {
			return
		}
	}
}

func (a *ServiceRequest) DecodeServiceRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ServiceRequestMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ServiceTypeAndNgksi.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.TMSI5GS.Len); err != nil {
		return err
	}
	a.TMSI5GS.SetLen(a.TMSI5GS.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.TMSI5GS.Octet); err != nil {
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
		case ServiceRequestUplinkDataStatusType:
			a.UplinkDataStatus = nasType.NewUplinkDataStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UplinkDataStatus.Len); err != nil {
				return err
			}
			a.UplinkDataStatus.SetLen(a.UplinkDataStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer[:a.UplinkDataStatus.GetLen()]); err != nil {
				return err
			}
		case ServiceRequestPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return err
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()]); err != nil {
				return err
			}
		case ServiceRequestAllowedPDUSessionStatusType:
			a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Len); err != nil {
				return err
			}
			a.AllowedPDUSessionStatus.SetLen(a.AllowedPDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer[:a.AllowedPDUSessionStatus.GetLen()]); err != nil {
				return err
			}
		case ServiceRequestNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len); err != nil {
				return err
			}
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer[:a.NASMessageContainer.GetLen()]); err != nil {
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
