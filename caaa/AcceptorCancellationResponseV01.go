package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.01 Document"`
	Message *AcceptorCancellationResponseV01 `xml:"AccptrCxlRspn"`
}

func (d *Document00600101) AddMessage() *AcceptorCancellationResponseV01 {
	d.Message = new(AcceptorCancellationResponseV01)
	return d.Message
}

// Scope
// The AcceptorCancellationResponse message is sent by the acquirer to inform the card acceptor of the outcome of the cancellation process. The message can be sent directly to the acceptor or through an agent.
// Usage
// The AcceptorCancellationResponse message is used to indicate one of the possible outcomes of a cancellation process:
// - a successful cancellation;
// - a rejection from the acquirer for financial reasons;
// - a rejection from the acquirer for technical reasons.
type AcceptorCancellationResponseV01 struct {

	// Cancellation response message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the cancellation response.
	CancellationResponse *model.AcceptorCancellationResponse1 `xml:"CxlRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationResponseV01) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorCancellationResponseV01) AddCancellationResponse() *model.AcceptorCancellationResponse1 {
	a.CancellationResponse = new(model.AcceptorCancellationResponse1)
	return a.CancellationResponse
}

func (a *AcceptorCancellationResponseV01) AddSecurityTrailer() *model.ContentInformationType3 {
	a.SecurityTrailer = new(model.ContentInformationType3)
	return a.SecurityTrailer
}
