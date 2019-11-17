package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.005.001.01 Document"`
	Message *AcceptorCancellationRequestV01 `xml:"AccptrCxlReq"`
}

func (d *Document00500101) AddMessage() *AcceptorCancellationRequestV01 {
	d.Message = new(AcceptorCancellationRequestV01)
	return d.Message
}

// Scope
// The AcceptorCancellationRequest message is sent by a card acceptor to cancel a successfully completed card payment transaction. The message can be sent directly to the acquirer or through an agent.
// Usage
// The AcceptorCancellationRequest message is used when the card acceptor is unaware of the cancellation of the transaction by the acquirer.
type AcceptorCancellationRequestV01 struct {

	// Cancellation request message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the cancellation request.
	CancellationRequest *model.AcceptorCancellationRequest1 `xml:"CxlReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationRequestV01) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorCancellationRequestV01) AddCancellationRequest() *model.AcceptorCancellationRequest1 {
	a.CancellationRequest = new(model.AcceptorCancellationRequest1)
	return a.CancellationRequest
}

func (a *AcceptorCancellationRequestV01) AddSecurityTrailer() *model.ContentInformationType3 {
	a.SecurityTrailer = new(model.ContentInformationType3)
	return a.SecurityTrailer
}
