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

type SecurityModeCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.SecurityModeCommandMessageIdentity
	nasType.SelectedNASSecurityAlgorithms
	nasType.SpareHalfOctetAndNgksi
	nasType.ReplayedUESecurityCapabilities
	*nasType.IMEISVRequest
	*nasType.SelectedEPSNASSecurityAlgorithms
	*nasType.Additional5GSecurityInformation
	*nasType.EAPMessage
	*nasType.ABBA
	*nasType.ReplayedS1UESecurityCapabilities
	*nasType.MasterSessionKey
}

func NewSecurityModeCommand(iei uint8) (securityModeCommand *SecurityModeCommand) {
	securityModeCommand = &SecurityModeCommand{}
	return securityModeCommand
}

const (
	SecurityModeCommandIMEISVRequestType                    uint8 = 0x0E
	SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType uint8 = 0x57
	SecurityModeCommandAdditional5GSecurityInformationType  uint8 = 0x36
	SecurityModeCommandEAPMessageType                       uint8 = 0x78
	SecurityModeCommandABBAType                             uint8 = 0x38
	SecurityModeCommandReplayedS1UESecurityCapabilitiesType uint8 = 0x19
	SecurityModeCommandMasterSessionKeyType                 uint8 = 0x55
)

func (a *SecurityModeCommand) EncodeSecurityModeCommand(buffer *bytes.Buffer) {
	if err := binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SecurityModeCommandMessageIdentity.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SelectedNASSecurityAlgorithms.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ReplayedUESecurityCapabilities.GetLen()); err != nil {
		return
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.ReplayedUESecurityCapabilities.Buffer); err != nil {
		return
	}
	if a.IMEISVRequest != nil {
		if err := binary.Write(buffer, binary.BigEndian, &a.IMEISVRequest.Octet); err != nil {
			return
		}
	}
	if a.SelectedEPSNASSecurityAlgorithms != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SelectedEPSNASSecurityAlgorithms.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.SelectedEPSNASSecurityAlgorithms.Octet); err != nil {
			return
		}
	}
	if a.Additional5GSecurityInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Additional5GSecurityInformation.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Additional5GSecurityInformation.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Octet); err != nil {
			return
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer); err != nil {
			return
		}
	}
	if a.ABBA != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ABBA.Buffer); err != nil {
			return
		}
	}
	if a.ReplayedS1UESecurityCapabilities != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.ReplayedS1UESecurityCapabilities.Buffer); err != nil {
			return
		}
	}
	if a.MasterSessionKey != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MasterSessionKey.GetIei()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MasterSessionKey.GetLen()); err != nil {
			return
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.MasterSessionKey.Buffer); err != nil {
			return
		}
	}
}

func (a *SecurityModeCommand) DecodeSecurityModeCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SecurityModeCommandMessageIdentity.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SelectedNASSecurityAlgorithms.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ReplayedUESecurityCapabilities.Len); err != nil {
		return err
	}
	a.ReplayedUESecurityCapabilities.SetLen(a.ReplayedUESecurityCapabilities.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, &a.ReplayedUESecurityCapabilities.Buffer); err != nil {
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
		case SecurityModeCommandIMEISVRequestType:
			a.IMEISVRequest = nasType.NewIMEISVRequest(ieiN)
			a.IMEISVRequest.Octet = ieiN
		case SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType:
			a.SelectedEPSNASSecurityAlgorithms = nasType.NewSelectedEPSNASSecurityAlgorithms(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SelectedEPSNASSecurityAlgorithms.Octet); err != nil {
				return err
			}
		case SecurityModeCommandAdditional5GSecurityInformationType:
			a.Additional5GSecurityInformation = nasType.NewAdditional5GSecurityInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Len); err != nil {
				return err
			}
			a.Additional5GSecurityInformation.SetLen(a.Additional5GSecurityInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Octet); err != nil {
				return err
			}
		case SecurityModeCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return err
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()]); err != nil {
				return err
			}
		case SecurityModeCommandABBAType:
			a.ABBA = nasType.NewABBA(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ABBA.Len); err != nil {
				return err
			}
			a.ABBA.SetLen(a.ABBA.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ABBA.Buffer[:a.ABBA.GetLen()]); err != nil {
				return err
			}
		case SecurityModeCommandReplayedS1UESecurityCapabilitiesType:
			a.ReplayedS1UESecurityCapabilities = nasType.NewReplayedS1UESecurityCapabilities(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ReplayedS1UESecurityCapabilities.Len); err != nil {
				return err
			}
			a.ReplayedS1UESecurityCapabilities.SetLen(a.ReplayedS1UESecurityCapabilities.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.Buffer[:a.ReplayedS1UESecurityCapabilities.GetLen()]); err != nil {
				return err
			}
		case SecurityModeCommandMasterSessionKeyType:
			a.MasterSessionKey = nasType.NewMasterSessionKey(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MasterSessionKey.Len); err != nil {
				return err
			}
			a.MasterSessionKey.SetLen(a.MasterSessionKey.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.MasterSessionKey.Buffer[:a.MasterSessionKey.GetLen()]); err != nil {
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
