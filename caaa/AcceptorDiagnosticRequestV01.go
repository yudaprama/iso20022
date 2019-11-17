package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.01 Document"`
	Message *AcceptorDiagnosticRequestV01 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300101) AddMessage() *AcceptorDiagnosticRequestV01 {
	d.Message = new(AcceptorDiagnosticRequestV01)
	return d.Message
}

// Scope
// The AcceptorDiagnosticRequest message is sent by the card acceptor to the acquirer to ensure the availability of the acquirer. An agent never forwards the message.
// Usage
// The AcceptorDiagnosticRequest message is used to:
// - test the availability of the acquirer;
// - validate the security of the exchanges with the acquirer;
// - validate the version of the configuration parameters.
type AcceptorDiagnosticRequestV01 struct {

	// Diagnostic request message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *model.AcceptorDiagnosticRequest1 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticRequestV01) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV01) AddDiagnosticRequest() *model.AcceptorDiagnosticRequest1 {
	a.DiagnosticRequest = new(model.AcceptorDiagnosticRequest1)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV01) AddSecurityTrailer() *model.ContentInformationType3 {
	a.SecurityTrailer = new(model.ContentInformationType3)
	return a.SecurityTrailer
}
