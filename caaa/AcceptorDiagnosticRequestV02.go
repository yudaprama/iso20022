package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300102 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.02 Document"`
	Message *AcceptorDiagnosticRequestV02 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300102) AddMessage() *AcceptorDiagnosticRequestV02 {
	d.Message = new(AcceptorDiagnosticRequestV02)
	return d.Message
}

// The AcceptorDiagnosticRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check the end-to-end communication, to test the availability of this acquirer, or to validate the security environment.
type AcceptorDiagnosticRequestV02 struct {

	// Diagnostic request message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *model.AcceptorDiagnosticRequest2 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticRequestV02) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV02) AddDiagnosticRequest() *model.AcceptorDiagnosticRequest2 {
	a.DiagnosticRequest = new(model.AcceptorDiagnosticRequest2)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
