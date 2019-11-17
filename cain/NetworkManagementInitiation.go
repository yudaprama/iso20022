package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:cain.009.001.01 Document"`
	Message *NetworkManagementInitiation `xml:"NtwkMgmtInitn"`
}

func (d *Document00900101) AddMessage() *NetworkManagementInitiation {
	d.Message = new(NetworkManagementInitiation)
	return d.Message
}

// The NetworkManagementInitiation message covers the range of activities to control the operating condition of the network and may be initiated by any party to an acquirer, an issuer or an agent.
type NetworkManagementInitiation struct {

	// Information related to the protocol management.
	Header *model.Header17 `xml:"Hdr"`

	// Information related to the network management.
	NetworkManagementInitiation *model.AcquirerNetworkManagementInitiation1 `xml:"NtwkMgmtInitn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (n *NetworkManagementInitiation) AddHeader() *model.Header17 {
	n.Header = new(model.Header17)
	return n.Header
}

func (n *NetworkManagementInitiation) AddNetworkManagementInitiation() *model.AcquirerNetworkManagementInitiation1 {
	n.NetworkManagementInitiation = new(model.AcquirerNetworkManagementInitiation1)
	return n.NetworkManagementInitiation
}

func (n *NetworkManagementInitiation) AddSecurityTrailer() *model.ContentInformationType15 {
	n.SecurityTrailer = new(model.ContentInformationType15)
	return n.SecurityTrailer
}
