package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400105 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.05 Document"`
	Message *AcceptorDiagnosticResponseV05 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400105) AddMessage() *AcceptorDiagnosticResponseV05 {
	d.Message = new(AcceptorDiagnosticResponseV05)
	return d.Message
}

// The AcceptorDiagnosticResponse message is sent by the acquirer (or its agent) to provide to the acceptor the result of the diagnostic request.
type AcceptorDiagnosticResponseV05 struct {

	// Diagnostic response message management information.
	Header *model.Header30 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *model.AcceptorDiagnosticResponse4 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorDiagnosticResponseV05) AddHeader() *model.Header30 {
	a.Header = new(model.Header30)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV05) AddDiagnosticResponse() *model.AcceptorDiagnosticResponse4 {
	a.DiagnosticResponse = new(model.AcceptorDiagnosticResponse4)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV05) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
