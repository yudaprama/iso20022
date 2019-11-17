package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.016.001.01 Document"`
	Message *ATMTransferRequestV01 `xml:"ATMTrfReq"`
}

func (d *Document01600101) AddMessage() *ATMTransferRequestV01 {
	d.Message = new(ATMTransferRequestV01)
	return d.Message
}

// The ATMTransferRequest message is sent by an ATM to an ATM manager to request the approval of a fund transfer at the ATM.
type ATMTransferRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMTransferRequest *model.ContentInformationType10 `xml:"PrtctdATMTrfReq,omitempty"`

	// Information related to the request of a fund transfer from an ATM.
	ATMTransferRequest *model.ATMTransferRequest1 `xml:"ATMTrfReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMTransferRequestV01) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMTransferRequestV01) AddProtectedATMTransferRequest() *model.ContentInformationType10 {
	a.ProtectedATMTransferRequest = new(model.ContentInformationType10)
	return a.ProtectedATMTransferRequest
}

func (a *ATMTransferRequestV01) AddATMTransferRequest() *model.ATMTransferRequest1 {
	a.ATMTransferRequest = new(model.ATMTransferRequest1)
	return a.ATMTransferRequest
}

func (a *ATMTransferRequestV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
