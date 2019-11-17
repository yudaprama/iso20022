package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:catp.006.001.01 Document"`
	Message *ATMInquiryRequestV01 `xml:"ATMNqryReq"`
}

func (d *Document00600101) AddMessage() *ATMInquiryRequestV01 {
	d.Message = new(ATMInquiryRequestV01)
	return d.Message
}

// The ATMInquiryRequest message is sent by an ATM to an ATM manager to request information about a customer (for example card, account).
type ATMInquiryRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMInquiryRequest *model.ContentInformationType10 `xml:"PrtctdATMNqryReq,omitempty"`

	// Information related to the request of an inquiry from an ATM.
	ATMInquiryRequest *model.ATMInquiryRequest1 `xml:"ATMNqryReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMInquiryRequestV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMInquiryRequestV01) AddProtectedATMInquiryRequest() *model.ContentInformationType10 {
	a.ProtectedATMInquiryRequest = new(model.ContentInformationType10)
	return a.ProtectedATMInquiryRequest
}

func (a *ATMInquiryRequestV01) AddATMInquiryRequest() *model.ATMInquiryRequest1 {
	a.ATMInquiryRequest = new(model.ATMInquiryRequest1)
	return a.ATMInquiryRequest
}

func (a *ATMInquiryRequestV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
