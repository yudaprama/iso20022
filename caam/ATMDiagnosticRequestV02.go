package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.005.001.02 Document"`
	Message *ATMDiagnosticRequestV02 `xml:"ATMDgnstcReq"`
}

func (d *Document00500102) AddMessage() *ATMDiagnosticRequestV02 {
	d.Message = new(ATMDiagnosticRequestV02)
	return d.Message
}

// The ATMDiagnosticRequest message is sent from an ATM to an acquirer to verify the availability of the acquirer. The acquirer will also validate that this ATM is a valid ATM for its particular network.
type ATMDiagnosticRequestV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDiagnosticRequest *model.ContentInformationType10 `xml:"PrtctdATMDgnstcReq,omitempty"`

	// Information related to the request of a diagnostic from an ATM.
	ATMDiagnosticRequest *model.ATMDiagnosticRequest2 `xml:"ATMDgnstcReq,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDiagnosticRequestV02) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMDiagnosticRequestV02) AddProtectedATMDiagnosticRequest() *model.ContentInformationType10 {
	a.ProtectedATMDiagnosticRequest = new(model.ContentInformationType10)
	return a.ProtectedATMDiagnosticRequest
}

func (a *ATMDiagnosticRequestV02) AddATMDiagnosticRequest() *model.ATMDiagnosticRequest2 {
	a.ATMDiagnosticRequest = new(model.ATMDiagnosticRequest2)
	return a.ATMDiagnosticRequest
}

func (a *ATMDiagnosticRequestV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
