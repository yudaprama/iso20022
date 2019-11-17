package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600102 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.02 Document"`
	Message *ATMInquiryRequestV02 `xml:"ATMNqryReq"`
}

func (d *Document00600102) AddMessage() *ATMInquiryRequestV02 {
	d.Message = new(ATMInquiryRequestV02)
	return d.Message
}

// The ATMInquiryRequest message is sent by an ATM to an ATM manager to request information about a customer (for example card, account).
type ATMInquiryRequestV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMInquiryRequest *model.ContentInformationType10 `xml:"PrtctdATMNqryReq,omitempty"`

	// Information related to the request of an inquiry from an ATM.
	ATMInquiryRequest *model.ATMInquiryRequest2 `xml:"ATMNqryReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMInquiryRequestV02) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMInquiryRequestV02) AddProtectedATMInquiryRequest() *model.ContentInformationType10 {
	a.ProtectedATMInquiryRequest = new(model.ContentInformationType10)
	return a.ProtectedATMInquiryRequest
}

func (a *ATMInquiryRequestV02) AddATMInquiryRequest() *model.ATMInquiryRequest2 {
	a.ATMInquiryRequest = new(model.ATMInquiryRequest2)
	return a.ATMInquiryRequest
}

func (a *ATMInquiryRequestV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
