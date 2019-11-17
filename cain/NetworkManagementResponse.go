package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:cain.010.001.01 Document"`
	Message *NetworkManagementResponse `xml:"NtwkMgmtRspn"`
}

func (d *Document01000101) AddMessage() *NetworkManagementResponse {
	d.Message = new(NetworkManagementResponse)
	return d.Message
}

// The NetworkManagementResponse message is sent by an acquirer, an issuer or an agent to answer to an NetworkManagementInitiation message.
type NetworkManagementResponse struct {

	// Information related to the protocol management.
	Header *model.Header17 `xml:"Hdr"`

	// Information related to the response to the network management.
	NetworkManagementResponse *model.AcquirerNetworkManagementResponse1 `xml:"NtwkMgmtRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (n *NetworkManagementResponse) AddHeader() *model.Header17 {
	n.Header = new(model.Header17)
	return n.Header
}

func (n *NetworkManagementResponse) AddNetworkManagementResponse() *model.AcquirerNetworkManagementResponse1 {
	n.NetworkManagementResponse = new(model.AcquirerNetworkManagementResponse1)
	return n.NetworkManagementResponse
}

func (n *NetworkManagementResponse) AddSecurityTrailer() *model.ContentInformationType15 {
	n.SecurityTrailer = new(model.ContentInformationType15)
	return n.SecurityTrailer
}
