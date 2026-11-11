// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasMessage"
)

type ProtocolOrContainerUnit struct {
	ProtocolOrContainerID uint16
	LengthOfContents      uint8
	Contents              []byte
}

type ProtocolConfigurationOptions struct {
	ProtocolOrContainerList []*ProtocolOrContainerUnit
}

type PCOReadingState int

const (
	ReadingID PCOReadingState = iota
	ReadingLength
	ReadingContent
)

func NewProtocolOrContainerUnit() (pcu *ProtocolOrContainerUnit) {
	pcu = &ProtocolOrContainerUnit{
		ProtocolOrContainerID: 0,
		LengthOfContents:      0,
		Contents:              []byte{},
	}
	return
}

func NewProtocolConfigurationOptions() (pco *ProtocolConfigurationOptions) {
	pco = &ProtocolConfigurationOptions{
		ProtocolOrContainerList: make([]*ProtocolOrContainerUnit, 0),
	}

	return
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) Marshal() []byte {
	var metaInfo uint8
	var extension uint8 = 1
	var spare uint8 = 0
	var configurationProtocol uint8 = 0
	buffer := new(bytes.Buffer)

	metaInfo = (extension << 7) | (spare << 6) | configurationProtocol
	if err := binary.Write(buffer, binary.BigEndian, &metaInfo); err != nil {
		logger.ConvertLog.Warnf("write metaInfo failed: %+v", err)
	}

	for _, containerUnit := range protocolConfigurationOptions.ProtocolOrContainerList {
		if err := binary.Write(buffer, binary.BigEndian, &containerUnit.ProtocolOrContainerID); err != nil {
			logger.ConvertLog.Warnf("write protocolOrContainerID failed: %+v", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &containerUnit.LengthOfContents); err != nil {
			logger.ConvertLog.Warnf("write length of contents failed: %+v", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, &containerUnit.Contents); err != nil {
			logger.ConvertLog.Warnf("write contents failed: %+v", err)
		}
	}

	return buffer.Bytes()
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) UnMarshal(data []byte) error {
	logger.ConvertLog.Debugln("in ProtocolConfigurationOptions UnMarshal")

	var Buf uint8
	numOfBytes := len(data)
	byteReader := bytes.NewReader(data)
	if err := binary.Read(byteReader, binary.BigEndian, &Buf); err != nil {
		return err
	}

	numOfBytes = numOfBytes - 1
	readingState := ReadingID
	var curContainer *ProtocolOrContainerUnit

	for numOfBytes > 0 {
		switch readingState {
		case ReadingID:
			curContainer = NewProtocolOrContainerUnit()
			if err := binary.Read(byteReader, binary.BigEndian, &curContainer.ProtocolOrContainerID); err != nil {
				return err
			}
			logger.ConvertLog.Debugln("reading ID:", curContainer.ProtocolOrContainerID)
			readingState = ReadingLength
			numOfBytes = numOfBytes - 2
		case ReadingLength:
			if err := binary.Read(byteReader, binary.BigEndian, &curContainer.LengthOfContents); err != nil {
				return err
			}
			logger.ConvertLog.Debugln("reading Length:", curContainer.LengthOfContents)
			readingState = ReadingContent
			numOfBytes = numOfBytes - 1
			if curContainer.LengthOfContents == 0 {
				protocolConfigurationOptions.ProtocolOrContainerList = append(
					protocolConfigurationOptions.ProtocolOrContainerList, curContainer)
				logger.ConvertLog.Debugln("for loop ProtocolOrContainerList:",
					protocolConfigurationOptions.ProtocolOrContainerList)
			}
		case ReadingContent:
			if curContainer.LengthOfContents > 0 {
				curContainer.Contents = make([]uint8, curContainer.LengthOfContents)
				if err := binary.Read(byteReader, binary.BigEndian, curContainer.Contents); err != nil {
					return err
				}
				protocolConfigurationOptions.ProtocolOrContainerList = append(
					protocolConfigurationOptions.ProtocolOrContainerList, curContainer)
				logger.ConvertLog.Debugln("for loop ProtocolOrContainerList:",
					protocolConfigurationOptions.ProtocolOrContainerList)
			}
			numOfBytes = numOfBytes - int(curContainer.LengthOfContents)
			readingState = ReadingID
		}
	}

	logger.ConvertLog.Infoln("protocolOrContainerList:", protocolConfigurationOptions.ProtocolOrContainerList)
	return nil
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddDNSServerIPv4AddressRequest() {
	protocolOrContainerUnit := NewProtocolOrContainerUnit()

	protocolOrContainerUnit.ProtocolOrContainerID = nasMessage.DNSServerIPv4AddressRequestUL
	protocolOrContainerUnit.LengthOfContents = 0

	protocolConfigurationOptions.ProtocolOrContainerList = append(protocolConfigurationOptions.ProtocolOrContainerList,
		protocolOrContainerUnit)
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddDNSServerIPv6AddressRequest() {
	protocolOrContainerUnit := NewProtocolOrContainerUnit()

	protocolOrContainerUnit.ProtocolOrContainerID = nasMessage.DNSServerIPv6AddressRequestUL
	protocolOrContainerUnit.LengthOfContents = 0

	protocolConfigurationOptions.ProtocolOrContainerList = append(protocolConfigurationOptions.ProtocolOrContainerList,
		protocolOrContainerUnit)
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddIPAddressAllocationViaNASSignallingUL() {
	protocolOrContainerUnit := NewProtocolOrContainerUnit()

	protocolOrContainerUnit.ProtocolOrContainerID = nasMessage.IPAddressAllocationViaNASSignallingUL
	protocolOrContainerUnit.LengthOfContents = 0

	protocolConfigurationOptions.ProtocolOrContainerList = append(protocolConfigurationOptions.ProtocolOrContainerList,
		protocolOrContainerUnit)
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddDNSServerIPv4Address(dnsIP net.IP) (err error) {
	if dnsIP.To4() == nil {
		return fmt.Errorf("the DNS IP should be IPv4 in AddDNSServerIPv4Address")
	}
	dnsIP = dnsIP.To4()

	if len(dnsIP) != net.IPv4len {
		return fmt.Errorf("the length of DNS IPv4 is wrong")
	}

	logger.ConvertLog.Debugln("in AddDNSServerIPv4Address")
	protocolOrContainerUnit := NewProtocolOrContainerUnit()

	protocolOrContainerUnit.ProtocolOrContainerID = nasMessage.DNSServerIPv4AddressDL
	protocolOrContainerUnit.LengthOfContents = uint8(net.IPv4len)
	logger.ConvertLog.Debugln("lengthOfContents:", protocolOrContainerUnit.LengthOfContents)
	protocolOrContainerUnit.Contents = append(protocolOrContainerUnit.Contents, dnsIP.To4()...)
	logger.ConvertLog.Debugln("contents:", protocolOrContainerUnit.Contents)

	protocolConfigurationOptions.ProtocolOrContainerList = append(protocolConfigurationOptions.ProtocolOrContainerList,
		protocolOrContainerUnit)
	return
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddPCSCFIPv4Address(pcscfIP net.IP) (err error) {
	if pcscfIP.To4() == nil {
		return fmt.Errorf("the P-CSCF IP should be IPv4")
	}
	pcscfIP = pcscfIP.To4()

	if len(pcscfIP) != net.IPv4len {
		return fmt.Errorf("the length of P-CSCF IP IPv4 is wrong")
	}

	logger.ConvertLog.Debugln("in AddDNSServerIPv4Address")
	protocolOrContainerUnit := NewProtocolOrContainerUnit()
	protocolOrContainerUnit.ProtocolOrContainerID = nasMessage.PCSCFIPv4AddressDL
	protocolOrContainerUnit.LengthOfContents = uint8(net.IPv4len)
	logger.ConvertLog.Debugln("lengthOfContents:", protocolOrContainerUnit.LengthOfContents)
	protocolOrContainerUnit.Contents = append(protocolOrContainerUnit.Contents, pcscfIP.To4()...)
	logger.ConvertLog.Debugln("contents:", protocolOrContainerUnit.Contents)

	protocolConfigurationOptions.ProtocolOrContainerList = append(protocolConfigurationOptions.ProtocolOrContainerList,
		protocolOrContainerUnit)
	return
}

// AddDNSServerIPv6Address adds a DNS Server IPv6 Address container (0x0003,
// TS 24.008 clause 10.5.6.3). dnsIP must be an IPv6 address; IPv4 addresses,
// including the IPv4-mapped form that net.ParseIP returns for dotted-quad
// input, are rejected.
func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddDNSServerIPv6Address(dnsIP net.IP) (err error) {
	if !isIPv6(dnsIP) {
		return fmt.Errorf("the DNS IP should be IPv6 in AddDNSServerIPv6Address")
	}

	protocolConfigurationOptions.addIPv6Container(nasMessage.DNSServerIPv6AddressDL, dnsIP)
	return
}

// AddPCSCFIPv6Address adds a P-CSCF IPv6 Address container (0x0001,
// TS 24.008 clause 10.5.6.3). pcscfIP must be an IPv6 address; IPv4 and
// IPv4-mapped addresses are rejected.
func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddPCSCFIPv6Address(pcscfIP net.IP) (err error) {
	if !isIPv6(pcscfIP) {
		return fmt.Errorf("the P-CSCF IP should be IPv6")
	}

	protocolConfigurationOptions.addIPv6Container(nasMessage.PCSCFIPv6AddressDL, pcscfIP)
	return
}

// isIPv6 reports whether ip is a 16-byte IPv6 address that is not an
// IPv4-mapped IPv4 address.
func isIPv6(ip net.IP) bool {
	return len(ip) == net.IPv6len && ip.To4() == nil
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) addIPv6Container(id uint16, ip net.IP) {
	protocolOrContainerUnit := NewProtocolOrContainerUnit()

	protocolOrContainerUnit.ProtocolOrContainerID = id
	protocolOrContainerUnit.LengthOfContents = uint8(net.IPv6len)
	protocolOrContainerUnit.Contents = append(protocolOrContainerUnit.Contents, ip...)

	protocolConfigurationOptions.ProtocolOrContainerList = append(protocolConfigurationOptions.ProtocolOrContainerList,
		protocolOrContainerUnit)
}

func (protocolConfigurationOptions *ProtocolConfigurationOptions) AddIPv4LinkMTU(mtu uint16) (err error) {
	logger.ConvertLog.Debugln("in AddIPv4LinkMTU")
	protocolOrContainerUnit := NewProtocolOrContainerUnit()

	protocolOrContainerUnit.ProtocolOrContainerID = nasMessage.IPv4LinkMTUDL
	protocolOrContainerUnit.LengthOfContents = 2
	logger.ConvertLog.Debugln("lengthOfContents:", protocolOrContainerUnit.LengthOfContents)
	protocolOrContainerUnit.Contents = append(protocolOrContainerUnit.Contents, []byte{uint8(mtu >> 8), uint8(mtu & 0xff)}...)
	logger.ConvertLog.Debugln("contents:", protocolOrContainerUnit.Contents)

	protocolConfigurationOptions.ProtocolOrContainerList = append(protocolConfigurationOptions.ProtocolOrContainerList, protocolOrContainerUnit)
	return
}
