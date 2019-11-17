package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.01 Document"`
	Message *AcceptorDiagnosticResponseV01 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400101) AddMessage() *AcceptorDiagnosticResponseV01 {
	d.Message = new(AcceptorDiagnosticResponseV01)
	return d.Message
}

// Scope
// The AcceptorDiagnosticResponse message is sent by the acquirer to the card acceptor to confirm the availability of the acquirer. An agent never forwards the message.
// Usage
// The AcceptorDiagnosticResponse message is used to:
// - confirm the availability of the acquirer;
// - validate the security of the exchanges with the acquirer;
// - validate the version of the configuration parameters.
type AcceptorDiagnosticResponseV01 struct {

	// Diagnostic response message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *model.AcceptorDiagnosticResponse1 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticResponseV01) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV01) AddDiagnosticResponse() *model.AcceptorDiagnosticResponse1 {
	a.DiagnosticResponse = new(model.AcceptorDiagnosticResponse1)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV01) AddSecurityTrailer() *model.ContentInformationType3 {
	a.SecurityTrailer = new(model.ContentInformationType3)
	return a.SecurityTrailer
}
