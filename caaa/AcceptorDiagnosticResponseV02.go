package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.02 Document"`
	Message *AcceptorDiagnosticResponseV02 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400102) AddMessage() *AcceptorDiagnosticResponseV02 {
	d.Message = new(AcceptorDiagnosticResponseV02)
	return d.Message
}

// The AcceptorDiagnosticResponse message is sent by the acquirer (or its agent) to provide to the acceptor the result of the diagnostic request.
type AcceptorDiagnosticResponseV02 struct {

	// Diagnostic response message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *model.AcceptorDiagnosticResponse2 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticResponseV02) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV02) AddDiagnosticResponse() *model.AcceptorDiagnosticResponse2 {
	a.DiagnosticResponse = new(model.AcceptorDiagnosticResponse2)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
