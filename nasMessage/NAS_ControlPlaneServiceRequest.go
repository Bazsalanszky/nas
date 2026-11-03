// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/omec-project/nas/v2/nasType"
)

// ControlPlaneServiceRequest 8.2.30
type ControlPlaneServiceRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.CONTROLPLANESERVICEREQUESTMessageIdentity
	nasType.ControlPlaneServiceTypeAndNgksi
	*nasType.PDUSessionStatus
	*nasType.UplinkDataStatus
	*nasType.NASMessageContainer
	*nasType.AllowedPDUSessionStatus
	*nasType.UERequestType
	*nasType.PagingRestriction
	*nasType.CIoTSmallDataContainer
	*nasType.PduSessionID2Value
	*nasType.AdditionalInformation
	*nasType.ReleaseAssistanceIndication
}

func NewControlPlaneServiceRequest(iei uint8) (controlPlaneServiceRequest *ControlPlaneServiceRequest) {
	controlPlaneServiceRequest = &ControlPlaneServiceRequest{}
	return controlPlaneServiceRequest
}

const (
	ControlPlaneServiceRequestPDUSessionStatusType            uint8 = 0x50
	ControlPlaneServiceRequestUplinkDataStatusType            uint8 = 0x40
	ControlPlaneServiceRequestNASMessageContainerType         uint8 = 0x71
	ControlPlaneServiceRequestAllowedPDUSessionStatusType     uint8 = 0x25
	ControlPlaneServiceRequestUERequestTypeType               uint8 = 0x29
	ControlPlaneServiceRequestPagingRestrictionType           uint8 = 0x28
	ControlPlaneServiceRequestCIoTSmallDataContainerType      uint8 = 0x6F
	ControlPlaneServiceRequestPduSessionIDType                uint8 = 0x12
	ControlPlaneServiceRequestAdditionalInformationType       uint8 = 0x24
	ControlPlaneServiceRequestReleaseAssistanceIndicationType uint8 = 0x0F
)

func (a *ControlPlaneServiceRequest) EncodeControlPlaneServiceRequest(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.CONTROLPLANESERVICEREQUESTMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ControlPlaneServiceTypeAndNgksi.Octet); err != nil {
		return
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
	if a.UERequestType != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UERequestType.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.UERequestType.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UERequestType.Buffer[:uint8(a.UERequestType.GetLen())]); err != nil {
			return
		}
	}
	if a.PagingRestriction != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PagingRestriction.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, uint8(a.PagingRestriction.GetLen())); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PagingRestriction.Buffer[:uint8(a.PagingRestriction.GetLen())]); err != nil {
			return
		}
	}
	if a.CIoTSmallDataContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.CIoTSmallDataContainer.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.CIoTSmallDataContainer.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.CIoTSmallDataContainer.Buffer); err != nil {
			return
		}
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
	if a.ReleaseAssistanceIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.ReleaseAssistanceIndication.Octet); err != nil {
			return
		}
	}
}

func (a *ControlPlaneServiceRequest) DecodeControlPlaneServiceRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.CONTROLPLANESERVICEREQUESTMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ControlPlaneServiceTypeAndNgksi.Octet); err != nil {
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
		case ControlPlaneServiceRequestPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return err
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
				return err
			}
		case ControlPlaneServiceRequestUplinkDataStatusType:
			a.UplinkDataStatus = nasType.NewUplinkDataStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.UplinkDataStatus.Len); err != nil {
				return err
			}
			a.UplinkDataStatus.SetLen(a.UplinkDataStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer); err != nil {
				return err
			}
		case ControlPlaneServiceRequestNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len); err != nil {
				return err
			}
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer); err != nil {
				return err
			}
		case ControlPlaneServiceRequestAllowedPDUSessionStatusType:
			a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Len); err != nil {
				return err
			}
			a.AllowedPDUSessionStatus.SetLen(a.AllowedPDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer); err != nil {
				return err
			}
		case ControlPlaneServiceRequestUERequestTypeType:
			a.UERequestType = nasType.NewUERequestType(ieiN)
			var lenN0 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN0); err != nil {
				return err
			}
			a.UERequestType.SetLen(uint16(lenN0))
			if err := binary.Read(buffer, binary.BigEndian, a.UERequestType.Buffer[:lenN0]); err != nil {
				return err
			}
		case ControlPlaneServiceRequestPagingRestrictionType:
			a.PagingRestriction = nasType.NewPagingRestriction(ieiN)
			var lenN1 uint8
			if err := binary.Read(buffer, binary.BigEndian, &lenN1); err != nil {
				return err
			}
			a.PagingRestriction.SetLen(uint16(lenN1))
			if err := binary.Read(buffer, binary.BigEndian, a.PagingRestriction.Buffer[:lenN1]); err != nil {
				return err
			}
		case ControlPlaneServiceRequestCIoTSmallDataContainerType:
			a.CIoTSmallDataContainer = nasType.NewCIoTSmallDataContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.CIoTSmallDataContainer.Len); err != nil {
				return err
			}
			a.CIoTSmallDataContainer.SetLen(a.CIoTSmallDataContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.CIoTSmallDataContainer.Buffer[:a.CIoTSmallDataContainer.GetLen()]); err != nil {
				return err
			}
		case ControlPlaneServiceRequestPduSessionIDType:
			a.PduSessionID2Value = nasType.NewPduSessionID2Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet); err != nil {
				return err
			}
		case ControlPlaneServiceRequestAdditionalInformationType:
			a.AdditionalInformation = nasType.NewAdditionalInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AdditionalInformation.Len); err != nil {
				return err
			}
			a.AdditionalInformation.SetLen(a.AdditionalInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AdditionalInformation.Buffer[:a.AdditionalInformation.GetLen()]); err != nil {
				return err
			}
		case ControlPlaneServiceRequestReleaseAssistanceIndicationType:
			a.ReleaseAssistanceIndication = nasType.NewReleaseAssistanceIndication(ieiN)
			a.ReleaseAssistanceIndication.Octet = ieiN
		default:
			if err := skipUnknownIE(buffer, ieiN); err != nil {
				return err
			}
		}
	}
	return nil
}
