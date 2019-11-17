package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300104 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.04 Document"`
	Message *AcceptorDiagnosticRequestV04 `xml:"AccptrDgnstcReq"`
}

func (d *Document01300104) AddMessage() *AcceptorDiagnosticRequestV04 {
	d.Message = new(AcceptorDiagnosticRequestV04)
	return d.Message
}

// The AcceptorDiagnosticRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check the end-to-end communication, to test the availability of this acquirer, or to validate the security environment.
type AcceptorDiagnosticRequestV04 struct {

	// Diagnostic request message management information.
	Header *model.Header10 `xml:"Hdr"`

	// Information related to the diagnostic request.
	DiagnosticRequest *model.AcceptorDiagnosticRequest4 `xml:"DgnstcReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticRequestV04) AddHeader() *model.Header10 {
	a.Header = new(model.Header10)
	return a.Header
}

func (a *AcceptorDiagnosticRequestV04) AddDiagnosticRequest() *model.AcceptorDiagnosticRequest4 {
	a.DiagnosticRequest = new(model.AcceptorDiagnosticRequest4)
	return a.DiagnosticRequest
}

func (a *AcceptorDiagnosticRequestV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
