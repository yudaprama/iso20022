package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300105 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.05 Document"`
	Message *AcceptorDiagnosticRequestV05 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300105) AddMessage() *AcceptorDiagnosticRequestV05 {
	d.Message = new(AcceptorDiagnosticRequestV05)
	return d.Message
}

// The AcceptorDiagnosticRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check the end-to-end communication, to test the availability of this acquirer, or to validate the security environment.
type AcceptorDiagnosticRequestV05 struct {

	// Diagnostic request message management information.
	Header *model.Header30 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *model.AcceptorDiagnosticRequest5 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorDiagnosticRequestV05) AddHeader() *model.Header30 {
	a.Header = new(model.Header30)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV05) AddDiagnosticRequest() *model.AcceptorDiagnosticRequest5 {
	a.DiagnosticRequest = new(model.AcceptorDiagnosticRequest5)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV05) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
