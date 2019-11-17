package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500101 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.005.001.01 Document"`
	Message *ATMDiagnosticRequestV01 `xml:"ATMDgnstcReq"`
}

func (d *Document00500101) AddMessage() *ATMDiagnosticRequestV01 {
	d.Message = new(ATMDiagnosticRequestV01)
	return d.Message
}

// The ATMDiagnosticRequest message is sent from an ATM to an acquirer to verify the availability of the acquirer. The acquirer will also validate that this ATM is a valid ATM for its particular network.
type ATMDiagnosticRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDiagnosticRequest *model.ContentInformationType10 `xml:"PrtctdATMDgnstcReq,omitempty"`

	// Information related to the request of a diagnostic from an ATM.
	ATMDiagnosticRequest *model.ATMDiagnosticRequest1 `xml:"ATMDgnstcReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDiagnosticRequestV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMDiagnosticRequestV01) AddProtectedATMDiagnosticRequest() *model.ContentInformationType10 {
	a.ProtectedATMDiagnosticRequest = new(model.ContentInformationType10)
	return a.ProtectedATMDiagnosticRequest
}

func (a *ATMDiagnosticRequestV01) AddATMDiagnosticRequest() *model.ATMDiagnosticRequest1 {
	a.ATMDiagnosticRequest = new(model.ATMDiagnosticRequest1)
	return a.ATMDiagnosticRequest
}

func (a *ATMDiagnosticRequestV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
