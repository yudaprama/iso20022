package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.006.001.01 Document"`
	Message *ATMDiagnosticResponseV01 `xml:"ATMDgnstcRspn"`
}

func (d *Document00600101) AddMessage() *ATMDiagnosticResponseV01 {
	d.Message = new(ATMDiagnosticResponseV01)
	return d.Message
}

// The ATMDiagnosticResponse message is sent by an acquirer to an ATM in response to an ATMDiagnosticRequest message ensuring the availability and the validity of the parameters.
type ATMDiagnosticResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDiagnosticResponse *model.ContentInformationType10 `xml:"PrtctdATMDgnstcRspn,omitempty"`

	// Information related to the response of a diagnostic from an ATM manager.
	ATMDiagnosticResponse *model.ATMDiagnosticResponse1 `xml:"ATMDgnstcRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDiagnosticResponseV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMDiagnosticResponseV01) AddProtectedATMDiagnosticResponse() *model.ContentInformationType10 {
	a.ProtectedATMDiagnosticResponse = new(model.ContentInformationType10)
	return a.ProtectedATMDiagnosticResponse
}

func (a *ATMDiagnosticResponseV01) AddATMDiagnosticResponse() *model.ATMDiagnosticResponse1 {
	a.ATMDiagnosticResponse = new(model.ATMDiagnosticResponse1)
	return a.ATMDiagnosticResponse
}

func (a *ATMDiagnosticResponseV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
