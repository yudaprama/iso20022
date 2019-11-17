package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.007.001.01 Document"`
	Message *ATMInquiryResponseV01 `xml:"ATMNqryRspn"`
}

func (d *Document00700101) AddMessage() *ATMInquiryResponseV01 {
	d.Message = new(ATMInquiryResponseV01)
	return d.Message
}

// The ATMInquiryResponse message is sent by an ATM manager or its agent to the ATM to provide the information and the outcome of the verifications requested in the ATMInquiryRequest.
type ATMInquiryResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMInquiryResponse *model.ContentInformationType10 `xml:"PrtctdATMNqryRspn,omitempty"`

	// Information related to the response of an ATM inquiry from an ATM manager.
	ATMInquiryResponse *model.ATMInquiryResponse1 `xml:"ATMNqryRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMInquiryResponseV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMInquiryResponseV01) AddProtectedATMInquiryResponse() *model.ContentInformationType10 {
	a.ProtectedATMInquiryResponse = new(model.ContentInformationType10)
	return a.ProtectedATMInquiryResponse
}

func (a *ATMInquiryResponseV01) AddATMInquiryResponse() *model.ATMInquiryResponse1 {
	a.ATMInquiryResponse = new(model.ATMInquiryResponse1)
	return a.ATMInquiryResponse
}

func (a *ATMInquiryResponseV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
