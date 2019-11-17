package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Document"`
	Message *HostToATMRequestV01 `xml:"HstToATMReq"`
}

func (d *Document00700101) AddMessage() *HostToATMRequestV01 {
	d.Message = new(HostToATMRequestV01)
	return d.Message
}

// The HostToATMRequest message is sent by a host to an ATM to request the ATM to contact a host by sending of a maintenance messages.
type HostToATMRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedHostToATMRequest *model.ContentInformationType10 `xml:"PrtctdHstToATMReq,omitempty"`

	// Information related to the request to an ATM to contact the ATM manager.
	HostToATMRequest *model.HostToATMRequest1 `xml:"HstToATMReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (h *HostToATMRequestV01) AddHeader() *model.Header20 {
	h.Header = new(model.Header20)
	return h.Header
}

func (h *HostToATMRequestV01) AddProtectedHostToATMRequest() *model.ContentInformationType10 {
	h.ProtectedHostToATMRequest = new(model.ContentInformationType10)
	return h.ProtectedHostToATMRequest
}

func (h *HostToATMRequestV01) AddHostToATMRequest() *model.HostToATMRequest1 {
	h.HostToATMRequest = new(model.HostToATMRequest1)
	return h.HostToATMRequest
}

func (h *HostToATMRequestV01) AddSecurityTrailer() *model.ContentInformationType15 {
	h.SecurityTrailer = new(model.ContentInformationType15)
	return h.SecurityTrailer
}
