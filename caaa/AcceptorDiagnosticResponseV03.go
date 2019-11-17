package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400103 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Document"`
	Message *AcceptorDiagnosticResponseV03 `xml:"AccptrDgnstcRspn"`
}

func (d *Document01400103) AddMessage() *AcceptorDiagnosticResponseV03 {
	d.Message = new(AcceptorDiagnosticResponseV03)
	return d.Message
}

// The AcceptorDiagnosticResponse message is sent by the acquirer (or its agent) to provide to the acceptor the result of the diagnostic request.
type AcceptorDiagnosticResponseV03 struct {

	// Diagnostic response message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to the diagnostic response.
	DiagnosticResponse *model.AcceptorDiagnosticResponse3 `xml:"DgnstcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorDiagnosticResponseV03) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorDiagnosticResponseV03) AddDiagnosticResponse() *model.AcceptorDiagnosticResponse3 {
	a.DiagnosticResponse = new(model.AcceptorDiagnosticResponse3)
	return a.DiagnosticResponse
}

func (a *AcceptorDiagnosticResponseV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
