package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.007.001.02 Document"`
	Message *ATMInquiryResponseV02 `xml:"ATMNqryRspn"`
}

func (d *Document00700102) AddMessage() *ATMInquiryResponseV02 {
	d.Message = new(ATMInquiryResponseV02)
	return d.Message
}

// The ATMInquiryResponse message is sent by an ATM manager or its agent to the ATM to provide the information and the outcome of the verifications requested in the ATMInquiryRequest.
type ATMInquiryResponseV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMInquiryResponse *model.ContentInformationType10 `xml:"PrtctdATMNqryRspn,omitempty"`

	// Information related to the response of an ATM inquiry from an ATM manager.
	ATMInquiryResponse *model.ATMInquiryResponse2 `xml:"ATMNqryRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMInquiryResponseV02) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMInquiryResponseV02) AddProtectedATMInquiryResponse() *model.ContentInformationType10 {
	a.ProtectedATMInquiryResponse = new(model.ContentInformationType10)
	return a.ProtectedATMInquiryResponse
}

func (a *ATMInquiryResponseV02) AddATMInquiryResponse() *model.ATMInquiryResponse2 {
	a.ATMInquiryResponse = new(model.ATMInquiryResponse2)
	return a.ATMInquiryResponse
}

func (a *ATMInquiryResponseV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
