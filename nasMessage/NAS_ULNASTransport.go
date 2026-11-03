// Copyright (C) 2026 Intel Corporation
// Copyright 2019 free5GC.org
// SPDX-License-Identifier: Apache-2.0

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/omec-project/nas/v2/nasType"
)

type ULNASTransport struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ULNASTRANSPORTMessageIdentity
	nasType.SpareHalfOctetAndPayloadContainerType
	nasType.PayloadContainer
	*nasType.PduSessionID2Value
	*nasType.OldPDUSessionID
	*nasType.RequestType
	*nasType.SNSSAI
	*nasType.DNN
	*nasType.AdditionalInformation
	*nasType.MAPDUSessionInformation
	*nasType.ReleaseAssistanceIndication
	*nasType.Non3GPPAccessPathSwitchingIndication
	AlternativeSNSSAI *nasType.SNSSAI
	*nasType.PayloadContainerInformation
}

func NewULNASTransport(iei uint8) (uLNASTransport *ULNASTransport) {
	uLNASTransport = &ULNASTransport{}
	return uLNASTransport
}

const (
	ULNASTransportPduSessionID2ValueType                   uint8 = 0x12
	ULNASTransportOldPDUSessionIDType                      uint8 = 0x59
	ULNASTransportRequestTypeType                          uint8 = 0x08
	ULNASTransportSNSSAIType                               uint8 = 0x22
	ULNASTransportDNNType                                  uint8 = 0x25
	ULNASTransportAdditionalInformationType                uint8 = 0x24
	ULNASTransportMAPDUSessionInformationType              uint8 = 0x0A
	ULNASTransportReleaseAssistanceIndicationType          uint8 = 0x0F
	ULNASTransportNon3GPPAccessPathSwitchingIndicationType uint8 = 0x4E
	ULNASTransportAlternativeSNSSAIType                    uint8 = 0x5A
	ULNASTransportPayloadContainerInformationType          uint8 = 0x09
)

func (a *ULNASTransport) EncodeULNASTransport(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ULNASTRANSPORTMessageIdentity.Octet); err != nil {
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
	if a.OldPDUSessionID != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.OldPDUSessionID.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.OldPDUSessionID.Octet); err != nil {
			return
		}
	}
	if a.RequestType != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.RequestType.Octet); err != nil {
			return
		}
	}
	if a.SNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
			return
		}
	}
	if a.DNN != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.DNN.Buffer); err != nil {
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
	if a.MAPDUSessionInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.MAPDUSessionInformation.Octet); err != nil {
			return
		}
	}
	if a.ReleaseAssistanceIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.ReleaseAssistanceIndication.Octet); err != nil {
			return
		}
	}
	if a.Non3GPPAccessPathSwitchingIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GPPAccessPathSwitchingIndication.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Non3GPPAccessPathSwitchingIndication.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Non3GPPAccessPathSwitchingIndication.Octet); err != nil {
			return
		}
	}
	if a.AlternativeSNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AlternativeSNSSAI.Octet[:a.AlternativeSNSSAI.GetLen()]); err != nil {
			return
		}
	}
	if a.PayloadContainerInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.PayloadContainerInformation.Octet); err != nil {
			return
		}
	}
}

func (a *ULNASTransport) DecodeULNASTransport(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ULNASTRANSPORTMessageIdentity.Octet); err != nil {
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
		case ULNASTransportPduSessionID2ValueType:
			a.PduSessionID2Value = nasType.NewPduSessionID2Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet); err != nil {
				return err
			}
		case ULNASTransportOldPDUSessionIDType:
			a.OldPDUSessionID = nasType.NewOldPDUSessionID(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.OldPDUSessionID.Octet); err != nil {
				return err
			}
		case ULNASTransportRequestTypeType:
			a.RequestType = nasType.NewRequestType(ieiN)
			a.RequestType.Octet = ieiN
		case ULNASTransportSNSSAIType:
			a.SNSSAI = nasType.NewSNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len); err != nil {
				return err
			}
			a.SNSSAI.SetLen(a.SNSSAI.GetLen())
			if a.SNSSAI.GetLen() > uint8(len(a.SNSSAI.Octet)) {
				a.SNSSAI = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: SNSSAI length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
				return err
			}
		case ULNASTransportDNNType:
			a.DNN = nasType.NewDNN(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.DNN.Len); err != nil {
				return err
			}
			a.DNN.SetLen(a.DNN.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.DNN.Buffer[:a.DNN.GetLen()]); err != nil {
				return err
			}
		case ULNASTransportAdditionalInformationType:
			a.AdditionalInformation = nasType.NewAdditionalInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AdditionalInformation.Len); err != nil {
				return err
			}
			a.AdditionalInformation.SetLen(a.AdditionalInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AdditionalInformation.Buffer[:a.AdditionalInformation.GetLen()]); err != nil {
				return err
			}
		case ULNASTransportMAPDUSessionInformationType:
			a.MAPDUSessionInformation = nasType.NewMAPDUSessionInformation(ieiN)
			a.MAPDUSessionInformation.Octet = ieiN
		case ULNASTransportReleaseAssistanceIndicationType:
			a.ReleaseAssistanceIndication = nasType.NewReleaseAssistanceIndication(ieiN)
			a.ReleaseAssistanceIndication.Octet = ieiN
		case ULNASTransportNon3GPPAccessPathSwitchingIndicationType:
			a.Non3GPPAccessPathSwitchingIndication = nasType.NewNon3GPPAccessPathSwitchingIndication(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Non3GPPAccessPathSwitchingIndication.Len); err != nil {
				return err
			}
			a.Non3GPPAccessPathSwitchingIndication.SetLen(a.Non3GPPAccessPathSwitchingIndication.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Non3GPPAccessPathSwitchingIndication.Octet); err != nil {
				return err
			}
		case ULNASTransportAlternativeSNSSAIType:
			a.AlternativeSNSSAI = nasType.NewSNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AlternativeSNSSAI.Len); err != nil {
				return err
			}
			a.AlternativeSNSSAI.SetLen(a.AlternativeSNSSAI.GetLen())
			if a.AlternativeSNSSAI.GetLen() > uint8(len(a.AlternativeSNSSAI.Octet)) {
				a.AlternativeSNSSAI = nil // discard the malformed IE so a later Encode cannot re-panic on it
				return fmt.Errorf("%w: AlternativeSNSSAI length exceeds maximum", ErrMalformedIE)
			}
			if err := binary.Read(buffer, binary.BigEndian, a.AlternativeSNSSAI.Octet[:a.AlternativeSNSSAI.GetLen()]); err != nil {
				return err
			}
		case ULNASTransportPayloadContainerInformationType:
			a.PayloadContainerInformation = nasType.NewPayloadContainerInformation(ieiN)
			a.PayloadContainerInformation.Octet = ieiN
		default:
			if err := skipUnknownIE(buffer, ieiN); err != nil {
				return err
			}
		}
	}
	return nil
}
